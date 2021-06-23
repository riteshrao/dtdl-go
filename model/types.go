package model

type (
	// EntityType with basic DTDL properties.
	EntityType interface {
		// @id of the entity
		ID() string
		// @types assigned to the entity.
		Types() []string
		// Entity display name.
		DisplayName() string
		// Entity description.
		Description() string
		// Developer comment.
		Comment() string
		// LocaleDisplayName returns a localized representation of the display name, if available.
		// If a value for the locale is not found, it returns an empty string.
		LocaleDisplayName(locale string) string
		// LocaleDescription returns a localized representation of the description, if available.
		// If a value for the locale is not found, it returns an empty string.
		LocaleDescription(locale string) string
	}

	// NamedEntityType is an entity that has a qualified name
	NamedEntityType interface {
		EntityType
		// Assigned name.
		Name() string
	}

	// CapabilityType is a named entity.
	CapabilityType interface {
		NamedEntityType
	}

	// SchemaType represents a primitive or complex schema
	SchemaType interface {
		EntityType
		Class() SchemaClass
	}

	// SchemaEntityType is a entity with a schema
	SchemaEntityType interface {
		EntityType
		// Schema assigned to the capability
		Schema() SchemaType
	}

	// Internal entity tracker interface used to track entities
	entityTracker interface {
		// Adds an entity to track.
		Add(entity EntityType) error
		// Gets an entity to track.
		Get(id string) (EntityType, bool)
	}
)
