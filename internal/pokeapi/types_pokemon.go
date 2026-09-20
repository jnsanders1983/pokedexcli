package pokeapi

type Pokemon struct {
	ID                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []AbilitySlot      `json:"abilities"`
	Forms                  []NamedAPIResource `json:"forms"`
	GameIndices            []GameIndex        `json:"game_indices"`
	HeldItems              []HeldItem         `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []MoveSlot         `json:"moves"`
	Species                NamedAPIResource   `json:"species"`
	Sprites                Sprites            `json:"sprites"`
	Stats                  []Stat             `json:"stats"`
	Types                  []Type             `json:"types"`
}

// NamedAPIResource is a reusable component for name/url pairs.
type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stat struct {
	BaseStat int              `json:"base_stat"`
	Stat     NamedAPIResource `json:"stat"`
}

type Type struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type AbilitySlot struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type HeldItem struct {
	Item           NamedAPIResource `json:"item"`
	VersionDetails []VersionDetail  `json:"version_details"`
}

type VersionDetail struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

type MoveSlot struct {
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	Order           int              `json:"order"`
}

type Sprites struct {
	BackDefault      string       `json:"back_default"`
	BackFemale       *string      `json:"back_female"` // Using pointers for fields that can be null
	BackShiny        string       `json:"back_shiny"`
	BackShinyFemale  *string      `json:"back_shiny_female"`
	FrontDefault     string       `json:"front_default"`
	FrontFemale      *string      `json:"front_female"`
	FrontShiny       string       `json:"front_shiny"`
	FrontShinyFemale *string      `json:"front_shiny_female"`
	Other            OtherSprites `json:"other"`
	// Note: Version-specific sprites omitted for brevity due to dynamic generation keys,
	// but can be unmarshaled using map[string]interface{} if needed.
}

type OtherSprites struct {
	DreamWorld      SpriteVariant `json:"dream_world"`
	Home            SpriteVariant `json:"home"`
	OfficialArtwork SpriteVariant `json:"official-artwork"`
	Showdown        SpriteVariant `json:"showdown"`
}

type SpriteVariant struct {
	BackDefault      string  `json:"back_default,omitempty"`
	BackFemale       *string `json:"back_female,omitempty"`
	BackShiny        string  `json:"back_shiny,omitempty"`
	BackShinyFemale  *string `json:"back_shiny_female,omitempty"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female,omitempty"`
	FrontShiny       string  `json:"front_shiny,omitempty"`
	FrontShinyFemale *string `json:"front_shiny_female,omitempty"`
}
