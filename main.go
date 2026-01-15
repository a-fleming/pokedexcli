package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleaned := cleanInput(userInput)
		fmt.Println("Your command was: " + cleaned[0])
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
