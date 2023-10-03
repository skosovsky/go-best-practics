package main

import (
	"fmt"
	"math/big"
	"strconv"
	"task1/bench"
	"task1/entrance"
	"task1/method"
	"task1/operation"
	"task1/output"
	"unsafe"
)

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
	human = method.Human{
		Title: "человек",
		Leg:   2,
	}
	dog = method.Dog{
		Title: "собака",
		Paw:   4,
		Tail:  1,
	}
	snail = method.Snail{
		Title: "улитка",
		Leg:   1,
	}
	StructMammals = method.Mammals{
		Human: method.Human{
			Title: human.Title,
			Leg:   human.Leg,
		},
		Dog: method.Dog{
			Title: dog.Title,
			Paw:   dog.Paw,
			Tail:  dog.Tail,
		},
	}
)

func main() {
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
	fmt.Printf("Var: %T, Size: %v\n", resultIntMap, unsafe.Sizeof(resultIntMap))

	resultIntSlice := []int{resultMulInt, resultDivInt, resultSubInt, resultSumInt}
	fmt.Printf("Var: %T, Size: %v\n", resultIntSlice, unsafe.Sizeof(resultIntSlice))

	resultIntStruct.resultMul = resultMulInt
	resultIntStruct.resultDiv = resultDivInt
	resultIntStruct.resultSub = resultSubInt
	resultIntStruct.resultSum = resultSumInt
	fmt.Printf("Var: %T, Size: %v\n\n", resultIntStruct, unsafe.Sizeof(resultIntStruct))

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
	fmt.Printf("Var: %T, Size: %v\n", resultUintMap, unsafe.Sizeof(resultUintMap))

	resultUintSlice := []uint64{resultMulUint, resultDivUint, resultSubUint, resultSumUint}
	fmt.Printf("Var: %T, Size: %v\n", resultUintSlice, unsafe.Sizeof(resultUintSlice))

	resultUintStruct.resultMul = resultMulUint
	resultUintStruct.resultDiv = resultDivUint
	resultUintStruct.resultSub = resultSubUint
	resultUintStruct.resultSum = resultSumUint
	fmt.Printf("Var: %T, Size: %v\n\n", resultUintStruct, unsafe.Sizeof(resultUintStruct))

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
	fmt.Printf("Var: %T, Size: %v\n", resultFloatMap, unsafe.Sizeof(resultFloatMap))

	resultFloatSlice := []float64{resultMulFloat, resultDivFloat, resultSubFloat, resultSumFloat}
	fmt.Printf("Var: %T, Size: %v\n", resultFloatSlice, unsafe.Sizeof(resultFloatSlice))

	resultFloatStruct.resultMul = resultMulFloat
	resultFloatStruct.resultDiv = resultDivFloat
	resultFloatStruct.resultSub = resultSubFloat
	resultFloatStruct.resultSum = resultSumFloat
	fmt.Printf("Var: %T, Size: %v\n\n", resultFloatStruct, unsafe.Sizeof(resultFloatStruct))

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
	fmt.Printf("Var: %T, Size: %v\n", resultStringMap, unsafe.Sizeof(resultStringMap))

	resultStringSlice := []string{resultMulString, resultDivString, resultSubString, resultSumString}
	fmt.Printf("Var: %T, Size: %v\n", resultStringSlice, unsafe.Sizeof(resultStringSlice))

	resultStringStruct.resultMul = resultMulString
	resultStringStruct.resultDiv = resultDivString
	resultStringStruct.resultSub = resultSubString
	resultStringStruct.resultSum = resultSumString
	fmt.Printf("Var: %T, Size: %v\n\n", resultStringStruct, unsafe.Sizeof(resultStringStruct))

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
	fmt.Printf("Var: %T, Size: %v\n", resultByteMap, unsafe.Sizeof(resultByteMap))

	resultByteSlice := [][]byte{resultMulByte, resultDivByte, resultSubByte, resultSumByte}
	fmt.Printf("Var: %T, Size: %v\n", resultByteSlice, unsafe.Sizeof(resultByteSlice))

	resultByteStruct.resultMul = resultMulByte
	resultByteStruct.resultDiv = resultDivByte
	resultByteStruct.resultSub = resultSubByte
	resultByteStruct.resultSum = resultSumByte
	fmt.Printf("Var: %T, Size: %v\n\n", resultByteStruct, unsafe.Sizeof(resultByteStruct))

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
	fmt.Printf("Var: %T, Size: %v\n", resultRuneMap, unsafe.Sizeof(resultRuneMap))

	resultRuneSlice := [][]rune{resultMulRune, resultDivRune, resultSubRune, resultSumRune}
	fmt.Printf("Var: %T, Size: %v\n", resultRuneSlice, unsafe.Sizeof(resultRuneSlice))

	resultRuneStruct.resultMul = resultMulRune
	resultRuneStruct.resultDiv = resultDivRune
	resultRuneStruct.resultSub = resultSubRune
	resultRuneStruct.resultSum = resultSumRune
	fmt.Printf("Var: %T, Size: %v\n\n", resultRuneStruct, unsafe.Sizeof(resultRuneStruct))

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
	fmt.Printf("Var: %T, Size: %v\n", resultBigIntMap, unsafe.Sizeof(resultBigIntMap))

	resultBigIntSlice := []big.Int{resultMulBigInt, resultDivBigInt, resultSubBigInt, resultSumBigInt}
	fmt.Printf("Var: %T, Size: %v\n", resultBigIntSlice, unsafe.Sizeof(resultBigIntSlice))

	resultBigIntStruct.resultMul = resultMulBigInt
	resultBigIntStruct.resultDiv = resultDivBigInt
	resultBigIntStruct.resultSub = resultSubBigInt
	resultBigIntStruct.resultSum = resultSumBigInt
	fmt.Printf("Var: %T, Size: %v\n\n", resultBigIntStruct, unsafe.Sizeof(resultBigIntStruct))

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
	fmt.Printf("Var: %T, Size: %v\n", resultBigFloatMap, unsafe.Sizeof(resultBigFloatMap))

	resultBigFloatSlice := []big.Float{resultMulBigFloat, resultDivBigFloat, resultSubBigFloat, resultSumBigFloat}
	fmt.Printf("Var: %T, Size: %v\n", resultBigFloatSlice, unsafe.Sizeof(resultBigFloatSlice))

	resultBigFloatStruct.resultMul = resultMulBigFloat
	resultBigFloatStruct.resultDiv = resultDivBigFloat
	resultBigFloatStruct.resultSub = resultSubBigFloat
	resultBigFloatStruct.resultSum = resultSumBigFloat
	fmt.Printf("Var: %T, Size: %v\n\n", resultBigFloatStruct, unsafe.Sizeof(resultBigFloatStruct))

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
	fmt.Printf("Var: %T, Size: %v\n", resultBigRatMap, unsafe.Sizeof(resultBigRatMap))

	resultBigRatSlice := []big.Rat{resultMulBigRat, resultDivBigRat, resultSubBigRat, resultSumBigRat}
	fmt.Printf("Var: %T, Size: %v\n", resultBigRatSlice, unsafe.Sizeof(resultBigRatSlice))

	resultBigRatStruct.resultMul = resultMulBigRat
	resultBigRatStruct.resultDiv = resultDivBigRat
	resultBigRatStruct.resultSub = resultSubBigRat
	resultBigRatStruct.resultSum = resultSumBigRat
	fmt.Printf("Var: %T, Size: %v\n\n", resultBigRatStruct, unsafe.Sizeof(resultBigRatStruct))

	// Вызов функции по указателям и без
	fmt.Println("result fot Sum (many value):")
	resultIntManyValue := operation.MathIntManyValue(resultMulInt, resultDivInt, resultSubInt, resultSumInt, int(resultMulUint), int(resultDivUint), int(resultSubUint), int(resultSumUint), int(resultMulFloat), int(resultDivFloat), int(resultSubFloat), int(resultSumFloat))
	fmt.Println(resultIntManyValue)

	fmt.Println("result fot Sum (many value) - pointers:")

	resultMulUintC := int(resultMulUint)
	resultDivUintC := int(resultDivUint)
	resultSubUintC := int(resultSubUint)
	resultSumUintC := int(resultSumUint)
	resultMulFloatC := int(resultMulFloat)
	resultDivFloatC := int(resultDivFloat)
	resultSubFloatC := int(resultSubFloat)
	resultSumFloatC := int(resultSumFloat)
	var resultIntManyValuePointer int

	operation.MathIntManyValuePointer(&resultIntManyValuePointer, &resultMulInt, &resultDivInt, &resultSubInt, &resultSumInt, &resultMulUintC, &resultDivUintC, &resultSubUintC, &resultSumUintC, &resultMulFloatC, &resultDivFloatC, &resultSubFloatC, &resultSumFloatC)
	fmt.Println(resultIntManyValuePointer)

	operation.AddOneNoPointer(resultMulInt, resultDivInt, resultSubInt, resultSumInt, int(resultMulUint), int(resultDivUint), int(resultSubUint), int(resultSumUint), int(resultMulFloat), int(resultDivFloat), int(resultSubFloat), int(resultSumFloat))
	operation.AddOnePointer(&resultMulInt, &resultDivInt, &resultSubInt, &resultSumInt, &resultMulUintC, &resultDivUintC, &resultSubUintC, &resultSumUintC, &resultMulFloatC, &resultDivFloatC, &resultSubFloatC, &resultSumFloatC)

	fmt.Println()
	fmt.Println("result for Compare:")
	human.CompareWithLegs(resultMulInt, resultDivInt, resultSubInt, resultSumInt, int(resultMulFloat), int(resultDivFloat), int(resultSubFloat), int(resultSumFloat))
	dog.CompareWithsPaws(resultMulInt, resultDivInt, resultSubInt, resultSumInt, int(resultMulFloat), int(resultDivFloat), int(resultSubFloat), int(resultSumFloat))
	snail.CompareWithLegs(resultMulInt, resultDivInt, resultSubInt, resultSumInt, int(resultMulFloat), int(resultDivFloat), int(resultSubFloat), int(resultSumFloat))
	StructMammals.CompareLimb()

	StructMammals.Human.Talk()
	StructMammals.Dog.Talk()

	operation.DivByZero()
	output.SaveToNewFile("result.txt", strconv.Itoa(resultMulInt), strconv.Itoa(resultDivInt), strconv.Itoa(resultSubInt), strconv.Itoa(resultSumInt), strconv.FormatUint(resultMulUint, 10), strconv.FormatUint(resultDivUint, 10), strconv.FormatUint(resultSubUint, 10), strconv.FormatUint(resultSumUint, 10), strconv.FormatFloat(resultMulFloat, 'f', -1, 64), strconv.FormatFloat(resultDivFloat, 'f', -1, 64), strconv.FormatFloat(resultSubFloat, 'f', -1, 64), strconv.FormatFloat(resultSumFloat, 'f', -1, 64), resultMulString, resultDivString, resultSubString, resultSumString, string(resultMulByte), string(resultDivByte), string(resultSubByte), string(resultSumByte), string(resultMulRune), string(resultDivRune), string(resultSubRune), string(resultSumRune))

	bench.Bench1()
	bench.Bench2()
	bench.Bench3()
	bench.Bench4()
}
