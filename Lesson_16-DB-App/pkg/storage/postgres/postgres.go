package postgres

import (
	"context"
	"fmt"
	"thinknetica_golang_core/Lesson_16-DB-App/pkg/storage"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PG struct {
	ctx context.Context
	db  *pgxpool.Pool
}

// Инициализирует подключение к серверу postgres. Принимает контекст и строку для подключения
func New(ctx context.Context, conn string) (*PG, error) {
	var pg PG

	db, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return &pg, err
	}
	pg = PG{
		ctx: ctx,
		db:  db,
	}
	return &pg, err
}

// Закрывает подключения к postgres
func (pg *PG) Close() {
	pg.db.Close()
}

// Films Возвращает список фильмов. Если studioId = 0 - все фильмы, иначе фильтр по id-студии
func (pg *PG) Films(studioId int) (films []storage.Film, err error) {
	sql := `
		Select films.id, films.title, films.year, studios.id, studios.title
		FROM films
		JOIN studios ON studios.id = films.studio_id`

	if studioId > 0 {
		sql = fmt.Sprintln(sql, " WHERE studio_id = ", studioId)
	}
	sql = fmt.Sprintln(sql, ";")

	rows, err := pg.db.Query(pg.ctx, sql)
	if err != nil {
		return films, err
	}
	defer rows.Close()

	for rows.Next() {
		var f storage.Film
		err := rows.Scan(
			&f.ID,
			&f.Title,
			&f.Year,
			&f.StudioId,
			&f.Studio,
		)
		if err != nil {
			return nil, err
		}
		films = append(films, f)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return films, nil
}

// AddFilms добавляет в БД массив фильмов одной транзакцией.
func (pg *PG) AddFilms(films []storage.Film) error {
	// начало транзакции
	tx, err := pg.db.Begin(pg.ctx)
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback(pg.ctx)

	// пакетный запрос
	var batch = &pgx.Batch{}
	// добавление заданий в пакет
	for _, film := range films {
		batch.Queue(`INSERT INTO films(title, year, studio_id) VALUES ($1, $2, $3)`, film.Title, film.Year, film.StudioId)
	}
	// отправка пакета в БД (может выполняться для транзакции или соединения)
	res := tx.SendBatch(pg.ctx, batch)
	// обязательная операция закрытия соединения
	err = res.Close()
	if err != nil {
		return err
	}
	// подтверждение транзакции
	return tx.Commit(pg.ctx)
}

// Удаляет фильм с указанным id
func (pg *PG) DelFilm(id int) (err error) {
	// начало транзакции
	tx, err := pg.db.Begin(pg.ctx)
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback(pg.ctx)

	// пакетный запрос
	var batch = &pgx.Batch{}
	// добавление заданий в пакет
	batch.Queue(`DELETE FROM films_actors WHERE film_id = $1`, id)
	batch.Queue(`DELETE FROM films_directors WHERE film_id = $1`, id)
	batch.Queue(`DELETE FROM films WHERE id = $1`, id)
	// отправка пакета в БД (может выполняться для транзакции или соединения)
	res := tx.SendBatch(pg.ctx, batch)
	// обязательная операция закрытия соединения
	err = res.Close()
	if err != nil {
		return err
	}
	// подтверждение транзакции
	return tx.Commit(pg.ctx)
}

func (pg *PG) UpdateFilm(id int, film storage.Film) error {
	sql := `UPDATE films SET title = $1, year = $2, studio_id = $3 WHERE id = $4;`
	_, err := pg.db.Exec(pg.ctx, sql, film.Title, film.Year, film.StudioId, id)
	return err
}
