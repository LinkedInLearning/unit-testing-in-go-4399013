package restaurant

import (
	"testing"
)

type restaurant struct {
	Name    string
	Cuisine string
}

func (r restaurant) Equal(other restaurant) bool {
	return r.Name == other.Name && r.Cuisine == other.Cuisine
}

func TestRestaurantWithCustomEquals(t *testing.T) {
	expected := restaurant{Name: "Jos Pizza", Cuisine: "Italian"}
	actual := restaurant{Name: "Mama Putt", Cuisine: "Indian"}

	if !expected.Equal(actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
