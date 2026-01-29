package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"github.com/Pratham-Creating/home/pratham/Personal/Golang/Projects/Pokedec/pokeapi"
)

type config struct {
	Next     string
	Previous string
}

//https://pokeapi.co/api/v2/location-area/{id or name}/

func main() {
	URL := baseURL
	scanner := bufio.NewScanner(os.Stdin)

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Errorf("Error fetching the location data")
	}

	defer res.Body.Close()

	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	commands := map[string]cliCommand{
		"exit": {
			name:        "Exit",
			description: "Exit from pokedex",
			callback:    exitCommand,
		},
		"help": {
			name:        "help",
			description: "Help regarding the usage of pokedex",
			callback:    helpCommand,
		},
	}

	for {

		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		isCommand := false
		for key, _ := range commands {
			if text == key {
				commands[key].callback()
				isCommand = true
				break
			}
		}

		if !isCommand {
			fmt.Println("Unknown command")
		}

	}
}
