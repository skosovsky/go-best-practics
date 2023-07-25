package main

import "fmt"

func main() {
	var userNum int

	fmt.Println("Добро пожаловать в приложение расчета числа Фиббоначи")
	fmt.Print("Введите число: ")
	fmt.Scan(&userNum)
	fmt.Printf("Число фиббоначи для %v: %v", userNum, fibbonachi(userNum))
}

var mapFibbonachi = map[int]int{
	0: 0,
	1: 1,
}

func fibbonachi(num int) (result int) {
	_, numExist := mapFibbonachi[num]
	if numExist == true {
		result = mapFibbonachi[num]
	} else {
		result = fibbonachi(num-1) + fibbonachi(num-2)
		mapFibbonachi[num] = result
	}
	return result
}
