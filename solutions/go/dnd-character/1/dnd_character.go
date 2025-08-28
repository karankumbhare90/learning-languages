package dndcharacter


import (
	"math"
    "math/rand"
	"sort"
	"time"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score

func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}	

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rand.Seed(time.Now().UnixNano())

	rolls := []int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}

	sort.Ints(rolls)              // sort rolls ascending
	return rolls[1] + rolls[2] + rolls[3] // drop lowest (rolls[0])
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	c.Hitpoints = 10 + Modifier(c.Constitution)
	return c
}