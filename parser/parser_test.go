package parser

import (
	"testing"

	"github.com/riteshrao/dtdl-go/model"
	"github.com/stretchr/testify/assert"
)

func TestParseSingleInterface(t *testing.T) {
	a := assert.New(t)
	p := NewModelParser()

	p.Parse([]byte(`
	{
		"@context": "dtmi:dtdl:context;2",
		"@id": "dtmi:interface;1",
		"@type": "Interface",
		"displayName": "Test Interface",
		"description": "Test Interface",
		"comment": "Test Interface",
		"contents": [
			{
				"@type": "Command",
				"name": "reboot",
				"displayName": "Reboot",
				"request": {
					"name": "delay",
					"displayName": "Delay",
					"schema": "double"
				},
				"response": {
					"name": "status",
					"displayName": "Status",
					"schema": "string"
				}
			}
		]
	}
	`))

	i, ok := p.GetInterface("dtmi:interface;1")
	a.True(ok)
	a.IsType(&model.Interface{}, i)
	a.Equal("Test Interface", i.DisplayName())
	a.Equal("Test Interface", i.Description())
	a.Equal("Test Interface", i.Comment())
	a.Equal(1, len(i.Contents()))

	cmd, ok := i.Contents()[0].(*model.Command)
	a.True(ok)
	a.Equal("reboot", cmd.Name())
	a.Equal("Reboot", cmd.DisplayName())
	a.Equal("delay", cmd.Request().Name())
	a.Equal("Delay", cmd.Request().DisplayName())
	a.Equal(model.DoubleSchema, cmd.Request().Schema())
	a.Equal("status", cmd.Response().Name())
	a.Equal("Status", cmd.Response().DisplayName())
	a.Equal(model.StringSchema, cmd.Response().Schema())
}

func TestParseMultipleInterface(t *testing.T) {
	a := assert.New(t)
	p := NewModelParser()

	err := p.Parse([]byte(`
	[
		{
			"@context": "dtmi:dtdl:context;2",
			"@id": "dtmi:interface;1",
			"@type": "Interface",
			"displayName": "Test Interface",
			"description": "Test Interface",
			"comment": "Test Interface",
			"extends": ["dtmi:shared;1"],
			"contents": [
				{
					"@type": "Telemetry",
					"name": "reboot",
					"displayName": "Reboot",
					"schema": "double"
				}
			]
		},
		{
			"@context": "dtmi:dtdl:context;2",
			"@id": "dtmi:shared;1",
			"@type": "Interface",
			"displayName": "Shared Interface",
			"description": "Shared Interface",
			"contents": [
				{
					"@type": "Property",
					"name": "enabled",
					"displayName": "Enabled",
					"schema": "boolean"
				}
			]
		}
	]
	`))

	a.Nil(err)

	root, ok := p.GetInterface("dtmi:interface;1")
	a.True(ok)

	shared, ok := p.GetInterface("dtmi:shared;1")
	a.True(ok)

	a.Same(shared, root.Extends()[0])
}

func TestParseMultiContext(t *testing.T) {
	a := assert.New(t)
	p := NewModelParser()

	err := p.Parse([]byte(`
	{
		"@context": ["dtmi:dtdl:context;2", "dtmi:iotcentral:context;2"],
		"@id": "dtmi:interface;1",
		"@type": "Interface",
		"displayName": "Test Interface",
		"description": "Test Interface",
		"comment": "Test Interface",
		"contents": [
			{
				"@type": "Command",
				"name": "reboot",
				"displayName": "Reboot",
				"request": {
					"name": "delay",
					"displayName": "Delay",
					"schema": "double"
				},
				"response": {
					"name": "status",
					"displayName": "Status",
					"schema": "string"
				}
			}
		]
	}
	`))

	a.Nil(err)

	i, ok := p.GetInterface("dtmi:interface;1")
	a.True(ok)
	a.IsType(&model.Interface{}, i)
	a.Equal("Test Interface", i.DisplayName())
	a.Equal("Test Interface", i.Description())
	a.Equal("Test Interface", i.Comment())
	a.Equal(1, len(i.Contents()))

	cmd, ok := i.Contents()[0].(*model.Command)
	a.True(ok)
	a.Equal("reboot", cmd.Name())
	a.Equal("Reboot", cmd.DisplayName())
	a.Equal("delay", cmd.Request().Name())
	a.Equal("Delay", cmd.Request().DisplayName())
	a.Equal(model.DoubleSchema, cmd.Request().Schema())
	a.Equal("status", cmd.Response().Name())
	a.Equal("Status", cmd.Response().DisplayName())
	a.Equal(model.StringSchema, cmd.Response().Schema())
}