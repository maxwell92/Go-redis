package main

import "fmt"

func main()  {
	scores := []int{1, 2, 3, 4, 5}
	fmt.Println(scores)
	removeAtIndex(scores, 2)

	// a cool exchange
	a := 1
	b := 2
	fmt.Println(a)
	fmt.Println(b)
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

func removeAtIndex(slice []int, index int) {
	lastIndex := len(slice) - 1
	//swap the value we want to remove and the last value
	slice[index] = slice[lastIndex]
	slice = slice[:lastIndex]
	fmt.Println(slice)

}
