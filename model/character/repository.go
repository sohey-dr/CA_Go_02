package character

import (
	"CA_Go/database"
)

func NewCharacter() *Character {
	return &Character{}
}

func (c *Character) FindByRarity(rarity string) {
	database.DB.Where("rarity = ?", rarity).First(&c)
}
