package area_test

import (
	"testing"

	"example.com/hello-world/area"
)

func TestCircle(t *testing.T) {
	circle := area.Circle{Radius: 10.5}
	rectangle := area.Rectangle{With: 10.5, Height: 5.5}

	if area.GetArea(circle) != 346.36059005827474 {
		t.Fatal("Wrong get area circle")
	}

	if area.GetArea(rectangle) != 57.750000 {
		t.Fatal("Wrong get area rectangle")
	}
}
