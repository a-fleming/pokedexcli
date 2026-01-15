package main

import (
	"fmt"
)

func commandMapf(cfg *config) error {
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}
	cfg.NextLocationsURL = locationsResponse.Next
	cfg.PrevLocationsURL = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.PrevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}
	cfg.NextLocationsURL = locationsResponse.Next
	cfg.PrevLocationsURL = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}
