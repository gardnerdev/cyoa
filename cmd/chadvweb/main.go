package main

import (
	"flag"
	"fmt"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)
}
