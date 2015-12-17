package main

import (
	"fmt"
)

func log (message string) {
}

func add (a int, b int) int {
	return a + b
}

func power (name string) (int, bool) {
	return 9000, true
}

func main ()  {
	value, exists := power("goku")

	if exists == false {
		// handle this error case
		fmt.Errorf("error")
	}

	if value < 9000 {
		fmt.Println("This is not Goku")
	}

	_, result := power("goku")

	if result == true {
		fmt.Println("true")
	}

	log("this is a test")
	fmt.Println(add(1, 2))
}

