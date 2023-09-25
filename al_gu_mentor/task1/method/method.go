package method

import "fmt"

type Creature struct {
	Creature string
	Body     int
	Leg      int
	Tail     int
}

type Creatures []Creature

func (c Creatures) CountLegs(i int) int {
	return c[i].Leg
}

func (c Creatures) CountBody(i int) int {
	return c[i].Body
}

func (c Creatures) CountLimb(i int) int {
	return c[i].Body + c[i].Tail
}

func CompareByMethod(values ...int) {
	man := Creature{
		Creature: "человек",
		Body:     1,
		Leg:      2,
		Tail:     0,
	}

	dog := Creature{
		Creature: "собака",
		Body:     1,
		Leg:      4,
		Tail:     1,
	}

	snail := Creature{
		Creature: "улитка",
		Body:     1,
		Leg:      0,
		Tail:     0,
	}

	arr := Creatures{man, dog, snail}

	for _, n := range values {
		for i := 0; i < len(arr); i++ {
			if arr[i].Creature == "улитка" {
				if n > arr.CountLimb(i) {
					fmt.Println(n, "больше чем ног у", arr[i].Creature)
				} else if n < arr.CountLimb(i) {
					fmt.Println(n, "меньше чем ног у", arr[i].Creature)
				} else {
					fmt.Println(n, "равен ногам у", arr[i].Creature)
				}
			} else {
				if n > arr.CountLegs(i) {
					fmt.Println(n, "больше чем ног у", arr[i].Creature)
				} else if n < arr.CountLegs(i) {
					fmt.Println(n, "меньше чем ног у", arr[i].Creature)
				} else {
					fmt.Println(n, "равен ногам у", arr[i].Creature)
				}
			}
		}
	}
}
