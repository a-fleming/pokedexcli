package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must specify a Pokemon to inspect")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	printStats(pokemon)
	return nil
}

func printStats(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}
}
