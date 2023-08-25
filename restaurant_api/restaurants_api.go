package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const serverPort = 8080

// Restaurant represents the
type Restaurant struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	AverageRating   float32 `json:"avg_rating"`
	HasVeganOptions bool    `json:"has_vegan_options"`
}

var restaurants = map[string]Restaurant{
	"123": {
		"1",
		"Papa Joe's Pizza",
		4.9,
		true,
	},
}

func main() {
	http.HandleFunc("/restaurants", handleRequest)
	log.Printf("Server Up and Running. Listening on Port: %d \n", serverPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRestaurantById(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getRestaurantById(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[2]
	if _, ok := restaurants[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	restaurant := restaurants[id]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(restaurant)
}
