package character

import (
	"github.com/jinzhu/gorm"
)

type Rarity string

const (
	RarityN  Rarity = "N"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityUR Rarity = "UR"
)

type Character struct {
	gorm.Model
	ID     int
	Name   string
	Rarity Rarity
}
