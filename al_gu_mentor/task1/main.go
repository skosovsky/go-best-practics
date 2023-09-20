package main

import (
	"errors"
	"fmt"
	"strconv"
	"task1/entrance"
)

func main() {
	var resultMulInt, resultDivInt, resultDifInt, resultSumInt int
	var resultMulUint, resultDivUint, resultDifUint, resultSumUint uint64
	var resultMulFloat, resultDivFloat, resultDifFloat, resultSumFloat float64
	var resultMulString, resultDivString, resultDifString, resultSumString string
	var err error

	valueOne, valueTwo, valueThree := entrance.GetValue()

	fmt.Println("result fot Int")
	// Int
	resultMulInt, _ = mathInt(valueOne, valueTwo, valueThree, "*")
	fmt.Println(resultMulInt)

	resultDivInt, err = mathInt(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivInt)
	}

	resultDifInt, _ = mathInt(valueOne, valueTwo, valueThree, "-")
	fmt.Println(resultDifInt)

	resultSumInt, _ = mathInt(valueOne, valueTwo, valueThree, "+")
	fmt.Println(resultSumInt)

	fmt.Println("result fot Uint")
	// Uint
	resultMulUint, err = mathUint(valueOne, valueTwo, valueThree, "*")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultMulUint)
	}

	resultDivUint, err = mathUint(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivUint)
	}

	resultDifUint, err = mathUint(valueOne, valueTwo, valueThree, "-")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDifUint)
	}

	resultSumUint, err = mathUint(valueOne, valueTwo, valueThree, "+")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultSumUint)
	}

	fmt.Println("result fot Float")
	// Float (если форматирование не требуется, нужно удалить "%.3f", если positiveOnly = true - только положительные числа)
	resultMulFloat, _ = mathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "*", false)
	fmt.Printf("%.2f", resultMulFloat)
	fmt.Println("")

	resultDivFloat, err = mathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "/", false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f", resultDivFloat)
		fmt.Println("")
	}

	resultDifFloat, _ = mathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "-", false)
	fmt.Printf("%.2f", resultDifFloat)
	fmt.Println("")

	resultSumFloat, _ = mathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "+", false)
	fmt.Printf("%.2f", resultSumFloat)
	fmt.Println("")

	fmt.Println("result fot String")
	// String (для перевода цифр в слова, используем numToWord, для обрезания cutFirstSymbol)
	resultMulString, _ = mathString(valueOne, valueTwo, valueThree, "*", true, true)
	fmt.Println(resultMulString)

	resultDivString, err = mathString(valueOne, valueTwo, valueThree, "/", true, false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivString)
	}

	resultDifString, _ = mathString(valueOne, valueTwo, valueThree, "-", true, false)
	fmt.Println(resultDifString)

	resultSumString, _ = mathString(valueOne, valueTwo, valueThree, "+", true, false)
	fmt.Println(resultSumString)

}

func mathInt(valueOne int, valueTwo int, valueThree int, mathOperation string) (result int, err error) {
	switch mathOperation {
	case "*":
		result = valueOne * valueTwo * valueThree
	case "/":
		if valueOne == 0 || valueTwo == 0 || valueThree == 0 {
			return 0, errors.New("division operation failed due to division by 0")
		}
		result = valueOne / valueTwo / valueThree
	case "-":
		result = valueOne - valueTwo - valueThree
	case "+":
		result = valueOne + valueTwo + valueThree
	}

	return result, err

}

func mathUint(valueOne int, valueTwo int, valueThree int, mathOperation string) (result uint64, err error) {
	switch mathOperation {
	case "*":
		result, err = strconv.ParseUint(fmt.Sprint(valueOne*valueTwo*valueThree), 10, 64)
		if err != nil {
			return 0, errors.New("Mul operation failed due to negative number")
		}
	case "/":
		if valueOne == 0 || valueTwo == 0 || valueThree == 0 {
			return 0, errors.New("division operation failed due to division by 0")
		}
		result, err = strconv.ParseUint(fmt.Sprint(valueOne/valueTwo/valueThree), 10, 64)
		if err != nil {
			return 0, errors.New("division operation failed due to negative number")
		}
	case "-":
		result, err = strconv.ParseUint(fmt.Sprint(valueOne-valueTwo-valueThree), 10, 64)
		if err != nil {
			return 0, errors.New("Diference operation failed due to negative number")
		}
	case "+":
		result, err = strconv.ParseUint(fmt.Sprint(valueOne+valueTwo+valueThree), 10, 64)
		if err != nil {
			return 0, errors.New("Sum operation failed due to negative number")
		}
	}

	return result, err

}

