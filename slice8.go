package main

import (
	"fmt"
	"sort"
	"math/rand"
)

func main() {
	scores := make([]int, 10)
	for i := 0; i < 10; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	fmt.Println(scores)


	worst := make([]int ,5)
	copy(worst, scores[:5])
	fmt.Println(worst)


	copy(worst[2:4], scores[:5])
	fmt.Println(worst)
	// worst[0] [1] [2] [3] [4]
	//scores        [0] [1]

	copy(worst, scores[1:3])
	fmt.Println(worst)
	// worst[0] [1] [2] [3] [4]
	//scores[1] [2]


}
