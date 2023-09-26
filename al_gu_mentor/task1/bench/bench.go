package bench

import "fmt"

func Bench1() {
	a := 1
	b := 2
	c := 3
	var result int

	pointerMath1(&result, &a, &b, &c)
	_ = fmt.Sprint(result)
}

func pointerMath1(result *int, values ...*int) {

	for _, n := range values {
		*result = *result + *n
	}
}

func Bench2() {
	a := 1
	b := 2
	c := 3
	var result int

	result = pointerMath2(a, b, c)
	_ = fmt.Sprint(result)
}

func pointerMath2(values ...int) int {
	var result int
	for _, n := range values {
		result = result + n
	}
	return result
}

func Bench3() {
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	var result int
	pointerMath3(&result, &myMap)
	_ = fmt.Sprint(result)
}

func pointerMath3(result *int, myMap *map[string]int) {
	for _, value := range *myMap {
		*result += value
	}
}

func Bench4() {
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	var result int
	result = pointerMath4(myMap)
	_ = fmt.Sprint(result)
}

func pointerMath4(myMap map[string]int) int {
	var result int
	for _, value := range myMap {
		result += value
	}
	return result
}
