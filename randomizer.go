package main

import (
	"flag"
	"fmt"
)

type Randomizer struct {
	// Core loadout struct
	loadout Loadout

	// Loaded profile, if one is present
	profile Profile

	// Runtime options set based on profile parsing
	saveProfile   bool
	modifyProfile bool

	// Flags parsed by flag package
	orbital   *bool
	eagle     *bool
	weapon    *bool
	util      *bool
	primary   *bool
	secondary *bool
	grenade   *bool

	slotInfo  *string
	rollMulti *string

	xOrbital   ExcludeFlag
	xEagle     ExcludeFlag
	xWeapon    ExcludeFlag
	xUtil      ExcludeFlag
	xPrimary   ExcludeFlag
	xSecondary ExcludeFlag
	xGrenade   ExcludeFlag

	loadProfile *string
	profSwitch  ProfileSwitch
}

func NewRandomizer() *Randomizer {
	r := &Randomizer{}

	r.loadout = *NewLoadout()
	r.profile = *NewProfile()

	r.orbital = flag.Bool("o", false, "Roll a new orbital strat")
	r.eagle = flag.Bool("e", false, "Roll a new eagle strat")
	r.weapon = flag.Bool("w", false, "Roll a new support weapon")
	r.util = flag.Bool("u", false, "Roll a new utility strat")
	r.primary = flag.Bool("p", false, "Roll a new primary")
	r.secondary = flag.Bool("s", false, "Roll a new secondary")
	r.grenade = flag.Bool("g", false, "Roll a new grenade")

	r.slotInfo = flag.String("i", "", "Print options for a specified slot")
	r.rollMulti = flag.String("r", "", "Roll multiple slots")

	flag.Var(&r.xOrbital, "xo", "Exclude these orbitals")
	flag.Var(&r.xEagle, "xe", "Exclude these eagles")
	flag.Var(&r.xWeapon, "xw", "Exclude these support weapons")
	flag.Var(&r.xUtil, "xu", "Exclude these utility items")
	flag.Var(&r.xPrimary, "xp", "Exclude these primary weapons")
	flag.Var(&r.xSecondary, "xs", "Exclude these secondary weapons")
	flag.Var(&r.xGrenade, "xg", "Exclude these grenades")

	r.loadProfile = flag.String("profile", "", "Options for profile: u(se), c(reate), e(dit), i(nfo)")
	flag.Var(&r.profSwitch, "pm", "Set the usage mode for the specified profile")

	r.saveProfile = false

	return r
}

func (r *Randomizer) InitProfile() (runRequired bool) {
	runRequired = true

	if *r.loadProfile != "" {
		r.profile.SetName(*r.loadProfile)

		// Default to "use" option if no switch is provided
		if r.profSwitch == "" {
			r.profSwitch = "u"
		}

		switch r.profSwitch {
		case "e": // Edit profile - overwrite specified settings based on other flags
			r.modifyProfile = true
			if !r.profile.ReadFromFile() {
				fmt.Println("Profile not found - please use create instead of edit for new profiles")
				runRequired = false
			}
			r.saveProfile = true

		case "d": // Delete profile (if present)
			if r.profile.Delete() {
				fmt.Println("Profile deleted:", r.profile.Name)
			}
			runRequired = false

		case "c": // Create profile based on specified settings
			r.modifyProfile = true
			r.saveProfile = true
			r.profile.DefaultRoll = "oeuwpsg" // Default to rolling everything unless overwritten below

		case "i": // Print info on the profile, but take no action based on it
			if !r.profile.ReadFromFile() {
				fmt.Println("Profile not found:", r.profile.Name)
			} else {
				r.profile.Describe()
			}
			runRequired = false

		case "u": // Use this profile for the roll - default option
			if !r.profile.ReadFromFile() {
				fmt.Println("Profile not found:", r.profile.Name)
				runRequired = false
			} else {
				r.loadout.PopulateExcludeListsFromProfile(&r.profile)
				if *r.rollMulti == "" {
					// If this command wasn't specified, replace it with the default
					*r.rollMulti = r.profile.DefaultRoll
				}
			}
		}
	}

	return
}

func (r *Randomizer) InitExcludeLists() {
	if len(r.xOrbital) > 0 {
		r.loadout.Orbital.ParseExcludeFromSlice(r.xOrbital)

		if r.modifyProfile {
			r.profile.XOrbital = r.loadout.Orbital.GetExcludeStringList()
		}
	}

	if len(r.xEagle) > 0 {
		r.loadout.Eagle.ParseExcludeFromSlice(r.xEagle)

		if r.modifyProfile {
			r.profile.XEagle = r.loadout.Eagle.GetExcludeStringList()
		}
	}

	if len(r.xWeapon) > 0 {
		r.loadout.Weapon.ParseExcludeFromSlice(r.xWeapon)

		if r.modifyProfile {
			r.profile.XWeapon = r.loadout.Weapon.GetExcludeStringList()
		}
	}

	if len(r.xUtil) > 0 {
		r.loadout.Utility.ParseExcludeFromSlice(r.xUtil)

		if r.modifyProfile {
			r.profile.XUtil = r.loadout.Utility.GetExcludeStringList()
		}
	}

	if len(r.xPrimary) > 0 {
		r.loadout.Primary.ParseExcludeFromSlice(r.xPrimary)

		if r.modifyProfile {
			r.profile.XPrimary = r.loadout.Primary.GetExcludeStringList()
		}
	}

	if len(r.xSecondary) > 0 {
		r.loadout.Secondary.ParseExcludeFromSlice(r.xSecondary)

		if r.modifyProfile {
			r.profile.XSecondary = r.loadout.Secondary.GetExcludeStringList()
		}
	}

	if len(r.xGrenade) > 0 {
		r.loadout.Grenade.ParseExcludeFromSlice(r.xGrenade)

		if r.modifyProfile {
			r.profile.XGrenade = r.loadout.Grenade.GetExcludeStringList()
		}
	}
}

func (r *Randomizer) Run() {

	if *r.slotInfo != "" {
		// User just wants info, we can skip the rest of the run
		r.loadout.ParseAndPrintInfo(r.slotInfo)
		return
	}

	if *r.rollMulti != "" {
		if r.modifyProfile {
			// Modify profile only, no need to roll
			r.profile.DefaultRoll = *r.rollMulti
			fmt.Println("Profile default roll updated to", *r.rollMulti)
		} else {
			r.loadout.RollMultipleSlots(r.rollMulti)
		}
		return
	}

	// Check individual roll flags
	rolled := false
	if *r.orbital {
		rolled = true
		fmt.Println("Orbital:", r.loadout.Orbital.Choose())
	}

	if *r.eagle {
		rolled = true
		fmt.Println("Eagle:", r.loadout.Eagle.Choose())
	}

	if *r.weapon {
		rolled = true
		fmt.Println("Weapon:", r.loadout.Weapon.Choose())
	}

	if *r.util {
		rolled = true
		fmt.Println("Utility:", r.loadout.Utility.Choose())
	}

	if *r.primary {
		rolled = true
		fmt.Println("Primary:", r.loadout.Primary.Choose())
	}

	if *r.secondary {
		rolled = true
		fmt.Println("Secondary:", r.loadout.Secondary.Choose())
	}

	if *r.grenade {
		rolled = true
		fmt.Println("Grenade:", r.loadout.Grenade.Choose())
	}

	// If we get to this point without having rolled anything, roll a full loadout
	if !rolled {
		r.loadout.ChooseAll()
	}
}

func (r *Randomizer) Cleanup() {
	if r.saveProfile {
		r.profile.WriteToFile()
	}

	InspireDemocracy()
}
