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
INSERT INTO films (id, title, year, box_office, studio_id, certification) VALUES (0, 'The Godfather', 1972, 268500000, 0, 'PG-10');
-- Фильм The Scarface
INSERT INTO films (id, title, year, box_office, studio_id, certification) VALUES (1, 'Scarface', 1983, 65884703, 1, 'PG-13');
ALTER SEQUENCE films_id_seq RESTART WITH 100;

-- Фильм The Godfather
-- Привязываем актеров
INSERT INTO films_actors (film_id, actor_id) VALUES (0, 1);
INSERT INTO films_actors (film_id, actor_id) VALUES (0, 2);
INSERT INTO films_actors (film_id, actor_id) VALUES (0, 3);
-- Привязываем режисера
INSERT INTO films_directors (film_id, director_id) VALUES (0, 0);

-- Фильм Scarface
-- Привязываем актеров
INSERT INTO films_actors (film_id, actor_id) VALUES (1, 2);
INSERT INTO films_actors (film_id, actor_id) VALUES (1, 5);
INSERT INTO films_actors (film_id, actor_id) VALUES (1, 6);
-- Привязываем режисера
INSERT INTO films_directors (film_id, director_id) VALUES (1, 4);



