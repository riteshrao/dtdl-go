package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseRelationship_Full(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:rel;1"
	}))

	r := parseRelationship(map[string]interface{}{
		"@id": "dtmi:rel;1",
		"@type": "Relationship",
		"name": "rel",
		"displayName": "Parent",
		"description": "Parent",
		"comment": "Parent",
		"maxMultiplicity": 100,
		"minMultiplicity": 1,
		"target": "dtmi:target;1",
		"writable": true,
	}, m)

	a.Equal("dtmi:rel;1", r.ID())
	a.Equal("rel", r.Name())
	a.Equal("Parent", r.DisplayName())
	a.Equal("Parent", r.Description())
	a.Equal("Parent", r.Comment())
	a.Equal(100, r.MaxMultiplicity())
	a.Equal(1, r.MinMultiplicity())
	a.Equal("dtmi:target;1", r.Target())
	a.Equal(true, r.Writable())
}