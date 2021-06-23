package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseInterface(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	ext := &Interface{}

	m.On("add", mock.Anything).Times(6)
	m.On("get", "dtmi:extends;1").Return(ext, true)

	i := ParseInterface(map[string]interface{}{
		"@id": "dtmi:interface;1",
		"@type": "Interface",
		"displayName": "Interface",
		"description": "Interface",
		"extends": []string{"dtmi:extends;1"},
		"contents": []interface{}{
			map[string]interface{}{
				"@id": "dtmi:command;1",
				"@type": "Command",
			},
			map[string]interface{}{
				"@id": "dtmi:component;1",
				"@type": "Component",
			},
			map[string]interface{}{
				"@id": "dtmi:property;1",
				"@type": []string{"Property", "DataRate"},
			},
			map[string]interface{}{
				"@id": "dtmi:rel;1",
				"@type": "Relationship",
			},
			map[string]interface{}{
				"@id": "dtmi:telemetry;1",
				"@type": []string{"Telemetry", "Temperature"},
			},
		},
	}, m)

	a.Equal("dtmi:interface;1", i.ID())
	a.Equal("Interface", i.DisplayName())
	a.Equal("Interface", i.Description())
	a.Equal(1, len(i.Extends()))
	a.Equal(5, len(i.Contents()))

	a.Same(ext, i.Extends()[0])
	a.IsType(&Command{}, i.Contents()[0])
	a.IsType(&Component{}, i.Contents()[1])
	a.IsType(&Property{}, i.Contents()[2])
	a.IsType(&Relationship{}, i.Contents()[3])
	a.IsType(&Telemetry{}, i.Contents()[4])
	m.AssertExpectations(t)
}