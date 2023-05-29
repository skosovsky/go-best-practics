CREATE VIEW courses_table AS SELECT
	courses.name,
	MAX(streams.stream_number) AS 'last_stream_course',
	MAX(streams.started_at) AS 'date_stream_course',
	AVG(grades.performance) AS 'average_performance_course'
FROM courses
	LEFT JOIN streams
		ON courses.course_id = streams.course_id
	LEFT JOIN grades
		ON streams.stream_id = grades.stream_id
GROUP BY courses.course_id;

BEGIN TRANSACTION;
	DELETE FROM grades WHERE teacher_id = 3;
	DELETE FROM teachers WHERE teacher_id = 3;
COMMIT;

CREATE TRIGGER check_grade BEFORE INSERT
ON grades
BEGIN
	SELECT CASE
	WHEN
		NEW.performance NOT BETWEEN 0 AND 5
	THEN
		RAISE(ABORT,'Wrong format for performance')
	END;
END;

CREATE TRIGGER check_stream BEFORE INSERT
ON streams
BEGIN
	SELECT CASE
	WHEN
		NEW.started_at <= DATE(CURRENT_TIMESTAMP)
		OR (NEW.stream_number <= (SELECT max(stream_number) FROM streams))
	THEN
		RAISE(ABORT, 'Wron start date or stream number')
	END;
END;