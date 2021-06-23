package model

import (
	"fmt"
	"reflect"
)

// parseCapability map and returns a CapabilityType.
func parseCapability(i map[string]interface{}, tracker entityTracker) (CapabilityType, error) {
	t, ok := i["@type"]
	if !ok {
		return nil, fmt.Errorf("Expected @type not found")
	}
	
	var types []string
	switch reflect.TypeOf(t).Kind() {
	case reflect.String:
		types = []string{t.(string)}
	case reflect.Slice:
		types = t.([]string)
	}

	for _, t := range types {
		switch (t) {
		case "Command":
			return parseCommand(i, tracker), nil
		case "Property":
			return parseProperty(i, tracker), nil
		case "Telemetry":
			return parseTelemetry(i, tracker), nil
		case "Component":
			return parseComponent(i, tracker), nil
		case "Relationship":
			return parseRelationship(i, tracker), nil
		}
	}

	return nil, fmt.Errorf("Unknown capability type found")
}
