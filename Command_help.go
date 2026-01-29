package main

import "fmt"

func helpCommand(cfg *config, waste string) error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:1
	
help: Displays a help message
exit: Exit the Pokedex`)
	return fmt.Errorf("")
}
