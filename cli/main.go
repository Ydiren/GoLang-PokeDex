package main

import (
	"bufio"
	"errors"
	"fmt"
	. "github.com/ydiren/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	command     string
	description string
	callback    func(*PokeLocations, *string) error
}

func main() {
	pokeData := PokeLocations{
		Next:     DefaultApiLocationsUri,
		Previous: nil,
	}
	commands := map[string]cliCommand{
		"help": {
			command:     "help",
			description: "Displays this help message",
			callback:    commandHelp,
		},
		"exit": {
			command:     "exit",
			description: "Exits the application",
			callback:    commandExit,
		},
		"map": {
			command:     "map",
			description: "Retrieves the next page of map locations",
			callback:    commandMap,
		},
		"mapb": {
			command:     "mapb",
			description: "Retrieves the previous page of map locations",
			callback:    commandMapB,
		},
		"explore": {
			command:     "explore",
			description: "Explore the map location",
			callback:    commandExplore,
		},
	}

	for {
		fmt.Print("\033[32mPokedex > \033[0m")
		scanner := bufio.NewScanner(os.Stdin)
		ok := scanner.Scan()

		var input string
		if ok {
			input = scanner.Text()

			command, ok := commands[input]
			if ok {
				firstArg := getFirstCommandArg(input)
				err := command.callback(&pokeData, firstArg)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func getFirstCommandArg(input string) *string {
	args := strings.Fields(input)
	if len(args) > 1 {
		return &args[1]
	}

	return nil
}

func commandHelp(_ *PokeLocations, _ *string) error {
	fmt.Println("PokeDex CLI")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("\tmap\tDisplay the next set of map locations")
	fmt.Println("\tmapb\tDisplay the previous set of map locations")
	fmt.Println("\thelp\tPrints this help text")
	fmt.Println("\texit\tExits the application")
	fmt.Println()
	return nil
}

func commandExit(_ *PokeLocations, _ *string) error {
	os.Exit(0)
	return nil
}

func commandMap(pokeData *PokeLocations, _ *string) error {
	err := pokeData.GetNextLocations()
	if err != nil {
		return err
	}

	printLocations(pokeData)
	return nil
}

func commandMapB(pokeData *PokeLocations, _ *string) error {
	err := pokeData.GetPreviousLocations()
	if err != nil {
		return err
	}

	printLocations(pokeData)
	return nil
}
func printLocations(locations *PokeLocations) {
	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}
}

func commandExplore(locations *PokeLocations, locationName *string) error {
	if locationName == nil {
		return errors.New("Command locationName is nil")
	}

	////err := locations.GetLocationPokemon(locationName)
	//if err != nil {
	//	return err
	//}

	return nil
}
