package main

import (
	"flag"
	"fmt"
)

func main() {

	// Flag setup

	// Core flags for selecting individual slots
	orbital := flag.Bool("o", false, "Roll a new orbital strat")
	eagle := flag.Bool("e", false, "Roll a new eagle strat")
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
	var xOrbital ExcludeFlag
	flag.Var(&xOrbital, "xo", "Exclude these orbitals")
	var xEagle ExcludeFlag
	flag.Var(&xEagle, "xe", "Exclude these eagles")
	var xWeapon ExcludeFlag
	flag.Var(&xWeapon, "xw", "Exclude these support weapons")
	var xUtil ExcludeFlag
	flag.Var(&xUtil, "xu", "Exclude these utility items")
	var xPrimary ExcludeFlag
	flag.Var(&xPrimary, "xp", "Exclude these primary weapons")
	var xSecondary ExcludeFlag
	flag.Var(&xSecondary, "xs", "Exclude these secondary weapons")
	var xGrenade ExcludeFlag
	flag.Var(&xGrenade, "xg", "Exclude these grenades")

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

		// Track if we have rolled anything - this is to account for the case
		// where we just get exclude lists and we want to roll everything
		rolled := false

		// Parse exclude lists first
		if len(xOrbital) > 0 {
			orbitalSlot.ParseExcludeFromFlag(xOrbital)
		}

		if len(xEagle) > 0 {
			eagleSlot.ParseExcludeFromFlag(xEagle)
		}

		if len(xWeapon) > 0 {
			weaponSlot.ParseExcludeFromFlag(xWeapon)
		}

		if len(xUtil) > 0 {
			utilitySlot.ParseExcludeFromFlag(xUtil)
		}

		if len(xPrimary) > 0 {
			primarySlot.ParseExcludeFromFlag(xPrimary)
		}

		if len(xSecondary) > 0 {
			secondarySlot.ParseExcludeFromFlag(xSecondary)
		}

		if len(xGrenade) > 0 {
			grenadeSlot.ParseExcludeFromFlag(xGrenade)
		}

		if *slotInfo != "" {
			rolled = true // Ok technically we haven't rolled anything, but assume the user just wants info
			loadout.ParseAndPrintInfo(slotInfo)
		}

		if *rollMulti != "" {
			rolled = true
			loadout.RollMultipleSlots(rollMulti)
		}

		if *orbital {
			rolled = true
			fmt.Println("Orbital:", orbitalSlot.Choose())
		}

		if *eagle {
			rolled = true
			fmt.Println("Eagle:", eagleSlot.Choose())
		}

		if *weapon {
			rolled = true
			fmt.Println("Weapon:", weaponSlot.Choose())
		}

		if *util {
			rolled = true
			fmt.Println("Utility:", utilitySlot.Choose())
		}

		if *primary {
			rolled = true
			fmt.Println("Primary:", primarySlot.Choose())
		}

		if *secondary {
			rolled = true
			fmt.Println("Secondary:", secondarySlot.Choose())
		}

		if *grenade {
			rolled = true
			fmt.Println("Grenade:", grenadeSlot.Choose())
		}

		if !rolled {
			loadout.ChooseAll()
		}
	}
}
