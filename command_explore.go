package main

import (
	"encoding/json"
	"fmt"
	papi "github.com/jjckrbbt/pokedex/internal/pokeapi"
	pcache "github.com/jjckrbbt/pokedex/internal/pokecache"
)

func commandExplore(cfg *config, cache *pcache.Cache, args []string) error {
	area := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)
	data, exists := cache.Get(url)
	if !exists {
		var err error
		data, err = papi.GetResponse(url)
		if err != nil {
			return err

		}
		cache.Add(url, data)
	}

	var pokeArea papi.PokeArea

	err := json.Unmarshal(data, &pokeArea)
	if err != nil {
		return err
	}

	var pokemon []string
	for _, encounter := range pokeArea.PokemonEncounters {
		name := fmt.Sprintf("%s\n", encounter.Pokemon.Name)
		pokemon = append(pokemon, name)
	}

	fmt.Println("Exploring", area)

	for _, name := range pokemon {
		fmt.Printf("%s", name)
	}

	return nil
}
