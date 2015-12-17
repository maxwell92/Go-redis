package main
import (
	"fmt"
)

type Saiya struct {
	Name string
	Power int
	Father *Alien
}

type Alien struct {
	Name string
	Power int
	Planet string
}

func NewSaiya(name string, power int) *Saiya {
	return &Saiya{
		Name: name,
		Power: power,
	}
}

func OldSaiya(name string, power int) Saiya {
	return Saiya{
		Name: name,
		Power: power,
	}
}

func main()  {
	goku := NewSaiya("Goku",9000)
	fmt.Println(goku.Name)
	gohan := OldSaiya("Gohan",5000)
	fmt.Println(gohan.Power)

	vegeta := new(Saiya)
	// same as
	trunks := &Saiya{ }

	fmt.Println(vegeta.Name)
	fmt.Println(trunks.Power)

	mushroom := &Saiya {
		Name: "Gohan",
		Power: 10000,
		Father: &Alien{
			Name: "Goku",
			Power: 9000,
			Planet: "alpha",
		},
	}
	fmt.Println(mushroom.Name)

}
