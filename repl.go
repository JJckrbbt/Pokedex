package main

import (
	"bufio"
	"fmt"
	papi "github.com/jjckrbbt/pokedex/internal/pokeapi"
	pcache "github.com/jjckrbbt/pokedex/internal/pokecache"
	"os"
	"strings"
	"time"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	cache := pcache.NewCache(5 * time.Second)

	cfg.Collection = make(map[string]papi.Pokemon)

	for {
		if cfg.Next == "" {
			cfg.Next = "https://pokeapi.co/api/v2/location-area"
		}

		fmt.Printf("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&cfg, cache, words[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pcache.Cache, []string) error
}

type config struct {
	Next       string
	Previous   string
	Collection map[string]papi.Pokemon
}

var cfg config

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Returns the next list of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Returns a list of previous locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Returns a list of Pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays a Pokemon's stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays list of caught Pokemon",
			callback:    commandPokedex,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}
