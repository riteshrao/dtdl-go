package model

// Telemetry maps to a DTDL telemetry defined in a model.
type Telemetry struct {
	NamedEntity
	schema SchemaType
}

// Schema definition of a telemetry.
func (t *Telemetry) Schema() SchemaType {
	if ref, ok := t.schema.(*SchemaReference); ok {
		s, ok := ref.Resolve()
		if ok {
			*&t.schema = s
		} else {
			*&t.schema = UnsupportedSchema
		}
	}

	return t.schema
}

func parseTelemetry(i map[string]interface{}, t entityTracker) *Telemetry {
	c := &Telemetry{
		NamedEntity: parseNamedEntity(i),
	}

	if s, ok := i["schema"]; ok {
		c.schema = parseSchema(s, t)
	}

	t.Add(c)
	return c
}
