package main

import (
	"fmt"

	"github.com/mbrockmandev/bootDev/pokedex/pokeapi"
)

type appCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
}

type Config struct {
	Next     string
	Previous string
}

func getAllAppCommands() map[string]appCommand {
	return map[string]appCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    cmdHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of locations in the Pokemon world",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of locations in the Pokemon world (previous)",
			callback:    cmdMapb,
		},
	}
}

func cmdHelp(cfg *Config) error {
	cmds := getAllAppCommands()
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range cmds {
		fmt.Printf("\n%s: %s", v.name, v.description)
	}
	fmt.Println()
	return nil
}

func cmdExit(cfg *Config) error {
	fmt.Println("Thanks for using the Pokedex!")
	return nil
}

func cmdMap(cfg *Config) error {
	res, err := pokeapi.GetNext(cfg.Next)
	if err != nil {
		return err
	}

	if res.Previous != nil {
		cfg.Previous = *res.Previous
	} else {
		cfg.Previous = ""
	}
	cfg.Next = res.Next

	return nil
}

func cmdMapb(cfg *Config) error {
	res, err := pokeapi.GetPrevious(cfg.Previous)
	if err != nil {
		return err
	}

	if res.Previous != nil {
		cfg.Previous = *res.Previous
	} else {
		cfg.Previous = ""
	}
	cfg.Next = res.Next

	return nil
}
