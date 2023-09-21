package operation

import (
	"errors"
	"fmt"
	"strconv"
)

func MathInt(valueOne int, valueTwo int, valueThree int, mathOperation string) (result int, err error) {
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

func MathUint(valueOne int, valueTwo int, valueThree int, mathOperation string) (result uint64, err error) {
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

func MathFloat(valueOne float64, valueTwo float64, valueThree float64, mathOperation string, positiveOnly bool) (result float64, err error) {
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

func MathString(valueOne int, valueTwo int, valueThree int, mathOperation string, numToWord bool, cutFirstSymbol bool) (result string, err error) {
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

func MathByte(valueOne int, valueTwo int, valueThree int, mathOperation string) (result []byte, err error) {
	switch mathOperation {
	case "*":
		result = []byte(strconv.Itoa(valueOne * valueTwo * valueThree))
	case "/":
		if valueOne == 0 || valueTwo == 0 || valueThree == 0 {
			return nil, errors.New("division operation failed due to division by 0")
		}
		result = []byte(strconv.Itoa(valueOne / valueTwo / valueThree))
	case "-":
		result = []byte(strconv.Itoa(valueOne - valueTwo - valueThree))
	case "+":
		result = []byte(strconv.Itoa(valueOne + valueTwo + valueThree))
	}
	return result, err

}

func MathRune(valueOne int, valueTwo int, valueThree int, mathOperation string) (result []rune, err error) {
	switch mathOperation {
	case "*":
		result = []rune(strconv.Itoa(valueOne * valueTwo * valueThree))
	case "/":
		if valueOne == 0 || valueTwo == 0 || valueThree == 0 {
			return nil, errors.New("division operation failed due to division by 0")
		}
		result = []rune(strconv.Itoa(valueOne / valueTwo / valueThree))
	case "-":
		result = []rune(strconv.Itoa(valueOne - valueTwo - valueThree))
	case "+":
		result = []rune(strconv.Itoa(valueOne + valueTwo + valueThree))
	}
	return result, err

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
