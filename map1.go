package main

import (
	"fmt"
)

type Saiya struct {
	Name string
	Friends map[string]*Saiya
}

func main() {

	lookup := make(map[string]int)
	lookup["goku"] = 9001
	fmt.Println(lookup)

	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)

	total := len(lookup)
	fmt.Println(total)

	delete(lookup, "goku")

	newmap := make(map[string]int, 10)
	fmt.Println(newmap)

	goku := &Saiya{
		Name: "Goku",
		Friends: make(map[string]*Saiya),
	}
	krillin := &Saiya{
		Name: "Krillin",
		Friends: make(map[string]*Saiya),
	}


	goku.Friends["Krillin"] = krillin
	krillin.Friends["Goku"] = goku

	fmt.Println(krillin.Name)
	fmt.Println(&krillin.Name)
	fmt.Println(goku.Friends)
	fmt.Println(goku.Friends["Krillin"].Name)












}
