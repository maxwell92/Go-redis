package main

import (
	"fmt"
	"encoding/json"
	"os"
	//"bufio"
	//"io"
	//"io/ioutil"
)

//func main()  {
//	filename := os.Args[1]
//	f, err := os.Open(filename, os.O_RDONLY)
//	if err != nil {
//		panic(err)
//	}
//
//
//	reader := bufio.NewReader(f)
//
//}

var repo struct {
	repositories []string `json: "repositories"`
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	//var byt []byte
	//r := io.Reader(f)
	//dec := json.NewDecoder(r)

	jsonParser := json.NewDecoder(f)

	if err = jsonParser.Decode(&repo); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", repo.repositories)
}
