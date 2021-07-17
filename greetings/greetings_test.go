package greetings_test

import (
	"fmt"
	"testing"

	"example.com/hello-world/greetings"
)

func TestGreeting(t *testing.T) {
	if greetings.GreertingName("Minh") != fmt.Sprintf("Hi %s", "Minh") {
		t.Fatal("Wrong")
	}
}
