package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a int64
	var i int64
	var result []int64

	fmt.Println("Добро пожаловать в программу поиска простых чисел от 0 до N")
	fmt.Printf("Введите число N: ")
	fmt.Scan(&a)

	for i = 0; i <= a; i++ {
		if big.NewInt(i).ProbablyPrime(0) == true {
			result = append(result, i)
		}
	}
	fmt.Printf("Список простых чисел от 0 до %v: ", a)
	fmt.Println(result)
}
