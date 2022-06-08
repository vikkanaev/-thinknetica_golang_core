package postgres

import (
	"context"
	"fmt"
	"thinknetica_golang_core/Lesson_16-DB-App/pkg/storage"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PG struct {
	db *pgxpool.Pool
}

// Инициализирует подключение к серверу postgres. Принимает контекст и строку для подключения
func New(conn string) (*PG, error) {
	var pg PG

	db, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}
	pg = PG{
		db: db,
	}
	return &pg, err
}

// Закрывает подключения к postgres
func (pg *PG) Close() {
	pg.db.Close()
}

// Films Возвращает список фильмов. Если studioId = 0 - все фильмы, иначе фильтр по id-студии
func (pg *PG) Films(ctx context.Context, studioId int) (films []storage.Film, err error) {
	sql := `
		Select films.id, films.title, films.year, studios.id, studios.title
		FROM films
		JOIN studios ON studios.id = films.studio_id`

	if studioId > 0 {
		sql = fmt.Sprintln(sql, " WHERE studio_id = ", studioId)
	}
	sql = fmt.Sprintln(sql, ";")

	rows, err := pg.db.Query(ctx, sql)
	if err != nil {
		return nil, err
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
func (pg *PG) AddFilms(ctx context.Context, films []storage.Film) error {
	// начало транзакции
	tx, err := pg.db.Begin(ctx)
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback(ctx)

	// пакетный запрос
	var batch = &pgx.Batch{}
	// добавление заданий в пакет
	for _, film := range films {
		batch.Queue(`INSERT INTO films(title, year, studio_id) VALUES ($1, $2, $3)`, film.Title, film.Year, film.StudioId)
	}
	// отправка пакета в БД (может выполняться для транзакции или соединения)
	res := tx.SendBatch(ctx, batch)
	// обязательная операция закрытия соединения
	err = res.Close()
	if err != nil {
		return err
	}
	// подтверждение транзакции
	return tx.Commit(ctx)
}

// Удаляет фильм с указанным id
func (pg *PG) DelFilm(ctx context.Context, id int) (err error) {
	// начало транзакции
	tx, err := pg.db.Begin(ctx)
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback(ctx)

	// пакетный запрос
	var batch = &pgx.Batch{}
	// добавление заданий в пакет
	// В данном слкчае транзакция не нужна, потому что я добавил в схему ON DELETE CASCADE
	// Но оставлю это для примера транзакции
	batch.Queue(`DELETE FROM films_actors WHERE film_id = $1`, id)
	batch.Queue(`DELETE FROM films_directors WHERE film_id = $1`, id)
	batch.Queue(`DELETE FROM films WHERE id = $1`, id)
	// отправка пакета в БД (может выполняться для транзакции или соединения)
	res := tx.SendBatch(ctx, batch)
	// обязательная операция закрытия соединения
	err = res.Close()
	if err != nil {
		return err
	}
	// подтверждение транзакции
	return tx.Commit(ctx)
}

func (pg *PG) UpdateFilm(ctx context.Context, id int, film storage.Film) error {
	sql := `UPDATE films SET title = $1, year = $2, studio_id = $3 WHERE id = $4;`
	_, err := pg.db.Exec(ctx, sql, film.Title, film.Year, film.StudioId, id)
	return err
}
