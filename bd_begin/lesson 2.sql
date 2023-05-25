/* create students.db */

.open students.db

CREATE TABLE courses(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    lessons_amount INTEGER
);

CREATE TABLE streams(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    number INTEGER NOT NULL UNIQUE,
    course_id INTEGER NOT NULL,
    start_date TEXT NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(id) /* Many to One - много потоков на один курс */
);

CREATE TABLE students(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    surname TEXT NOT NULL,
    name TEXT NOT NULL
);

CREATE TABLE grades(
    student_id INTEGER NOT NULL,
    stream_id INTEGER NOT NULL,
    grade REAL NOT NULL,
    PRIMARY KEY (student_id, stream_id),
    FOREIGN KEY (student_id) REFERENCES students(id), /* Many to One - много оценок на одного студента */
    FOREIGN KEY (stream_id) REFERENCES streams(id) /* Many to One - много оценок на один поток */
);

/* create teachers.db */

.open teachers.db

CREATE TABLE teachers(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	surname TEXT NOT NULL,
	email TEXT
);

CREATE TABLE courses(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
);

CREATE TABLE streams(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	course_id INTEGER NOT NULL,
	number INTEGER NOT NULL,
	date_start TEXT NOT NULL,
	count_students INTEGER,
	FOREIGN KEY (course_id) REFERENCES courses(id) /* Many to One - много потоков на один курс */
);

CREATE TABLE grades(
	teachers_id INTEGER NOT NULL,
	stream_id INTEGER NOT NULL,
	stream_grade INTEGER NOT NULL,
	PRIMARY KEY (teachers_id, stream_id),
	FOREIGN KEY (teachers_id) REFERENCES teachers(id), /* Many to One - много оценок на одного учителя */
	FOREIGN KEY (stream_id) REFERENCES streams(id) /* One to One - одна оценка на один поток */
);

/* create services.db */

.open services.db

CREATE TABLE users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	surname TEXT NOT NULL,
	age INTEGER
);

CREATE TABLE services(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE
);

CREATE TABLE statuses(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	number INTEGER NOT NULL UNIQUE,
	description TEXT NOT NULL
);

CREATE TABLE involvements(
	user_id INTEGER NOT NULL,
	service_id INTEGER NOT NULL,
	status_id INTEGER NOT NULL,
	PRIMARY KEY (user_id, service_id),
	FOREIGN KEY (user_id) REFERENCES users(id), /* One to Many - один пользователь на много вовлечений */
	FOREIGN KEY (service_id) REFERENCES services(id), /* Many to Many - много сервисов на много вовлечений */
	FOREIGN KEY (status_id) REFERENCES statuses(id) /* Many to One - много статусов на много вовлечений */
);
