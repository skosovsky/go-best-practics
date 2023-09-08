package areacalculate

import (
	"fmt"
	"testing"
)

func TestAreaCalculate(t *testing.T) {
	testCases := []struct {
		name     string
		lenSide1 int
		lenSide2 int
		output   int
	}{
		{
			name:     "Case with len side 5 and 5",
			lenSide1: 5,
			lenSide2: 5,
			output:   25,
		},
		{
			name:     "Case with len side 10 and 15",
			lenSide1: 10,
			lenSide2: 15,
			output:   150,
		},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.name, func(t *testing.T) {
			result := AreaCalculate(cse.lenSide1, cse.lenSide2)
			if result != cse.output {
				t.Errorf("function returns wrong answer: expected %d, got: %d", cse.output, result)
			}
		})

	}
}

func BenchmarkAreacalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreaCalculate(5, 10)
	}
}

func ExampleAreacalculate() {
	lenSide1 := 5
	lenSede2 := 10
	result := AreaCalculate(lenSide1, lenSede2)

	fmt.Println(result)
	// output: 50
}
