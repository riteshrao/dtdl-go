package model

// InterfaceReference allows resolving a reference to an interface
type InterfaceReference struct {
	id string
	tracker entityTracker
}

// Resolve resolves the referenced interface
func (r *InterfaceReference) Resolve() (*Interface, bool) {
	e, ok := r.tracker.Get(r.id)
	if !ok {
		return nil, false
	}

	i, ok := e.(*Interface)
	if !ok {
		return nil, false
	}

	return i, true
}