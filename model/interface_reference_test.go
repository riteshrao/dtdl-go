package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceReferenceResolve_WhenNotFound(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Get", "dtmi:ref;1").Return(&Entity{}, false)
	ref := InterfaceReference{
		id: "dtmi:ref;1",
		tracker: m,
	}

	res, ok := ref.Resolve()
	a.Nil(res)
	a.Equal(false, ok)
}

func TestInterfaceReferenceResolve_WhenNotInterface(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}

	m.On("Get", "dtmi:ref;1").Return(&Entity{}, true)
	ref := InterfaceReference{
		id: "dtmi:ref;1",
		tracker: m,
	}

	res, ok := ref.Resolve()
	a.Nil(res)
	a.Equal(false, ok)
}

func TestInterfaceReferenceResolve_WhenFound(t *testing.T) {
	a := assert.New(t)
	m := &mockTracker{}
	i := &Interface{}

	m.On("Get", "dtmi:ref;1").Return(i, true)
	ref := InterfaceReference{
		id: "dtmi:ref;1",
		tracker: m,
	}

	res, ok := ref.Resolve()
	a.Equal(true, ok)
	a.Same(i, res)
}