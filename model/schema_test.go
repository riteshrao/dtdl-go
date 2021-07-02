package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseSchema_Standard(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	a.Same(BooleanSchema, parseSchema("boolean", m), "Expected boolean schema reference")
	a.Same(DateSchema, parseSchema("date", m), "Expected date schema reference")
	a.Same(DateTimeSchema, parseSchema("dateTime", m), "Expected datetime schema reference")
	a.Same(DoubleSchema, parseSchema("double", m), "Expected double schema reference")
	a.Same(GeopointSchema, parseSchema("geopoint", m), "Expected geopoint schema reference")
	a.Same(IntegerSchema, parseSchema("integer", m), "Expected integer schema reference")
	a.Same(LongSchema, parseSchema("long", m), "Expected long schema reference")
	a.Same(StringSchema, parseSchema("string", m), "Expected string schema reference")
	a.Same(VectorSchema, parseSchema("vector", m), "Expected vector schema reference")
}

func TestParseSchema_Custom(t *testing.T) {
	a := assert.New(t)
	cases := map[string]SchemaType{
		"Array": &ArraySchema{},
		"Enum": &EnumSchema{},
		"Map": &MapSchema{},
		"Object": &ObjectSchema{},
	}

	
	for n, st := range cases {
		m := &mockTracker{}
		m.On("Add", mock.MatchedBy(func(e EntityType) bool {
			return e.ID() == "dtmi:schema;1"
		}))

		s := parseSchema(map[string]interface{} {
			"@id": "dtmi:schema;1",
			"@type": n,
		}, m)

		a.IsType(st, s)
		m.AssertExpectations(t)
	}
}