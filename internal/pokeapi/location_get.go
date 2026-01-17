package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	data, exists := c.pokeCache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Location{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Location{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Location{}, err
		}
		c.pokeCache.Add(url, data)
	}

	location := Location{}
	err := json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}
