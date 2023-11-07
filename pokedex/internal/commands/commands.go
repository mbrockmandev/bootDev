package commands

import (
	"fmt"

	"github.com/mbrockmandev/bootDev/pokedex/internal/pokeapi"
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
			callback:    Help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    Exit,
		},
		"map": {
			name:        "map",
			description: "Display the names of locations in the Pokemon world",
			callback:    Map,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of locations in the Pokemon world (previous)",
			callback:    Mapb,
		},
	}
}

func Help(cfg *Config) error {
	cmds := getAllAppCommands()
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	for _, v := range cmds {
		fmt.Printf("\n > %s: %s", v.name, v.description)
	}
	fmt.Println()
	fmt.Println()
	return nil
}

func Exit(cfg *Config) error {
	fmt.Println("Thanks for using the Pokedex!")
	return nil
}

func Map(cfg *Config) error {
	res, err := pokeapi.GetNext(cfg.Next)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, v := range res.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()

	if res.Previous != nil {
		cfg.Previous = *res.Previous
	} else {
		cfg.Previous = ""
	}
	cfg.Next = res.Next

	return nil
}

func Mapb(cfg *Config) error {
	res, err := pokeapi.GetPrevious(cfg.Previous)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, v := range res.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()

	if res.Previous != nil {
		cfg.Previous = *res.Previous
	} else {
		cfg.Previous = ""
	}
	cfg.Next = res.Next

	return nil
}
