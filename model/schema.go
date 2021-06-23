package model

import "reflect"

type (
	// Schema maps to a schema definition, built-in or custom
	Schema struct {
		Entity
		class SchemaClass
	}

	// SchemaClass - the class type of a schema.
	SchemaClass int
)

const (
	// ArraySchemaClass applied to array schema types.
	ArraySchemaClass SchemaClass = iota
	// BooleanSchemaClass applied to boolean schema types.
	BooleanSchemaClass
	// DateSchemaClass applied to date schema types.
	DateSchemaClass
	// DateTimeSchemaClass applied to date time schema types.
	DateTimeSchemaClass
	// DoubleSchemaClass applied to double schema types.
	DoubleSchemaClass
	// EnumSchemaClass applied to enum schema types.
	EnumSchemaClass
	// GeopointSchemaClass applied to geopoint schema types.
	GeopointSchemaClass
	// IntegerSchemaClass applied to integer schema types.
	IntegerSchemaClass
	// LongSchemaClass applied to long schema types.
	LongSchemaClass
	// ObjectSchemaClass applied to object schema types.
	ObjectSchemaClass
	// MapSchemaClass applied to map schema types.
	MapSchemaClass
	// StringSchemaClass class applied to string schema types.
	StringSchemaClass
	// UnsupportedSchemaClass type applied to unsupported schema types.
	UnsupportedSchemaClass
)

// Class gets the schema class of the schema
func (p *Schema) Class() SchemaClass {
	return p.class
}

// parseSchema parses schemas
func parseSchema(input interface{}, tracker entityTracker) SchemaType {
	switch reflect.TypeOf(input).Kind() {
	case reflect.String:
		switch input.(string) {
		case "boolean":
			return BooleanSchema
		case "date":
			return DateSchema
		case "dateTime":
			return DateTimeSchema
		case "double":
			return DoubleSchema
		case "geopoint":
			return GeopointSchema
		case "integer":
			return IntegerSchema
		case "long":
			return LongSchema
		case "string":
			return StringSchema
		default:
			return &SchemaReference{
				id:      input.(string),
				tracker: tracker,
			}
		}
	case reflect.Map:
		if t, ok := input.(map[string]interface{})["@type"]; ok {
			st := ""
			if reflect.TypeOf(t).Kind() == reflect.String {
				st = t.(string)
			}

			if reflect.TypeOf(t).Kind() == reflect.Array {
				st = t.([]string)[0]
			}

			switch st {
			case "Array":
				return parseArraySchema(input.(map[string]interface{}), tracker)
			case "Enum":
				return parseEnumSchema(input.(map[string]interface{}), tracker)
			case "Map":
				return parseMapSchema(input.(map[string]interface{}), tracker)
			case "Object":
				return parseObjectSchema(input.(map[string]interface{}), tracker)
			}
		}
	}
	return UnsupportedSchema
}
