package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Slot struct {
	options []string        // Possible options for this slot
	exclude map[string]bool // List of options to exclude - runtime populated
}

func NewSlot(options []string) *Slot {
	return &Slot{
		options: options,
		exclude: make(map[string]bool),
	}
}

func (slot *Slot) Choose() string {

	// Edge case - have we excluded the whole slot?
	if len(slot.exclude) >= len(slot.options) {
		return "no valid options"
	}

	// Pick an entry for the slot
	choice := slot.options[rand.Intn(len(slot.options))]

	// Check that it's not in the exclude list
	// If it is, pick new entires until we get a non-excluded one
	for _, present := slot.exclude[choice]; present; {
		choice = slot.options[rand.Intn(len(slot.options))]
	}

	return choice
}

// This function takes a list that can be a mix of strings and integers
// Integers should correspond to an index to exclude
// Strings will be fuzzy-matched against the list
func (slot *Slot) ParseExcludeFromFlag(exl []string) {

	for _, ex := range exl {
		// Try to parse to an int
		exi, err := strconv.Atoi(ex)
		if err == nil {
			if exi <= len(slot.options) {
				slot.exclude[slot.options[exi-1]] = true
				fmt.Println("Excluding:", slot.options[exi-1])
			} else {
				fmt.Printf("Arg %d is not valid for this slot\n", exi)
			}
		} else {
			// Item could not be parsed to an int, assume it's a string and try to match
			for _, option := range slot.options {
				match := fuzzy.MatchFold(ex, option)
				if match {
					fmt.Printf("Excluding %s on arg %s\n", option, ex)
					slot.exclude[option] = true
				}
			}
		}
	}
}

func (slot *Slot) PrintWithNumbers() {
	for ind, val := range slot.options {
		fmt.Printf("%d: %s\n", ind+1, val)
	}
}

func (slot *Slot) GetExcludeStringList() []string {
	var s []string

	for str := range slot.exclude {
		s = append(s, str)
	}

	return s
}
