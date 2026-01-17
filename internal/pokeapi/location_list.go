package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationListResponse, error) {
	url := baseURL + "/location-area/?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}
	data, exists := c.pokeCache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationListResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationListResponse{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationListResponse{}, err
		}
		c.pokeCache.Add(url, data)
	}

	locationsRes := LocationListResponse{}
	err := json.Unmarshal(data, &locationsRes)
	if err != nil {
		return LocationListResponse{}, err
	}
	return locationsRes, nil
}
