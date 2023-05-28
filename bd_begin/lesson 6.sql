SELECT
	streams.stream_number,
	courses.name,
	streams.started_at
FROM streams
	INNER JOIN courses
		ON streams.course_id = courses.course_id;

SELECT
	courses.name,
	SUM(streams.count_students) AS 'total_students'
FROM streams
	INNER JOIN courses
		ON courses.course_id = streams.course_id
GROUP BY courses.name;

SELECT
	teachers.teacher_id,
	teachers.surname,
	teachers.name,
	AVG(grades.performance) AS 'average_performance'
FROM teachers
	LEFT JOIN grades
		ON grades.teacher_id = teachers.teacher_id
GROUP BY teachers.teacher_id;

SELECT
	teachers.name,
	teachers.surname,
	MIN(grades.performance) AS 'min_performance',
	courses.name AS 'course_name_min_performance',
	MAX(grades.performance) AS 'max_performance',
	courses.name AS 'course_name_max_performance'
FROM teachers
	LEFT JOIN grades
		ON grades.teacher_id = teachers.teacher_id
	LEFT JOIN courses
		ON courses.course_id = (SELECT streams.course_id FROM streams WHERE streams.stream_id =
			(SELECT grades.teacher_id FROM grades WHERE grades.teacher_id = teachers.teacher_id) )
GROUP BY teachers.teacher_id;