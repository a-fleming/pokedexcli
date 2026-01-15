package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapf(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	if config.Next != nil {
		url = *config.Next
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var jsonData locationAreaQuery
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&jsonData)
	if err != nil {
		return err
	}
	for _, location := range jsonData.Results {
		fmt.Println(location.Name)
	}
	config.Next = jsonData.Next
	config.Previous = jsonData.Previous
	return nil
}

func commandMapb(config *Config) error {
	if config.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}
	config.Next = config.Previous
	return commandMapf(config)
}
