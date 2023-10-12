package main

import "fmt"

type appCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func cmdHelp() error {
	cmds := getAllAppCommands()
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range cmds {
		fmt.Printf("\n%s: %s", v.name, v.description)
	}
	fmt.Println()
	return nil
}

func cmdExit() error {
	fmt.Println("Thanks for using the Pokedex!")
	return nil
}
