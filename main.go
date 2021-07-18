package main

import (
	"fmt"

	"example.com/hello-world/greetings"
	"example.com/hello-world/person"
)

func main() {

	fmt.Println(greetings.GreertingName("Minh"))

	// struct
	person := person.Person{"Minh", "Nguyen", "Da Nang", 25}

	fmt.Println(person.GetPerson())
}
