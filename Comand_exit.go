package main

import (
	"fmt"
	"os"
)

func exitCommand(cfg *config, waste string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
