package main

import "fmt"

func fibbonachi(num int) (result int) {
	if num == 0 {
		result = 0
	} else if num == 1 {
		result = 1
	} else {
		result = fibbonachi(num-1) + fibbonachi(num-2)
	}
	return result
}

func main() {
	var userNum int

	fmt.Println("Добро пожаловать в приложение расчета числа Фиббоначи")
	fmt.Print("Введите число: ")
	fmt.Scan(&userNum)
	fmt.Printf("Число фиббоначи для %v: %v", userNum, fibbonachi(userNum))
}
