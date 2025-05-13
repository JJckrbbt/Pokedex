package main

import (
	"fmt"
	pcache "github.com/jjckrbbt/pokedex/internal/pokecache"
)

func commandInspect(cfg *config, cache *pcache.Cache, args []string) error {
	target := args[0]

	pokemon, exists := cfg.Collection[target]
	if exists {

		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, ptype := range pokemon.Types {
			fmt.Printf("  - %s\n", ptype.Type.Name)
		}
	} else {
		fmt.Println("You haven't caught any Pokemon yet, go get'em")
	}
	return nil
}
