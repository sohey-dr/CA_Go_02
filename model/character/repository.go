package character

import (
	"CA_Go/database"
)

func NewCharacter() *Character {
	return &Character{}
}

func (c *Character) FindByRarityOrderByRand(rarity Rarity) {
	database.DB.Where("rarity = ?", rarity).Order("RAND()").First(&c)
}

func (c *Character) FindById(id string) {
	database.DB.Where("id = ?", id).First(&c)
}
