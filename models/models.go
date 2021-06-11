package models

type Results struct {
	Pokemon []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemons struct {
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
}

type Sprites struct {
	Back_default  string `json:"back_default"`
	Front_default string `json:"front_default"`
}

type PokemonInfo struct {
	Base_experience int           `json:"base_experience"`
	Id              int           `json:"id"`
	Name            string        `json:"name"`
	Order           int           `json:"order"`
	Height          int           `json:"height"`
	Weight          int           `json:"weight"`
	Sprites         Sprites       `json:"sprites"`
	Species         Species       `json:"species"`
	Abilities       []Abilities   `json:"abilities"`
	Forms           []Forms       `json:"forms"`
	GameIndices     []GameIndices `json:"game_indices"`
	Moves           []Moves       `json:"moves"`
	Types           []Types       `json:"types"`
	Stats           []Stats       `json:"stats"`
}

type Abilities struct {
	Ability Ability `json:"ability"`
}

type Ability struct {
	Name string `json:"name"`
}

type Forms struct {
	Name string `json:"name"`
}

type GameIndices struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

type Version struct {
	Name string `json:"name"`
}

type Moves struct {
	Move                Move                  `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}

type Move struct {
	Name string `json:"name"`
}

type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	VersionGroup    VersionGroup    `json:"version_group"`
}

type MoveLearnMethod struct {
	Name string `json:"name"`
}

type VersionGroup struct {
	Name string `json:"name"`
}

type Types struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}

type Species struct {
	Name string `json:"name"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}
