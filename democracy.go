package main

import (
	"fmt"
	"math/rand"
)

var inspiration = []string{
	"Though we are the most powerful fighting force the universe has ever seen, we are naught but humble servants before Managed Democracy",

	"Remain ever-vigilant for deceit and treachery. They can take root in the smallest of cracks",

	"Tyranny is a cancer, and Managed Democracy is the cure. And you helldiver, shall administer the antidote",

	"There is but one assured path to peace... And that path is war",

	"The enemy may anticipate but one outcome: total annillation",

	"Ours are the hands that will tear asunder the binds of tyranny",

	"Here at war, you may cast your vote many times: Once on Election Day, and another with every bullet berried in an enemy combatant",

	"Dissidents sow their lives in the smallest of fractures, Helldiver. Stay vigilant",

	"The people of Super Earth look to us for hope. And we will deliver it- one enemy corpse at a time",

	"Innocents perish with every wasted second. We must act",
}

func InspireDemocracy() {
	fmt.Println(inspiration[rand.Intn(len(inspiration))])
}
