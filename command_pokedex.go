package main

import (
	"fmt"
)

func commandPokedex(_ *string) error {
	if len(pokedex.CaughtPokemon) == 0 {
		fmt.Println("You have not caught any Pokemon yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedex.CaughtPokemon {
		fmt.Printf(" - %s\n", properTitle(pokemon.Name))
	}
	return nil
}
