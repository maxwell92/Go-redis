package main

import (
	"fmt"
)

type Mushroom struct {
	Name string
	Power int
}

func main () {
	goku := Mushroom{
		Name: "Goku",
		Power: 9000,
	}

//	gohan := Mushroom{Name: "Gohan"}
//	gohan.Power = 5000
//
//	vegeta := Mushroom{"Vegeta", 8500}

	Super(goku)
	fmt.Println(goku.Power)

	RealSuper(&goku)
	fmt.Println(goku.Power)

}

func Super(m Mushroom) {
	m.Power += 10000
}

func RealSuper(m *Mushroom) {
	m.Power += 10000
}

