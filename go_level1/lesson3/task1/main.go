package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var result float64
	var statusResult bool
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, min, max): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result = a + b
		statusResult = true
	case "-":
		result = a - b
		statusResult = true
	case "*":
		result = a * b
		statusResult = true
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя")
		} else {
			result = a / b
			statusResult = true
		}
	case "^":
		result = math.Pow(a, b)
		statusResult = true
	case "min":
		result = math.Min(a, b)
		statusResult = true
	case "max":
		result = math.Max(a, b)
		statusResult = true
	default:
		fmt.Println("Операция выбрана неверно")
	}

	if statusResult == true {
		fmt.Printf("Результат выполнения операции %s: %.2f\n", operator, result)
	}
}
