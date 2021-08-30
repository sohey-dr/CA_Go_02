package character

type Rarity string

const (
	RarityN  Rarity = "N"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityUR Rarity = "UR"
)

type Character struct {
	gorm.Model
	name   string
	Rarity Rarity
}
