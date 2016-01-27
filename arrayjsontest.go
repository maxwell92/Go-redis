package main

import (
	"fmt"
	"encoding/json"
	"os"
)

var tarray struct {
	Names []string `json: "names"`
	//Page int `json:"page"`
	//Stage string `json:"stage"`
}


func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}


	//Parser := json.NewDecoder(f)

	//tarray := &struct {
	//	name: []string `json: "name"`
	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&tarray); err != nil {
		panic(err)
	}

	//if err = Parser.Decode(&tarray); err != nil {
	//	panic(err)
	//}
	//fmt.Println(tarray.Page)
	//fmt.Println(tarray.Stage)
	fmt.Println(tarray.Names)
}


