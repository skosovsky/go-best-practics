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

func (c Creatures) CompareByLegs(i int, values ...int) {
	for _, n := range values {
		if i == 2 {
			if n > Arr[i].Leg+Arr[i].Body {
				fmt.Println(n, "больше чем ног у", Arr[i].Creature)
			} else if n < Arr[i].Leg+Arr[i].Body {
				fmt.Println(n, "меньше чем ног у", Arr[i].Creature)
			} else {
				fmt.Println(n, "равен ногам у", Arr[i].Creature)
			}
		} else {
			if n > Arr[i].Leg {
				fmt.Println(n, "больше чем ног у", Arr[i].Creature)
			} else if n < Arr[i].Leg {
				fmt.Println(n, "меньше чем ног у", Arr[i].Creature)
			} else {
				fmt.Println(n, "равен ногам у", Arr[i].Creature)
			}
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
