package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must specify a location to explore")
	}
	locationName := args[0]
	location, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		fmt.Printf("Unknown location '%s'\n", locationName)
		return err
	}
	fmt.Printf("Exploring %s...\n", locationName)
	if len(location.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
	} else {
		fmt.Println("No Pokemon found.")
	}
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
