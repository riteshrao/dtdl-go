package model

// CommandPayload maps to the request and response definition of a command.
type CommandPayload struct {
	NamedEntity
	schema SchemaType
}

// Command maps to a DTDL command defined in a model.
type Command struct {
	NamedEntity
	request  *CommandPayload
	response *CommandPayload
}


// Schema of a command request or response payload.
func (p *CommandPayload) Schema() SchemaType {
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

// Request definition of the command.
func (c *Command) Request() *CommandPayload {
	return c.request
}

// Response definition of a command.
func (c *Command) Response() *CommandPayload {
	return c.response
}

func parseCommand(i map[string]interface{}, t entityTracker) *Command {
	c := &Command{
		NamedEntity: parseNamedEntity(i),
	}

	if req, ok := i["request"].(map[string]interface{}); ok {
		c.request = &CommandPayload{
			NamedEntity: parseNamedEntity(req),
		}

		if s, ok := req["schema"]; ok {
			c.request.schema = parseSchema(s, t)
		}
	}

	if res, ok := i["response"].(map[string]interface{}); ok {
		c.response = &CommandPayload{
			NamedEntity: parseNamedEntity(res),
		}

		if s, ok := res["schema"]; ok {
			c.response.schema = parseSchema(s, t)
		}
	}

	t.Add(c)
	return c
}
