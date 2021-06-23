package model

// SchemaReference -
type SchemaReference struct {
	id      string
	tracker entityTracker
}

// ID returns the id of the referenced schema
func (d *SchemaReference) ID() string {
	if e, ok := d.Resolve(); ok {
		return e.ID()
	}

	return UnsupportedSchema.ID()
}

// Types returns the type annotations of the schema.
func (d *SchemaReference) Types() []string {
	if e, ok := d.Resolve(); ok {
		return e.Types()
	}

	return UnsupportedSchema.Types()
}

// DisplayName returns the display name of the schema.
func (d *SchemaReference) DisplayName() string {
	if e, ok := d.Resolve(); ok {
		return e.DisplayName()
	}

	return UnsupportedSchema.DisplayName()
}

// Description returns the description of the schema.
func (d *SchemaReference) Description() string {
	if e, ok := d.Resolve(); ok {
		return e.Description()
	}

	return UnsupportedSchema.Description()
}

// Comment returns the comment of the schema.
func (d *SchemaReference) Comment() string {
	if e, ok := d.Resolve(); ok {
		return e.Comment()
	}

	return UnsupportedSchema.Comment()
}

// LocaleDisplayName returns a localized display name of the schema.
// If no locale specific display name matching the specified locale is found, then an empty string is returned.
func (d *SchemaReference) LocaleDisplayName(locale string) string {
	if e, ok := d.Resolve(); ok {
		return e.LocaleDisplayName(locale)
	}

	return UnsupportedSchema.LocaleDisplayName(locale)
}

// LocaleDescription returns a localized display name of the schema.
// If no locale specific display name matching the specified locale is found, then an empty string is returned.
func (d *SchemaReference) LocaleDescription(locale string) string {
	if e, ok := d.Resolve(); ok {
		return e.LocaleDescription(locale)
	}

	return UnsupportedSchema.LocaleDescription(locale)
}

// Class returns the class type of the schema.
func (d *SchemaReference) Class() SchemaClass {
	if e, ok := d.Resolve(); ok {
		if s, ok := e.(SchemaType); ok {
			return s.Class()
		}
	}

	return UnsupportedSchema.class
}

// Resolve resolves the reference to a shared schema definition in the model.
func (d *SchemaReference) Resolve() (SchemaType, bool) {
	e, ok := d.tracker.Get(d.id)
	if !ok {
		return nil, false
	}

	s, ok := e.(SchemaType)
	if !ok {
		return nil, false
	}

	return s, true
}