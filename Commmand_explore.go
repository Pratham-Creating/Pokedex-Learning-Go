package main

import (
	"fmt"

	"github.com/Pratham-Creating/Pokedex-Learning-Go/pokeapi"
)

func exploreCommand(cfg *config, name string) error {

	pokname, err := pokeapi.ResponseNames(name)
	if err != nil {
		fmt.Errorf("Error getting the names")
	}
	fmt.Printf("Exploring %v\n", name)
	fmt.Println("Found Pokemon:")
	for _, i := range pokname.PokemonEncounter {
		fmt.Println(i.Pokemon.Name)
	}
	return nil
}
