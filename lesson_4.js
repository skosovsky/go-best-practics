// Используем http://jsbasics.geekbrains.ru/

/*
	Вынос кода вычисления факториала в функцию
*/

// Сначала сделаем это без функции

let value1 = prompt('Введите первое число');
let value2 = prompt('Введите второе число');

let result = 1;

for (let n = 2; n <= value1; n = n + 1) {
	result1 = result1 * n;
}

let result2 = 1;

for (let n = 2; n <= value2; n = n + 1) {
	result2 = result2 * n;
}

let finalResult = result1 + result2;

console.log('Сумма факториала чисел ' + value1 + ' и ' + value2 + ' равна ' + finalResult);

// Тоже самое, но реализация с функцией

function factorial(value) {
	let result = 1;

	for (let n = 2; n <= value; n = n + 1) {
		result = result * n;
	}

	return result;
}

let value1 = prompt('Введите первое число');
let value2 = prompt('Введите второе число');
let result = factorial(value1) + factorial(value2);

console.log('Сумма факториала чисел ' + value1 + ' и ' + value2 + ' равна ' + result);

// Тоже самое, но реализация с функцией и использования рекурсии - и для 1 числа

function factorial(value) {
	if (value <= 1) {
		return 1;
	}
	return value * factorial(value - 1);
}

let value = prompt('Введите число');

console.log('Факториал числа ' + value + ' равен ' + factorial(value));

// Пример функции console.log

function log (value) {
	console.log(value);
}

log('Test');

// Пример функции генерации чисел от 0 до переданного аргумента

function genarateRandomNumber(max) {
	return Math.round(Math.random() * max);
}

console.log(genarateRandomNumber(1000));

// Нарисуем с помощью функций узор двойной шахматной доски

function drawChessboard (size, left, top, cellSize) {
	let isBlack = true;

	for (let y = 0; y < size; y = y + 1) {
		for (let x = 0; x < size; x = x + 1) {
			let color = 'white';
			if (isBlack == true) {
				color = 'black';
				isBlack = false;
			} else {
				isBlack = true;
			}
			drawRect(left + x * cellSize, top + y * cellSize, cellSize, cellSize, color);
		}
	}
}

drawChessboard(3, 100, 100, 100);
drawChessboard(3, 200, 200, 100 / 3);

// Домашнее задание 4
// Программа для вычисления сложных процентов по банковскому вкладу с помесячной детализацией, на обычный срок, а затем в 2 и в 3 раза больше - функцией

function calcIncome(depositAmount, depositMonthPercent, depositMonths, multiply) {
	let depositIncome = parseFloat(depositAmount);
	for (let n = 1; n <= depositMonths * multiply; n = n + 1) {
		depositIncome = depositIncome + depositIncome * depositMonthPercent;
	}	
	return (depositIncome - depositAmount).toFixed(3);
}

let depositAmount = parseFloat(prompt('Введите сумму вклада'));
let depositMonthPercent = parseFloat(prompt('Введите годовой процент') / 100 / 12);
let depositMonths = parseFloat(prompt('Введите количество месяцев'));

console.log ('Итого доход за ' + depositMonths * 1 + ' месяцев: ' + calcIncome(depositAmount, depositMonthPercent, depositMonths, 1));
console.log ('Итого доход за ' + depositMonths * 2 + ' месяцев: ' + calcIncome(depositAmount, depositMonthPercent, depositMonths, 2));
console.log ('Итого доход за ' + depositMonths * 3 + ' месяцев: ' + calcIncome(depositAmount, depositMonthPercent, depositMonths, 3));

// Программа для рисования доски для игры "Крестики-нолики" вписанная в себя 3 раза с уменьшением - функцией
// Прошлая реализация с квадратами здесь не реализуема, будем рисовать полосками

function drawTicTacBoard (left, top, cellSize) {
	for (let y = 1; y <= 2; y = y + 1) {
		for (let x = 1; x <= 2; x = x + 1) {
			drawRect(left, top + cellSize / 3 * x, cellSize, 1, 'black'); // рисуем горизонтальные полоски
		}
		drawRect(left + cellSize / 3 * y, top, 1, cellSize, 'black'); // рисуем вертикальные полоски
	}
}

drawTicTacBoard(10, 100, 300); // большой
drawTicTacBoard(10 + 300 / 6, 100 + 300 / 6, 300 * 2 / 3); // средний
drawTicTacBoard(10 + 300 / 3, 100 + 300 / 3, 300 * 1 / 3); // маленький

