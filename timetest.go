package main

import (
	"fmt"
	"time"
)

func main()  {

	timeLayout := "2006-01-02 15:04:05"

	fmt.Println(time.Now())
	fmt.Println(time.Stamp)

	timestamp :=  time.Now().Unix()
	fmt.Println(timestamp)

	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format(timeLayout))

	fmt.Println(time.UnixDate)
	fmt.Println(time.Now().Format(timeLayout))

	fmt.Println("nano")
	fmt.Println(time.Now().Nanosecond() % 3)
	time.Sleep(time.Second)
	fmt.Println(time.Now().Nanosecond() % 3)

	fmt.Println("Unix")
	fmt.Println(time.Now().Unix() % 3)
	time.Sleep(time.Second)
	fmt.Println(time.Now().Unix() % 3)
}

