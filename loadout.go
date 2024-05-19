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

func (l *Loadout) RollMultipleSlots(args *string) {
	for _, arg := range *args {
		switch arg {
		case 'o':
			fmt.Println("Orbital:", l.Orbital.Choose())
		case 'e':
			fmt.Println("Eagle:", l.Eagle.Choose())
		case 'w':
			fmt.Println("Weapon:", l.Weapon.Choose())
		case 'u':
			fmt.Println("Utility:", l.Utility.Choose())
		case 'p':
			fmt.Println("Primary:", l.Primary.Choose())
		case 's':
			fmt.Println("Secondary:", l.Secondary.Choose())
		case 'g':
			fmt.Println("Grenade:", l.Grenade.Choose())
		}
	}
}
