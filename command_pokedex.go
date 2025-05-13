package main

import (
	"fmt"
	pcache "github.com/jjckrbbt/pokedex/internal/pokecache"
)

func commandPokedex(cfg *config, cache *pcache.Cache, args []string) error {
	if len(cfg.Collection) < 1 {
		fmt.Println("You haven't caught any Pokemon yet, go get'em")
	}

	for pokemon, _ := range cfg.Collection {
		fmt.Println("Your Pokedex:")
		fmt.Println("  -", pokemon)
	}
	return nil
}
