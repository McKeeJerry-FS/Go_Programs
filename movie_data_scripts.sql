-- Joining Tables

/*
    How to use UNION and UNION ALL
	Learn about Intersect and Except
*/

-- SQL Joins Explained

/*
	Join queries allow us to select data from multiple tables
	There needs to be a related column of data between tables
	These columns are usually Primary and Foreign Key columns
	Create a join between tables on the related columns
	Different types of joins include: Inner Joins, Left Joins, Right Joins, Full Joins
*/

-- INNER JOIN: Will only return rows of data with matching values in both tables
-- LEFT JOIN: Will return all rows of data in the left table and ONLY MATCHING ROWS of data in the right table
-- RIGHT JOIN: (RIGHT OUTER JOIN) Will return all rows of data from the right table and ONLY MATCHING ROWS of data from the left table, not very common
-- FULL JOIN: (Full Outer Join) Will only return all rows of data

-- INNER JOIN: Part 1
/*
SELECT table1.col1, table1.col2, table2.col1 FROM table1
INNER JOIN table2 table1.col3 = table2.col3;
*/

SELECT * FROM directors;
SELECT * FROM movies;

INSERT INTO directors(first_name, last_name, date_of_birth, nationality)
VALUES ('Christopher', 'Nolan', '1970-07-30', 'British');

SELECT directors.director_id, directors.first_name, directors.Last_name, movies.movie_name
FROM directors
INNER JOIN movies ON directors.director_id = movies.director_id;


SELECT directors.director_id, directors.first_name, directors.Last_name, movies.movie_name
FROM directors
INNER JOIN movies ON directors.director_id = movies.director_id
WHERE movies.movie_lang = 'Japanese';


SELECT directors.director_id, directors.first_name, directors.Last_name, movies.movie_name
FROM directors
INNER JOIN movies ON directors.director_id = movies.director_id
WHERE movies.movie_lang = 'Japanese'
ORDER BY movies.movie_length;

-- INNER JOIN: Part 2
/*
SELECT table1.col1, table1.col2, table2.col1 FROM table1
INNER JOIN table2 ON table1.col3 = table2.col3;

SELECT t1.col1, t1.col2, t2.col1 FROM table1 t1
JOIN table2 t2 ON t1.col3 = t2.col3;
*/

SELECT directors.director_id, directors.first_name, directors.Last_name, movies.movie_name
FROM directors
INNER JOIN movies ON directors.director_id = movies.director_id
WHERE movies.movie_lang = 'Japanese'
ORDER BY movies.movie_length;

-- Shortened way to write the query
SELECT t1.director_id, t1.first_name, t1.last_name, t2.movie_name
FROM directors AS t1
JOIN movies AS t2 ON t1.director_id = t2.director_id
WHERE t2.movie_lang = 'Japanese'
ORDER BY t2.movie_length;



SELECT mo.movie_name, mr.domestic_takings, mr.international_takings FROM movies mo
JOIN movie_revenues mr ON mo.movie_id = mr.movie_id
ORDER BY mr.domestic_takings;

-- INNER JOINS: PART 3
-- Only when the joining columns have the same name! (USING)

/*
SELECT t1.col1, t1.col2, t2.col1 FROM table t1
JOIN table2 t2 USING (col3);
*/









