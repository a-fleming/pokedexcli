package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResponseLocations, error) {
	url := baseURL + "/location-area/?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocations{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return ResponseLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocations{}, err
	}
	locationsRes := ResponseLocations{}
	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return ResponseLocations{}, err
	}
	return locationsRes, nil
}
