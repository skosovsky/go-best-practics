package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"task1/entrance"
)

func main() {
	var resultMultiInt, resultDivInt, resultDiffInt, resultSummInt int
	var resultMultiFloat, resultDivFloat, resultDiffFloat, resultSummFloat float64
	var err error

	valueOne, valueTwo, valueThree := entrance.GetValue()

	// Int
	resultMultiInt, _ = mathInt(valueOne, valueTwo, valueThree, "*")
	fmt.Println(resultMultiInt)

	resultDivInt, err = mathInt(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivInt)
	}

	resultDiffInt, _ = mathInt(valueOne, valueTwo, valueThree, "-")
	fmt.Println(resultDiffInt)

	resultSummInt, _ = mathInt(valueOne, valueTwo, valueThree, "+")
	fmt.Println(resultSummInt)

	// Float (в CutNum указываем количество знаков после запятой, -1 если не обрезать)
	resultMultiFloat, _ = mathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "*", -1)
	fmt.Println(resultMultiFloat)

	resultDivFloat, err = mathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "/", -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivFloat)
	}

	resultDiffFloat, _ = mathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "-", -1)
	fmt.Println(resultDiffFloat)

	resultSummFloat, _ = mathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "+", -1)
	fmt.Println(resultSummFloat)

}

func mathInt(valueOne int, valueTwo int, valueThree int, mathOperation string) (resultInt int, err error) {
	switch mathOperation {
	case "*":
		resultInt = valueOne * valueTwo * valueThree
	case "/":
		if valueOne == 0 || valueTwo == 0 || valueThree == 0 {
			return 0, errors.New("division operation failed due to division by 0")
		}
		resultInt = valueOne / valueTwo / valueThree
	case "-":
		resultInt = valueOne - valueTwo - valueThree
	case "+":
		resultInt = valueOne + valueTwo + valueThree
	}

	return resultInt, err

}

func mathFloat(valueOne float64, valueTwo float64, valueThree float64, mathOperation string, cutNum int) (resultInt float64, err error) {
	switch mathOperation {
	case "*":
		resultInt = cutFloat(valueOne*valueTwo*valueThree, cutNum)
	case "/":
		if valueOne == 0 || valueTwo == 0 || valueThree == 0 {
			return 0, errors.New("division operation failed due to division by 0")
		}
		resultInt = cutFloat(valueOne/valueTwo/valueThree, cutNum)
	case "-":
		resultInt = cutFloat(valueOne-valueTwo-valueThree, cutNum)
	case "+":
		resultInt = cutFloat(valueOne+valueTwo+valueThree, cutNum)
	}

	return resultInt, err

}

func cutFloat(floatNumber float64, cutNumber int) (cuttingFloat float64) {
	if cutNumber == -1 {
		return floatNumber
	}

	floatNumberString := fmt.Sprint(floatNumber)
	numberToPoint := len(strings.SplitAfterN(floatNumberString, ".", 2)) + cutNumber
	floatNumberString = floatNumberString[:numberToPoint]

	cuttingFloat, err := strconv.ParseFloat(floatNumberString, 64)
	if err != nil {
		errors.New("unknown error from ParseFloat")
	}

	return cuttingFloat
}

//- умножить
//- поделить
//- отнять
//- прибавить

//- целого (a)
//- дробного (b)
//- целого, но только положительного числа (c)
//- дробного положительного (d)
//- текста
//- в виде цифры
//- текста в виде слова (вместо 3 что б было "три")
//- текста в виде слова, у которого удалишь первый символ
//- байт
//- рун
//- большого цельного
//- сложить их
//- отнять их
//- большого дробного
//- сложить их
//- отнять их
