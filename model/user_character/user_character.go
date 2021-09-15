package user_character

type UserCharacter struct {
	ID          string `json:"userCharacterID"`
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}
type CharacterListResponse struct {
	Characters []UserCharacter `json:"characters"`
}
