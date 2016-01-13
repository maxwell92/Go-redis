package main

import (
	"fmt"
)

type Saiya struct {
	Name int
	Power int
}

func main()  {
	sy := make([]Saiya, 5)
	var powers []int
	spawn(sy)
	powers = extractPowers(sy)
	fmt.Println(powers)


	slice := sy[2:4]
	slice[0].Name = 100
	slice[0].Power = 100
	fmt.Println(slice)

	// the original array doesn't change. the slice is complete new copy of the array.
	// but in golang, it's not the same
	fmt.Println(sy)
}


func spawn(saiyas []Saiya) {
	for i := 0; i < len(saiyas); i++ {
		saiyas[i].Name = i
		saiyas[i].Power = i * 1000
	}

}

func extractPowers(saiyans []Saiya) []int {
	powers := make([]int, 0, len(saiyans))
	for _, saran := range saiyans {
		powers = append(powers, saran.Power)
	}
	return powers
}



