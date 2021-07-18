package person_test

import (
	"testing"

	"example.com/hello-world/person"
)

func TestPerson(t *testing.T) {
	person := person.Person{"Minh", "Nguyen", "Da Nang", 25}

	if person.GetPerson() != "Hello, My name is Minh Nguyen. I living in the Da Nang city and I am 25 years old" {
		t.Fatal("Not matching person")
	}
}
