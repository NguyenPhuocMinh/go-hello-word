package main

import (
	"fmt"

	"example.com/hello-world/area"
	"example.com/hello-world/greetings"
	"example.com/hello-world/person"
)

func main() {

	fmt.Println(greetings.GreertingName("Minh"))

	// struct
	person := person.Person{"Minh", "Nguyen", "Da Nang", "Javascript", 25}
	person.HasLanguage()

	// interfaces
	circle := area.Circle{Radius: 10.5}
	rectangle := area.Rectangle{With: 10.5, Height: 5.5}

	fmt.Printf("Circle %f\n", area.GetArea(circle))
	fmt.Printf("Rectangle %f\n", area.GetArea(rectangle))

	fmt.Println(person.GetPerson())
}
