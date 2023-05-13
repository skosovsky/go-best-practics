// Используем http://jsbasics.geekbrains.ru/

/*
	Эта программа играет с пользователем в игру "Угадай число"
*/

let target = Math.round(Math.random() * 10);

while (true) {
	variant = prompt('Угадай число'); // в примере тут указано let variant, но у меня после break завершается программа, без вывода сообщения
  
	console.log('Ты попробовал число ' + variant);

	if (variant > target) {
		console.log('Я загадал число меньше, чем ' + variant);
	} else if (variant < target) {
		console.log('Я загадал число больше, чем ' + variant);
	} else {
      	break;
	}
}

console.log('Ты угадал, я загадал число ' + variant);

// Вариант программы, но с циклом с постусловием

let target = Math.round(Math.random() * 10);
let variant;

do { // аналогично, вместо do, здесь можно указать сразу while
	variant = prompt('Угадай число');
  
	console.log('Ты попробовал число ' + variant);

	if (variant > target) {
		console.log('Я загадал число меньше, чем ' + variant);
	} else if (variant < target) {
		console.log('Я загадал число больше, чем ' + variant);
	} 
} while (variant != target);

console.log('Ты угадал, я загадал число ' + variant);

// Считаем факториал числа

let value = prompt('Введите число для расчета факториала');
let result = 1; // присваиваем результату 1, чтобы рассчитывался и факториал чисел 0 и 1, равной 1
let n = 2; // присваиваем n 2, чтобы цикл срабатывал с минимального 2 (а 0! и 1! учетны ранее)

while (n <= value) {
	result = result * n; // увеличиваем result на следующее значение n
	n = n + 1; // увеличиваем счетчик на 1
}

console.log('Факториал числа ' + value + ' равен ' + result);

// Вариант программы, но с циклом for (переменная; условия выхода; действие при выполнении цикла)

let value = prompt('Введите число для расчета факториала');
let result = 1; // присваиваем результату 1, чтобы рассчитывался и факториал чисел 0 и 1, равной 1

for (let n = 2; n <= value; n = n + 1) {
	result = result * n; // увеличиваем result на следующее значение n
}

console.log('Факториал числа ' + value + ' равен ' + result);

// Рисуем шахматную доску 3х3

let isBlack = true;

for (let y = 0; y < 3; y = y + 1) {
	console.log('Новая строка');
	for (let x = 0; x < 3; x = x + 1) {
		let color = 'white';
		if (isBlack == true) { // если тут добавить проверку на четность, то можно было и 4х4 доску сделать
			color = 'black';
			isBlack = false;
		} else {
			isBlack = true;
		}
		console.log ('y = ' + y + ', x = ' + x + ', color = ' + color);
		drawRect(x * 50, y * 50, 50, 50, color);
	}
}

// Домашнее задание 3
// Программа для вычисления сложных процентов по банковскому вкладу с помесячной детализацией
// Основной геморой с типами данных, в консоле числа хотят быть текстом, не хватает дебага

let depositAmount = parseFloat(prompt('Введите сумму вклада'));
let depositMonthPercent = parseFloat(prompt('Введите годовой процент') / 100 / 12); // сразу рассчитываем месячный процент
let depositMonths = parseFloat(prompt('Введите количество месяцев'));
let depositIncome = parseFloat(depositAmount);

for (let n = 1; n <= depositMonths; n = n + 1) {
	console.log (n + '. ' + n + '-й месяц: ' + depositIncome.toFixed(3) + ' + ' + (depositIncome * depositMonthPercent).toFixed(3) + ' = ' + (depositIncome + depositIncome * depositMonthPercent).toFixed(3));
	depositIncome = depositIncome + depositIncome * depositMonthPercent; // почему тут не хочет работать toFixed(3) - хз
}

console.log (''); // пустая строка
console.log ('Итого доход: ' + (depositIncome - depositAmount).toFixed(3));

// Программа для рисования доски для игры "Крестики-нолики"
// Очень костыльный вариант, но пока знаний не хватает более гибко рисовать
// Хотелось бы добавить нормальные отступы и возможность генерировать большего размера поля (но чем больше полей, тем больше будут видны меньшего размера квадраты по краям), поэтому пока оставлю так

drawRect(0, 0, 150, 150, 'black'); // рисуем черный квадрат с запасом и отступами

for (let y = 0; y < 3; y = y + 1) {
	for (let x = 0; x < 3; x = x + 1) {
		drawRect(x * 51, y * 51, 48, 48, 'white'); // рисуем белые квадраты поверх
	}
}




