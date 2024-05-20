package main

import (
	"fmt"
)

type ProfileSwitch string

func (f *ProfileSwitch) String() string {
	return fmt.Sprint(*f)
}

// Validate the switch provided
func (f *ProfileSwitch) Set(value string) error {
	if len(*f) > 0 {
		return fmt.Errorf("profile switch already set")
	}

	switch value {
	case "e", "error":
		*f = "e"
	case "d", "delete":
		*f = "d"
	case "c", "create":
		*f = "c"
	case "i", "info":
		*f = "i"
	case "u", "use":
		*f = "u"
	default:
		return fmt.Errorf("invalid profile switch")
	}

	return nil
}
