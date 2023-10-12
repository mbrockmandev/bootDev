package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mainLoop()
}

func mainLoop() {
	cfg := &Config{Previous: "", Next: ""}
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		s.Scan()

		text := s.Text()
		if len(text) == 0 {
			fmt.Println("Please enter a command.")
			continue
		} else if text == "help" {
			cmdHelp(cfg)
		} else if text == "exit" {
			cmdExit(cfg)
			break
		} else if text == "map" {
			cmdMap(cfg)
		} else if text == "mapb" {
			cmdMapb(cfg)
		}
	}
}
