package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapf(config *Config) error {
	locationsResponse, err := queryAPI(config.Next)
	if err != nil {
		return err
	}
	config.Next = locationsResponse.Next
	config.Previous = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}
	locationsResponse, err := queryAPI(config.Previous)
	if err != nil {
		return err
	}
	config.Next = locationsResponse.Next
	config.Previous = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func queryAPI(pageURL *string) (ResponseLocations, error) {
	baseURL := "https://pokeapi.co/api/v2"
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
