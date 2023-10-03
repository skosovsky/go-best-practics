package operation

import (
	"math/big"
	"reflect"
	"testing"
)

func BenchmarkMathIntManyValue(b *testing.B) {
	var a1, a2, a3, a4 int = 10, 100, -200, -10
	var b1, b2, b3, b4 uint64 = 10, 100, 200, 10
	var c1, c2, c3, c4 float64 = 10.15, 100.25, -200.35, -10.45

	for i := 0; i < b.N; i++ {
		MathIntManyValue(a1, a2, a3, a4, int(b1), int(b2), int(b3), int(b4), int(c1), int(c2), int(c3), int(c4))
	}
}

func BenchmarkMathIntManyValuePointer(b *testing.B) {
	var a1, a2, a3, a4 int = 10, 100, -200, -10
	var b1, b2, b3, b4 uint64 = 10, 100, 200, 10
	var c1, c2, c3, c4 float64 = 10.15, 100.25, -200.35, -10.45

	b1c := int(b1)
	b2c := int(b2)
	b3c := int(b3)
	b4c := int(b4)
	c1c := int(c1)
	c2c := int(c2)
	c3c := int(c3)
	c4c := int(c4)

	for i := 0; i < b.N; i++ {
		MathIntManyValuePointer(&a1, &a2, &a3, &a4, &b1c, &b2c, &b3c, &b4c, &c1c, &c2c, &c3c, &c4c)
	}
}

func BenchmarkMathIntManyValueNoInt(b *testing.B) {
	var b1, b2, b3, b4 uint64 = 10, 100, 200, 10
	var c1, c2, c3, c4 float64 = 10.15, 100.25, -200.35, -10.45

	for i := 0; i < b.N; i++ {
		MathIntManyValue(int(b1), int(b2), int(b3), int(b4), int(c1), int(c2), int(c3), int(c4))
	}
}

func BenchmarkMathIntManyValuePointerNoInt(b *testing.B) {
	var b1, b2, b3, b4 uint64 = 10, 100, 200, 10
	var c1, c2, c3, c4 float64 = 10.15, 100.25, -200.35, -10.45

	b1c := int(b1)
	b2c := int(b2)
	b3c := int(b3)
	b4c := int(b4)
	c1c := int(c1)
	c2c := int(c2)
	c3c := int(c3)
	c4c := int(c4)

	for i := 0; i < b.N; i++ {
		MathIntManyValuePointer(&b1c, &b2c, &b3c, &b4c, &c1c, &c2c, &c3c, &c4c)
	}
}

//BenchmarkMathIntManyValue-10                    236451531                4.985 ns/op
//BenchmarkMathIntManyValuePointer-10             147734468                8.141 ns/op
//BenchmarkMathIntManyValueNoInt-10               320508895                3.729 ns/op
//BenchmarkMathIntManyValuePointerNoInt-10        237164148                5.071 ns/op

