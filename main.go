package main

import (
	"github.com/phanpak/fizzbuzz-server/fizzbuzz"
)

func main() {
	s := fizzbuzz.NewServer()
	s.Serve()
}
