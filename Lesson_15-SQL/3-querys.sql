-- выборка фильмов с названием студии;
Select films.title, studios.title
FROM films
JOIN studios ON studios.id = films.studio_id;

--      title     |       title
-- ---------------+--------------------
--  The Godfather | Paramount Pictures
--  Scarface      | Universal Pictures
-- (2 rows)

-- выборка фильмов для некоторого актёра;
Select films.title, persons.first_name, persons.last_name
FROM films
JOIN films_actors ON films.id = films_actors.film_id
JOIN persons ON persons.id = films_actors.actor_id
WHERE persons.id = (Select id from persons where first_name ='Al' and last_name='Pacino');

--      title     | first_name | last_name
-- ---------------+------------+-----------
--  The Godfather | Al         | Pacino
--  Scarface      | Al         | Pacino
-- (2 rows)

-- подсчёт количества фильмов для актёра;
Select persons.first_name, persons.last_name, count(films.id) as films_count 
FROM films
JOIN films_actors ON films.id = films_actors.film_id
JOIN persons ON persons.id = films_actors.actor_id
GROUP BY persons.id;

--  first_name | last_name | films_count
-- ------------+-----------+-------------
--  Steven     | Bauer     |           1
--  Al         | Pacino    |           2
--  James      | Caan      |           1
--  Michelle   | Pfeiffer  |           1
--  Marlon     | Brando    |           1
-- (5 rows)