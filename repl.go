package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokedex          map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		commandArgs := []string{}
		if length := len(cleaned); length > 1 {
			commandArgs = cleaned[1:length]
		}
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, commandArgs...)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "List information about a Pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 map locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map locations",
			callback:    commandMapb,
		},
	}
}
