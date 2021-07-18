package person_test

import (
	"testing"

	"example.com/hello-world/person"
)

func TestPerson(t *testing.T) {
	person1 := person.Person{"Minh", "Nguyen", "Da Nang", "No", 25}
	person1.HasLanguage()

	if person1.GetPerson() != "Hello, My name is Minh Nguyen. I living in the Da Nang city and I am 25 years old. Language often used : N/a" {
		t.Fatal("Not matching person 1")
	}

	person2 := person.Person{"Tri", "Chau", "Dak Lak", "Javascript", 30}
	if person2.GetPerson() != "Hello, My name is Tri Chau. I living in the Dak Lak city and I am 30 years old. Language often used : Javascript" {
		t.Fatal("Not matching person 2")
	}
}
