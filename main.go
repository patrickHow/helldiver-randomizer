package main

import (
	"flag"
)

func main() {

	randomizer := NewRandomizer()

	flag.Parse()

	if randomizer.InitProfile() {
		randomizer.InitExcludeLists()
		randomizer.Run()
		randomizer.Cleanup()
	}
}
