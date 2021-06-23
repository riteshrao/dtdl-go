package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID_ReturnUnsupported_WhenNotFound(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("get","dtmi:contoso:schema;1").Return(&Entity{}, false)
	s := SchemaReference{
		id: "dtmi:contoso:schema;1",
		tracker: m,
	}

	a.Equal(UnsupportedSchema.ID(), s.ID())
	m.AssertExpectations(t)
}

func TestID_ReturnUnsupported_WhenNotSchema(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("get","dtmi:contoso:schema;1").Return(&Entity{}, true)
	s := SchemaReference{
		id: "dtmi:contoso:schema;1",
		tracker: m,
	}

	a.Equal(UnsupportedSchema.ID(), s.ID())
	m.AssertExpectations(t)
}

func TestID_ReturnReferencedID(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	r := &Schema{
		Entity: Entity{
			id: "dtmi:contoso:schema;1",
		},
	}

	m.On("get", "dtmi:contoso:schema;1").Return(r, true)
	
	s := SchemaReference{
		id: "dtmi:contoso:schema;1",
		tracker: m,
	}

	a.Equal(r.ID(), s.ID())
	m.AssertExpectations(t)
}

func TestTypes_ReturnUnsupported_WhenNotFound(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("get","dtmi:contoso:schema;1").Return(&Entity{}, false)
	s := SchemaReference{
		id: "dtmi:contoso:schema;1",
		tracker: m,
	}

	a.ElementsMatch(UnsupportedSchema.Types(), s.Types())
	m.AssertExpectations(t)
}

func TestTypes_ReturnUnsupported_WhenNotSchema(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("get","dtmi:contoso:schema;1").Return(&Entity{}, true)
	s := SchemaReference{
		id: "dtmi:contoso:schema;1",
		tracker: m,
	}

	a.Equal(UnsupportedSchema.Types(), s.Types())
	m.AssertExpectations(t)
}

func TestTypes_ReturnReferencedTypes(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	r := &Schema{
		Entity: Entity{
			types: []string{"TypeA", "TypeB"},
		},
	}

	m.On("get", "dtmi:contoso:schema;1").Return(r, true)
	
	s := SchemaReference{
		id: "dtmi:contoso:schema;1",
		tracker: m,
	}

	a.ElementsMatch(r.Types(), s.Types())
	m.AssertExpectations(t)
}

func TestDisplayName_ReturnUnsupported_WhenNotFound(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("get","dtmi:contoso:schema;1").Return(&Entity{}, true)
	s := SchemaReference{
		id: "dtmi:contoso:schema;1",
		tracker: m,
	}

	a.Equal(UnsupportedSchema.DisplayName(), s.DisplayName())
	m.AssertExpectations(t)
}