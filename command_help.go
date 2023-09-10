package main

import "fmt"

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
