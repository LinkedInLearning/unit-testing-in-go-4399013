package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func NewClient(baseURL string) *Client {
	return &Client{
		client:  &http.Client{Timeout: time.Duration(time.Duration.Seconds(4))},
		baseURL: baseURL,
	}
}

type Client struct {
	client  *http.Client
	baseURL string
}

// makeRequest makes a request to the API and unmarshals the response into the given interface.
func (c *Client) makeRequest(r *http.Request, v interface{}) error {
	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	switch status := resp.StatusCode; {
	case status >= 200 && status < 300: // success response codes.
		if err := json.Unmarshal(data, v); err != nil {
			return fmt.Errorf("error unmarshalling response: %w", err)
		}
		return nil
	default: // 5xx
		return fmt.Errorf("error response: %s", string(data))
	}
}

// GetRestaurantById returns a restaurant by ID.
func (c *Client) GetRestaurantById(id int) (*Restaurant, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/restaurants/%d", c.baseURL, id), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	var restaurant Restaurant
	if err := c.makeRequest(req, &restaurant); err != nil {
		return nil, err
	}

	return &restaurant, nil
}

// CreateRestaurant creates a restaurant.
func (c *Client) CreateRestaurant(name string, hasVeganOptions bool, avgRating float32) (*Restaurant, error) {
	// Create a new restaurant
	newRestaurant := Restaurant{
		Name:            name,
		AverageRating:   avgRating,
		HasVeganOptions: hasVeganOptions,
	}

	payload, err := json.Marshal(newRestaurant)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/restaurants", c.baseURL), bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	var restaurant Restaurant
	if err := c.makeRequest(req, &restaurant); err != nil {
		return nil, err
	}

	return &restaurant, nil
}
