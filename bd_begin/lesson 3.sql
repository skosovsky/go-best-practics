-- table teachers.db

.open teachers.db

-- rename key "По возможности не используйте id в качестве первичного идентификатора таблицы."

ALTER TABLE streams RENAME COLUMN id TO teacher_id;
ALTER TABLE streams RENAME COLUMN id TO course_id;
ALTER TABLE streams RENAME COLUMN id TO stream_id;

-- rename and insert new columns

ALTER TABLE streams RENAME COLUMN date_start TO started_at;
ALTER TABLE streams ADD COLUMN finished_at TEXT;
ALTER TABLE grades RENAME COLUMN stream_grade TO performance;

-- but, because column order not comfortable, i'm delete this table and create new

DROP TABLE streams;
CREATE TABLE streams(
	stream_id INTEGER PRIMARY KEY AUTOINCREMENT,
	course_id INTEGER NOT NULL,
	stream_number INTEGER NOT NULL, -- because number is reserved words "Убедитесь в том, что имени нет в списке зарезервированных ключевых слов."
	started_at TEXT NOT NULL,
	finished_at TEXT,
	count_students INTEGER,
	FOREIGN KEY (course_id) REFERENCES courses(id) /* Many to One - много потоков на один курс */
);

-- then create new data to teachers

INSERT INTO teachers (name, surname, email) VALUES
('Николай', 'Савельев', 'saveliev.n@mai.ru'),
('Наталья', 'Петрова', 'petrova.n@yandex.ru'),
('Елена', 'Малышева', 'malisheva.e@google.com');

-- test data for teachers

.mode column
.header on

SELECT * FROM teachers;

-- then create new data to courses

INSERT INTO courses (name) VALUES
('Базы данных'),
('Основы Python'),
('Linux. Рабочая станция');

-- test data for courses

SELECT * FROM courses;

-- then create new data to streams

INSERT INTO streams (course_id, stream_number, started_at, count_students) VALUES
(3, 165, '18.08.2020', 34),
(2, 178, '02.10.2020', 37),
(1, 203, '12.11.2020', 35),
(1, 210, '03.12.2020', 41);

-- test data for streams

SELECT * FROM streams;

-- then create new data to grades

INSERT INTO grades (teacher_id, stream_id, performance) VALUES
(3, 1, 4.7),
(2, 2, 4.9),
(1, 3, 4.8),
(1, 4, 4.9);

-- test data for grades

SELECT * FROM grades;

-- change type stream_id from INTEGER to REAL in grades

ALTER TABLE grades RENAME TO grades_old;
CREATE TABLE grades(
	teacher_id INTEGER NOT NULL,
	stream_id REAL NOT NULL,
	performance INTEGER NOT NULL,
	PRIMARY KEY (teacher_id, stream_id),
	FOREIGN KEY (teacher_id) REFERENCES teachers(id), /* Many to One - много оценок на одного учителя */
	FOREIGN KEY (stream_id) REFERENCES streams(id) /* One to One - одна оценка на один поток */
);

INSERT INTO grades (teacher_id, stream_id, performance)
  SELECT teacher_id, 
         stream_id, 
         performance
  FROM grades_old;

DROP TABLE grades_old;

-- test data for grades

SELECT * FROM grades;

