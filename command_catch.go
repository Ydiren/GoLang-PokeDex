package main

import (
	"errors"
	"fmt"
	"github.com/ydiren/pokedexcli/internal/pokeapi"
	"math/rand"
)

func commandCatch(pokemonName *string) error {
	if pokemonName == nil {
		return errors.New("Please provide a pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at: %v...\n", *pokemonName)

	pokemonDetails, err := pokeapi.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon Details: %v\n", pokemonDetails.BaseExperience)

	random := rand.Intn(pokemonDetails.BaseExperience)

	if random > 100 {
		fmt.Printf("%v ran away!\n", *pokemonName)
		return nil
	} else {
		fmt.Printf("%v was caught!\n", *pokemonName)
	}

	pokedex.Pokemon[*pokemonName] = *pokemonDetails

	return nil
}
