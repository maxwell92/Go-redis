package main

import (
	"fmt"
)

func main() {
	var test[10]int
	test[0] = 1
	fmt.Println(test[0])
	fmt.Println(test[1])


	array := [4]int {1, 2, 3, 4}
	fmt.Println(array[3])

	fmt.Print("The length of array is: ")
	fmt.Println(len(array ))

}


