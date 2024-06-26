package main

import "fmt"

type Loadout struct {
	Orbital   Slot
	Eagle     Slot
	Weapon    Slot
	Utility   Slot
	Primary   Slot
	Secondary Slot
	Grenade   Slot
}

func NewLoadout() *Loadout {
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

	return &loadout
}

func (l *Loadout) ChooseAll() {
	fmt.Println("Orbital:", l.Orbital.Choose())
	fmt.Println("Eagle:", l.Eagle.Choose())
	fmt.Println("Weapon:", l.Weapon.Choose())
	fmt.Println("Utility:", l.Utility.Choose())
	fmt.Println("Primary:", l.Primary.Choose())
	fmt.Println("Secondary:", l.Secondary.Choose())
	fmt.Println("Grenade:", l.Grenade.Choose())
}

func (l *Loadout) ParseAndPrintInfo(args *string) {
	for _, arg := range *args {
		switch arg {
		case 'o':
			l.Orbital.PrintWithNumbers()
		case 'e':
			l.Eagle.PrintWithNumbers()
		case 'w':
			l.Weapon.PrintWithNumbers()
		case 'u':
			l.Utility.PrintWithNumbers()
		case 'p':
			l.Primary.PrintWithNumbers()
		case 's':
			l.Secondary.PrintWithNumbers()
		case 'g':
			l.Grenade.PrintWithNumbers()
		}
	}
}

// Once an option is chosen, add it to the exclude list
// This allows flags such as `-r="-ooew"` to not produce duplicates
func (l *Loadout) RollMultipleSlots(args *string) {
	for _, arg := range *args {
		switch arg {
		case 'o':
			option := l.Orbital.Choose()
			fmt.Println("Orbital:", option)
			l.Orbital.exclude[option] = true
		case 'e':
			option := l.Eagle.Choose()
			fmt.Println("Eagle:", option)
			l.Eagle.exclude[option] = true
		case 'w':
			option := l.Weapon.Choose()
			fmt.Println("Weapon:", option)
			l.Weapon.exclude[option] = true
		case 'u':
			option := l.Utility.Choose()
			fmt.Println("Utility:", option)
			l.Utility.exclude[option] = true
		case 'p':
			option := l.Primary.Choose()
			fmt.Println("Primary:", option)
			l.Primary.exclude[option] = true
		case 's':
			option := l.Secondary.Choose()
			fmt.Println("Secondary:", option)
			l.Secondary.exclude[option] = true
		case 'g':
			option := l.Grenade.Choose()
			fmt.Println("Grenade:", option)
			l.Grenade.exclude[option] = true
		default:
			fmt.Println("Invalid slot option:", string(arg))
		}
	}
}

func (l *Loadout) PopulateExcludeListsFromProfile(prof *Profile) {
	l.Orbital.ParseExcludeFromSlice(prof.XOrbital)
	l.Eagle.ParseExcludeFromSlice(prof.XEagle)
	l.Utility.ParseExcludeFromSlice(prof.XUtil)
	l.Weapon.ParseExcludeFromSlice(prof.XWeapon)
	l.Primary.ParseExcludeFromSlice(prof.XPrimary)
	l.Secondary.ParseExcludeFromSlice(prof.XSecondary)
	l.Grenade.ParseExcludeFromSlice(prof.XGrenade)
}
