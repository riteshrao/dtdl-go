package model

// Property maps to a DTDL property defined in a model.
type Property struct {
	NamedEntity
	schema   SchemaType
	writable bool
}


// Schema definition of a property.
func (p *Property) Schema() SchemaType {
	if ref, ok := p.schema.(*SchemaReference); ok {
		s, ok := ref.Resolve()
		if ok {
			*&p.schema = s
		} else {
			*&p.schema = UnsupportedSchema
		}
	}

	return p.schema
}

// Writable flag indicating if a property is writable.
func (p *Property) Writable() bool {
	return p.writable
}

func parseProperty(i map[string]interface{}, t entityTracker) *Property {
	p := &Property{
		NamedEntity: parseNamedEntity(i),
	}

	if s, ok := i["schema"]; ok {
		p.schema = parseSchema(s, t)
	}

	if w, ok := i["writable"].(bool); ok {
		p.writable = w
	}

	t.Add(p)

	return p
}