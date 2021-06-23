package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseEnumSchema(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:contoso:enum;1"
	}))

	s := parseEnumSchema(map[string]interface{}{
		"@id":         "dtmi:contoso:enum;1",
		"@type":       "Enum",
		"displayName": "Test enum",
		"description": "Test enum",
		"comment":     "Test enum",
		"valueSchema": "integer",
		"enumValues": []interface{}{
			map[string]interface{}{
				"displayName": "open",
				"description": "open",
				"comment":     "open",
				"enumValue":   1,
			},
			map[string]interface{}{
				"displayName": "closed",
				"description": "closed",
				"comment":     "closed",
				"enumValue":   0,
			},
		},
	}, m)

	a.IsType(&EnumSchema{}, s)
	a.Equal("dtmi:contoso:enum;1", s.ID())
	a.Equal("Test enum", s.DisplayName())
	a.Equal("Test enum", s.Description())
	a.Equal("Test enum", s.Comment())

	a.Same(IntegerSchema, s.valueSchema)
	a.Equal(2, len(s.enumValues))
	a.Equal("open", s.enumValues[0].DisplayName())
	a.Equal("open", s.enumValues[0].Description())
	a.Equal("open", s.enumValues[0].Comment())
	a.Equal("closed", s.enumValues[1].DisplayName())
	a.Equal("closed", s.enumValues[1].Description())
	a.Equal("closed", s.enumValues[1].Comment())

	val1, _ := s.enumValues[0].IntegerValue()
	val2, _ := s.enumValues[1].IntegerValue()
	a.Equal(1, val1)
	a.Equal(0, val2)
	m.AssertExpectations(t)
}