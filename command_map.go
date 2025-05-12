package main

import (
	"encoding/json"
	"fmt"
	papi "github.com/jjckrbbt/pokedex/internal/pokeapi"
	pcache "github.com/jjckrbbt/pokedex/internal/pokecache"
)

func commandMap(cfg *config, cache *pcache.Cache, args []string) error {

	data, exists := cache.Get(cfg.Next)
	if !exists {
		var err error
		data, err = papi.GetResponse(cfg.Next)
		if err != nil {
			return err

		}
		cache.Add(cfg.Next, data)
	}

	var pokeMap papi.Pokemap

	err := json.Unmarshal(data, &pokeMap)
	if err != nil {
		return err
	}

	next := pokeMap.Next
	previous := pokeMap.Previous

	var locations []string
	for _, pokeMap := range pokeMap.Results {
		loc := fmt.Sprintf("%s\n", pokeMap.Name)
		locations = append(locations, loc)
	}

	for _, location := range locations {
		fmt.Printf("%s", location)
	}

	cfg.Next = next
	cfg.Previous = previous

	return nil
}

func commandMapb(cfg *config, cache *pcache.Cache, args []string) error {
	data, exists := cache.Get(cfg.Previous)
	if !exists {
		var err error
		data, err = papi.GetResponse(cfg.Previous)
		if err != nil {
			return err

		}
		cache.Add(cfg.Next, data)
	}

	var pokeMap papi.Pokemap

	err := json.Unmarshal(data, &pokeMap)
	if err != nil {
		return err
	}

	next := pokeMap.Next
	previous := pokeMap.Previous

	var locations []string
	for _, pokeMap := range pokeMap.Results {
		loc := fmt.Sprintf("%s\n", pokeMap.Name)
		locations = append(locations, loc)
	}

	for _, location := range locations {
		fmt.Printf("%s", location)
	}

	cfg.Next = next
	cfg.Previous = previous

	return nil
}
