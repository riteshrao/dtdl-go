package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseProperty_Primitive(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:property;1"
	}))

	p := parseProperty(map[string]interface{}{
		"@id": "dtmi:property;1",
		"@type": []string{"Property", "DataRate"},
		"name": "dataRate",
		"displayName": "Data Rate",
		"description": "Data Rate",
		"comment": "Data Rate",
		"writable": true,
		"schema": "integer",
	}, m)

	a.Equal("dtmi:property;1", p.ID())
	a.True(p.IsType("DataRate"))
	a.Equal("dataRate", p.Name())
	a.Equal("Data Rate", p.DisplayName())
	a.Equal("Data Rate", p.Description())
	a.Equal("Data Rate", p.Comment())
	a.Equal(true, p.Writable())
	a.Same(IntegerSchema, p.Schema())
	m.AssertExpectations(t)
}


func TestParseProperty_SchemaReference(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	r := &Schema{
		Entity: Entity{
			id: "dtmi:schema:ref;1",
		},
	}

	m.On("add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:property;1"
	}))

	m.On("get", "dtmi:schema:ref;1").Return(r, true)
	c := parseProperty(map[string]interface{} {
		"@id": "dtmi:property;1",
		"@type": "Property",
		"name": "temp",
		"schema": "dtmi:schema:ref;1",
	}, m)

	a.IsType(&SchemaReference{}, c.schema)
	a.Same(r, c.Schema())
	m.AssertNumberOfCalls(t, "get", 1)
	m.AssertExpectations(t)
}

func TestParseProperty_SchemaReference_Unsupported_WhenNotFound(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	
	m.On("add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:property;1"
	}))

	m.On("get", "dtmi:schema:ref;1").Return(&Entity{}, true)
	c := parseProperty(map[string]interface{} {
		"@id": "dtmi:property;1",
		"@type": "Property",
		"name": "temp",
		"schema": "dtmi:schema:ref;1",
	}, m)

	a.IsType(&SchemaReference{}, c.schema)
	a.Same(UnsupportedSchema, c.Schema())
	m.AssertNumberOfCalls(t, "get", 1)
	m.AssertExpectations(t)
}