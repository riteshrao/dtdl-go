package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseCommand_Primitive(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Add", mock.MatchedBy(func (e EntityType) bool {
		return e.ID() == "dtmi:cmd;1"
	}))

	c := parseCommand(map[string]interface{}{
		"@id": "dtmi:cmd;1",
		"@type": "Command",
		"name": "cmd",
		"displayName": "Command",
		"description": "Command",
		"comment": "Command",
		"request": map[string]interface{}{
			"name": "req",
			"displayName": "Request",
			"description": "Request",
			"schema": "integer",
		},
		"response": map[string]interface{}{
			"name": "res",
			"displayName": "Response",
			"description": "Response",
			"schema": "string",
		},
	}, m)

	a.Equal("dtmi:cmd;1", c.ID())
	a.Equal("cmd", c.Name())
	a.Equal("Command", c.DisplayName())
	a.Equal("Command", c.Description())
	a.Equal("Command", c.Comment())
	a.NotNil(c.Request())
	a.Equal("req", c.Request().Name())
	a.Equal("Request", c.Request().DisplayName())
	a.Equal("Request", c.Request().Description())
	a.Same(IntegerSchema, c.Request().Schema())
	a.Equal("res", c.Response().Name())
	a.Equal("Response", c.Response().DisplayName())
	a.Equal("Response", c.Response().Description())
	a.Same(StringSchema, c.Response().Schema())
	m.AssertExpectations(t)
}