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

	vegeta := &Saiya{"Vegeta", 8000}
	err := vegeta.Magic(vegeta.Power)

	err = goku.Magic(goku.Power)
	if err != nil {
		fmt.Println("Something is wrong")
	}
}

func (s *Saiya) Magic (power int) (err error) {
	if s.Name != "Goku" {
		fmt.Println("It's not Sun Goku")
	} else {
		s.Power *= 2
		fmt.Println(s.Power)
	}

	return nil
}

func Super(s *Saiya) {
	s = &Saiya{"Gohan", 3000}
}

func Add(s *Saiya) {
	s.Power += 10000
}

