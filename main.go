package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profile {
	switch name {
	case "Sasuke":
		return Profile{
			Name:   "Sasuke",
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	case "Goku":
		return Profile{
			Name:   "Goku",
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	case "Naruto":
		return Profile{
			Name:   "Naruto",
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	default:
		return Profile{}
	}
}

func PowerUp(profile Profile, multiplier int) Profile {
	profile.Health = profile.Health + (profile.Health * multiplier)
	profile.Power = profile.Power + (profile.Power * multiplier)
	profile.Exp = profile.Exp + (profile.Exp * multiplier)
	return profile
}

func main() {
	ProfileNaruto := MakeProfile("Naruto")
	fmt.Println("Name:", ProfileNaruto.Name)
	fmt.Println("Health:", ProfileNaruto.Health)
	fmt.Println("Power:", ProfileNaruto.Power)
	fmt.Println("Exp:", ProfileNaruto.Exp)

	fmt.Println("===HEROES POWERED UP===")

	powerUpNaruto := PowerUp(ProfileNaruto, 4)
	fmt.Println("Name:", powerUpNaruto.Name)
	fmt.Println("Health:", powerUpNaruto.Health)
	fmt.Println("Power:", powerUpNaruto.Power)
	fmt.Println("Exp:", powerUpNaruto.Exp)
}
