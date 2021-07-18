package person

import (
	"strconv"
)

type Person struct {
	FirstName string
	LastName  string
	City      string
	Language  string
	Age       int
}

func (p Person) GetPerson() string {
	return "Hello, " + "My name is " + p.FirstName + " " + p.LastName +
		". I living in the " + p.City + " city " + "and I am " +
		strconv.Itoa(p.Age) + " years old" + ". Language often used : " + p.Language
}

func (p *Person) HasLanguage() {
	if p.Language == "No" {
		p.Language = "N/a"
	}
}
