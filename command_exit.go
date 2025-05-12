package main

import (
	"fmt"
	pcache "github.com/jjckrbbt/pokedex/internal/pokecache"
	"os"
)

func commandExit(cfg *config, cache *pcache.Cache, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
