package main

import (
	"log"

	"github.com/phanpak/fizzbuzz-server/fizzbuzz"
)

func main() {
	s := fizzbuzz.NewServer()

	err := s.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
