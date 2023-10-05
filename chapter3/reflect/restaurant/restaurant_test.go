package restaurant

import (
	"reflect"
	"testing"
)

type restaurant struct {
	Name    string
	Cuisine string
}

func TestRestaurantWithReflect(t *testing.T) {
	expected := restaurant{Name: "Jos Pizza", Cuisine: "Italian"}
	actual := restaurant{Name: "Mama Putt", Cuisine: "Indian"}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
