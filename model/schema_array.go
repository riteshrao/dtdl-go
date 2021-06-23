package model

// ArraySchema maps an Array schema definition in DTDL
type ArraySchema struct {
	Schema
	elementSchema SchemaType
}

// ElementSchema returns the schema definition of the array elements.
func (a *ArraySchema) ElementSchema() SchemaType {
	if ref, ok := a.elementSchema.(*SchemaReference); ok {
		s, ok := ref.Resolve()
		if ok {
			*&a.elementSchema = s
		} else {
			*&a.elementSchema = UnsupportedSchema
		}
	}

	return a.elementSchema
}

func parseArraySchema(i map[string]interface{}, t entityTracker) *ArraySchema {
	s := &ArraySchema{
		Schema: Schema{
			class:  ArraySchemaClass,
			Entity: parseEntity(i),
		},
	}

	if e, ok := i["elementSchema"]; ok {
		s.elementSchema = parseSchema(e, t)
	}

	t.Add(s)
	return s
}
