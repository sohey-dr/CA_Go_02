package gacha

import (
	"math/rand"

	"CA_Go/model/character"
)

type GachaResult struct {
	CharacterID int    `json:"characterID"`
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
	num := rand.Intn(100)

	switch {
	case num < 80:
		return &character.Character{Rarity: character.RarityN, Name: "スライム"}
	case num < 95:
		return &character.Character{Rarity: character.RarityR, Name: "オーク"}
	case num < 99:
		return &character.Character{Rarity: character.RaritySR, Name: "ドラゴン"}
	default:
		return &character.Character{Rarity: character.RarityUR, Name: "イフリート"}
	}
}
