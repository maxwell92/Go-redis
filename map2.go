package main

import (
	"fmt"
)


func main()  {
	lookup := map[string]int {
		"Goku": 30,
		"Gohan": 7,
		"Vegeta": 35,
		"Trunks": 3,
	}

	for key, value := range lookup {
		fmt.Println(key)
		fmt.Println(value)
	}
}

