UPDATE
	streams
SET started_at = SUBSTR(started_at, 7, 4) || '-' || SUBSTR(started_at, 4, 2) || '-' || SUBSTR(started_at, 1, 2);

SELECT
	stream_id, stream_number, MAX(started_at) AS 'latest_date'
FROM streams;

SELECT
	DISTINCT(SUBSTR(started_at, 1, 4)) AS 'year'
FROM streams;

SELECT
	COUNT(*) AS 'total_teachers'
FROM teachers;

SELECT
	started_at
FROM streams
	ORDER BY started_at
	DESC LIMIT 2;

SELECT
	AVG(performance) AS 'average_performance'
FROM grades
	WHERE teacher_id = 1
	GROUP BY teacher_id;

SELECT
	teacher_id, AVG(performance) AS 'average_performance'
FROM grades
	GROUP BY teacher_id
	HAVING average_performance < 4.8;