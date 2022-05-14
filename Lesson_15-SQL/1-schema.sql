DROP TABLE IF EXISTS films_certifications;
DROP TABLE IF EXISTS films_studios;
DROP TABLE IF EXISTS films_directors;
DROP TABLE IF EXISTS films_actors;
DROP TABLE IF EXISTS certifications;
DROP TABLE IF EXISTS studios;
DROP TABLE IF EXISTS films;
DROP TABLE IF EXISTS persons;

-- persons - люди
CREATE TABLE persons (
    id SERIAL PRIMARY KEY, -- первичный ключ
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    year_of_birth INTEGER NOT NULL DEFAULT 0,
    UNIQUE(first_name, last_name, year_of_birth) 
);

-- films - фильмы
CREATE TABLE films (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    title TEXT NOT NULL, -- название
    year INTEGER DEFAULT 0, -- Год выхода не может быть меньше 1800
    box_office INTEGER DEFAULT 0,
    UNIQUE(title, year) -- В один год не может быть двух фильмов с одинаковым названием.
);

-- Студии
CREATE TABLE studios (
  id BIGSERIAL PRIMARY KEY, -- первичный ключ
  title TEXT NOT NULL, -- название
  UNIQUE(title)
);

-- рейтинги фильмов
CREATE TABLE certifications (
  id BIGSERIAL PRIMARY KEY, -- первичный ключ
  title TEXT NOT NULL, -- название рейтинга
  UNIQUE(title)
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

-- связь между фильмами и студиями
CREATE TABLE films_studios (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    film_id BIGINT NOT NULL REFERENCES films(id),
    studio_id INTEGER NOT NULL REFERENCES studios(id),
    UNIQUE(film_id, studio_id)
);

-- связь между фильмами и рейтингами
CREATE TABLE films_certifications (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    film_id BIGINT NOT NULL REFERENCES films(id),
    certification_id INTEGER NOT NULL REFERENCES certifications(id),
    UNIQUE(film_id, certification_id)
);

