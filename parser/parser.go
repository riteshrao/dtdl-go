package parser

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/riteshrao/dtdl-go/model"
)

// ModelParser allows parsing DTDL inputs.
type ModelParser struct {
	entities map[string]model.EntityType
}

// NewModelParser creates a new parser instance.
func NewModelParser() *ModelParser {
	return &ModelParser{
		entities: make(map[string]model.EntityType, 0),
	}
}

// Parse parses an input DTDL.
// Parsed input is available as entities
func (p *ModelParser) Parse(input []byte) error {
	var parsed interface{}
	if err := json.Unmarshal(input, &parsed); err != nil {
		return err
	}

	switch reflect.TypeOf(parsed).Kind() {
	case reflect.Slice:
		for _, item := range parsed.([]interface{}) {
			if err := p.parseInterface(item.(map[string]interface{})); err != nil {
				return err
			}
		}
		return nil
	case reflect.Map:
		return p.parseInterface(parsed.(map[string]interface{}))
	default:
		return fmt.Errorf("Unrecognized content")
	}
}

// GetCommand returns a command by its id.
func (p *ModelParser) GetCommand(id string) (*model.Command, bool) {
	e, ok := p.entities[id]
	if !ok {
		return nil, false
	}

	cmd, ok := e.(*model.Command)
	if !ok {
		return nil, false
	}

	return cmd, true
}

// GetComponent returns a component by its id.
func (p *ModelParser) GetComponent(id string) (*model.Component, bool) {
	c, ok := p.entities[id]
	if !ok {
		return nil, false
	}

	cmp, ok := c.(*model.Component)
	if !ok {
		return nil, false
	}

	return cmp, true
}

// GetInterface returns a interface by its id.
func (p *ModelParser) GetInterface(id string) (*model.Interface, bool) {
	e, ok := p.entities[id]
	if !ok {
		return nil, false
	}

	i, ok := e.(*model.Interface)
	if !ok {
		return nil, false
	}

	return i, true
}

// GetProperty gets a property by its id
func (p *ModelParser) GetProperty(id string) (*model.Property, bool) {
	e, ok := p.entities[id]
	if !ok {
		return nil, false
	}

	prop, ok := e.(*model.Property)
	if !ok {
		return nil, false
	}

	return prop, true
}

// GetRelationship gets a relationship by its id
func (p *ModelParser) GetRelationship(id string) (*model.Relationship, bool) {
	e, ok := p.entities[id]
	if !ok {
		return nil, false
	}

	rel, ok := e.(*model.Relationship)
	if !ok {
		return nil, false
	}

	return rel, true
}

// GetTelemetry returns a command by its id.
func (p *ModelParser) GetTelemetry(id string) (*model.Telemetry, bool) {
	e, ok := p.entities[id]
	if !ok {
		return nil, false
	}

	tel, ok := e.(*model.Telemetry)
	if !ok {
		return nil, false
	}

	return tel, true
}

// Add adds an entity to the parser
func (p *ModelParser) Add(entity model.EntityType) error {
	if entity.ID() == "" {
		return nil
	}

	if _, ok := p.entities[entity.ID()]; ok {
		return fmt.Errorf("A duplicate definition for %s was found", entity.ID())
	}

	p.entities[entity.ID()] = entity
	return nil
}

// Get returns an entity by its id.
func (p *ModelParser) Get(id string) (model.EntityType, bool) {
	e, ok := p.entities[id]
	return e, ok
}

func (p *ModelParser) parseInterface(i map[string]interface{}) error {
	if _, ok := i["@context"]; !ok {
		return fmt.Errorf("Expected @context not found in input")
	}

	if _, ok := i["@type"]; !ok {
		return fmt.Errorf("Expected @type not found in input")
	}

	if _, ok := i["@id"]; !ok {
		return fmt.Errorf("Expected @id not found in input")
	}

	if i["@context"] != "dtmi:dtdl:context;2" {
		return fmt.Errorf("Unrecognized @context")
	}

	if i["@type"] != "Interface" {
		return fmt.Errorf("Unrecognized @type. Expected Interface")
	}

	model.ParseInterface(i, p)
	return nil
}
