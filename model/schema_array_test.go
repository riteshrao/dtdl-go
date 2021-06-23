package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseArraySchema_Primitive(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:test:array;1"
	}))

	s := parseArraySchema(map[string]interface{}{
		"@id":           "dtmi:test:array;1",
		"@type":         "Array",
		"displayName":   "Test array",
		"description":   "Test array",
		"comment":       "Test array",
		"elementSchema": "string",
	}, m)

	a.IsType(&ArraySchema{}, s)
	a.Equal("dtmi:test:array;1", s.ID())
	a.Equal(ArraySchemaClass, s.Class())
	a.Equal("Test array", s.DisplayName())
	a.Equal("Test array", s.Description())
	a.Equal("Test array", s.Comment())
	a.Same(StringSchema, s.elementSchema)
	m.AssertExpectations(t)
}

func TestParseArraySchema_Reference(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	r := &Schema{
		Entity: Entity{
			id:    "dtmi:schema:ref;1",
			types: []string{"TypeA", "TypeB"},
		},
	}

	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:test:array;1"
	}))

	m.On("Get", "dtmi:schema:ref;1").Return(r, true)

	s := parseArraySchema(map[string]interface{}{
		"@id":           "dtmi:test:array;1",
		"@type":         "Array",
		"displayName":   "Test array",
		"description":   "Test array",
		"comment":       "Test array",
		"elementSchema": "dtmi:schema:ref;1",
	}, m)

	a.IsType(&SchemaReference{}, s.elementSchema)
	a.Same(r, s.ElementSchema())
	a.Equal(r.ID(), s.ElementSchema().ID())
	a.ElementsMatch(r.Types(), s.ElementSchema().Types())
	m.AssertNumberOfCalls(t, "Get", 1)
}

func TestParseArraySchema_Enum(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:test:enum;1" || e.ID() == "dtmi:test:array;1"
	}))

	s := parseArraySchema(map[string]interface{}{
		"@id": "dtmi:test:array;1",
		"@type": "Array",
		"elementSchema": map[string]interface{}{
			"@id": "dtmi:test:enum;1",
			"@type": "Enum",
		},
	}, m)

	a.IsType(&EnumSchema{}, s.ElementSchema());
	a.Equal("dtmi:test:enum;1", s.ElementSchema().ID())
	m.AssertExpectations(t)
}

func TestParseArraySchema_Map(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:test:map;1" || e.ID() == "dtmi:test:array;1"
	}))

	s := parseArraySchema(map[string]interface{}{
		"@id": "dtmi:test:array;1",
		"@type": "Array",
		"elementSchema": map[string]interface{}{
			"@id": "dtmi:test:map;1",
			"@type": "Map",
		},
	}, m)

	a.IsType(&MapSchema{}, s.ElementSchema());
	a.Equal("dtmi:test:map;1", s.ElementSchema().ID())
	m.AssertExpectations(t)
}

func TestParseArraySchema_Object(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:test:obj;1" || e.ID() == "dtmi:test:array;1"
	}))

	s := parseArraySchema(map[string]interface{}{
		"@id": "dtmi:test:array;1",
		"@type": "Array",
		"elementSchema": map[string]interface{}{
			"@id": "dtmi:test:obj;1",
			"@type": "Object",
		},
	}, m)

	a.IsType(&ObjectSchema{}, s.ElementSchema());
	a.Equal("dtmi:test:obj;1", s.ElementSchema().ID())
	m.AssertExpectations(t)
}