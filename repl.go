package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRpl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		"exit": {
			name:        "Exit",
			description: "Exit from pokedex",
			callback:    exitCommand,
		},
		"help": {
			name:        "Help",
			description: "Help regarding the usage of pokedex",
			callback:    helpCommand,
		},
		"map": {
			name:        "Location Mapping",
			description: "List the 20 Location Areas",
			callback:    mapCommandforward,
		},
		"mapb": {
			name:        "Location Mapping",
			description: "List the 20 previous Location Areas",
			callback:    mapCommandbackword,
		},
		"explore": {
			name:        "Explore Pokemon",
			description: "List the name of all Pokemon in that region",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "Catch Pokemon",
			description: "Catch a pokemon if you can",
			callback:    catchCommand,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(text)

		if len(parts) == 0 {
			continue
		}

		commandName := parts[0]
		arg := ""
		if len(parts) > 1 {
			arg = parts[1]
		}

		isCommand := false
		for key, _ := range commands {
			if commandName == key {
				if err := commands[key].callback(cfg, arg); err != nil {
					fmt.Println(err)
				}
				isCommand = true
				break
			}
		}

		if !isCommand {
			fmt.Println("Unknown command")
		}

	}
}
