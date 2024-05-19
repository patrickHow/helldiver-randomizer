package main

import (
	"fmt"
	"math/rand"
)

type Slot struct {
	options []string
	exclude map[string]bool
}

func NewSlot(options []string) *Slot {
	return &Slot{
		options: options,
		exclude: make(map[string]bool),
	}
}

func (slot *Slot) Choose() string {
	// Pick an entry for the slot
	choice := slot.options[rand.Intn(len(slot.options))]

	// Check that it's not in the exclude list - if it is, pick new entries
	for _, present := slot.exclude[choice]; present; {
		choice = slot.options[rand.Intn(len(slot.options))]
	}

	return choice
}

func (slot *Slot) AddToExcludeList(exc string) {
	slot.exclude[exc] = true
}

func (slot *Slot) PrintWithNumbers() {
	for ind, val := range slot.options {
		fmt.Printf("%d: %s\n", ind+1, val)
	}
}
