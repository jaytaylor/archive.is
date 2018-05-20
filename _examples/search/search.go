package main

import (
	"fmt"
	"time"

	"github.com/jaytaylor/archive.is"
)

var searchURL = "https://jaytaylor.com/"

func main() {
	snapshots, err := archiveis.Search(searchURL, 10*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", snapshots)
}
