package main

import "fmt"

// In this case,
// At the first, the length of scores is 0 and the capacity is 5
// then we append scores with i (0..25), every time append 1 number
// Once the length is larger than capacity, the capacity will grow at 2x times of the former capacity



func main() {
	scores := make([]int, 0, 5)
	c := cap(scores) //get the capacity
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}

	fmt.Println(scores)
}
