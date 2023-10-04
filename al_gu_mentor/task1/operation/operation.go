package operation

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

func MathInt(valueOne int, valueTwo int, valueThree int, mathOperation string) (result int, err error) {
	switch mathOperation {
	case "*":
		result = valueOne * valueTwo * valueThree
	case "/":
		if valueTwo == 0 || valueThree == 0 {
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
			return 0, errors.New("mul operation failed due to negative number")
		}
	case "/":
		if valueTwo == 0 || valueThree == 0 {
			return 0, errors.New("division operation failed due to division by 0")
		}
		result, err = strconv.ParseUint(fmt.Sprint(valueOne/valueTwo/valueThree), 10, 64)
		if err != nil {
			return 0, errors.New("division operation failed due to negative number")
		}
	case "-":
		result, err = strconv.ParseUint(fmt.Sprint(valueOne-valueTwo-valueThree), 10, 64)
		if err != nil {
			return 0, errors.New("subtract operation failed due to negative number")
		}
	case "+":
		result, err = strconv.ParseUint(fmt.Sprint(valueOne+valueTwo+valueThree), 10, 64)
		if err != nil {
			return 0, errors.New("sum operation failed due to negative number")
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
		if valueTwo == 0 || valueThree == 0 {
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
		if valueTwo == 0 || valueThree == 0 {
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
		if valueTwo == 0 || valueThree == 0 {
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
		if valueTwo == 0 || valueThree == 0 {
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

func MathBigInt(valueOneBigInt *big.Int, valueTwoBigInt *big.Int, valueThreeBigInt *big.Int, mathOperation string) (result *big.Int, err error) {
	switch mathOperation {
	case "*":
		result = valueOneBigInt.Mul(valueOneBigInt, valueTwoBigInt.Mul(valueTwoBigInt, valueThreeBigInt))
	case "/":
		if valueTwoBigInt.Cmp(big.NewInt(0)) == 0 || valueThreeBigInt.Cmp(big.NewInt(0)) == 0 {
			return big.NewInt(0), errors.New("division operation failed due to division by 0")
		}
		result = valueOneBigInt.Div(valueOneBigInt, valueTwoBigInt.Div(valueTwoBigInt, valueThreeBigInt))
	case "-":
		result = valueOneBigInt.Sub(valueOneBigInt, valueTwoBigInt.Sub(valueTwoBigInt, valueThreeBigInt))
	case "+":
		result = valueOneBigInt.Add(valueOneBigInt, valueTwoBigInt.Add(valueTwoBigInt, valueThreeBigInt))
	}

	return result, err

}

func MathBigFloat(valueOneBigFloat *big.Float, valueTwoBigFloat *big.Float, valueThreeBigFloat *big.Float, mathOperation string) (result *big.Float, err error) {
	switch mathOperation {
	case "*":
		result = valueOneBigFloat.Mul(valueOneBigFloat, valueTwoBigFloat.Mul(valueTwoBigFloat, valueThreeBigFloat))
	case "/":
		if valueTwoBigFloat.Cmp(big.NewFloat(0)) == 0 || valueThreeBigFloat.Cmp(big.NewFloat(0)) == 0 {
			return big.NewFloat(0), errors.New("division operation failed due to division by 0")
		}
		result = valueOneBigFloat.Quo(valueOneBigFloat, valueTwoBigFloat.Quo(valueTwoBigFloat, valueThreeBigFloat))
	case "-":
		result = valueOneBigFloat.Sub(valueOneBigFloat, valueTwoBigFloat.Sub(valueTwoBigFloat, valueThreeBigFloat))
	case "+":
		result = valueOneBigFloat.Add(valueOneBigFloat, valueTwoBigFloat.Add(valueTwoBigFloat, valueThreeBigFloat))
	}

	return result, err

}

func MathBigRat(valueOneBigRat *big.Rat, valueTwoBigRat *big.Rat, valueThreeBigRat *big.Rat, mathOperation string) (result *big.Rat, err error) {
	switch mathOperation {
	case "*":
		result = valueOneBigRat.Mul(valueOneBigRat, valueTwoBigRat.Mul(valueTwoBigRat, valueThreeBigRat))
	case "/":
		if valueTwoBigRat.Cmp(big.NewRat(0, 1)) == 0 || valueThreeBigRat.Cmp(big.NewRat(0, 1)) == 0 {
			return big.NewRat(0, 0), errors.New("division operation failed due to division by 0")
		}
		result = valueOneBigRat.Quo(valueOneBigRat, valueTwoBigRat.Quo(valueTwoBigRat, valueThreeBigRat))
	case "-":
		result = valueOneBigRat.Sub(valueOneBigRat, valueTwoBigRat.Sub(valueTwoBigRat, valueThreeBigRat))
	case "+":
		result = valueOneBigRat.Add(valueOneBigRat, valueTwoBigRat.Add(valueTwoBigRat, valueThreeBigRat))
	}

	return result, err

}

func MathIntManyValue(values ...int) (result int) {
	for _, n := range values {
		result = result + n
	}
	return result
}

func MathIntManyValuePointer(result *int, values ...*int) {
	for _, n := range values {
		*result = *result + *n
	}
}

// Функция не меняет значение переменной, т.к. работает с новыми переменными в своей области видимости
func AddOneNoPointer(values ...int) {
	for _, n := range values {
		n = n + 1
	}
}

// Функция изменяет значения переменной, т.к. обращается к самим переменным по поинтерам
func AddOnePointer(values ...*int) {
	for _, n := range values {
		*n = *n + 1
	}
}

func DivByZero() {
	var a int
	defer func() {
		if err := recover(); err != nil {
			err = errors.New("divide by zero")
			log.Println("ooops:", err)
		}
	}()

	a = a / a
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
