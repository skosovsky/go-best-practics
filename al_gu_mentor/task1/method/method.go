package method

import "fmt"

type Human struct {
	Title string
	Leg   int
}

type Dog struct {
	Title string
	Paw   int
	Tail  int
}

type Snail struct {
	Title string
	Leg   int
}

type Mammals struct {
	Human
	Dog
}

//type Talker interface {
//	Talk()
//}

func (h Human) CompareWithLegs(values ...int) {
	for _, n := range values {
		if n > h.Leg {
			fmt.Println(n, "больше чем ног у человека")
		} else if n < h.Leg {
			fmt.Println(n, "меньше чем ног у человека")
		} else {
			fmt.Println(n, "равно ногам у человека")
		}
	}
}

func (d Dog) CompareWithsPaws(values ...int) {
	for _, n := range values {
		if n > d.Paw {
			fmt.Println(n, "больше чем лап у собаки")
		} else if n < d.Paw {
			fmt.Println(n, "меньше чем лап у собаки")
		} else {
			fmt.Println(n, "равно лапам у собаки")
		}
	}
}

func (s Snail) CompareWithLegs(values ...int) {
	for _, n := range values {
		if n > s.Leg {
			fmt.Println(n, "больше чем ног у улитки")
		} else if n < s.Leg {
			fmt.Println(n, "меньше чем ног у улитки")
		} else {
			fmt.Println(n, "равно ногам у улитки")
		}
	}
}

func (m Mammals) CompareLimb() {
	if m.Human.Leg > m.Dog.Paw+m.Dog.Tail {
		fmt.Println("у человека конечностей больше чем у собаки")
	} else if m.Human.Leg < m.Dog.Paw+m.Dog.Tail {
		fmt.Println("у человека конечностей меньше чем у собаки")
	} else {
		fmt.Println("у человека столько же конечностей как и у собаки")
	}
}

func (h Human) Talk() {
	fmt.Println("человек говорит привет")
}

func (d Dog) Talk() {
	fmt.Println("собака гавкает гав")
}