func mathFloat(valueOne float64, valueTwo float64, valueThree float64, mathOperation string, positiveOnly bool) (result float64, err error) {
	switch mathOperation {
	case "*":
		result = valueOne * valueTwo * valueThree
		if positiveOnly && result < 0 {
			result = 0
		}
	case "/":
		if valueOne == 0 || valueTwo == 0 || valueThree == 0 {
			return 0, errors.New("division operation failed due to division by 0")
		}
		result = valueOne / valueTwo / valueThree
		if positiveOnly && result < 0 {
			result = 0
		}
	case "-":
		result = valueOne - valueTwo - valueThree
		if positiveOnly && result < 0 {
			result = 0
		}
	case "+":
		result = valueOne + valueTwo + valueThree
		if positiveOnly && result < 0 {
			result = 0
		}
	}

	return result, err

}

func mathString(valueOne int, valueTwo int, valueThree int, mathOperation string, numToWord bool, cutFirstSymbol bool) (result string, err error) {
	switch mathOperation {
	case "*":
		result = numbersToWords(strconv.Itoa(valueOne*valueTwo*valueThree), numToWord, cutFirstSymbol)
	case "/":
		if valueOne == 0 || valueTwo == 0 || valueThree == 0 {
			return "", errors.New("division operation failed due to division by 0")
		}
		result = numbersToWords(strconv.Itoa(valueOne/valueTwo/valueThree), numToWord, cutFirstSymbol)
	case "-":
		result = numbersToWords(strconv.Itoa(valueOne-valueTwo-valueThree), numToWord, cutFirstSymbol)
	case "+":
		result = numbersToWords(strconv.Itoa(valueOne+valueTwo+valueThree), numToWord, cutFirstSymbol)
	}

	return result, err

}

func numbersToWords(value string, numToWords bool, cutFirstSymbol bool) (result string) {
	if numToWords == false {
		return value
	}

	result = ""
	for i := 0; i < len(value); i++ {
		numberToWord := fmt.Sprint(value[i : i+1])

		switch numberToWord {
		case "0":
			numberToWord = "ноль"
		case "1":
			numberToWord = "один"
		case "2":
			numberToWord = "два"
		case "3":
			numberToWord = "три"
		case "4":
			numberToWord = "четыре"
		case "5":
			numberToWord = "пять"
		case "6":
			numberToWord = "шесть"
		case "7":
			numberToWord = "семь"
		case "8":
			numberToWord = "восемь"
		case "9":
			numberToWord = "девять"
		case "-":
			numberToWord = "минус"
		}

		if cutFirstSymbol == true {
			numberToWord = fmt.Sprint(numberToWord[2:])
		}

		if i == len(value)-1 {
			result = result + numberToWord
		} else {
			result = result + numberToWord + " "
		}
	}

	return result
}

//func cutFloat(floatNumber float64, cutNumber int) (cuttingFloat float64) {
//	if cutNumber == -1 {
//		return floatNumber
//	}
//	//cuttingFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", floatNumber), 64)
//
//	strconv.FormatFloat(floatNumber, 'f', -1, 64)
//
//	floatNumberString := strconv.FormatFloat(floatNumber, 'f', -1, 64)
//	numberToPoint := len(strings.SplitAfterN(floatNumberString, ".", 2)) + cutNumber
//	floatNumberString = floatNumberString[:numberToPoint]
//
//	cuttingFloat, err := strconv.ParseFloat(floatNumberString, 64)
//	if err != nil {
//		errors.New("unknown error from ParseFloat")
//	}
//
//	return cuttingFloat
//}

//- умножить
//- поделить
//- отнять
//- прибавить

//- целого (a) +
//- дробного (b) +
//- целого, но только положительного числа (c) +
//- дробного положительного (d) +
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
