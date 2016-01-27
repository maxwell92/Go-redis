package main

import (
	"fmt"
	"encoding/json"
	"time"
)

type people struct {
	name string
	birthday string
	age int
}

//type logtype struct {
//	time map[string]string
//	event map[string]string
//	error map[string]string
//}

func main()  {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	bolB, _ = json.Marshal(1)
	fmt.Println(string(bolB))

	bolB, _ = json.Marshal("hello world")
	fmt.Println(string(bolB))

	bolB, _ = json.Marshal(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(string(bolB))

	magic := &people{
		name: "magic",
		birthday: "1992-03-15",
		age: 23,
	}

	bolB, _ = json.Marshal(magic)
	fmt.Println(string(bolB))

	saiya := make(map[string]string, 5)
	saiya["goku"]="sun"
	saiya["gohan"]="sun"
	saiya["krillin"]="None"

	bolB, _ = json.Marshal(saiya)
	fmt.Println(string(bolB))

	//loga := new(logtype)
	//loga.time = make(map[string]string, 1)
	//loga.event = make(map[string]string, 1)
	//loga.error = make(map[string]string, 1)
	//loga.time["time"]=time.Now().Format("2006-01-02 15:04:05")
	//loga.event["event"]="Segment Fault"
	//loga.error["error"]="Fatal"
	//
	//bolB, _ = json.Marshal(loga)
	//fmt.Println(string(bolB))


}

