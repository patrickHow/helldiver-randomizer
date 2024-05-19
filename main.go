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

	// Flags related to profile
	loadProfile := flag.String("profile", "", "Specify the profile to use, create, edit, or view")
	var profSwitch ProfileSwitch = "u"
	flag.Var(&profSwitch, "um", "Set the usage mode for the specified profile")

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
		fmt.Println("No flags set, choosing entire loadout")
		loadout.ChooseAll()

	} else {
		// Check individual flags

		// Track if we have rolled anything - this is to account for the case
		// where we just get exclude lists and we want to roll everything
		rolled := false

		// If this flag is true, parsing exclude lists will also update the current profile
		modifyProfile := false

		// Load profile if one is specified
		profile := NewProfile()

		if *loadProfile != "" {
			// Profile not empty,
			profile.SetName(*loadProfile)

			switch profSwitch {
			case "e": // Edit profile - overwrite specified settings based on other flags
				modifyProfile = true
				profile.ReadFromFile()
				defer profile.WriteToFile()
			case "d": // Delete profile - remove from the file system
				profile.Delete()
			case "c": // Create profile based on specified settings
				modifyProfile = true
				profile.DefaultRoll = "oeuwpsg" // Default to rolling everything unless overwritten below
				defer profile.WriteToFile()
			case "i": // Print info on the profile, but take no action based on it
				rolled = true
				profile.ReadFromFile()
				profile.Describe()
			case "u": // Use this profile for the roll - default option
				profile.ReadFromFile()
				loadout.PopulateExcludeListsFromProfile(profile)
				if *rollMulti == "" {
					// If this command wasn't specified, replace it with the default
					*rollMulti = profile.DefaultRoll
				}
			}
		}

		// Parse exclude lists
		if len(xOrbital) > 0 {
			orbitalSlot.ParseExcludeFromSlice(xOrbital)

			if modifyProfile {
				profile.XOrbital = orbitalSlot.GetExcludeStringList()
			}
		}

		if len(xEagle) > 0 {
			eagleSlot.ParseExcludeFromSlice(xEagle)

			if modifyProfile {
				profile.XEagle = eagleSlot.GetExcludeStringList()
			}
		}

		if len(xWeapon) > 0 {
			weaponSlot.ParseExcludeFromSlice(xWeapon)

			if modifyProfile {
				profile.XWeapon = weaponSlot.GetExcludeStringList()
			}
		}

		if len(xUtil) > 0 {
			utilitySlot.ParseExcludeFromSlice(xUtil)

			if modifyProfile {
				profile.XUtil = utilitySlot.GetExcludeStringList()
			}
		}

		if len(xPrimary) > 0 {
			primarySlot.ParseExcludeFromSlice(xPrimary)

			if modifyProfile {
				profile.XPrimary = primarySlot.GetExcludeStringList()
			}
		}

		if len(xSecondary) > 0 {
			secondarySlot.ParseExcludeFromSlice(xSecondary)

			if modifyProfile {
				profile.XSecondary = secondarySlot.GetExcludeStringList()
			}
		}

		if len(xGrenade) > 0 {
			grenadeSlot.ParseExcludeFromSlice(xGrenade)

			if modifyProfile {
				profile.XGrenade = grenadeSlot.GetExcludeStringList()
			}
		}

		if *slotInfo != "" {
			// Ok technically we haven't rolled anything, but assume the user just wants info
			rolled = true
			loadout.ParseAndPrintInfo(slotInfo)
		}

		if *rollMulti != "" {
			rolled = true
			loadout.RollMultipleSlots(rollMulti)

			if modifyProfile {
				profile.DefaultRoll = *rollMulti
			}
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
