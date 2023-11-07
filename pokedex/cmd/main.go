package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mbrockmandev/bootDev/pokedex/internal/commands"
)

func main() {
	mainLoop()
}

func mainLoop() {
	cfg := &commands.Config{Previous: "", Next: ""}
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		s.Scan()

		text := s.Text()
		if len(text) == 0 {
			fmt.Println("Please enter a command.")
			continue
		} else if text == "help" {
			commands.Help(cfg)
		} else if text == "exit" {
			commands.Exit(cfg)
			break
		} else if text == "map" {
			commands.Map(cfg)
		} else if text == "mapb" {
			commands.Mapb(cfg)
		} else {
			fmt.Println("Command not recognized. Try again. Please enter \"help\" for a list of commands. ")
		}
	}
}
