package main

import (
	"bufio"
	"fmt"
	. "github.com/ydiren/pokedexcli/internal/cliProcessor"
	. "github.com/ydiren/pokedexcli/internal/pokeapi"
	"os"
)

type Pokedex struct {
	CurrentLocation PokeLocations
	CaughtPokemon   map[string]PokemonDetails
}

var pokedex = Pokedex{
	CurrentLocation: PokeLocations{
		Next:     DefaultApiLocationsUri,
		Previous: nil,
	},
	CaughtPokemon: make(map[string]PokemonDetails),
}

func main() {
	commands := GetCommands()

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

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Command:     "help",
			Description: "Displays this help message",
			Callback:    commandHelp,
		},
		"exit": {
			Command:     "exit",
			Description: "Exits the application",
			Callback:    commandExit,
		},
		"map": {
			Command:     "map",
			Description: "Retrieves the next page of map locations",
			Callback:    commandMap,
		},
		"mapb": {
			Command:     "mapb",
			Description: "Retrieves the previous page of map locations",
			Callback:    commandMapB,
		},
		"explore": {
			Command:     "explore <area name>",
			Description: "Explore the map location",
			Callback:    commandExplore,
		},
		"catch": {
			Command:     "catch <pokemon name>",
			Description: "Catch the named pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Command:     "inspect <pokemon name>",
			Description: "Inspect the named pokemon",
			Callback:    commandInspect,
		},
	}
}
