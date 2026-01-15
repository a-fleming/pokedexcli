package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Next     *string
	Previous *string
}

type ResponseLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := Config{
		Next:     nil,
		Previous: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}
		commandStr := cleaned[0]
		command, ok := getCommands()[commandStr]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(&config)
		if err != nil {
			fmt.Println(err)
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
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
