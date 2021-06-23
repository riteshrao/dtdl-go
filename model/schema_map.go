package model

// MapValueSchema maps to a
type MapValueSchema struct {
	Entity
	schema SchemaType
}

// MapSchema maps
type MapSchema struct {
	Schema
	mapValue MapValueSchema
}

// Schema gets the schema of the map value
func (mv *MapValueSchema) Schema() SchemaType {
	if ref, ok := mv.schema.(*SchemaReference); ok {
		s, ok := ref.Resolve()
		if ok {
			*&mv.schema = s
		} else {
			*&mv.schema = UnsupportedSchema
		}
	}

	return mv.schema
}

// MapValue gets the map value definition of the map
func (m *MapSchema) MapValue() *MapValueSchema {
	return &m.mapValue
}

func parseMapSchema(i map[string]interface{}, t entityTracker) *MapSchema {
	mv := MapValueSchema{}
	if v, ok := i["mapValue"].(map[string]interface{}); ok {
		mv.Entity = parseEntity(v)

		if mvs, ok := v["schema"]; ok {
			mv.schema = parseSchema(mvs, t)
		}
	}

	s := &MapSchema{
		Schema: Schema{
			class:  MapSchemaClass,
			Entity: parseEntity(i),
		},
		mapValue: mv,
	}

	t.Add(s)
	return s
}
