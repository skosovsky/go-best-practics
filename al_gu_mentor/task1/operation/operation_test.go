package operation

import "testing"

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
