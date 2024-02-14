package main

import (
	"log"
	"os"

	"github.com/kenjitheman/url-probe/core"
)

func main() {
	if err := core.Run(os.Args[1:]); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
