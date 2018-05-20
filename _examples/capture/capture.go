package main

import (
	"fmt"

	"github.com/jaytaylor/archive.is"
)

var captureURL = "https://jaytaylor.com/"

func main() {
	archiveURL, err := archiveis.Capture(captureURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully archived %v via archive.is: %v\n", captureURL, archiveURL)
}
