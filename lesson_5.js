// Используем http://jsbasics.geekbrains.ru/

/*
	Массив и его методы
*/

let array = [100, 120, 130]; // объявлем массив и присваиваем ему значения
console.log(array); // вывести все значения массива, через запятую
console.log(array[0]); // вывести первый элемент массива (с 0 начинается)
console.log(array[array.lenght -1]); // узнать длину массива и через это узнать значение последнего элемента массива

array.push(200); // добавить 200 в конец массива
console.log(array);
array.pop(); // удалить последний элемент массива
console.log(array);
array.shift(); // удалить первый элемент массива
console.log(array);
array.unshift(77); // добавить 77 в начало массива
console.log(array);

console.log(array[2]); // вывести значение элемента с индексом 2 (3-й)

// массив может содержать любые данные и типы данных

let array = ['Василий', 'съел', 7, 'яблок', 'это', true, 'или', false, '?'];
console.log(array.join(' ')); // join используется для замены , на другое значение
console.log(array.reverse().join(' ')); // reverse выводит элементы массива в обратном порядке

// найдем максимальное число из 3-х чисел (считаем что все числа разные)

let a = prompt('Введите число a');
let b = prompt('Введите число b');
let c = prompt('Введите число c');

if (a > b && a > c) { // Логические операторы: && - И, || - ИЛИ, ! - НЕ
	console.log('Максимальное число: a');
} else if (b > a && b > c) {
	console.log('Максимальное число: b');
} else { 
	console.log('Максимальное число: c');
}

// тип переменной

let input = prompt('Введите числа через запятую');

console.log(typeof input); // typeof выводит тип переменной

// поиск максимального числа из множества чисел с помощью массива

let input = prompt('Введите числа через запятую');
let array = input.split(',');

for (let i = 0; i < array.length; i++) { // ++ означает +1 
	array[i] = +array[i]; // проиводим к числовому типу каждое значение массива
}

let result = array[0];

for (let i = 1; i < array.length; i++) {
	if (array[i] > result) {
		result = array[i];
	}
}

console.log('Максимальное число: ' + result);

// Выведем график изменения цены из массива

function findMax(array) { // сделаем функцию для вычисления максимального числа массива
	let result = array[0];

	for (let i = 1; i < array.length; i++) {
		if (array[i] > result) {
			result = array[i];
		}
	}
	return result;
}


let prices = [100, 110, 140, 100, 150, 200, 130, 120];
let max = findMax(prices); 

for (let i = 1; i < prices.length; i++) {
	drawLine(i * 10, max - prices[i - 1], i * 10 + 10, max - prices[i], 'black', 1);
}

// Домашнее задание 5
// Программа для подсчета сумм покупок пользователя

let input = prompt('Введите суммы покупок из чека через запятую');
let array = input.split(',');

for (let i = 0; i < array.length; i++) {
	array[i] = +array[i];
}

let result = array[0];

for (let i = 1; i < array.length; i++) {
	result = result + array[i];
}

console.log('Сумма ваших покупок ' + result + ' рублей');

// Рисуем график функцией с возможность задавать свои цены и выводить высоту графика

function findMax(array) { // фунция поиска максимального значения в массиве, для вывода рекомендуемого значения высоты
	let result = array[0];

	for (let i = 1; i < array.length; i++) {
		if (array[i] > result) {
			result = array[i];
		}
	}
	return result;
}

let input = prompt('Введите список цен через запятую');
let prices = input.split(',');

for (let i = 0; i < prices.length; i++) { // приводим массив из строки в цифры
	prices[i] = +prices[i];
}

let height = prompt('Укажите максимальную высоту графика, рекомендуемое значение ' + findMax(prices));

for (let i = 1; i < prices.length; i++) {
	drawLine(i * 10, height - prices[i - 1], i * 10 + 10, height - prices[i], 'black', 1);
}

// Рисуем график функцией с новой возможностью, указывать ширину графика и вывода сетки графика
// В постановке задачи ошибка - шаг вычисляется как кол-во цен / ширина, по факту расчет делается наоборот ширина / кол-во цен
// Заодно немного исправил рисование графика, на видео не корректно сделали, первая точка не с 0 начиналась

