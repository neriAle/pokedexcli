package pokeapi

import(
	"fmt"
	"io"
	"net/http"
)

// Makes a request to the given url and returns the body of the response as a slice of bytes
// if something went wrong, an empty slice and an error is returned instead
func Get_api_data(url string) ([]byte, error) {
	// Execute the GET request
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("error requesting the data from APIs: %w", err)
	}

	defer res.Body.Close()

	// If the request wasn't successful, return an error
	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	// Transform the body of the response to a slice of bytes
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error parsing the data: %w", err)
	}

	return body, nil
}