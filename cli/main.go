package main

import (
	"bufio"
	"fmt"
	. "github.com/ydiren/pokedexcli/internal/pokeapi"
	//. "github.com/ydiren/pokedexcli/internal/pokecache"
	"os"
)

type cliCommand struct {
	command     string
	description string
	callback    func(*PokeLocations) error
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
				err := command.callback(&pokeData)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func commandHelp(_ *PokeLocations) error {
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

func commandExit(_ *PokeLocations) error {
	os.Exit(0)
	return nil
}

func commandMap(pokeData *PokeLocations) error {
	err := pokeData.GetNextLocations()
	if err != nil {
		return err
	}

	printLocations(pokeData)
	return nil
}

func commandMapB(pokeData *PokeLocations) error {
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
