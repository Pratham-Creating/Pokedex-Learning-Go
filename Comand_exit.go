package main

import (
	"fmt"
	"os"
)

func exitCommand(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("")
}
