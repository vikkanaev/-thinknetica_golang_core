package storage

type Film struct {
	ID       int
	Title    string
	Year     int
	StudioId int
	Studio   string
}

type Interface interface {
	// Закрывает подключения к базе
	Close()

	// Возвращает список фильмов. Если studioId = 0 - все фильмы, иначе фильтр по id-студии
	Films(int) ([]Film, error)

	// Добавляет в базу фильмы из переданного массива
	AddFilms([]Film) error

	// Удаляет фильм с указанным id
	DelFilm(int) error

	// Обновляет данные для фильма с указанным id
	UpdateFilm(int, Film) error
}
