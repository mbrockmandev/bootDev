package main

import (
	"os"
)

func handleDebugMode() bool {
	// set through command line args
	if len(os.Args) > 1 && os.Args[1] == "--debug" {
		return true
	}

	// set through env vars
	if value := os.Getenv("DEBUG"); value == "true" {
		return true
	}

	return false
}
