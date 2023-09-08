// Package fibbonachi calculate fibbonachi numbers
package fibbonachi

import (
	"errors"
	"fmt"
)

func InputFibbonachi() {
	var userNum int

	fmt.Println("Добро пожаловать в приложение расчета числа Фиббоначи")
	fmt.Print("Введите число: ")
	fmt.Scan(&userNum)

	result, err := Fibbonachi(userNum, true)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Число фиббоначи для %v: %v", userNum, result)
	}
}

var ErrNegativeNumbers = errors.New("cant't use negative number")

// Fibbonachi calculate fibbonachi with given number
func Fibbonachi(n int, useOptimized bool) (int, error) {
	if n < 0 {
		return 0, ErrNegativeNumbers
	}

	if useOptimized == true {
		return CalcFibbonachi(n), nil
	} else {
		return CalcFibbonachiSimple(n), nil
	}
}

var mapFibbonachi = map[int]int{
	0: 0,
	1: 1,
}

func CalcFibbonachi(num int) (result int) {

	_, numExist := mapFibbonachi[num]
	if numExist == true {
		result = mapFibbonachi[num]
	} else {
		result = CalcFibbonachi(num-1) + CalcFibbonachi(num-2)
		mapFibbonachi[num] = result
	}
	return result
}

func CalcFibbonachiSimple(num int) (result int) {
	if num == 0 {
		result = 0
	} else if num == 1 {
		result = 1
	} else {
		result = CalcFibbonachiSimple(num-1) + CalcFibbonachiSimple(num-2)
	}
	return result
}
