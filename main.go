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

func GetCommands() []CliCommand {
	return []CliCommand{
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
			Command:     "explore <area name>",
			Description: "Explore the map location",
			Callback:    commandExplore,
		},
		{
			Command:     "catch <pokemon name>",
			Description: "Catch the named pokemon",
			Callback:    commandCatch,
		},
		{
			Command:     "inspect <pokemon name>",
			Description: "Inspect the named pokemon",
			Callback:    commandInspect,
		},
	}
}
