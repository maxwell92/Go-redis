package main
import "fmt"

// Case 1:
//type Saiya struct {
//	Name string
//	Power int
//}
//
//func main() {
//	goku := Saiya{
//		Name: "Goku",
//		Power: 9000,
//	}
//	Display(goku)
//}
//
//
//func Display(m Saiya) (err error) {
//	fmt.Println(m.Name)
//	fmt.Println(m.Power)
//	return nil
//}

// Case 2:
//type Saiya struct {
//	Name string
//	Power int
//}
//
//func main() {
//	//goku := Saiya{9000, "Goku"}// cannot use 9000 (type int) as type string in field value
//	//goku := Saiya{"Goku" 9000}// non-declaration statement outside function body
//
//	goku := &Saiya{"Goku", 9000}
//	fmt.Println(goku.Name)
//	fmt.Println(goku.Power)
//}

// Case 3:
//type Saiya struct {
//	Name string
//	Power int
//}
//
//func main() {
//	goku := newSaiya("Goku", 9000)
//	fmt.Println(goku.Name)
//	fmt.Println(goku.Power)
//
//	gohan := anotherSaiya("Gohan", 5000)
//	fmt.Println(gohan.Name)
//	fmt.Println(gohan.Power)
//
//
//	//Show1(goku) // cannot use goku (type *Saiya) as type Saiya in arguement to Show1
//	Show1(gohan)
//	Show2(goku)
//	//Show2(gohan) // cannot use gohan (type Saiya) as type *Saiay in arguement to Show2
//}
////Constructor newSaiya()
//func newSaiya(name string, power int) *Saiya {
//	return &Saiya{
//		Name: name,
//		Power: power,
//	}
//	//return &Saiya{name, power}
//}
//
////Constructor anotherSaiya()
//func anotherSaiya(name string, power int) Saiya {
//	return Saiya{
//		Name: name,
//		Power: power,
//	}
//}
//
//
//func Show1(m Saiya) {
//	fmt.Println(m.Name)
//}
//
//func Show2(m *Saiya) {
//	fmt.Println(m.Name)
//}

// Case 4:
//type Saiya struct {
//	Name string
//	Power int
//}
//
//func main() {
//	goku := new(Saiya)
//	goku.Name = "Goku"
//	goku.Power = 9000
//	fmt.Println(goku.Name)
//	fmt.Println(goku.Power)
//}

//-------------------------------- Struct Compose ------------------------------------
//type Person struct {
//	Name string
//	Age int
//}
//type Saiya struct {
//	*Person
//	Name string
//	Power int
//}
//
//func main() {
//	goku := &Saiya {
//		Person: &Person {
//			Name: "Sun",
//			Age: 30,
//		},
//		Name: "Goku",
//		Power: 9000,
//	}
//
//	fmt.Println(goku.Name)
//	fmt.Println(goku.Person.Name)
//	fmt.Println(goku.Age)
//	fmt.Println(goku.Power)
//}

type Person struct {
	Name string
	Age int
}
type Alient struct {
	Planet string
	Name string
}
type Saiya struct {
	*Person
	Name string
	Power int
	Father *Alient
}

func main() {
	goku := &Saiya {
		Person: &Person {
			Name: "Sun",
			Age: 30,
		},
		Name: "Goku",
		Power: 9000,
		Father: &Alient {
			Planet: "Kuru",
			Name: "Alient",
		},
	}

	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)
	fmt.Println(goku.Age)
	fmt.Println(goku.Power)

	//fmt.Println(goku.Planet)
	fmt.Println(goku.Father.Name)
	fmt.Println(goku.Father.Planet)

	goku.show()
	goku.Person.show()
}
// Because we didn't give Person an explicit field name, we can impcility access the fields and functions of the composed type.
// If not specifed it will be 'overload'.
// Unlike the Alien case that we have to access Father's fields and functions.

func (s *Saiya) show() {
	fmt.Println(s.Name)
}

func (p *Person) show() {
	fmt.Println(p.Name)
}