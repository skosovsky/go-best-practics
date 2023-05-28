SELECT
	stream_number,
	(SELECT name FROM courses WHERE courses.course_id = streams.course_id) AS course_name,
	count_students
FROM streams
WHERE count_students >= 40;

SELECT
	(SELECT name FROM courses WHERE courses.course_id = (SELECT course_id FROM streams WHERE streams.stream_id = grades.stream_id)) AS course_name,
	(SELECT surname FROM teachers WHERE teachers.teacher_id = grades.teacher_id) || " " || (SELECT name FROM teachers WHERE teachers.teacher_id = grades.teacher_id) AS teacher,
	performance
FROM grades
	ORDER BY performance
	LIMIT 2;

SELECT
	teacher_id,
	AVG(performance) AS 'average_performance'
FROM grades
WHERE teacher_id = (SELECT teacher_id FROM teachers WHERE name = 'Николай' AND surname = "Савельев");

SELECT
		stream_id,
		(SELECT surname FROM teachers WHERE name = 'Наталья' AND surname = "Петрова") AS 'surname',
		(SELECT name FROM teachers WHERE name = 'Наталья' AND surname = "Петрова") AS 'name'
FROM streams
WHERE stream_id = (SELECT stream_id FROM grades WHERE grades.teacher_id = (SELECT teacher_id FROM teachers WHERE name = 'Наталья' AND surname = "Петрова") )
UNION
SELECT
	stream_id,
	(SELECT surname FROM teachers WHERE teachers.teacher_id = grades.teacher_id) AS 'surname',
	(SELECT name FROM teachers WHERE teachers.teacher_id = grades.teacher_id) AS 'name'
FROM grades
WHERE performance < 4.8;

SELECT
	(SELECT AVG(performance) AS 'average_performance' FROM grades GROUP BY teacher_id ORDER BY average_performance DESC LIMIT 1)
-
	(SELECT AVG(performance) AS 'average_performance' FROM grades GROUP BY teacher_id ORDER BY average_performance ASC LIMIT 1)
AS 'difference_between_two_teachers';