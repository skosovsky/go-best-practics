// Для http://jsbasics.geekbrains.ru

let userName = prompt('Введите ваше имя');

console.log('Привет, ' + userName + '!');


let userName = 'Василий';
let userAge = 23;
let isStudent = true;


console.log(userName);
console.log(userAge);
console.log(isStudent);

let a = 3;
let b = 2;


console.log(a + b); // 5
console.log(a - b); // 1
console.log(a / b); // 1.5
console.log(a * b); // 6


let a = prompt('Введите число a');
let b = prompt('Введите число b');

console.log('Результат: ' + a * b);


// Домашнее задание 1
// Вывод возвраcта в консоль
let userName = prompt('Введите ваше имя'); // запрашиваем и сохраняем имя
let userYear = parseInt(prompt('Введите ваш год рождения')); // запрашиваем и сохраняем год
let currentYear = parseInt('2023'); // текущий год
let currentAge = currentYear - userYear + 1; // рассчитываем возвраст

console.log(userName + ', уже в следующем году вам будет ' + currentAge + ' :-('); // выводим в консоль сообщение о возрасте

// Рисуем смайл в консоль
drawCircle(200, 200, 200, 'black');
drawCircle(120, 120, 50, 'white');
drawCircle(280, 120, 50, 'white');
drawCircle(120, 120, 20, 'black');
drawCircle(280, 120, 20, 'black');
drawRect(120, 250, 180, 40, 'white');
