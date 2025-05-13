package main

import (
	"encoding/json"
	"fmt"
	papi "github.com/jjckrbbt/pokedex/internal/pokeapi"
	pcache "github.com/jjckrbbt/pokedex/internal/pokecache"
	rand "math/rand"
	"time"
)

func commandCatch(cfg *config, cache *pcache.Cache, args []string) error {
	target := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", target)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", target)
	data, exists := cache.Get(url)
	if !exists {
		var err error
		data, err = papi.GetResponse(url)
		if err != nil {
			return err

		}
		cache.Add(url, data)
	}
	var targetPokemon papi.Pokemon

	err := json.Unmarshal(data, &targetPokemon)
	if err != nil {
		return err
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	throw := r.Float64()

	var targetXP float64
	targetXP = float64(targetPokemon.BaseExperience) / 300

	if throw > targetXP {
		fmt.Printf("%s was caught!\n", target)
		cfg.Collection[targetPokemon.Name] = targetPokemon
	} else {
		fmt.Printf("%s escaped!\n", target)
	}
	return nil
}
