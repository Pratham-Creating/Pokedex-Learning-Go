package main

import (
	"fmt"
	"math/rand"

	"github.com/Pratham-Creating/Pokedex-Learning-Go/pokeapi"
)

func catchCommand(cfg *config, name string) error {

	rescatch, err := pokeapi.ReponseCatch(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	luck := rand.Intn(rescatch.BaseExperience)
	fmt.Print(luck)
	if luck > rescatch.BaseExperience/2 {
		fmt.Printf("%v was caught!\n", name)
	} else {
		fmt.Printf("%v escaped!\n", name)
	}
	return nil
}
