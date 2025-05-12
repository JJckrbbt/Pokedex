package main

import (
	"fmt"
	pcache "github.com/jjckrbbt/pokedex/internal/pokecache"
)

func commandHelp(cfg *config, cache *pcache.Cache, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
