# Helldiver-Randomizer

A CLI loadout randomizer for Helldivers 2!

# Requirements To Build
Go: https://go.dev/

# Usage
## Building
Once go is installed, build the package using `go install`

## Running 
To run the basic program, use `helldiver-randomizer` with no arguments. This should print a randomized loadout for you that looks something like this:

```
> helldiver-randomizer
Orbital: smoke
Eagle: napalm
Weapon: eat
Utility: autocannon sentry
Primary: spray and pray
Secondary: senator
Grenade: impact
```


## Options
To list all available options, run `helldiver-randomizer -h` 

# Detailed Usage
## No Arguments
With no arguments, the program will print an entire random loadout as seen above.

## Rolling Specific Slots 
There are two ways to roll specific slots.

You can specify slots to roll with the flags (just type the first letter):  
-o(rbital)  
-e(agle)  
-w(eapon)  
-u(til)  
-p(rimary)  
-s(econdary)  
-g(renade)  

For example, `helldiver-randomizer -o -s` will output someting like:
```
> helldiver-randomizer -o -s
Orbital: 380
Secondary: verdict
```

You can also specify a group of rolls with the `-r` flag, using the same letters as above.

For example, `-r=psg` will roll the primary, secondary, and grenade slots:
```
> helldiver-randomizer -r=psg
Primary: scythe
Secondary: peacemaker
Grenade: incendiary
```

The `-r` flag can be used to specify custom loadout sets. For example:  
- Roll only stratagems with `-r=oweu`
- Roll only a loadout with `-r=psg`
- Roll strategems with a duplicate slot (such as two eagles) with `-oeew` (this will automatically prevent duplicates)

## Listing Options for a Slot
The `-i` flag can be used to list all the options for one or more slots. 
For example `-i=oe` will print all the options for the orbital and eagle slots.

```
> helldiver-randomizer -i=oe`
1: gatling
2: airburst
3: 120
4: 380
5: walking
6: railcannon
7: ems
8: gas
9: smoke
10: precision
11: laser
1: strafing
2: airstrike
3: cluster
4: napalm
5: smokes
6: rockets
7: 500kg
```

The numbers returned are important - they indicate the numbers to use when you want to exclude certain entries from a slot roll.

## Excluding Options for a Slot

Excluding certain options is useful if you want to exclude weapons you haven't unlocked, or you just dislike some item so much you don't event want to use it randomly. 

Using the numbers provided by the `-i` option for a slot, you can specify a comma-separated, quoted list of integers to exclude. Each slot has an exclude command consisting of the letter flag prefaced with `x`, i.e. `-xo` to exclude orbitals, `-xp` to exlude primaries, etc.

For example, to set up an exclude list for orbitals that prevents rolling the gatling barrage, orbital smoke, and orbital precision strike, you can use the switch
 `-xo="1,9,10"`

Running the program with only exclude options will roll every slot while respecting your exclude options:
```
> helldiver-randomizer -xo="1,9,10" -xe="1,3,6"
Excluding: gatling
Excluding: smoke
Excluding: precision
Excluding: strafing
Excluding: cluster
Excluding: rockets
Orbital: walking
Eagle: napalm
Weapon: spear
Utility: mortar sentry
Primary: diligence counter sniper
Secondary: dagger
Grenade: thermite
```
Exclude lists can also be combined with individual slot selection. This command will roll the orbital, primary, secondary, and grenade slots while excluding severa orbitals and primaries: 

```
> helldiver-randomizer -xo="1,9,10" -xp="4,23,15" -r=opsg
Excluding: gatling
Excluding: smoke
Excluding: precision
Excluding: tenderizer
Excluding: purifier
Excluding: exploding crossbow
Orbital: railcannon
Primary: adjudicator
Secondary: verdict
Grenade: incendiary
```

### Fuzzy Matching 
The `-x` commands can also do fuzzy string matching using the name of the item to exclude. The command takes an argument of comma-separated strings and will attempt to match them against the list for the slot. 

```
> helldiver-randomizer -xu="guard dog,mort,shield gen"
Excluding guard dog mg on arg guard dog
Excluding guard dog laser on arg guard dog
Excluding mortar sentry on arg mort
Excluding shield generator relay on arg shield gen
Orbital: gas
Eagle: napalm
Weapon: mg
Utility: exosuit
Primary: crossbow
Secondary: dagger
Grenade: smoke
```

Notice that one argument may match multiple items. 

```
> helldiver-randomizer -xp="liberator" -p
Excluding liberator on arg liberator
Excluding liberator pen on arg liberator
Excluding liberator concussive on arg liberator
Primary: scorcher
```

This can be useful for exluding multiple similar items. This can be worked around by using more specific strings:

```
helldiver-randomizer -xp="conc" -p
Excluding liberator concussive on arg conc
Primary: blitzer
```
