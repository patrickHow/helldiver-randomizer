package main

import (
	"fmt"
	"strings"
)

// Custom ExcludeFlag type - parse into a slice of strings
type ExcludeFlag []string

// Implement the flag.Value interface

func (f *ExcludeFlag) String() string {
	return fmt.Sprint(*f)
}

// Set parses a comma-separated string and adds integers to the slice
func (f *ExcludeFlag) Set(value string) error {
	if len(*f) > 0 {
		return fmt.Errorf("exclude flag already in use")
	}

	for _, v := range strings.Split(value, ",") {
		*f = append(*f, v)
	}
	return nil
}
