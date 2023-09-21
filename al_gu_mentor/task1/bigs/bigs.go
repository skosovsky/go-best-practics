package bigs

import (
	"errors"
	"fmt"
)

var resultMulInt64, resultDivInt64, resultDifInt64, resultSumInt64 int64
var resultMulFloat64, resultDivFloat64, resultDifFloat64, resultSumFloat64 float64
var valueOneInt64 int64 = 41.3e12
var valueTwoInt64 int64 = 5
var valueThreeInt64 int64 = 4
var valueOneFloat64 float64 = 41.3e12
var valueTwoFloat64 float64 = 5.555555
var valueThreeFloat64 float64 = 4.5555555555
var err error

func MathBigs() {
	fmt.Println("result fot Int64")
	// Int64
	resultMulInt64, _ = mathInt64(valueOneInt64, valueTwoInt64, valueThreeInt64, "*")
	fmt.Println(resultMulInt64)

	resultDivInt64, err = mathInt64(valueOneInt64, valueTwoInt64, valueThreeInt64, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivInt64)
	}

	resultDifInt64, _ = mathInt64(valueOneInt64, valueTwoInt64, valueThreeInt64, "-")
	fmt.Println(resultDifInt64)

	resultSumInt64, _ = mathInt64(valueOneInt64, valueTwoInt64, valueThreeInt64, "+")
	fmt.Println(resultSumInt64)

	fmt.Println("result fot Float64")
	// Float64
	resultMulFloat64, _ = mathFloat64(valueOneFloat64, valueTwoFloat64, valueThreeFloat64, "*")
	fmt.Println(resultMulFloat64)

	resultDivFloat64, err = mathFloat64(valueOneFloat64, valueTwoFloat64, valueThreeFloat64, "/")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultDivFloat64)
	}

	resultDifFloat64, _ = mathFloat64(valueOneFloat64, valueTwoFloat64, valueThreeFloat64, "-")
	fmt.Println(resultDifFloat64)

	resultSumFloat64, _ = mathFloat64(valueOneFloat64, valueTwoFloat64, valueThreeFloat64, "+")
	fmt.Println(resultSumFloat64)
}

func mathInt64(valueOneInt64 int64, valueTwoInt64 int64, valueThreeInt64 int64, mathOperation string) (result int64, err error) {
	switch mathOperation {
	case "*":
		result = valueOneInt64 * valueTwoInt64 * valueThreeInt64
	case "/":
		if valueOneInt64 == 0 || valueTwoInt64 == 0 || valueThreeInt64 == 0 {
			return 0, errors.New("division operation failed due to division by 0")
		}
		result = valueOneInt64 / valueTwoInt64 / valueThreeInt64
	case "-":
		result = valueOneInt64 - valueTwoInt64 - valueThreeInt64
	case "+":
		result = valueOneInt64 + valueTwoInt64 + valueThreeInt64
	}

	return result, err

}

func mathFloat64(valueOneFloat64 float64, valueTwoFloat64 float64, valueThreeFloat64 float64, mathOperation string) (result float64, err error) {
	switch mathOperation {
	case "*":
		result = valueOneFloat64 * valueTwoFloat64 * valueThreeFloat64
	case "/":
		if valueOneFloat64 == 0 || valueTwoFloat64 == 0 || valueThreeFloat64 == 0 {
			return 0, errors.New("division operation failed due to division by 0")
		}
		result = valueOneFloat64 / valueTwoFloat64 / valueThreeFloat64
	case "-":
		result = valueOneFloat64 - valueTwoFloat64 - valueThreeFloat64
	case "+":
		result = valueOneFloat64 + valueTwoFloat64 + valueThreeFloat64
	}

	return result, err

}
