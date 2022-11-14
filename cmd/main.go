package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <path>", os.Args[0])
	}

}
