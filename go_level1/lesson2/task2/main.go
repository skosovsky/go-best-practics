package main

import (
	"fmt"
	"math"
)

func main() {
	var area float64     // площадь
	var diameter float64 // диаметр
	var length float64   // длина

	fmt.Println("Добро пожаловать в программу расчета диаметра и длины окружности")
	fmt.Println("Введите площадь круга")
	fmt.Scanln(&area)
	diameter = 2 * math.Sqrt(area/math.Pi)
	length = math.Pi * diameter
	fmt.Println("Диаметр круга равен:", diameter, ", а длина:", length)
}
