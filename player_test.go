package daft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSkill(t *testing.T) {
	s, err := NewSkill(101, 102, 253, "Hello")

	assert.Nil(t, err)
	assert.Equal(t, 101, s.Attribute())
	assert.Equal(t, 102, s.ID())
	assert.Equal(t, 253, s.Level())

	s, err = NewSkill(101, 102, 256, "Hello")
	assert.NotNil(t, err)

	s, err = NewSkill(101, 300, 2, "Hello")
	assert.NotNil(t, err)

	s, err = NewSkill(301, 2, 2, "Hello")
	assert.NotNil(t, err)
}
