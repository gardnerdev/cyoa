package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gardnerdev/cyoa"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
}
