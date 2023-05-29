SELECT DISTINCT
	courses.name,
	SUM(streams.count_students) OVER (window_course) AS total_count
FROM streams
	LEFT JOIN courses
		ON courses.course_id = streams.course_id
WINDOW window_course AS (PARTITION BY streams.course_id);

SELECT DISTINCT
	teachers.teacher_id,
	teachers.surname,
	teachers.name,
	AVG(performance) OVER (window_teacher) AS average_performance
FROM teachers
	LEFT JOIN grades
		ON grades.teacher_id = teachers.teacher_id
WINDOW window_teacher AS (PARTITION BY grades.teacher_id);

/* Для максимально быстрого выполнения запроса нужно создать индексы на все поля, что и сделано :)
Очевидно, что задача некорректно поставлена, потому что нам незвестно какие запросы являются частыми.
Кроме того, нужно учитывать ресурсы на поддержание индексов - т.е. корректнее было бы попросить
"для оптимального выполнения запроса, по скорости базы данных в целом" */
CREATE INDEX teachers_idx ON teachers(surname, name);
CREATE INDEX performance_idx ON academic_performance(performance);
CREATE UNIQUE INDEX streams_uq ON streams(number); 

/* В этом и предыдущем уроке, похожие задачи, в которых кроме прочего требуется указать "дату начала следующего потока",
но не указана логика, дату следующего потока, по отношению к чему, и потока этого преподавателя, этого курса или вообще,
поэтому не реализовано (ну и если честно не очень понятно как это реализовать - типа взять дату максимального и следующего за ним?) */
SELECT DISTINCT
	teachers.name,
	teachers.surname,
	MIN(grades.performance) OVER (window_teacher) AS 'min_performance',
	courses.name AS 'course_name_min_performance',
	MAX(grades.performance) OVER (window_teacher) AS 'max_performance',
	courses.name AS 'course_name_max_performance'
FROM teachers
	LEFT JOIN grades
		ON grades.teacher_id = teachers.teacher_id
	LEFT JOIN courses
		ON courses.course_id = (SELECT streams.course_id FROM streams WHERE streams.stream_id =
			(SELECT grades.teacher_id FROM grades WHERE grades.teacher_id = teachers.teacher_id) )
WINDOW window_teacher AS (PARTITION BY grades.teacher_id);