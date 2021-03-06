package gacha

import (
	"math/rand"
	"strconv"

	"CA_Go/database"
	"CA_Go/model/character"
	"CA_Go/model/user"
)

type GachaDrawRequest struct {
	Times int `json:"times"`
}

type GachaResult struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

type GachaDrawResponse struct {
	Results []GachaResult `json:"results"`
}

func Draws(n int, xToken string) []GachaResult {
	characters := make([]*character.Character, n)
	for i := 0; i < n; i++ {
		characters[i] = draw()
	}

	var results []GachaResult

	u := user.NewUser()
	u.FindByToken(xToken)

	for _, character := range characters {
		result := GachaResult{CharacterID: strconv.FormatInt(character.ID, 10), Name: character.Name}

		//TODO:リファクタリングで別のパッケージに移行する
		//TODO:同じキャラクターを複数体持てるようにする
		database.DB.Model(u).Association("Characters").Append(character)

		results = append(results, result)
	}

	return results
}

func draw() *character.Character {
	c := character.NewCharacter()

	// NOTE:ランダム性を出すために100までの数字を出す
	num := rand.Intn(100)

	// TODO:数字がマジックナンバーだから持たせて定数
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
