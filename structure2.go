package main

import (
	"fmt"
)

type Saiya struct {
	Name string
	Power int
}


func main() {
	goku := &Saiya{"Goku", 9000}
	Super(goku)
	fmt.Println(goku.Power)
	Add(goku)
	fmt.Println(goku.Power)
}

func Super(s *Saiya) {
	s = &Saiya{"Gohan", 3000}
}

func Add(s *Saiya) {
	s.Power += 10000
}

