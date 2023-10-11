package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

}

func mainLoop() {
	s := bufio.NewScanner(os.Stdin)
	cmds := getAllAppCommands()

	for {
		fmt.Print("Pokedex > ")

		s.Scan()

		text := s.Text()
		if len(text) == 0 {
			fmt.Println("Please enter a command.")
			continue
		}

		if  {

		}
	}
}
