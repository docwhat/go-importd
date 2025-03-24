package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("go-importd: ")

	config, err := parseFlags(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = serve(config)
	if err != nil {
		log.Fatal(err)
	}
}
