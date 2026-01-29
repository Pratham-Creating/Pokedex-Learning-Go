package main

import (
	"errors"
	"fmt"

	"github.com/Pratham-Creating/Pokedex-Learning-Go/pokeapi"
)

func mapCommandforward(cfg *config, waste string) error {
	location, err := pokeapi.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = location.Next
	cfg.prevLocationsURL = location.Prev

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func mapCommandbackword(cfg *config, waste string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	location, err := pokeapi.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = location.Next
	cfg.prevLocationsURL = location.Prev

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
