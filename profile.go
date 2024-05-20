package main

import (
	"fmt"
	"os"
	"strings"
)

type Profile struct {
	Name        string
	DefaultRoll string

	// Exclude lists
	XOrbital   []string
	XEagle     []string
	XWeapon    []string
	XUtil      []string
	XPrimary   []string
	XSecondary []string
	XGrenade   []string
}

// Write the whole profile to a string
func (prof *Profile) ToString() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Name: %s\n", prof.Name))
	sb.WriteString(fmt.Sprintf("DefaultRoll: %s\n", prof.DefaultRoll))

	sb.WriteString(fmt.Sprintf("XOrbital: %s\n", strings.Join(prof.XOrbital, ",")))
	sb.WriteString(fmt.Sprintf("XEagle: %s\n", strings.Join(prof.XEagle, ",")))
	sb.WriteString(fmt.Sprintf("XWeapon: %s\n", strings.Join(prof.XWeapon, ",")))
	sb.WriteString(fmt.Sprintf("XUtil: %s\n", strings.Join(prof.XUtil, ",")))
	sb.WriteString(fmt.Sprintf("XPrimary: %s\n", strings.Join(prof.XPrimary, ",")))
	sb.WriteString(fmt.Sprintf("XSecondary: %s\n", strings.Join(prof.XSecondary, ",")))
	sb.WriteString(fmt.Sprintf("XGrenade: %s\n", strings.Join(prof.XGrenade, ",")))

	return sb.String()
}

// Parse the string from the file into the struct
func (prof *Profile) FromString(data string) {
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) < 2 {
			continue
		}
		key := parts[0]
		value := parts[1]

		// Workaround - splitting an empty string results in a string slice
		// containing only an empty string - but this slice has length 1!
		// Which means it will get passed to the ignore function
		// And the empty string will fuzzy match everything - ignoring the entire slot
		if len(value) == 0 {
			continue
		}

		switch key {
		case "Name":
			prof.Name = value
		case "DefaultRoll":
			prof.DefaultRoll = value
		case "XOrbital":
			prof.XOrbital = strings.Split(value, ",")
		case "XEagle":
			prof.XEagle = strings.Split(value, ",")
		case "XWeapon":
			prof.XWeapon = strings.Split(value, ",")
		case "XUtil":
			prof.XUtil = strings.Split(value, ",")
		case "XPrimary":
			prof.XPrimary = strings.Split(value, ",")
		case "XSecondary":
			prof.XSecondary = strings.Split(value, ",")
		case "XGrenade":
			prof.XGrenade = strings.Split(value, ",")
		}
	}
}

func (prof *Profile) ReadFromFile() bool {
	raw, err := os.ReadFile(prof.Name + ".profile")

	if err != nil {
		fmt.Println("Could not open file:", err)
		return false
	} else {
		prof.FromString(string(raw))
		fmt.Println("Loaded profile:", prof.Name)
		return true
	}
}

func (prof *Profile) WriteToFile() {
	data := prof.ToString()
	err := os.WriteFile(prof.Name+".profile", []byte(data), 0664)
	if err != nil {
		fmt.Println("Error writing to profile:", err)
	} else {
		fmt.Println("Writing profile:", prof.Name)
	}
}

func (prof *Profile) SetName(name string) {
	// Trim spaces from the name since it will be used as a filename
	prof.Name = strings.ReplaceAll(name, " ", "")
}

func NewProfile() *Profile {
	return &Profile{}
}

func (prof *Profile) Delete() {
	_ = os.Remove(prof.Name + ".profile")
}

func (prof *Profile) Describe() {
	fmt.Println("Profile:", prof.Name)
	fmt.Println("Default roll:", prof.DefaultRoll)
	fmt.Println("Exclude list:")

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s\n", strings.Join(prof.XOrbital, ", ")))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Join(prof.XEagle, ", ")))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Join(prof.XWeapon, ", ")))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Join(prof.XUtil, ", ")))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Join(prof.XPrimary, ", ")))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Join(prof.XSecondary, ", ")))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Join(prof.XGrenade, ", ")))

	fmt.Println(sb.String())

}
