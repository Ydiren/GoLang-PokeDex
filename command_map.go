package main

import (
	"fmt"
	"github.com/ydiren/pokedexcli/internal/pokeapi"
)

func commandMap(_ *string) error {
	err := pokedex.CurrentLocation.GetNextLocations()
	if err != nil {
		return err
	}

	printLocations(&pokedex.CurrentLocation)
	return nil
}

func commandMapB(_ *string) error {
	err := pokedex.CurrentLocation.GetPreviousLocations()
	if err != nil {
		return err
	}

	printLocations(&pokedex.CurrentLocation)
	return nil
}

func printLocations(locations *pokeapi.PokeLocations) {
	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}
}
