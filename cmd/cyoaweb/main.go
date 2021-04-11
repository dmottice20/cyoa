package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dmottice20/cyoa"
)

func main() {
	// Run our web application from here when we have it...
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.", *filename)

	// open the file
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	// pass in io.read --i.e. a file
	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