function findMax(array) { // фунция поиска максимального значения в массиве, для вывода рекомендуемого значения высоты
	let result = array[0];

	for (let i = 1; i < array.length; i++) {
		if (array[i] > result) {
			result = array[i];
		}
	}
	return result;
}

function drawGrid (pricesQuantity, pricesMax, height, width) { // сделаем функцию для рисования сетки графика
	for (let y = 0; y < pricesQuantity; y++) {
		for (let x = 0; x <= 4; x++) {
			drawLine(0, height / 4 * x, width, height / 4 * x, 'gray', 1); // рисуем горизонтальные полоски
		}
		drawLine(width / (pricesQuantity - 1) * y, 0, width / (pricesQuantity - 1) * y, height, 'gray', 1); // рисуем вертикальные полоски
	}
}

let input = prompt('Введите список цен через запятую');
let prices = input.split(',');

for (let i = 0; i < prices.length; i++) { // приводим массив из строки в цифры
	prices[i] = +prices[i];
}

let height = +prompt('Укажите высоту графика, рекомендуемое значение ' + findMax(prices));
let width = +prompt('Укажите ширину графика, минимальное значение ' + prices.length * 10);

drawGrid(prices.length, findMax(prices), height, width); // большой

for (let i = 0; i < prices.length; i++) {
	drawLine(i * width / (prices.length - 1), height - prices[i], i * width / (prices.length - 1) + width / (prices.length - 1), height - prices[i + 1], 'black', 3);
}

// Реализуем текстовый вариант игры "Рулетка"

let balance = 1000;
let bet;
let betOn;
let spinResult;
let betOnResult;
alert('Добро пожаловать в игру "Рулетка". На вашем счету 1000$.');

function findNumber(spinResult) { // считаем черное или красное
	let result;
	for (let i = 0; i < [2, 4, 6, 8, 10, 11, 13, 15, 17, 20, 22, 24, 26, 28, 29, 31, 33, 35].length - 1; i++) {
		if ([2, 4, 6, 8, 10, 11, 13, 15, 17, 20, 22, 24, 26, 28, 29, 31, 33, 35][i] == spinResult) {
			result = 'black';
		}
	}
	if (result == undefined) {
		result = 'red';
	}
	return result;
}

while (balance > 0) {

	while (true) { // проверяем корректность суммы
		bet = prompt('Текущий баланс: ' + balance + '$ . Введите сумму ставки. Для выхода введите quit');
		if (bet == 'quit') {
			break;
		} else if (bet <= balance && bet > 0) {
			break;
		} 
	}

	if (bet == 'quit') { // окончательный выход
		break;
	}

	while (true) { // проверяем корректность ставки
		betOn = prompt('Делайте вашу ставку. Выберите число от 1 до 36 или укажите цвет black или red');
		if (betOn == 'black' || betOn == 'red') {
			break;
		} else if (betOn >= 1 && betOn <=36) {
			break;
		}
	}

	spinResult = Math.round(Math.random() * 35) + 1; // крутим рулетку

	if (betOn == 'black') { // проверяем результаты
		if (betOn == findNumber(spinResult)) {
			betOnResult = true;
		}
	} else if (betOn == 'red') {
		if (betOn == findNumber(spinResult)) {
			betOnResult = true;
		}
	} else if (betOn == spinResult) {
		betOnResult = true;
	} else {
		betOnResult = false;
	}

	if (betOnResult == true) { // выводим результаты
		if (betOn == 'black' || betOn == 'red') {
			alert('Ваша ставка ' + bet + '$ на ' + betOn + ' сыграла! Ваш выигрыш: ' + bet * 2 + '$' );
			balance = balance + bet * 2;
			betOnResult = false;
		} else {
			alert('Ваша ставка ' + bet + '$ на ' + betOn + ' сыграла! Ваш выигрыш: ' + bet * 35 + '$' );
			balance = balance + bet * 35;
			betOnResult = false;
		}
	} else {
		alert('Ваша ставка ' + bet + '$ на ' + betOn + ' не сыграла! Ваш проигрыш: ' + bet + '$' );
		balance = balance - bet;
	}
}

alert('Ждем вас вновь. На вашем счету ' + balance + '$');





