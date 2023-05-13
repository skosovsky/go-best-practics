// Используем http://jsbasics.geekbrains.ru/

/*
	Эта программа определяет максимальное число из двух чисел
*/

let a = prompt('Введите число a');
let b = prompt('Введите число b');

if (a > b) { // если
	console.log('Максимальное число: a');
} else if (a == b) { // но если
	console.log('Числа равны');
} else { // в противном  случае
	console.log('Максимальное число: b');
}

/*
	Операторы сравнения:
	> - больше
	>= - больше или равно
	< - меньше
	<= - меньше или равно
	== - равно
	!= - не равно
*/

let a = prompt('Введите число a');
let result = a % 2; // у четных чисел остаток равен 0 

if (result == 0) { // если
	console.log('Число четное');
} else { // в противном  случае
	console.log('Число не четное');
}

// Упрощаем

let a = prompt('Введите число a');

if (a % 2 == 0) { // если
	console.log('Число четное');
} else { // в противном  случае
	console.log('Число не четное');
}

// Рисуем квадрат

let size = prompt('Укажите размер квадрата');

if (size > 200) {
	size = 200;
}

drawRect(10, 10, size, size, 'black');

// Домашнее задание 2
// Вывод в консоль ответа на вопрос - високосный ли год
// {} - это извращенец придумал
let userYear = prompt('Введите год');

if (userYear % 4 == 0) {
	if (userYear % 100 == 0) {
		if (userYear % 400 == 0) {
        	console.log('Год високосный');
        } else {
        	console.log('Год не високосный');
        }
	} else {
		console.log('Год високосный');
    }
} else {
	console.log('Год не високосный');
}

// Нарисовать груг или квадрат по параметрам пользователя
let userForm = prompt('Введите 1, если риусем круг или 2, если - квадрат'); // выбор фигуры

if (userForm == 1) {
	let formSize = prompt('Введите радиус круга');
	drawCircle(formSize * 1.1, formSize * 1.1, formSize, 'black'); // выводим круг с умными отступами
} else if (userForm == 2) {
	let formSize = prompt('Введите размер квадрата');
	drawRect(formSize * 0.1, formSize * 0.1, formSize, formSize, 'black'); // выводим квадрат с умными отступами
} else {
	console.log('Указано неверное число на первом шаге'); // выводим ошибку если выбор фигуры неверный
}
