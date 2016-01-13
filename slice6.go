package main

import (
	"fmt"
	"strings"
)

func main()  {
	slogan := "Hello Mushrom from the Golang"
	fmt.Println(slogan)

	fmt.Println(strings.Index(slogan[3:], " ")) // how many " " frm the [3] of slogan
	// [x:] : from x to
	// [:x] : from   to x
	// go doesn't support negative values

}
