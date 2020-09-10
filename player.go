package daft

import "github.com/pkg/errors"

// Player of the game
type Player struct {
	id         int
	level      int
	name       string
	attributes []Attribute
	skills     []Skill
}

// Attribute of the player
type Attribute struct {
	id    int
	level int
	name  string
}

// Skill of the player
type Skill struct {
	id    int
	level int
	name  string
}

// NewSkill creates a skill by the given attribute ID, skill ID and level number
// this function will wrap all number infos into ONE ID
//
func NewSkill(attr int, skill int, level int, desc string) (*Skill, error) {
	if attr > 0xff {
		return nil, errors.Errorf("attribute ID can not be greater than 255: %d", attr)
	}

	if skill > 0xff {
		return nil, errors.Errorf("skill ID can not be greater than 255: %d", skill)
	}

	if level > 0xff {
		return nil, errors.Errorf("level num can not be greater than 255: %d", level)
	}

	id := (attr << 16) | (skill << 8) | level

	s := &Skill{id: id}
	return s, nil
}

// Level of the skill
func (s *Skill) Level() int {
	return s.id & 0xFF
}

// ID of the skill
func (s *Skill) ID() int {
	return s.id >> 8 & 0xFF
}

// Attribute 's ID of the skill'
func (s *Skill) Attribute() int {
	return s.id >> 16 & 0xFF
}
