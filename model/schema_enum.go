package model

import "reflect"

// EnumSchemaValue maps to an enum value defined in a custom enum schema
type EnumSchemaValue struct {
	NamedEntity
	enumValue interface{}
}

// EnumSchema maps to an Enums schema definition in DTDL
type EnumSchema struct {
	Schema
	valueSchema SchemaType
	enumValues  []EnumSchemaValue
}

// Kind gets the value kind of the enum value.
func (v *EnumSchemaValue) Kind() reflect.Kind {
	return reflect.TypeOf(v.enumValue).Kind()
}

// IntegerValue returns the value of the enum as an int
func (v *EnumSchemaValue) IntegerValue() (int, bool) {
	if reflect.TypeOf(v.enumValue).Kind() == reflect.Int {
		return v.enumValue.(int), true
	}
	return 0, false
}

// StringValue returns the value of the enum as a string.
func (v *EnumSchemaValue) StringValue() (string, bool) {
	if reflect.TypeOf(v.enumValue).Kind() == reflect.String {
		return v.enumValue.(string), true
	}

	return "", false
}

// ValueSchemaClass returns the schema class of the enum values
func (e *EnumSchema) ValueSchemaClass() SchemaClass {
	return e.valueSchema.Class()
}

// Values returns the enum schema values.
func (e *EnumSchema) Values() []EnumSchemaValue {
	return e.enumValues
}

func parseEnumSchema(i map[string]interface{}, t entityTracker) *EnumSchema {
	evs := make([]EnumSchemaValue, 0)
	if ev, ok := i["enumValues"]; ok {
		for _, i := range ev.([]interface{}) {
			enumValue := i.(map[string]interface{})
			evs = append(evs, EnumSchemaValue{
				NamedEntity: parseNamedEntity(enumValue),
				enumValue:   enumValue["enumValue"],
			})
		}
	}

	var val *Schema
	if valSchema, ok := i["valueSchema"]; ok {
		switch valSchema.(string) {
		case "string":
			val = StringSchema
		case "integer":
			val = IntegerSchema
		default:
			val = UnsupportedSchema
		}
	}

	s := &EnumSchema{
		Schema: Schema{
			class:  EnumSchemaClass,
			Entity: parseEntity(i),
		},
		valueSchema: val,
		enumValues:  evs,
	}

	t.Add(s)
	return s
}
