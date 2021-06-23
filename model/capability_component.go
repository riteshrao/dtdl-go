package model

// Component maps to a component type in DTDL
type Component struct {
	NamedEntity
	schema string
	tracker entityTracker
}

// Schema returns the interface schema of the component
func (c *Component) Schema() *Interface {
	if c.schema == "" {
		return nil
	}

	if ref, ok := c.tracker.Get(c.schema); ok {
		if i, ok := ref.(*Interface); ok {
			return i
		}
	}

	return nil
}

func parseComponent(i map[string]interface{}, t entityTracker) *Component {
	c := &Component{
		NamedEntity: parseNamedEntity(i),
		tracker: t,
	}

	if s, ok := i["schema"]; ok {
		c.schema = s.(string)
	}

	t.Add(c)
	return c
}