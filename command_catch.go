package main

import (
	"fmt"
	"math/rand/v2"
)

const catchThresholdPct float64 = 0.6

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must specify a Pokemon to catch")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		//fmt.Printf("Unknown Pokemon '%s'\n", pokemonName)
		return fmt.Errorf("Unknown Pokemon '%s'", pokemonName)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	catchThresholdVal := float64(pokemon.BaseExperience) * catchThresholdPct
	catchVal := rand.Float64() * float64(pokemon.BaseExperience)
	if catchVal >= catchThresholdVal {
		fmt.Printf("%s was caught!\n", pokemonName)
		if _, exists := cfg.pokedex[pokemonName]; !exists {
			cfg.pokedex[pokemonName] = pokemon
			fmt.Printf("%s was added to the Pokedex.\n", pokemonName)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}
