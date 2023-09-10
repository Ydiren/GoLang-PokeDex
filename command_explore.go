package main

import (
	"errors"
	"fmt"
	"github.com/ydiren/pokedexcli/internal/pokeapi"
)

func commandExplore(locationName *string) error {
	if locationName == nil {
		return errors.New("Please provide a location name")
	}

	fmt.Printf("Exploring: %v...\n", *locationName)
	fmt.Println("Found CaughtPokemon:")

	pokemon, err := pokeapi.GetPokemonAtLocation(locationName)
	if err != nil {
		return err
	}

	printPokemon(pokemon)
	return nil
}

func printPokemon(pokemon []pokeapi.Pokemon) {
	for i := 0; i < len(pokemon); i++ {
		fmt.Printf(" - %v\n", pokemon[i].Name)
	}
}
