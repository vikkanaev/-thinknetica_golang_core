DROP TABLE IF EXISTS films_directors;
DROP TABLE IF EXISTS films_actors;
DROP TABLE IF EXISTS films;
DROP TABLE IF EXISTS studios;
DROP TABLE IF EXISTS persons;
DROP TYPE IF EXISTS certification;

-- persons - люди
CREATE TABLE persons (
    id SERIAL PRIMARY KEY, -- первичный ключ
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    year_of_birth INTEGER NOT NULL DEFAULT 0,
    UNIQUE(first_name, last_name, year_of_birth) 
);

-- рейтинги фильмов
CREATE TYPE certification AS ENUM ('PG-10', 'PG-13', 'PG-18');

-- Студии
CREATE TABLE studios (
  id BIGSERIAL PRIMARY KEY, -- первичный ключ
  title TEXT NOT NULL, -- название
  UNIQUE(title)
);

-- films - фильмы
CREATE TABLE films (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    title TEXT NOT NULL, -- название
    year INTEGER DEFAULT 0, -- Год выхода не может быть меньше 1800
    box_office INTEGER DEFAULT 0,
    studio_id INTEGER REFERENCES studios(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    certification certification,
    UNIQUE(title, year) -- В один год не может быть двух фильмов с одинаковым названием.
);

-- связь между фильмами и актерами
CREATE TABLE films_actors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    film_id BIGINT NOT NULL REFERENCES films(id),
    actor_id INTEGER NOT NULL REFERENCES persons(id),
    UNIQUE(film_id, actor_id)
);

-- связь между фильмами и режисерами
CREATE TABLE films_directors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    film_id BIGINT NOT NULL REFERENCES films(id),
    director_id INTEGER NOT NULL REFERENCES persons(id),
    UNIQUE(film_id, director_id)
);

