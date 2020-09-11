package daft

import "github.com/pelletier/go-toml"

type Config struct {
	Attributes []Info
	Players    []Player
}

type Info struct {
	ID          int
	Name        string
	Description string
}

// Player of the game
type Player struct {
	ID         int
	Level      int
	Name       string
	Attributes []Attribute
	Skills     []Skill
}

// Attribute of the player
type Attribute struct {
	ID    int
	Level int
}

// Skill of the player
type Skill struct {
	ID    int
	Level int
}

// NewNewPlayerFromToml create a player from toml input
func NewConfigFromToml(src []byte) (*Config, error) {
	c := &Config{}
	err := toml.Unmarshal(src, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
