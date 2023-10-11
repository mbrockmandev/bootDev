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
		fmt.Printf("%s:%s", v.name, v.description)
	}
	return nil
}

func cmdExit() error {
	return nil
}
