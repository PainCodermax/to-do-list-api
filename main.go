package main

import (
	"log"

	"github.com/PainCodermax/to-do-list-api/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
