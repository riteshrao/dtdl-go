package model

// ObjectField maps to an field definition in an custom DTDL object schema
type ObjectField struct {
	NamedEntity
	schema SchemaType
}

// ObjectSchema maps to a custom Object schema definition in DTDL.
type ObjectSchema struct {
	Schema
	fields []ObjectField
}

// Schema returns the schema definition of the field.
func (of *ObjectField) Schema() SchemaType {
	if ref, ok := of.schema.(*SchemaReference); ok {
		s, ok := ref.Resolve()
		if ok {
			*&of.schema = s
		} else {
			*&of.schema = UnsupportedSchema
		}
	}

	return of.schema
}

// Fields returns the fields defined for the object schema.
func (o *ObjectSchema) Fields() []ObjectField {
	return o.fields
}

func parseObjectSchema(i map[string]interface{}, t entityTracker) *ObjectSchema {
	f := make([]ObjectField, 0)
	if fields, ok := i["fields"].([]interface{}); ok {
		for _, i := range fields {
			f = append(f, ObjectField{
				NamedEntity: parseNamedEntity(i.(map[string]interface{})),
				schema:      parseSchema(i.(map[string]interface{})["schema"], t),
			})
		}
	}

	s := &ObjectSchema{
		Schema: Schema{
			class:  ObjectSchemaClass,
			Entity: parseEntity(i),
		},
		fields: f,
	}

	t.Add(s)
	return s
}
