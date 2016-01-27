package main

import (
	"fmt"
	"time"
	//"os"
	//"bufio"
)


func main()  {
	interval := 1
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("logA running...")
			time.Sleep(time.Second * time.Duration(interval))
		}
	}()
	time.Sleep(time.Second * time.Duration(interval * 5))
}

