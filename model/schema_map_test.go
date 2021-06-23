package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseMapSchema_Primitive(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:com:map;1"
	}))

	s := parseMapSchema(map[string]interface{}{
		"@id":         "dtmi:com:map;1",
		"@type":       "Map",
		"displayName": "Test map",
		"description": "Test map",
		"comment":     "Test map",
		"mapValue": map[string]interface{}{
			"displayName": "Test map value",
			"description": "Test map value",
			"comment":     "Test map value",
			"schema":      "string",
		},
	}, m)

	a.Equal("dtmi:com:map;1", s.ID())
	a.Equal("Test map", s.DisplayName())
	a.Equal("Test map", s.Description())
	a.Equal("Test map", s.Comment())
	a.Equal("Test map value", s.MapValue().DisplayName())
	a.Equal("Test map value", s.MapValue().Description())
	a.Equal("Test map value", s.MapValue().Comment())
	a.Same(StringSchema, s.MapValue().Schema())
	m.AssertExpectations(t)
}

func TestParseMapSchema_Array(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:test:array;1" || e.ID() == "dtmi:com:map;1"
	}))

	s := parseMapSchema(map[string]interface{}{
		"@id":         "dtmi:com:map;1",
		"@type":       "Map",
		"displayName": "Test map",
		"description": "Test map",
		"comment":     "Test map",
		"mapValue": map[string]interface{}{
			"displayName": "Test map value",
			"description": "Test map value",
			"comment":     "Test map value",
			"schema": map[string]interface{}{
				"@id":   "dtmi:test:array;1",
				"@type": "Array",
			},
		},
	}, m)

	a.IsType(&MapSchema{}, s)
	a.Equal("dtmi:com:map;1", s.ID())
	a.Equal("Test map", s.DisplayName())
	a.Equal("Test map", s.Description())
	a.Equal("Test map", s.Comment())
	a.Equal("Test map value", s.MapValue().DisplayName())
	a.Equal("Test map value", s.MapValue().Description())
	a.Equal("Test map value", s.MapValue().Comment())
	a.IsType(&ArraySchema{}, s.MapValue().Schema())
	m.AssertExpectations(t)
}

func TestParseMapSchema_Map(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:test:map;1" || e.ID() == "dtmi:com:map;1"
	}))

	s := parseMapSchema(map[string]interface{}{
		"@id":         "dtmi:com:map;1",
		"@type":       "Map",
		"displayName": "Test map",
		"description": "Test map",
		"comment":     "Test map",
		"mapValue": map[string]interface{}{
			"displayName": "Test map value",
			"description": "Test map value",
			"comment":     "Test map value",
			"schema": map[string]interface{}{
				"@id":   "dtmi:test:map;1",
				"@type": "Map",
			},
		},
	}, m)

	a.IsType(&MapSchema{}, s)
	a.Equal("dtmi:com:map;1", s.ID())
	a.Equal("Test map", s.DisplayName())
	a.Equal("Test map", s.Description())
	a.Equal("Test map", s.Comment())
	a.Equal("Test map value", s.MapValue().DisplayName())
	a.Equal("Test map value", s.MapValue().Description())
	a.Equal("Test map value", s.MapValue().Comment())
	a.IsType(&MapSchema{}, s.MapValue().Schema())
	m.AssertExpectations(t)
}

func TestParseMapSchema_Object(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:test:obj;1" || e.ID() == "dtmi:com:map;1"
	}))

	s := parseMapSchema(map[string]interface{}{
		"@id":         "dtmi:com:map;1",
		"@type":       "Map",
		"displayName": "Test map",
		"description": "Test map",
		"comment":     "Test map",
		"mapValue": map[string]interface{}{
			"displayName": "Test map value",
			"description": "Test map value",
			"comment":     "Test map value",
			"schema": map[string]interface{}{
				"@id":   "dtmi:test:obj;1",
				"@type": "Object",
			},
		},
	}, m)

	a.IsType(&MapSchema{}, s)
	a.Equal("dtmi:com:map;1", s.ID())
	a.Equal("Test map", s.DisplayName())
	a.Equal("Test map", s.Description())
	a.Equal("Test map", s.Comment())
	a.Equal("Test map value", s.MapValue().DisplayName())
	a.Equal("Test map value", s.MapValue().Description())
	a.Equal("Test map value", s.MapValue().Comment())
	a.IsType(&ObjectSchema{}, s.MapValue().Schema())
	m.AssertExpectations(t)
}

func TestParseMapSchema_Reference(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	r := &Schema{
		Entity: Entity{
			id: "dtmi:schema:ref;1",
		},
	}

	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:schema:ref;1" || e.ID() == "dtmi:com:map;1"
	}))

	m.On("Get", "dtmi:schema:ref;1").Return(r, true)

	s := parseMapSchema(map[string]interface{}{
		"@id":   "dtmi:com:map;1",
		"@type": "Map",
		"mapValue": map[string]interface{}{
			"schema": "dtmi:schema:ref;1",
		},
	}, m)

	a.IsType(&SchemaReference{}, s.MapValue().schema)
	a.Same(r, s.MapValue().Schema())
	a.Equal(r.ID(), s.MapValue().Schema().ID())
	m.AssertNumberOfCalls(t, "Get", 1)
}
