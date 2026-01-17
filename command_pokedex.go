package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You have not caught any Pokemon!")
	} else {
		fmt.Println("Your Pokedex:")
		for name := range cfg.pokedex {
			fmt.Printf(" - %s\n", name)
		}
	}
	return nil
}
