package main

import "fmt"

type Saiya struct {
	Name int
	Power int
}
func main() {
	sy := make([]Saiya, 5)
	var powers []int
	spawn(sy)
	powers = extractPowers(sy)
	fmt.Println(powers)
}


func spawn(saiyas []Saiya) {
	for i := 0; i < len(saiyas); i++ {
		saiyas[i].Name = i
		saiyas[i].Power = i * 1000
	}

}

func extractPowers(saiyans []Saiya) []int {
	powers := make([]int, len(saiyans))
	for index, saran := range saiyans {
		powers[index] = saran.Power
	}
	return powers
}

// if it's []*Saiya type, panic: runtime error: invalid memory address or nil pointer dereference
// but it's fine when []Saiya type.