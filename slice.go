package main

import (
	"fmt"
)

//func main()  {
//	scores := []int{1,4,222,210}
//	sls := make([]int, 10)
//	tst := make([]int , 0, 10)
//	fmt.Println(scores[1])
//	fmt.Println(sls[0])
//	//the setence below create a slice with length of 0 but with capacity of 10
//	//so there will be a panic: runtime error: index out of range
//	//fmt.Println(tst[0])
//}

func main() {
//// crash
//	scores := make([]int, 0, 10)
//	scores[5] = 9033
//	fmt.Println(scores)

	scores := make([]int, 0, 10)
	// explicitly expand the slice via append()
	scores = append(scores, 5)
	fmt.Println(scores)

	// resize the slice. the max length of slice is it's capacity
	scores = scores[0:6]
	scores[5] = 9033
	fmt.Println(scores)


	scores[1] = 1
	scores[2] = 2
	scores[3] = 3
	scores[4] = 4
	scores[5] = 5
	fmt.Println(scores)
	scores = append(scores, 6)

	fmt.Println(scores)

	stu := []string{"maxwell", "sheep", "mushroom"}
	fmt.Println(stu)
	stu = append(stu, "strawberry")
	fmt.Println(stu)


}
