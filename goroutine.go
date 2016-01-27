package main

import (
	"fmt"
	"time"
)

//func main()  {
//	go func() {
//		for {
//			time.Sleep(time.Second)
//			fmt.Println("This is a go routine")
//		}
//	}()
//}

func say(s string) {
	for i := 0 ; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("world")
}
