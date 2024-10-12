SELECT mo.movie_name, mr.domestic_takings, mr.international_takings FROM movies mo
JOIN movie_revenues mr ON mo.movie_id = mr.movie_id
ORDER BY mr.domestic_takings;


SELECT mo.movie_name, mr.domestic_takings FROM movies mo
JOIN movie_revenues mr USING (movie_id);


SELECT d.first_name, d.last_name, mo.movie_name, mo.release_date FROM directors d
JOIN movies mo ON d.director_id = mo.director_id
WHERE mo.movie_lang IN ('Chinese','Japanese','Korean');


SELECT mo.movie_name, mo.release_date, mr.international_takings FROM movies mo 
JOIN movie_revenues mr ON mo.movie_id = mr.movie_id
WHERE mo.movie_lang = 'English';


SELECT mo.movie_name, mr.international_takings, mr.domestic_takings FROM movies mo
JOIN movie_revenues mr ON mo.movie_id = mr.movie_id 
WHERE mr.domestic_takings IS NULL
OR mr.international_takings IS NULL
ORDER BY mo.movie_name;

-- LEFT JOINS

/*
SELECT t1.col1, t1.col2, t2.col1 FROM table1 t1
LEFT JOIN table2 t2 ON t1.col3 = t2.col3;
*/

SELECT d.director_id, d.first_name, d.last_name, mo.movie_name FROM directors d
LEFT JOIN movies mo ON d.director_id = mo.director_id;

DELETE FROM directors
WHERE director_id = '39';


SELECT * FROM directors;


SELECT d.director_id, d.first_name, d.last_name, mo.movie_name FROM directors d
FULL JOIN movies mo ON d.director_id = mo.director_id;








