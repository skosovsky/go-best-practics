package areacalculate

import "fmt"

func InputArea() {
	var lenSide1 int // сторона 1
	var lenSide2 int // сторона 2
	var area int     // площадь

	fmt.Println("Добро пожаловать в программу для вычисления площади прямоугольника")
	fmt.Println("Введите длину стороны 1")
	fmt.Scanln(&lenSide1)
	fmt.Println("Введите длину стороны 2")
	fmt.Scanln(&lenSide2)
	area = AreaCalculate(lenSide1, lenSide2)
	fmt.Println("----------")
	fmt.Println("Площадь прямоугольника равна: ", area)
}

func AreaCalculate(lenSide1 int, lenSide2 int) int {
	return lenSide1 * lenSide2
}
