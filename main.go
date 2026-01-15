package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var commands map[string]cliCommand

type Config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type locationAreaQuery struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func init() {
	commands = map[string]cliCommand{
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map locations",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for name, command := range commands {
		fmt.Println(name + ": " + command.description)
	}
	return nil
}

func commandMap(config *Config) error {
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
		fmt.Println("you're on the first page")
		return nil
	}
	config.Next = config.Previous
	return commandMap(config)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := Config{
		Next:     "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0",
		Previous: "",
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleaned := cleanInput(userInput)
		if len(cleaned) == 0 {
			continue
		}
		commandStr := cleaned[0]
		command, ok := commands[commandStr]
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
