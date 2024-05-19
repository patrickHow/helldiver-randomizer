package main

import (
	"flag"
	"fmt"
)

func main() {

	// Flag setup

	// Core flags for selecting individual slots
	eagle := flag.Bool("e", false, "Roll a new eagle strat")
	orbital := flag.Bool("o", false, "Roll a new orbital strat")
	weapon := flag.Bool("w", false, "Roll a new support weapon")
	util := flag.Bool("u", false, "Roll a new utility strat")
	primary := flag.Bool("p", false, "Roll a new primary")
	secondary := flag.Bool("s", false, "Roll a new secondary")
	grenade := flag.Bool("g", false, "Roll a new grenade")

	// Flag to print list for specific slots
	slotInfo := flag.String("i", "", "Print options for a slot")

	// Flag to select multiple slots for individual rolls
	rollMulti := flag.String("r", "", "Roll multiple slots")

	// Flags to set exclude values for specific slots

	flag.Parse()

	// Set up slot variables
	orbitalSlot := NewSlot(OrbitalList)
	eagleSlot := NewSlot(EagleList)
	weaponSlot := NewSlot(WeaponList)
	utilitySlot := NewSlot(UtilityList)
	primarySlot := NewSlot(PrimaryList)
	secondarySlot := NewSlot(SecondaryList)
	grenadeSlot := NewSlot(GrenadeList)

	// Loadout variable to hold everything
	loadout := Loadout{
		*orbitalSlot,
		*eagleSlot,
		*weaponSlot,
		*utilitySlot,
		*primarySlot,
		*secondarySlot,
		*grenadeSlot}

	// If no args, print everything
	if flag.NFlag() == 0 {
		loadout.ChooseAll()

	} else {
		// Check individual flags

		if *slotInfo != "" {
			loadout.ParseAndPrintInfo(slotInfo)
		}

		if *rollMulti != "" {
			loadout.RollMultipleSlots(rollMulti)
		}

		if *orbital {
			fmt.Println("Orbital:", orbitalSlot.Choose())
		}

		if *eagle {
			fmt.Println("Eagle:", eagleSlot.Choose())
		}

		if *weapon {
			fmt.Println("Weapon:", weaponSlot.Choose())
		}

		if *util {
			fmt.Println("Utility:", utilitySlot.Choose())
		}

		if *primary {
			fmt.Println("Primary:", primarySlot.Choose())
		}

		if *secondary {
			fmt.Println("Secondary:", secondarySlot.Choose())
		}

		if *grenade {
			fmt.Println("Grenade:", grenadeSlot.Choose())
		}
	}
}
