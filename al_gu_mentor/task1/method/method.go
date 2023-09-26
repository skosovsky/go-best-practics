package method

import "fmt"

var (
	man = Creature{
		Creature: "человек",
		Body:     1,
		Leg:      2,
		Tail:     0,
	}

	dog = Creature{
		Creature: "собака",
		Body:     1,
		Leg:      4,
		Tail:     1,
	}

	snail = Creature{
		Creature: "улитка",
		Body:     1,
		Leg:      0,
		Tail:     0,
	}
	Arr = Creatures{man, dog, snail}
)

type Creature struct {
	Creature string
	Body     int
	Leg      int
	Tail     int
}

type Creatures []Creature

func (c Creatures) CompareLegsByMan(values ...int) {
	for _, n := range values {
		if n > Arr[0].Leg {
			fmt.Println(n, "больше чем ног у человека")
		} else if n < Arr[0].Leg {
			fmt.Println(n, "меньше чем ног у человека")
		} else {
			fmt.Println(n, "равен ногам у человека")
		}
	}
}

func (c Creatures) CompareLegsByDog(values ...int) {
	for _, n := range values {
		if n > Arr[1].Leg {
			fmt.Println(n, "больше чем ног у собаки")
		} else if n < Arr[1].Leg {
			fmt.Println(n, "меньше чем ног у собаки")
		} else {
			fmt.Println(n, "равен ногам у собаки")
		}
	}
}

func (c Creatures) CompareLegsBySnail(values ...int) {
	for _, n := range values {
		if n > Arr[2].Leg+Arr[2].Body {
			fmt.Println(n, "больше чем ног у улитки")
		} else if n < Arr[2].Leg+Arr[2].Body {
			fmt.Println(n, "меньше чем ног у улитки")
		} else {
			fmt.Println(n, "равен ногам у улитки")
		}
	}
}

func (c Creatures) CountLimb(i int) int {
	return c[i].Leg + c[i].Tail
}

func CompareLImb() {
	if Arr.CountLimb(0) > Arr.CountLimb(1) {
		fmt.Println("у человека конечностей больше чем у собаки")
	} else if Arr.CountLimb(0) < Arr.CountLimb(1) {
		fmt.Println("у человека конечностей меньше чем у собаки")
	} else {
		fmt.Println("у человека столько же конечностей как и у собаки")
	}
}
