package main

import "fmt"

func main() {
	var number int
	var number100 int
	var number10 int
	var number1 int

	fmt.Println("Добро пожаловать в программу расчета сотен, десяток и единиц")
	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&number)
	number100 = number / 100
	number10 = (number % 100) / 10
	number1 = ((number % 100) % 10) / 1
	fmt.Println("В числе", number, "-", number100, "сотен", number10, "десяток", number1, "единиц")
}
