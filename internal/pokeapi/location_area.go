package pokeapi

import(
	"fmt"
	"io"
	"encoding/json"
	"net/http"
)

type Location_Areas struct {
	Count    int    	`json:"count"`
	Next     *string 	`json:"next"`
	Previous *string	`json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
} 

// Returns a slice of 20 location areas names, as well as the URLs of the previous and next 20, and an error
func Get_location_areas(prev, next string) ([]string, string, string, error) {
	result := []string{}
	var url string
	
	// If config points to a Next_area, GET that, otherwise GET the base URL
	if next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = next
	}

	// Execute the GET request
	res, err := http.Get(url)
	if err != nil {
		return result, prev, next, fmt.Errorf("error retrieving the location areas: %w", err)
	}

	defer res.Body.Close()

	// If the request wasn't successful, return an error
	if res.StatusCode > 299 {
		return result, prev, next, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	// Transform the body of the response to a slice of bytes
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return result, prev, next, fmt.Errorf("error reading the location areas: %w", err)
	}

	// Unmarshal the response into a Location_Areas struct
	var locations = Location_Areas{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return result, prev, next, fmt.Errorf("error unmarshalling the location areas: %w", err)
	}

	// Update the next and previous location of the config struct
	if locations.Next != nil {
		next = *locations.Next
	}
	if locations.Previous != nil {
		prev = *locations.Previous
	} else {
		prev = ""
	}

	for _, loc := range locations.Results {
		result = append(result, loc.Name)
	}
	return result, prev, next, nil
}