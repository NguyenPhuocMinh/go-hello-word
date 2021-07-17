package greetings

import (
	"fmt"
)

func GreertingName(name string) string {
	message := fmt.Sprintf("Hi %s", name)

	return message
}
