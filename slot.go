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

	// Edge case - have we excluded the whole slot?
	if len(slot.exclude) >= len(slot.options) {
		return "no valid options"
	}

	// Pick an entry for the slot
	choice := slot.options[rand.Intn(len(slot.options))]

	// Check that it's not in the exclude list - if it is, pick new entries
	for _, present := slot.exclude[choice]; present; {
		choice = slot.options[rand.Intn(len(slot.options))]
	}

	return choice
}

func (slot *Slot) ParseExcludeFromFlag(exl []int) {
	for _, exi := range exl {
		if exi <= len(slot.options) {
			slot.exclude[slot.options[exi-1]] = true
			fmt.Println("Excluding:", slot.options[exi-1])
		} else {
			fmt.Printf("Arg %d is not valid for this slot\n", exi)
		}
	}
}

func (slot *Slot) PrintWithNumbers() {
	for ind, val := range slot.options {
		fmt.Printf("%d: %s\n", ind+1, val)
	}
}
