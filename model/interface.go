package model

import "reflect"

// Interface represents an interface model
type Interface struct {
	Entity
	contents []CapabilityType
	extends []InterfaceReference
	tracker  entityTracker
}

// Contents returns the contents of the interface
func (i *Interface) Contents() []CapabilityType {
	return i.contents
}

// Extends returns the parent interfaces this interface extends from.
func (i *Interface) Extends() []*Interface {
	extends := make([]*Interface, 0)
	for _, ref := range i.extends {
		if res, ok := ref.Resolve(); ok {
			extends = append(extends, res)
		}
	}
	return extends
}

// GetCapability returns a capability in the interface with the specified type.
func (i *Interface) GetCapability(id string) CapabilityType {
	for _, cap := range i.contents {
		if (cap.ID() == id) {
			return cap
		}
	}

	return nil
}

func ParseInterface(i map[string]interface{}, t entityTracker) *Interface {
	res := &Interface{
		Entity: parseEntity(i),
	}

	extends := make([]InterfaceReference, 0)
	if items, ok := i["extends"].([]string); ok {
		for _, item := range items {
			extends = append(extends, InterfaceReference{
				id: item,
				tracker: t,
			})
		}
	}

	contents := make([]CapabilityType, 0)
	if items, ok := i["contents"].([]interface{}); ok {
		for _, item := range items {
			if cap, ok := item.(map[string]interface{}); ok {
				if types, ok := cap["@type"]; ok {
					if reflect.TypeOf(types).Kind() == reflect.String {
						switch types.(string) {
						case "Command":
							contents = append(contents, parseCommand(cap, t))
						case "Component":
							contents = append(contents, parseComponent(cap, t))
						case "Property":
							contents = append(contents, parseProperty(cap, t))
						case "Relationship":
							contents = append(contents, parseRelationship(cap, t))
						case "Telemetry":
							contents = append(contents, parseTelemetry(cap, t))
						}
					}

					if reflect.TypeOf(types).Kind() == reflect.Slice {
						for _, ct := range types.([]string) {
							switch ct {
							case "Command":
								contents = append(contents, parseCommand(cap, t))
							case "Component":
								contents = append(contents, parseComponent(cap, t))
							case "Property":
								contents = append(contents, parseProperty(cap, t))
							case "Relationship":
								contents = append(contents, parseRelationship(cap, t))
							case "Telemetry":
								contents = append(contents, parseTelemetry(cap, t))
							}
						}
					}
				}
			}
		}
	}

	res.extends = extends
	res.contents = contents
	
	t.Add(res)
	return res
}
