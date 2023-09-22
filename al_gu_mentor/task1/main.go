package main

import (
	"fmt"
	"math/big"
	"strconv"
	"task1/entrance"
	"task1/operation"
)

func main() {
	var (
		resultMulInt, resultDivInt, resultSubInt, resultSumInt             int
		resultMulUint, resultDivUint, resultSubUint, resultSumUint         uint64
		resultMulFloat, resultDivFloat, resultSubFloat, resultSumFloat     float64
		resultMulString, resultDivString, resultSubString, resultSumString string
		resultMulByte, resultDivByte, resultSubByte, resultSumByte         []byte
		resultMulRune, resultDivRune, resultSubRune, resultSumRune         []rune
		err                                                                error
		resultIntStruct                                                    struct {
			resultMul int
			resultDiv int
			resultSub int
			resultSum int
		}
		resultUintStruct struct {
			resultMul uint64
			resultDiv uint64
			resultSub uint64
			resultSum uint64
		}
		resultFloatStruct struct {
			resultMul float64
			resultDiv float64
			resultSub float64
			resultSum float64
		}
		resultStringStruct struct {
			resultMul string
			resultDiv string
			resultSub string
			resultSum string
		}
		resultByteStruct struct {
			resultMul []byte
			resultDiv []byte
			resultSub []byte
			resultSum []byte
		}
		resultRuneStruct struct {
			resultMul []rune
			resultDiv []rune
			resultSub []rune
			resultSum []rune
		}
		resultBigIntStruct struct {
			resultMul big.Int
			resultDiv big.Int
			resultSub big.Int
			resultSum big.Int
		}
		resultBigFloatStruct struct {
			resultMul big.Float
			resultDiv big.Float
			resultSub big.Float
			resultSum big.Float
		}
		resultBigRatStruct struct {
			resultMul big.Rat
			resultDiv big.Rat
			resultSub big.Rat
			resultSum big.Rat
		}
	)

	valueOne, valueTwo, valueThree := entrance.GetValue()

	// Int
	resultMulInt, _ = operation.MathInt(valueOne, valueTwo, valueThree, "*")
	resultDivInt, err = operation.MathInt(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	}
	resultSubInt, _ = operation.MathInt(valueOne, valueTwo, valueThree, "-")
	resultSumInt, _ = operation.MathInt(valueOne, valueTwo, valueThree, "+")

	fmt.Println("result fot Int:")
	resultIntMap := map[string]int{
		"resultMulInt": resultMulInt,
		"resultDivInt": resultDivInt,
		"resultSubInt": resultSubInt,
		"resultSumInt": resultSumInt,
	}
	fmt.Println(resultIntMap)

	resultIntSlice := []int{resultMulInt, resultDivInt, resultSubInt, resultSumInt}
	fmt.Println(resultIntSlice)

	resultIntStruct.resultMul = resultMulInt
	resultIntStruct.resultDiv = resultDivInt
	resultIntStruct.resultSub = resultSubInt
	resultIntStruct.resultSum = resultSumInt
	fmt.Println(resultIntStruct)

	// Uint
	resultMulUint, err = operation.MathUint(valueOne, valueTwo, valueThree, "*")
	if err != nil {
		fmt.Println(err)
	}
	resultDivUint, err = operation.MathUint(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	}
	resultSubUint, err = operation.MathUint(valueOne, valueTwo, valueThree, "-")
	if err != nil {
		fmt.Println(err)
	}
	resultSumUint, err = operation.MathUint(valueOne, valueTwo, valueThree, "+")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("result fot Uint:")
	resultUintMap := map[string]uint64{
		"resultMulUint": resultMulUint,
		"resultDivUint": resultDivUint,
		"resultSubUint": resultSubUint,
		"resultSumUint": resultSumUint,
	}
	fmt.Println(resultUintMap)

	resultUintSlice := []uint64{resultMulUint, resultDivUint, resultSubUint, resultSumUint}
	fmt.Println(resultUintSlice)

	resultUintStruct.resultMul = resultMulUint
	resultUintStruct.resultDiv = resultDivUint
	resultUintStruct.resultSub = resultSubUint
	resultUintStruct.resultSum = resultSumUint
	fmt.Println(resultUintStruct)

	// Float (если форматирование не требуется, нужно удалить "%.3f", если positiveOnly = true - только положительные числа)
	resultMulFloat, _ = operation.MathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "*", false)
	resultMulFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", resultMulFloat), 64)

	resultDivFloat, err = operation.MathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "/", false)
	if err != nil {
		fmt.Println(err)
	}
	resultDivFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", resultDivFloat), 64)

	resultSubFloat, _ = operation.MathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "-", false)
	resultSubFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", resultSubFloat), 64)

	resultSumFloat, _ = operation.MathFloat(float64(valueOne), float64(valueTwo), float64(valueThree), "+", false)
	resultSumFloat, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", resultSumFloat), 64)

	fmt.Println("result fot Float:")
	resultFloatMap := map[string]float64{
		"resultMulFloat": resultMulFloat,
		"resultDivFloat": resultDivFloat,
		"resultSubFloat": resultSubFloat,
		"resultSumFloat": resultSumFloat,
	}
	fmt.Println(resultFloatMap)

	resultFloatSlice := []float64{resultMulFloat, resultDivFloat, resultSubFloat, resultSumFloat}
	fmt.Println(resultFloatSlice)

	resultFloatStruct.resultMul = resultMulFloat
	resultFloatStruct.resultDiv = resultDivFloat
	resultFloatStruct.resultSub = resultSubFloat
	resultFloatStruct.resultSum = resultSumFloat
	fmt.Println(resultFloatStruct)

	// String (для перевода цифр в слова, используем numToWord, для обрезания cutFirstSymbol)
	resultMulString, _ = operation.MathString(valueOne, valueTwo, valueThree, "*", false, false)
	resultDivString, err = operation.MathString(valueOne, valueTwo, valueThree, "/", true, false)
	if err != nil {
		fmt.Println(err)
	}

	resultSubString, _ = operation.MathString(valueOne, valueTwo, valueThree, "-", true, true)
	resultSumString, _ = operation.MathString(valueOne, valueTwo, valueThree, "+", true, false)

	fmt.Println("result fot String:")
	resultStringMap := map[string]string{
		"resultMulString": resultMulString,
		"resultDivString": resultDivString,
		"resultSubString": resultSubString,
		"resultSumString": resultSumString,
	}
	fmt.Println(resultStringMap)

	resultStringSlice := []string{resultMulString, resultDivString, resultSubString, resultSumString}
	fmt.Println(resultStringSlice)

	resultStringStruct.resultMul = resultMulString
	resultStringStruct.resultDiv = resultDivString
	resultStringStruct.resultSub = resultSubString
	resultStringStruct.resultSum = resultSumString
	fmt.Println(resultStringStruct)

	// Byte
	resultMulByte, _ = operation.MathByte(valueOne, valueTwo, valueThree, "*")
	resultDivByte, err = operation.MathByte(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	}

	resultSubByte, _ = operation.MathByte(valueOne, valueTwo, valueThree, "-")
	resultSumByte, _ = operation.MathByte(valueOne, valueTwo, valueThree, "+")

	fmt.Println("result fot Byte:")
	resultByteMap := map[string][]byte{
		"resultMulByte": resultMulByte,
		"resultDivByte": resultDivByte,
		"resultSubByte": resultSubByte,
		"resultSumByte": resultSumByte,
	}
	fmt.Println(resultByteMap)

	resultByteSlice := [][]byte{resultMulByte, resultDivByte, resultSubByte, resultSumByte}
	fmt.Println(resultByteSlice)

	resultByteStruct.resultMul = resultMulByte
	resultByteStruct.resultDiv = resultDivByte
	resultByteStruct.resultSub = resultSubByte
	resultByteStruct.resultSum = resultSumByte
	fmt.Println(resultByteStruct)

	// Rune
	resultMulRune, _ = operation.MathRune(valueOne, valueTwo, valueThree, "*")

	resultDivRune, err = operation.MathRune(valueOne, valueTwo, valueThree, "/")
	if err != nil {
		fmt.Println(err)
	}

	resultSubRune, _ = operation.MathRune(valueOne, valueTwo, valueThree, "-")
	resultSumRune, _ = operation.MathRune(valueOne, valueTwo, valueThree, "+")

	fmt.Println("result fot Rune:")
	resultRuneMap := map[string][]rune{
		"resultMulRune": resultMulRune,
		"resultDivRune": resultDivRune,
		"resultSubRune": resultSubRune,
		"resultSumRune": resultSumRune,
	}
	fmt.Println(resultRuneMap)

	resultRuneSlice := [][]rune{resultMulRune, resultDivRune, resultSubRune, resultSumRune}
	fmt.Println(resultRuneSlice)

	resultRuneStruct.resultMul = resultMulRune
	resultRuneStruct.resultDiv = resultDivRune
	resultRuneStruct.resultSub = resultSubRune
	resultRuneStruct.resultSum = resultSumRune
	fmt.Println(resultRuneStruct)

	// BigInt
	valueOneBigInt, _ := new(big.Int).SetString("25000000000000000000", 10)
	valueTwoBigInt, _ := new(big.Int).SetString("2400000000000000000", 10)
	valueThreeBigInt, _ := new(big.Int).SetString("230000000000000000", 10)

	resultMulBigInt, _ := operation.MathBigInt(*valueOneBigInt, *valueTwoBigInt, *valueThreeBigInt, "*")
	resultDivBigInt, err := operation.MathBigInt(*valueOneBigInt, *valueTwoBigInt, *valueThreeBigInt, "/")
	if err != nil {
		fmt.Println(err)
	}

	resultSubBigInt, _ := operation.MathBigInt(*valueOneBigInt, *valueTwoBigInt, *valueThreeBigInt, "-")
	resultSumBigInt, _ := operation.MathBigInt(*valueOneBigInt, *valueTwoBigInt, *valueThreeBigInt, "+")

	fmt.Println("result fot BigInt:")
	resultBigIntMap := map[string]big.Int{
		"resultMulBigInt": resultMulBigInt,
		"resultDivBigInt": resultDivBigInt,
		"resultSubBigInt": resultSubBigInt,
		"resultSumBigInt": resultSumBigInt,
	}
	fmt.Println(resultBigIntMap)

	resultBigIntSlice := []big.Int{resultMulBigInt, resultDivBigInt, resultSubBigInt, resultSumBigInt}
	fmt.Println(resultBigIntSlice)

	resultBigIntStruct.resultMul = resultMulBigInt
	resultBigIntStruct.resultDiv = resultDivBigInt
	resultBigIntStruct.resultSub = resultSubBigInt
	resultBigIntStruct.resultSum = resultSumBigInt
	fmt.Println(resultBigIntStruct)

	// BigFloat
	valueOneBigFloat, _ := new(big.Float).SetString("25000000000000000000.555")
	valueTwoBigFloat, _ := new(big.Float).SetString("2400000000000000000.555")
	valueThreeBigFloat, _ := new(big.Float).SetString("230000000000000000.555")

	resultMulBigFloat, _ := operation.MathBigFloat(*valueOneBigFloat, *valueTwoBigFloat, *valueThreeBigFloat, "*")
	resultDivBigFloat, err := operation.MathBigFloat(*valueOneBigFloat, *valueTwoBigFloat, *valueThreeBigFloat, "/")
	if err != nil {
		fmt.Println(err)
	}

	resultSubBigFloat, _ := operation.MathBigFloat(*valueOneBigFloat, *valueTwoBigFloat, *valueThreeBigFloat, "-")
	resultSumBigFloat, _ := operation.MathBigFloat(*valueOneBigFloat, *valueTwoBigFloat, *valueThreeBigFloat, "+")

	fmt.Println("result fot BigFloat:")
	resultBigFloatMap := map[string]big.Float{
		"resultMulBigFloat": resultMulBigFloat,
		"resultDivBigFloat": resultDivBigFloat,
		"resultSubBigFloat": resultSubBigFloat,
		"resultSumBigFloat": resultSumBigFloat,
	}
	fmt.Println(resultBigFloatMap)

	resultBigFloatSlice := []big.Float{resultMulBigFloat, resultDivBigFloat, resultSubBigFloat, resultSumBigFloat}
	fmt.Println(resultBigFloatSlice)

	resultBigFloatStruct.resultMul = resultMulBigFloat
	resultBigFloatStruct.resultDiv = resultDivBigFloat
	resultBigFloatStruct.resultSub = resultSubBigFloat
	resultBigFloatStruct.resultSum = resultSumBigFloat
	fmt.Println(resultBigFloatStruct)

	// BigRat
	valueOneBigRat := big.NewRat(1, 3)
	valueTwoBigRat := big.NewRat(2, 3)
	valueThreeBigRat := big.NewRat(3, 4)

	resultMulBigRat, _ := operation.MathBigRat(*valueOneBigRat, *valueTwoBigRat, *valueThreeBigRat, "*")
	resultDivBigRat, err := operation.MathBigRat(*valueOneBigRat, *valueTwoBigRat, *valueThreeBigRat, "/")
	if err != nil {
		fmt.Println(err)
	}

	resultSubBigRat, _ := operation.MathBigRat(*valueOneBigRat, *valueTwoBigRat, *valueThreeBigRat, "-")
	resultSumBigRat, _ := operation.MathBigRat(*valueOneBigRat, *valueTwoBigRat, *valueThreeBigRat, "+")

	fmt.Println("result fot BigRat:")
	resultBigRatMap := map[string]big.Rat{
		"resultMulBigRat": resultMulBigRat,
		"resultDivBigRat": resultDivBigRat,
		"resultSubBigRat": resultSubBigRat,
		"resultSumBigRat": resultSumBigRat,
	}
	fmt.Println(resultBigRatMap)

	resultBigRatSlice := []big.Rat{resultMulBigRat, resultDivBigRat, resultSubBigRat, resultSumBigRat}
	fmt.Println(resultBigRatSlice)

	resultBigRatStruct.resultMul = resultMulBigRat
	resultBigRatStruct.resultDiv = resultDivBigRat
	resultBigRatStruct.resultSub = resultSubBigRat
	resultBigRatStruct.resultSum = resultSumBigRat
	fmt.Println(resultBigRatStruct)

}
