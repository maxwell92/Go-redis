package main
import "fmt"

type Saiya struct {
	Name int
	Power int
}

//
//func main() {
//	scores := make([]int, 5)
//	scores = append(scores, 9999)
//	fmt.Println(scores)
//// the result is [0 0 0 0 0 9999] but not [9999 0 0 0 0 0]
//
//// there are 4 ways to initialize the slice
//	checks := make([]bool, 5)
//	name := []string{"Gohan", "Goku", "Vegeta", "Trunks"}
//	//name := []string{"Goku", "Gohan", "Vegeta", "Trunks"}
//	score := make([]int, 0, 10)
//	var times []string
//	fmt.Println(checks)
//	fmt.Println(name)
//	fmt.Println(score)
//	fmt.Println(times)
//
//	//var esf *[]Saiya
//	esf := make([]Saiya, 5)
//	spawn(esf)
//	fmt.Println("sometimes:")
//	showPower(esf)
//	extract(esf)
//	fmt.Println("sometimes:")
//	showPower(esf)
//}
//
//func showPower(m []Saiya) {
//	for i := 0; i < len(m); i++ {
//		fmt.Println(m[i].Name)
//		fmt.Println(m[i].Power)
//	}
//}
//
//
//// if m is slice or array. it's direct (address) refer.
//// So not need to use point.
//
//func spawn(m []Saiya) {
//	for i := 0; i < len(m); i++ {
//		m[i].Name = i
//		m[i].Power = 5000
//	}
//}
//
//func extract(m []Saiya) {
//	for i := 0; i < len(m); i++ {
//		m[i].Power = 2 * m[i].Power
//	}
//}

//func showPower(m *[]Saiya) {
//	for index, m := range m {
//		fmt.Println(m[index].Name)
//		fmt.Println(m.Powet)
//	}
//}
//
//func spawn(m *[]Saiya) {
//	for index, m := range m {
//		m.Name = "1"
//		m.Power = 5000
//	}
//}
//
//func extract(m *[]Saiya) {
//	for index, m := range m {
//		m.Power = 2 * m.Power
//	}
//}

func main() {
	sy := make([]*Saiya, 5)
	var powers []int
	spawn(sy)
	powers = extractPowers(sy)
	fmt.Println(powers)
}

func spawn(saiyans []*Saiya) {
	for index, n := range saiyans{
		n.Name = index
		n.Power = 1000
	}
}


func extractPowers(saiyans []*Saiya) []int {
	powers := make([]int, len(saiyans))
	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}
