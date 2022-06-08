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
    film_id BIGINT NOT NULL REFERENCES films(id) ON DELETE CASCADE,
    actor_id INTEGER NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    UNIQUE(film_id, actor_id)
);

-- связь между фильмами и режисерами
CREATE TABLE films_directors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    film_id BIGINT NOT NULL REFERENCES films(id) ON DELETE CASCADE,
    director_id INTEGER NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    UNIQUE(film_id, director_id)
);


-- Студии
INSERT INTO studios (id, title) VALUES (0, 'Paramount Pictures');
INSERT INTO studios (id, title) VALUES (1, 'Universal Pictures');
ALTER SEQUENCE studios_id_seq RESTART WITH 100;

-- Фильм The Godfather
-- Режисер 
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (0, 'Francis Ford', 'Coppola', 1939);
-- Актеры 
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (1, 'Marlon', 'Brando', 1924);
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (2, 'Al', 'Pacino', 1940);
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (3, 'James', 'Caan', 1940);

-- Фильм Scarface
-- Режисер
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (4, 'Brian', 'De Palma', 1939);
-- Актеры
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (5, 'Michelle', 'Pfeiffer', 1924);
INSERT INTO persons (id, first_name, last_name, year_of_birth) VALUES (6, 'Steven', 'Bauer', 1940);
ALTER SEQUENCE persons_id_seq RESTART WITH 100;


-- Фильм The Godfather
INSERT INTO films (id, title, year, box_office, studio_id, certification) VALUES (1, 'The Godfather', 1972, 268500000, 0, 'PG-10');
-- Фильм The Scarface
INSERT INTO films (id, title, year, box_office, studio_id, certification) VALUES (2, 'Scarface', 1983, 65884703, 1, 'PG-13');
ALTER SEQUENCE films_id_seq RESTART WITH 100;

-- Фильм The Godfather
-- Привязываем актеров
INSERT INTO films_actors (film_id, actor_id) VALUES (1, 1);
INSERT INTO films_actors (film_id, actor_id) VALUES (1, 2);
INSERT INTO films_actors (film_id, actor_id) VALUES (1, 3);
-- Привязываем режисера
INSERT INTO films_directors (film_id, director_id) VALUES (1, 0);

-- Фильм Scarface
-- Привязываем актеров
INSERT INTO films_actors (film_id, actor_id) VALUES (2, 2);
INSERT INTO films_actors (film_id, actor_id) VALUES (2, 5);
INSERT INTO films_actors (film_id, actor_id) VALUES (2, 6);
-- Привязываем режисера
INSERT INTO films_directors (film_id, director_id) VALUES (2, 4);
