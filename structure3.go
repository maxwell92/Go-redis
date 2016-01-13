package main

import "fmt"

type Person struct {
	Name string
	age int
}

type Saiya struct {
	*Person
	Name string
	Power int
}

func main()  {
	goku := &Saiya{
		Person: &Person {
			Name: "EarthMan",
			age: 30,
		},
		Name: "Goku",
		Power: 9000,
	}
	vegeta := Saiya{
		Name: "Vegeta",
		Power: 8000,
	}

	gohan := Saiya{
		Name: "Gohan",
		Person: &Person{
			Name: "PiccoMan",
			age: 23,
		},
		Power: 5000,
	}

	fmt.Println(goku.Name) //Person.Name
	fmt.Println(goku.Name) //Saiya.Name
	fmt.Println(goku.Person.Name)
	fmt.Println(goku.Power)
	fmt.Println(goku.age)


	fmt.Println(vegeta.Name)
	fmt.Println(vegeta.Power)

	fmt.Println(gohan.Name)
	fmt.Println(gohan.Person.Name)
	fmt.Println(gohan.age)
	fmt.Println(gohan.Power)
}

// what's the difference between goku's definiton & vegeta's definition?

// Notice that if not specifed, goku's name is "Goku" not goku.Person's "EarthMan", but goku's age does.

// If not specifed, the member in one structure should be created follow the order.