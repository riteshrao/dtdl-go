package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseComponent(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	i := &Interface{}

	m.On("add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:component;1"
	}))

	m.On("get", "dtmi:interface;1").Return(i, true)

	c := parseComponent(map[string]interface{}{
		"@id": "dtmi:component;1",
		"@type": "Component",
		"name": "component",
		"displayName": "Component",
		"description": "Component",
		"comment": "Component",
		"schema": "dtmi:interface;1",
	}, m)

	a.Equal("dtmi:component;1", c.ID())
	a.Equal("component", c.Name())
	a.Equal("Component", c.DisplayName())
	a.Equal("Component", c.Description())
	a.Equal("Component", c.Comment())
	a.Same(i, c.Schema())
}