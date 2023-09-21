package main

import (
	"fmt"
	"task1/bigs"
	"task1/entrance"
	"task1/operation"
)

func main() {
	var resultMulInt, resultDivInt, resultDifInt, resultSumInt int
	var resultMulUint, resultDivUint, resultDifUint, resultSumUint uint64
	var resultMulFloat, resultDivFloat, resultDifFloat, resultSumFloat float64
	var resultMulString, resultDivString, resultDifString, resultSumString string
	var resultMulByte, resultDivByte, resultDifByte, resultSumByte []byte
	var resultMulRune, resultDivRune, resultDifRune, resultSumRune []rune
	var err error

	valueOne, valueTwo, valueThree := entrance.GetValue()

	fmt.Println("result fot Int")
	// Int
	resultMulInt, _ = operation.MathInt(valueOne, valueTwo, valueThree, "*")
	fmt.Println(resultMulInt)

	resultDivInt, err = operation.MathInt(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivInt)
	}

	resultDifInt, _ = operation.MathInt(valueOne, valueTwo, valueThree, "-")
	fmt.Println(resultDifInt)

	resultSumInt, _ = operation.MathInt(valueOne, valueTwo, valueThree, "+")
	fmt.Println(resultSumInt)

	fmt.Println("result fot Uint")
	// Uint
	resultMulUint, err = operation.MathUint(valueOne, valueTwo, valueThree, "*")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultMulUint)
	}

	resultDivUint, err = operation.MathUint(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivUint)
	}

	resultDifUint, err = operation.MathUint(valueOne, valueTwo, valueThree, "-")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDifUint)
	}

	resultSumUint, err = operation.MathUint(valueOne, valueTwo, valueThree, "+")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultSumUint)
	}

	fmt.Println("result fot Float")
	// Float (если форматирование не требуется, нужно удалить "%.3f", если positiveOnly = true - только положительные числа)
	resultMulFloat, _ = operation.MathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "*", false)
	fmt.Printf("%.2f", resultMulFloat)
	fmt.Println("")

	resultDivFloat, err = operation.MathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "/", false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f", resultDivFloat)
		fmt.Println("")
	}

	resultDifFloat, _ = operation.MathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "-", false)
	fmt.Printf("%.2f", resultDifFloat)
	fmt.Println("")

	resultSumFloat, _ = operation.MathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "+", false)
	fmt.Printf("%.2f", resultSumFloat)
	fmt.Println("")

	fmt.Println("result fot String")
	// String (для перевода цифр в слова, используем numToWord, для обрезания cutFirstSymbol)
	resultMulString, _ = operation.MathString(valueOne, valueTwo, valueThree, "*", true, true)
	fmt.Println(resultMulString)

	resultDivString, err = operation.MathString(valueOne, valueTwo, valueThree, "/", true, false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivString)
	}

	resultDifString, _ = operation.MathString(valueOne, valueTwo, valueThree, "-", true, false)
	fmt.Println(resultDifString)

	resultSumString, _ = operation.MathString(valueOne, valueTwo, valueThree, "+", true, false)
	fmt.Println(resultSumString)

	fmt.Println("result fot Byte")
	// Byte
	resultMulByte, _ = operation.MathByte(valueOne, valueTwo, valueThree, "*")
	fmt.Println(resultMulByte)

	resultDivByte, err = operation.MathByte(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivByte)
	}

	resultDifByte, _ = operation.MathByte(valueOne, valueTwo, valueThree, "-")
	fmt.Println(resultDifByte)

	resultSumByte, _ = operation.MathByte(valueOne, valueTwo, valueThree, "+")
	fmt.Println(resultSumByte)

	fmt.Println("result fot Rune")
	// Byte
	resultMulRune, _ = operation.MathRune(valueOne, valueTwo, valueThree, "*")
	fmt.Println(resultMulRune)

	resultDivRune, err = operation.MathRune(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivRune)
	}

	resultDifRune, _ = operation.MathRune(valueOne, valueTwo, valueThree, "-")
	fmt.Println(resultDifRune)

	resultSumRune, _ = operation.MathRune(valueOne, valueTwo, valueThree, "+")
	fmt.Println(resultSumRune)

	// Int64 & Float64
	bigs.MathBigs()
}
