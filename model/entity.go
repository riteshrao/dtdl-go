package model

import "reflect"

// Entity is the root of all DTDL model entities.
type Entity struct {
	id           string
	displayNames map[string]interface{}
	descriptions map[string]interface{}
	comment      string
	types        []string
}

// ID returns the @id of the entity.
func (e *Entity) ID() string {
	return e.id
}

// Types return the @types assigned to the entity.
func (e *Entity) Types() []string {
	return e.types
}

// DisplayName returns the display name of the entity in the default locale, i.e. en
func (e *Entity) DisplayName() string {
	if v, ok := e.displayNames["en"]; ok {
		return v.(string)
	}
	return ""
}

// Description returns the description of the entity in the default locale. i.e. en
func (e *Entity) Description() string {
	if v, ok := e.descriptions["en"]; ok {
		return v.(string)
	}

	return ""
}

// Comment returns the developer comment
func (e *Entity) Comment() string {
	return e.comment
}

// LocaleDescription returns a localized representation of the description, if available.
// If a value for the locale is not found, it returns an empty string.
func (e *Entity) LocaleDescription(locale string) string {
	if desc, ok := e.descriptions[locale]; ok {
		return desc.(string)
	}

	return ""
}

// LocaleDisplayName returns a localized representation of the display name, if available.
// If a value for the locale is not found, it returns an empty string.
func (e *Entity) LocaleDisplayName(locale string) string {
	if disp, ok := e.displayNames[locale]; ok {
		return disp.(string)
	}

	return ""
}

// IsType checks if a entity is of a specific type
func (e *Entity) IsType(typeName string) bool {
	for _, t := range e.types {
		if t == typeName {
			return true
		}
	}

	return false
}

// NamedEntity that contains a user specified name.
type NamedEntity struct {
	Entity
	name string
}

// Name returns the name of the named entity.
func (n *NamedEntity) Name() string {
	return n.name
}

// parseEntity parses an entity.
func parseEntity(input map[string]interface{}) Entity {
	e := Entity{}
	if i, ok := input["@id"]; ok {
		e.id = i.(string)
	}

	if types, ok := input["@type"]; ok {
		switch reflect.TypeOf(types).Kind() {
		case reflect.Slice:
			t := []string{}
			for _, i := range types.([]interface{}) {
				t = append(t, i.(string))
			}
			e.types = t
		case reflect.String:
			e.types = []string{types.(string)}
		}
	} else {
		e.types = make([]string, 0)
	}

	if dispName, ok := input["displayName"]; ok {
		switch reflect.TypeOf(dispName).Kind() {
		case reflect.Map:
			e.displayNames = dispName.(map[string]interface{})
		case reflect.String:
			e.displayNames = map[string]interface{}{
				"en": dispName.(string),
			}
		}
	} else {
		e.displayNames = make(map[string]interface{})
	}

	if desc, ok := input["description"]; ok {
		switch reflect.TypeOf(desc).Kind() {
		case reflect.Map:
			e.descriptions = desc.(map[string]interface{})
		case reflect.String:
			e.descriptions = map[string]interface{}{
				"en": desc.(string),
			}
		}
	} else {
		e.descriptions = make(map[string]interface{})
	}

	if cmnt, ok := input["comment"]; ok {
		e.comment = cmnt.(string)
	}

	return e
}

// parseNamedEntity parses a named entity.
func parseNamedEntity(input map[string]interface{}) NamedEntity {
	e := parseEntity(input)
	n := NamedEntity{Entity: e}

	if nm, ok := input["name"]; ok {
		n.name = nm.(string)
	}

	return n
}
