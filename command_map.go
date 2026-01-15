package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapf(config *Config) error {
	res, err := http.Get(config.Next)
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
	if config.Previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	config.Next = config.Previous
	return commandMapf(config)
}
