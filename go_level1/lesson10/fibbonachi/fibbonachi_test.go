package fibbonachi

import (
	"errors"
	"fmt"
	"testing"
)

func TestFibbonachi(t *testing.T) {
	testCases := []struct {
		name          string
		input         int
		useOptimized  bool
		output        int
		expectedError error
	}{
		{
			name:         "simple test with 10",
			input:        10,
			useOptimized: false,
			output:       55,
		},
		{
			name:         "simple test with 11",
			input:        11,
			useOptimized: true,
			output:       89,
		},
		{
			name:   "case with zero",
			input:  0,
			output: 0,
		},
		{
			name:          "case with negative number",
			input:         -1,
			expectedError: ErrNegativeNumbers,
		},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.name, func(t *testing.T) {
			result, err := Fibbonachi(cse.input, cse.useOptimized)
			if errors.Is(err, cse.expectedError) == false {
				t.Errorf("function didn't return error")
			}

			if cse.expectedError == nil && result != cse.output {
				t.Errorf("function returns wrong answer: expected %d, got: %d", cse.output, result)
			}
		})

	}
}

func BenchmarkFibbonachi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibbonachi(20, true)
	}
}

func BenchmarkFibbonachiSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibbonachi(20, false)
	}
}

func ExampleFibbonachi() {
	input := 10
	useOptimezed := true
	result, err := Fibbonachi(input, useOptimezed)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// output: 55
}
