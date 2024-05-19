package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Custom ExcludeFlag type - parse into a slice of ints
type ExcludeFlag []int

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
		num, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*f = append(*f, num)
	}
	return nil
}
