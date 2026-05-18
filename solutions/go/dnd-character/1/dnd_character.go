package dndcharacter

import (
    "math"
    "math/rand"
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
    val := (float64(score) - 10.0) / 2.0
    if val < 0.0 {
        val = math.Round(val)
    }
	return int(val)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
    rolls := [4]int{
        rand.Intn(6) + 1, 
        rand.Intn(6) + 1,
        rand.Intn(6) + 1,
        rand.Intn(6) + 1,
    }
    
    min := rolls[0]
    for _, roll := range rolls[1:] {
        if roll < min {
            min = roll
        }
    }
    
    sum := 0
    for _, roll := range rolls {
        sum += roll
    }
    
    return sum - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
    constitution := Ability() 
	return Character{
        Strength:     Ability(),
        Dexterity:    Ability(),
        Constitution: constitution,
        Intelligence: Ability(),
        Wisdom:       Ability(),
        Charisma:     Ability(),
        Hitpoints:    10 + Modifier(constitution),
    }
}
