package main

import (
	"fmt"
	"encoding/json"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type people struct {
	Name string
	Birthday string
	Age int
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main()  {
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	magic := &people{
		Name: "magic",
		Birthday: "1992-03-15",
		Age: 23,
	}
	bolB, _ := json.Marshal(magic)
	fmt.Println(string(bolB))

	res2D := &Response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"page": 1, "fruits": ["apple", "peach", "pear"]}`)
	var dat map[string]interface{}

	err := json.Unmarshal(byt, &dat)
	if err != nil {
		panic(nil)
	}
	fmt.Println(dat)


	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}
