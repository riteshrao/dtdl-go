package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseTelemetry_Primitive(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:telemetry;1"
	}))

	c := parseTelemetry(map[string]interface{} {
		"@id": "dtmi:telemetry;1",
		"@type": []string{"Telemetry", "Temperature"},
		"name": "temp",
		"displayName": "Temperature",
		"description": "Temperature",
		"comment": "Temperature",
		"schema": "double",
	}, m)

	a.Equal("dtmi:telemetry;1", c.ID())
	a.True(c.IsType("Temperature"))
	a.Equal("temp", c.Name())
	a.Equal("Temperature", c.DisplayName())
	a.Equal("Temperature", c.Description())
	a.Equal("Temperature", c.Comment())
	a.Same(DoubleSchema, c.Schema())
	m.AssertExpectations(t)
}

func TestParseTelemetry_SchemaReference(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	r := &Schema{
		Entity: Entity{
			id: "dtmi:schema:ref;1",
		},
	}

	m.On("Add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:telemetry;1"
	}))

	m.On("Get", "dtmi:schema:ref;1").Return(r, true)
	c := parseTelemetry(map[string]interface{} {
		"@id": "dtmi:telemetry;1",
		"@type": []string{"Telemetry", "Temperature"},
		"name": "temp",
		"schema": "dtmi:schema:ref;1",
	}, m)

	a.IsType(&SchemaReference{}, c.schema)
	a.Same(r, c.Schema())
	m.AssertNumberOfCalls(t, "Get", 1)
	m.AssertExpectations(t)
}

func TestParseTelemetry_SchemaReference_Unsupported_WhenNotFound(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	
	m.On("Add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:telemetry;1"
	}))

	m.On("Get", "dtmi:schema:ref;1").Return(&Entity{}, true)
	c := parseTelemetry(map[string]interface{} {
		"@id": "dtmi:telemetry;1",
		"@type": []string{"Telemetry", "Temperature"},
		"name": "temp",
		"schema": "dtmi:schema:ref;1",
	}, m)

	a.IsType(&SchemaReference{}, c.schema)
	a.Same(UnsupportedSchema, c.Schema())
	m.AssertNumberOfCalls(t, "Get", 1)
	m.AssertExpectations(t)
}