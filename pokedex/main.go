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
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		s.Scan()

		text := s.Text()
		if len(text) == 0 {
			fmt.Println("Please enter a command.")
			continue
		} else if text == "help" {
			cmdHelp()
		} else if text == "exit" {
			cmdExit()
			break
		}
	}
}
