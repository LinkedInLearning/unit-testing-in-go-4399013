package restaurant

import (
	"testing"
)

type restaurant struct {
	Name    string
	Cuisine string
}

func TestRestaurant(t *testing.T) {

	expected := restaurant{Name: "Jos Pizza", Cuisine: "Italian"}
	actual := restaurant{Name: "Mama Putt", Cuisine: "Indian"}

	if expected.Name != actual.Name || expected.Cuisine != actual.Cuisine {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
