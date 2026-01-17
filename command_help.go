package main

import (
	"fmt"
	"maps"
	"slices"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()
	for _, commandName := range slices.Sorted(maps.Keys(commands)) {
		command := commands[commandName]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
