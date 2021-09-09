package character

import (
	"time"
)

type Rarity string

const (
	RarityN  Rarity = "N"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityUR Rarity = "UR"
)

type Character struct {
	ID        int64
	Name      string
	Rarity    Rarity
	CreatedAt time.Time
	UpdatedAt time.Time
}
