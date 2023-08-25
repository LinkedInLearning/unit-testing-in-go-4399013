package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetRestaurantById(t *testing.T) {
	// Create a sample restaurant
	restaurant := Restaurant{
		ID:              "123",
		Name:            "Test Restaurant",
		AverageRating:   3.5,
		HasVeganOptions: true,
	}

	// Create a new request with the restaurant ID in the URL path
	req, err := http.NewRequest("GET", "/restaurants/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function with the response recorder and request
	handler := http.HandlerFunc(handleRequest)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	var retrievedRestaurant Restaurant
	err = json.Unmarshal(rr.Body.Bytes(), &retrievedRestaurant)
	if err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	// Compare the retrieved restaurant with the expected restaurant
	if reflect.DeepEqual(retrievedRestaurant, restaurant) {
		t.Errorf("retrieved restaurant does not match the expected restaurant")
	}
}
