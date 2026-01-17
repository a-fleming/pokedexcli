package main

import (
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsResponse.Next
	cfg.prevLocationsURL = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsResponse.Next
	cfg.prevLocationsURL = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}
