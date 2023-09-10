package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(pokemonName *string) error {
	key := strings.ToLower(*pokemonName)
	pokemonDetails, ok := pokedex.CaughtPokemon[key]
	if !ok {
		return errors.New(fmt.Sprintf("You have not caught that pokemon!"))
	}

	fmt.Println(fmt.Sprintf("Name: %s", properTitle(pokemonDetails.Name)))
	fmt.Println(fmt.Sprintf("Height: %d", pokemonDetails.Height))
	fmt.Println(fmt.Sprintf("Weight: %d", pokemonDetails.Weight))
	fmt.Println(fmt.Sprintf("Stats:"))
	for _, stat := range pokemonDetails.Stats {
		fmt.Println(fmt.Sprintf("\t- %s: %d", properTitle(stat.Stat.Name), stat.BaseStat))
	}
	fmt.Println(fmt.Sprintf("Types:"))
	for _, type_ := range pokemonDetails.Types {
		fmt.Println(fmt.Sprintf("\t- %s", properTitle(type_.Type.Name)))
	}

	return nil
}
