package main

import (
	"fmt"
)

func main () {
	var power int
	power = 9000
//	panic := 100000
	panic := getPower()
	//	fmt.Println("It's Over %d\n", power)
	fmt.Printf("It's Over %d\n", power)
	fmt.Printf("It's Panic %d\n", panic)
}

func getPower() int{
	return 1000
}

