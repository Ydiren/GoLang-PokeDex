package main

import (
	"bufio"
	"errors"
	"fmt"
	. "github.com/ydiren/pokedexcli/internal/cliProcessor"
	. "github.com/ydiren/pokedexcli/internal/pokeapi"
	"math/rand"
	"os"
)

type Pokedex struct {
	CurrentLocation PokeLocations
	Pokemon         map[string]PokemonDetails
}

var pokedex = Pokedex{
	CurrentLocation: PokeLocations{
		Next:     DefaultApiLocationsUri,
		Previous: nil,
	},
	Pokemon: make(map[string]PokemonDetails),
}

func main() {
	commands := []CliCommand{
		{
			Command:     "help",
			Description: "Displays this help message",
			Callback:    commandHelp,
		},
		{
			Command:     "exit",
			Description: "Exits the application",
			Callback:    commandExit,
		}, {
			Command:     "map",
			Description: "Retrieves the next page of map locations",
			Callback:    commandMap,
		},
		{
			Command:     "mapb",
			Description: "Retrieves the previous page of map locations",
			Callback:    commandMapB,
		},
		{
			Command:     "explore",
			Description: "Explore the map location",
			Callback:    commandExplore,
		},
		{
			Command:     "catch",
			Description: "Catch the pokemon",
			Callback:    commandCatch,
		},
	}

	cli := NewCliProcessor(commands)

	for {
		fmt.Print("\033[32mPokedex > \033[0m")
		scanner := bufio.NewScanner(os.Stdin)
		ok := scanner.Scan()

		if ok {
			input := scanner.Text()
			if len(input) == 0 {
				continue
			}

			err := cli.ProcessCommand(input)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func commandHelp(_ *string) error {
	fmt.Println("PokeDex CLI")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("\tmap\tDisplay the next set of map locations")
	fmt.Println("\tmapb\tDisplay the previous set of map locations")
	fmt.Println("\thelp\tPrints this help text")
	fmt.Println("\texit\tExits the application")
	fmt.Println("\texplore <area name>\tExplore the map location")
	fmt.Println("\tcatch <pokemon name>\tAttempt to catch the pokemon")
	fmt.Println()
	return nil
}

func commandExit(_ *string) error {
	os.Exit(0)
	return nil
}

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

func printLocations(locations *PokeLocations) {
	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}
}

func commandExplore(locationName *string) error {
	if locationName == nil {
		return errors.New("Please provide a location name")
	}

	fmt.Printf("Exploring: %v...\n", *locationName)
	fmt.Println("Found Pokemon:")

	pokemon, err := GetPokemonAtLocation(locationName)
	if err != nil {
		return err
	}

	printPokemon(pokemon)
	return nil
}

func commandCatch(pokemonName *string) error {
	if pokemonName == nil {
		return errors.New("Please provide a pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at: %v...\n", *pokemonName)

	pokemonDetails, err := GetPokemonDetails(pokemonName)
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

func printPokemon(pokemon []Pokemon) {
	for i := 0; i < len(pokemon); i++ {
		fmt.Printf(" - %v\n", pokemon[i].Name)
	}
}
