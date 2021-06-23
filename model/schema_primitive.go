package model

var (
	// BooleanSchema definition
	BooleanSchema = &Schema{
		Entity: Entity{
			id:           "boolean",
			displayNames: map[string]string{"en": "Boolean"},
			descriptions: map[string]string{"en": "Boolean"},
			types:        []string{"Boolean"},
		},
		class: BooleanSchemaClass,
	}
	// DateSchema definition
	DateSchema = &Schema{
		Entity: Entity{
			id:           "date",
			displayNames: map[string]string{"en": "Date"},
			descriptions: map[string]string{"en": "Date"},
			types:        []string{"Date"},
		},
		class: DateSchemaClass,
	}
	// DateTimeSchema definition
	DateTimeSchema = &Schema{
		Entity: Entity{
			id:           "dateTime",
			displayNames: map[string]string{"en": "DateTime"},
			descriptions: map[string]string{"en": "DateTime"},
			types:        []string{"DateTime"},
		},
		class: DateTimeSchemaClass,
	}
	// DoubleSchema definition
	DoubleSchema = &Schema{
		Entity: Entity{
			id:           "double",
			displayNames: map[string]string{"en": "Double"},
			descriptions: map[string]string{"en": "Double"},
			types:        []string{"Double"},
		},
		class: DoubleSchemaClass,
	}
	// GeopointSchema definition
	GeopointSchema = &Schema{
		Entity: Entity{
			id:           "geopoint",
			displayNames: map[string]string{"en": "Geopoint"},
			descriptions: map[string]string{"en": "Geopoint"},
			types:        []string{"Geopoint"},
		},
		class: GeopointSchemaClass,
	}
	// IntegerSchema defnition
	IntegerSchema = &Schema{
		Entity: Entity{
			id:           "integer",
			displayNames: map[string]string{"en": "Integer"},
			descriptions: map[string]string{"en": "Integer"},
			types:        []string{"Integer"},
		},
	}
	// LongSchema definition
	LongSchema = &Schema{
		Entity: Entity{
			id:           "long",
			displayNames: map[string]string{"en": "Long"},
			descriptions: map[string]string{"en": "Long"},
			types:        []string{"Boolean"},
		},
		class: IntegerSchemaClass,
	}
	// StringSchema definition
	StringSchema = &Schema{
		Entity: Entity{
			id:           "string",
			displayNames: map[string]string{"en": "String"},
			descriptions: map[string]string{"en": "String"},
			types:        []string{"String"},
		},
		class: StringSchemaClass,
	}
	// UnsupportedSchema definition
	UnsupportedSchema = &Schema{
		Entity: Entity{
			id:           "unsupported",
			displayNames: map[string]string{"en": "Unsupported"},
			descriptions: map[string]string{"en": "Unsupported"},
			types:        []string{"Unsupported"},
		},
		class: UnsupportedSchemaClass,
	}
)