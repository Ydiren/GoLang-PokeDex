package main

import (
	"errors"
	"fmt"
	"github.com/ydiren/pokedexcli/internal/pokeapi"
	"math/rand"
	"strings"
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

	fmt.Printf("Pokemon Details: %d\n", pokemonDetails.BaseExperience)

	random := rand.Intn(pokemonDetails.BaseExperience)

	if random > 100 {
		fmt.Printf("%s ran away!\n", *pokemonName)
		return nil
	} else {
		fmt.Printf("%s was caught!\n", *pokemonName)
	}

	key := strings.ToLower(*pokemonName)
	pokedex.CaughtPokemon[key] = *pokemonDetails

	return nil
}
