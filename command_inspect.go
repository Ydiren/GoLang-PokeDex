package main

import (
	"errors"
	"fmt"
)

func commandInspect(pokemonName *string) error {
	pokemonDetails, ok := pokedex.CaughtPokemon[*pokemonName]
	if !ok {
		return errors.New(fmt.Sprintf("You have not caught that pokemon!"))
	}

	fmt.Println(fmt.Sprintf("Name: %s", pokemonDetails.Name))
	fmt.Println(fmt.Sprintf("Height: %d", pokemonDetails.Height))
	fmt.Println(fmt.Sprintf("Weight: %d", pokemonDetails.Weight))
	fmt.Println(fmt.Sprintf("Stats:"))
	for _, stat := range pokemonDetails.Stats {
		fmt.Println(fmt.Sprintf("\t- %s: %d", stat.Stat.Name, stat.BaseStat))
	}
	fmt.Println(fmt.Sprintf("Types:"))
	for _, type_ := range pokemonDetails.Types {
		fmt.Println(fmt.Sprintf("\t- %s", type_.Type.Name))
	}

	return nil
}