func TestMathInt(t *testing.T) {
	type args struct {
		valueOne      int
		valueTwo      int
		valueThree    int
		mathOperation string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
		wantErr    bool
	}{
		{
			name: "* 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "*",
			},
			wantResult: 0,
			wantErr:    false,
		}, {
			name: "* 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "*",
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "* -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "*",
			},
			wantResult: -1,
			wantErr:    false,
		},
		{
			name: "/ 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "/",
			},
			wantResult: 0,
			wantErr:    true,
		}, {
			name: "/ 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "/",
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "/ -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "/",
			},
			wantResult: -1,
			wantErr:    false,
		},
		{
			name: "- 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "-",
			},
			wantResult: 0,
			wantErr:    false,
		}, {
			name: "- 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "-",
			},
			wantResult: -1,
			wantErr:    false,
		},
		{
			name: "- -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "-",
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "+ 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "+",
			},
			wantResult: 0,
			wantErr:    false,
		}, {
			name: "+ 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "+",
			},
			wantResult: 3,
			wantErr:    false,
		},
		{
			name: "+ -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "+",
			},
			wantResult: -3,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MathInt(tt.args.valueOne, tt.args.valueTwo, tt.args.valueThree, tt.args.mathOperation)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("MathInt() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMathUint(t *testing.T) {
	type args struct {
		valueOne      int
		valueTwo      int
		valueThree    int
		mathOperation string
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint64
		wantErr    bool
	}{
		{
			name: "* 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "*",
			},
			wantResult: 0,
			wantErr:    false,
		}, {
			name: "* 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "*",
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "* -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "*",
			},
			wantResult: 0,
			wantErr:    true,
		},
		{
			name: "/ 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "/",
			},
			wantResult: 0,
			wantErr:    true,
		}, {
			name: "/ 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "/",
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "/ -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "/",
			},
			wantResult: 0,
			wantErr:    true,
		},
		{
			name: "- 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "-",
			},
			wantResult: 0,
			wantErr:    false,
		}, {
			name: "- 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "-",
			},
			wantResult: 0,
			wantErr:    true,
		},
		{
			name: "- -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "-",
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "+ 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "+",
			},
			wantResult: 0,
			wantErr:    false,
		}, {
			name: "+ 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "+",
			},
			wantResult: 3,
			wantErr:    false,
		},
		{
			name: "+ -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "+",
			},
			wantResult: 0,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MathUint(tt.args.valueOne, tt.args.valueTwo, tt.args.valueThree, tt.args.mathOperation)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("MathUint() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMathFloat(t *testing.T) {
	type args struct {
		valueOne      float64
		valueTwo      float64
		valueThree    float64
		mathOperation string
		positiveOnly  bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
		wantErr    bool
	}{
		{
			name: "* 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "*",
			},
			wantResult: 0,
			wantErr:    false,
		}, {
			name: "* 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "*",
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "* -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "*",
			},
			wantResult: -1,
			wantErr:    false,
		},
		{
			name: "/ 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "/",
			},
			wantResult: 0,
			wantErr:    true,
		}, {
			name: "/ 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "/",
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "/ -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "/",
			},
			wantResult: -1,
			wantErr:    false,
		},
		{
			name: "- 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "-",
			},
			wantResult: 0,
			wantErr:    false,
		}, {
			name: "- 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "-",
			},
			wantResult: -1,
			wantErr:    false,
		},
		{
			name: "- -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "-",
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "+ 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "+",
			},
			wantResult: 0,
			wantErr:    false,
		}, {
			name: "+ 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "+",
			},
			wantResult: 3,
			wantErr:    false,
		},
		{
			name: "+ -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "+",
			},
			wantResult: -3,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MathFloat(tt.args.valueOne, tt.args.valueTwo, tt.args.valueThree, tt.args.mathOperation, tt.args.positiveOnly)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("MathFloat() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMathString(t *testing.T) {
	type args struct {
		valueOne       int
		valueTwo       int
		valueThree     int
		mathOperation  string
		numToWord      bool
		cutFirstSymbol bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{
			name: "* 0",
			args: args{
				valueOne:       0,
				valueTwo:       0,
				valueThree:     0,
				mathOperation:  "*",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "0",
			wantErr:    false,
		},
		{
			name: "* 1",
			args: args{
				valueOne:       1,
				valueTwo:       1,
				valueThree:     1,
				mathOperation:  "*",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "1",
			wantErr:    false,
		},
		{
			name: "* -1",
			args: args{
				valueOne:       -1,
				valueTwo:       -1,
				valueThree:     -1,
				mathOperation:  "*",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "-1",
			wantErr:    false,
		},
		{
			name: "* 5 word",
			args: args{
				valueOne:       5,
				valueTwo:       5,
				valueThree:     5,
				mathOperation:  "*",
				numToWord:      true,
				cutFirstSymbol: false,
			},
			wantResult: "один два пять",
			wantErr:    false,
		},
		{
			name: "* 5 word without first",
			args: args{
				valueOne:       5,
				valueTwo:       5,
				valueThree:     5,
				mathOperation:  "*",
				numToWord:      true,
				cutFirstSymbol: true,
			},
			wantResult: "дин ва ять",
			wantErr:    false,
		},
		{
			name: "* -5 word",
			args: args{
				valueOne:       -5,
				valueTwo:       -5,
				valueThree:     -5,
				mathOperation:  "*",
				numToWord:      true,
				cutFirstSymbol: false,
			},
			wantResult: "минус один два пять",
			wantErr:    false,
		},
		{
			name: "* -5 word without first",
			args: args{
				valueOne:       -5,
				valueTwo:       -5,
				valueThree:     -5,
				mathOperation:  "*",
				numToWord:      true,
				cutFirstSymbol: true,
			},
			wantResult: "инус дин ва ять",
			wantErr:    false,
		},
		{
			name: "/ 0",
			args: args{
				valueOne:       0,
				valueTwo:       0,
				valueThree:     0,
				mathOperation:  "/",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "",
			wantErr:    true,
		},
		{
			name: "/ 1",
			args: args{
				valueOne:       1,
				valueTwo:       1,
				valueThree:     1,
				mathOperation:  "/",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "1",
			wantErr:    false,
		},
		{
			name: "/ -1",
			args: args{
				valueOne:       -1,
				valueTwo:       -1,
				valueThree:     -1,
				mathOperation:  "/",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "-1",
			wantErr:    false,
		},
		{
			name: "- 0",
			args: args{
				valueOne:       0,
				valueTwo:       0,
				valueThree:     0,
				mathOperation:  "-",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "0",
			wantErr:    false,
		},
		{
			name: "- 1",
			args: args{
				valueOne:       1,
				valueTwo:       1,
				valueThree:     1,
				mathOperation:  "-",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "-1",
			wantErr:    false,
		},
		{
			name: "- -1",
			args: args{
				valueOne:       -1,
				valueTwo:       -1,
				valueThree:     -1,
				mathOperation:  "-",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "1",
			wantErr:    false,
		},
		{
			name: "+ 0",
			args: args{
				valueOne:       0,
				valueTwo:       0,
				valueThree:     0,
				mathOperation:  "+",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "0",
			wantErr:    false,
		},
		{
			name: "+ 1",
			args: args{
				valueOne:       1,
				valueTwo:       1,
				valueThree:     1,
				mathOperation:  "+",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "3",
			wantErr:    false,
		},
		{
			name: "+ -1",
			args: args{
				valueOne:       -1,
				valueTwo:       -1,
				valueThree:     -1,
				mathOperation:  "+",
				numToWord:      false,
				cutFirstSymbol: false,
			},
			wantResult: "-3",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MathString(tt.args.valueOne, tt.args.valueTwo, tt.args.valueThree, tt.args.mathOperation, tt.args.numToWord, tt.args.cutFirstSymbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("MathString() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_numbersToWords(t *testing.T) {
	type args struct {
		value          string
		numToWords     bool
		cutFirstSymbol bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "1",
			args: args{
				value:          "1",
				numToWords:     false,
				cutFirstSymbol: false,
			},
			wantResult: "1",
		},
		{
			name: "12 words",
			args: args{
				value:          "12",
				numToWords:     true,
				cutFirstSymbol: false,
			},
			wantResult: "один два",
		},
		{
			name: "-12 words cut",
			args: args{
				value:          "-12",
				numToWords:     true,
				cutFirstSymbol: true,
			},
			wantResult: "инус дин ва",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := numbersToWords(tt.args.value, tt.args.numToWords, tt.args.cutFirstSymbol); gotResult != tt.wantResult {
				t.Errorf("numbersToWords() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMathByte(t *testing.T) {
	type args struct {
		valueOne      int
		valueTwo      int
		valueThree    int
		mathOperation string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []byte
		wantErr    bool
	}{
		{
			name: "* 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "*",
			},
			wantResult: []byte{48},
			wantErr:    false,
		},
		{
			name: "/ 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "/",
			},
			wantResult: nil,
			wantErr:    true,
		},
		{
			name: "* 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "*",
			},
			wantResult: []byte{49},
			wantErr:    false,
		},
		{
			name: "* -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "*",
			},
			wantResult: []byte{45, 49},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MathByte(tt.args.valueOne, tt.args.valueTwo, tt.args.valueThree, tt.args.mathOperation)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MathByte() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMathRune(t *testing.T) {
	type args struct {
		valueOne      int
		valueTwo      int
		valueThree    int
		mathOperation string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []rune
		wantErr    bool
	}{
		{
			name: "* 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "*",
			},
			wantResult: []rune{48},
			wantErr:    false,
		},
		{
			name: "/ 0",
			args: args{
				valueOne:      0,
				valueTwo:      0,
				valueThree:    0,
				mathOperation: "/",
			},
			wantResult: nil,
			wantErr:    true,
		},
		{
			name: "* 1",
			args: args{
				valueOne:      1,
				valueTwo:      1,
				valueThree:    1,
				mathOperation: "*",
			},
			wantResult: []rune{49},
			wantErr:    false,
		},
		{
			name: "* -1",
			args: args{
				valueOne:      -1,
				valueTwo:      -1,
				valueThree:    -1,
				mathOperation: "*",
			},
			wantResult: []rune{45, 49},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MathRune(tt.args.valueOne, tt.args.valueTwo, tt.args.valueThree, tt.args.mathOperation)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MathRune() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMathBigInt(t *testing.T) {
	valueOneBigInt, _ := new(big.Int).SetString("1", 10)
	valueTwoBigInt, _ := new(big.Int).SetString("1", 10)
	valueThreeBigInt, _ := new(big.Int).SetString("1", 10)
	resultBigInt, _ := new(big.Int).SetString("1", 10)

	type args struct {
		valueOneBigInt   big.Int
		valueTwoBigInt   big.Int
		valueThreeBigInt big.Int
		mathOperation    string
	}
	tests := []struct {
		name       string
		args       args
		wantResult big.Int
		wantErr    bool
	}{
		{
			name: "* 1",
			args: args{
				valueOneBigInt:   *valueOneBigInt,
				valueTwoBigInt:   *valueTwoBigInt,
				valueThreeBigInt: *valueThreeBigInt,
				mathOperation:    "*",
			},
			wantResult: *resultBigInt,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MathBigInt(tt.args.valueOneBigInt, tt.args.valueTwoBigInt, tt.args.valueThreeBigInt, tt.args.mathOperation)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathBigInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MathBigInt() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMathBigFloat(t *testing.T) {
	valueOneBigFloat, _ := new(big.Float).SetString("1.1")
	valueTwoBigFloat, _ := new(big.Float).SetString("1.0")
	valueThreeBigFloat, _ := new(big.Float).SetString("1.0")
	resultBigFloat, _ := new(big.Float).SetString("1.1")

	type args struct {
		valueOneBigInt   big.Float
		valueTwoBigInt   big.Float
		valueThreeBigInt big.Float
		mathOperation    string
	}
	tests := []struct {
		name       string
		args       args
		wantResult big.Float
		wantErr    bool
	}{
		{
			name: "* 1",
			args: args{
				valueOneBigInt:   *valueOneBigFloat,
				valueTwoBigInt:   *valueTwoBigFloat,
				valueThreeBigInt: *valueThreeBigFloat,
				mathOperation:    "*",
			},
			wantResult: *resultBigFloat,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := MathBigFloat(tt.args.valueOneBigInt, tt.args.valueTwoBigInt, tt.args.valueThreeBigInt, tt.args.mathOperation)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathBigFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MathBigFloat() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
