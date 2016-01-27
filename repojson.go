package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type repo struct {
	Repositories []string `json: "Repositories"`
}

func main()  {
	var r repo
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	repoParser := json.NewDecoder(f)
	if err = repoParser.Decode(&r); err != nil {
		panic(err)
	}

	for i := range r.Repositories {
		fmt.Println(i)
		fmt.Println(r.Repositories[i])

	}
	//fmt.Println(r.Repositories)
}
