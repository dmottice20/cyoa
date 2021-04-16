package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dmottice20/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
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

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
