package gacha

import (
	"math/rand"

	"CA_Go/model/character"
)

type GachaResult struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

type GachaDrawResponse struct {
	Results []GachaResult `json:"results"`
}

func Draws(n int) []*character.Character {
	results := make([]*character.Character, n)
	for i := 0; i < n; i++ {
		results[i] = draw()
	}

	return results
}

func draw() *character.Character {
	c := character.NewCharacter()

	// NOTE:ランダム性を出すために100までの数字を出す
	num := rand.Intn(100)

	switch {
	case num < 80:
		c.FindByRarityOrderByRand(character.RarityN)
		return c
	case num < 95:
		c.FindByRarityOrderByRand(character.RarityR)
		return c
	case num < 99:
		c.FindByRarityOrderByRand(character.RaritySR)
		return c
	default:
		c.FindByRarityOrderByRand(character.RarityUR)
		return c
	}
}
