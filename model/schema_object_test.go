package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestParseObjectSchema_Primitive(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:com:object;1"
	}))

	o := parseObjectSchema(map[string]interface{}{
		"@id":         "dtmi:com:object;1",
		"@type":       "Object",
		"displayName": "Test object",
		"description": "Test object",
		"comment":     "Test object",
		"fields": []interface{}{
			map[string]interface{}{
				"name":        "field1",
				"displayName": "field1",
				"description": "field1",
				"schema":      "string",
			},
			map[string]interface{}{
				"name":        "field2",
				"displayName": "field2",
				"description": "field2",
				"schema":      "double",
			},
		},
	}, m)

	a.IsType(&ObjectSchema{}, o)
	a.Equal("dtmi:com:object;1", o.ID())
	a.Equal("Test object", o.DisplayName())
	a.Equal("Test object", o.Description())
	a.Equal("Test object", o.Comment())
	a.Equal(2, len(o.fields))

	a.Equal("field1", o.fields[0].Name())
	a.Equal("field1", o.fields[0].DisplayName())
	a.Equal("field1", o.fields[0].Description())
	a.Same(StringSchema, o.fields[0].Schema())

	a.Equal("field2", o.fields[1].Name())
	a.Equal("field2", o.fields[1].DisplayName())
	a.Equal("field2", o.fields[1].Description())
	a.Same(DoubleSchema, o.fields[1].Schema())

	m.AssertExpectations(t)
}

func TestParseObjectSchema_Reference(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	r := &Schema{
		Entity: Entity{
			id: "dtmi:schema:ref;1",
		},
	}
	
	m.On("Add", mock.MatchedBy(func(e EntityType) bool {
		return e.ID() == "dtmi:com:object;1" || e.ID() == "dtmi:schema:ref;"
	}))

	m.On("Get", "dtmi:schema:ref;1").Return(r, true)

	o := parseObjectSchema(map[string]interface{}{
		"@id":         "dtmi:com:object;1",
		"@type":       "Object",
		"displayName": "Test object",
		"description": "Test object",
		"comment":     "Test object",
		"fields": []interface{}{
			map[string]interface{}{
				"name":        "field1",
				"displayName": "field1",
				"description": "field1",
				"schema":      "dtmi:schema:ref;1",
			},
			map[string]interface{}{
				"name":        "field2",
				"displayName": "field2",
				"description": "field2",
				"schema":      "dtmi:schema:ref;1",
			},
		},
	}, m)

	a.IsType(&SchemaReference{}, o.Fields()[0].schema)
	a.Same(r, o.Fields()[0].Schema())
	a.Equal(r.ID(), o.Fields()[0].Schema().ID())

	a.IsType(&SchemaReference{}, o.Fields()[1].schema)
	a.Same(r, o.Fields()[1].Schema())
	a.Equal(r.ID(), o.Fields()[1].Schema().ID())
	
	m.AssertNumberOfCalls(t, "Get", 2)
}