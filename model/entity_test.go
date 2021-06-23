package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEntity(t *testing.T) {
	input := map[string]interface{}{
		"@id":         "dtmi:com:contoso:test;1",
		"@type":       "Interface",
		"displayName": "Test Interface",
		"description": "Test Interface",
		"comment":     "Test Interface",
	}

	e := parseEntity(input)
	a := assert.New(t)

	a.Equal(e.ID(), "dtmi:com:contoso:test;1", "id should be equal")
	a.Equal(len(e.Types()), 1, "Expected types to have a single element")
	a.Equal(e.Types()[0], "Interface", "Expected")
	a.Equal(e.DisplayName(), "Test Interface", "Expected default locale display name")
	a.Equal(e.Description(), "Test Interface", "Expected default locale description")
	a.Equal(e.Comment(), "Test Interface", "Expected comment to be equal")
}

func TestParseEntityMultipleTypes(t *testing.T) {
	input := map[string]interface{}{
		"@id":   "dtmi:com:contoso:test;1",
		"@type": []string{"Telemetry", "Temperature"},
	}

	e := parseEntity(input)
	a := assert.New(t)
	a.Equal(len(e.Types()), 2, "Expected two types to be defined")
	a.Equal(e.Types()[0], "Telemetry", "Expected first type to be telemetry")
	a.Equal(e.Types()[1], "Temperature", "Expected second type to be telemetry")
}

func TestParseLocaleSpecificDisplayName(t *testing.T) {
	input := map[string]interface{}{
		"@id":   "dtmi:com:contoso:test;1",
		"@type": "Interface",
		"displayName": map[string]string{
			"en": "Test interface",
			"fr": "Test interface (fr)",
		},
	}

	e := parseEntity(input)
	a := assert.New(t)

	a.Equal(e.DisplayName(), "Test interface", "Expected default to still map to en")
	a.Equal(e.LocaleDisplayName("fr"), "Test interface (fr)", "Expected locale value to be equal")
}

func TestParseLocaleSpecificDescription(t *testing.T) {
	input := map[string]interface{}{
		"@id":   "dtmi:com:contoso:test;1",
		"@type": "Interface",
		"description": map[string]string{
			"en": "Test interface",
			"fr": "Test interface (fr)",
		},
	}

	e := parseEntity(input)
	a := assert.New(t)

	a.Equal(e.Description(), "Test interface", "Expected default to still map to en")
	a.Equal(e.LocaleDescription("fr"), "Test interface (fr)", "Expected locale value to be equal")
}

func TestParseNamedEntity(t *testing.T) {
	input := map[string]interface{}{
		"@id":         "dtmi:com:contoso:test;1",
		"@type":       "Interface",
		"displayName": "Test Interface",
		"description": "Test Interface",
		"comment":     "Test Interface",
		"name": "test",
	}

	e := parseNamedEntity(input)
	a := assert.New(t)

	a.Equal(e.ID(), "dtmi:com:contoso:test;1", "id should be equal")
	a.Equal(len(e.Types()), 1, "Expected types to have a single element")
	a.Equal(e.Types()[0], "Interface", "Expected")
	a.Equal(e.DisplayName(), "Test Interface", "Expected default locale display name")
	a.Equal(e.Description(), "Test Interface", "Expected default locale description")
	a.Equal(e.Comment(), "Test Interface", "Expected comment to be equal")
	a.Equal(e.Name(), "test", "Expected name to be equal")
}