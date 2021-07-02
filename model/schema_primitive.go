package model

var (
	// BooleanSchema definition
	BooleanSchema = &Schema{
		Entity: Entity{
			id:           "boolean",
			displayNames: map[string]interface{}{"en": "Boolean"},
			descriptions: map[string]interface{}{"en": "Boolean"},
			types:        []string{"Boolean"},
		},
		class: BooleanSchemaClass,
	}
	// DateSchema definition
	DateSchema = &Schema{
		Entity: Entity{
			id:           "date",
			displayNames: map[string]interface{}{"en": "Date"},
			descriptions: map[string]interface{}{"en": "Date"},
			types:        []string{"Date"},
		},
		class: DateSchemaClass,
	}
	// DateTimeSchema definition
	DateTimeSchema = &Schema{
		Entity: Entity{
			id:           "dateTime",
			displayNames: map[string]interface{}{"en": "DateTime"},
			descriptions: map[string]interface{}{"en": "DateTime"},
			types:        []string{"DateTime"},
		},
		class: DateTimeSchemaClass,
	}
	// DoubleSchema definition
	DoubleSchema = &Schema{
		Entity: Entity{
			id:           "double",
			displayNames: map[string]interface{}{"en": "Double"},
			descriptions: map[string]interface{}{"en": "Double"},
			types:        []string{"Double"},
		},
		class: DoubleSchemaClass,
	}
	// GeopointSchema definition
	GeopointSchema = &Schema{
		Entity: Entity{
			id:           "geopoint",
			displayNames: map[string]interface{}{"en": "Geopoint"},
			descriptions: map[string]interface{}{"en": "Geopoint"},
			types:        []string{"Geopoint"},
		},
		class: GeopointSchemaClass,
	}
	// IntegerSchema defnition
	IntegerSchema = &Schema{
		Entity: Entity{
			id:           "integer",
			displayNames: map[string]interface{}{"en": "Integer"},
			descriptions: map[string]interface{}{"en": "Integer"},
			types:        []string{"Integer"},
		},
	}
	// LongSchema definition
	LongSchema = &Schema{
		Entity: Entity{
			id:           "long",
			displayNames: map[string]interface{}{"en": "Long"},
			descriptions: map[string]interface{}{"en": "Long"},
			types:        []string{"Boolean"},
		},
		class: IntegerSchemaClass,
	}
	// StringSchema definition
	StringSchema = &Schema{
		Entity: Entity{
			id:           "string",
			displayNames: map[string]interface{}{"en": "String"},
			descriptions: map[string]interface{}{"en": "String"},
			types:        []string{"String"},
		},
		class: StringSchemaClass,
	}
	// UnsupportedSchema definition
	UnsupportedSchema = &Schema{
		Entity: Entity{
			id:           "unsupported",
			displayNames: map[string]interface{}{"en": "Unsupported"},
			descriptions: map[string]interface{}{"en": "Unsupported"},
			types:        []string{"Unsupported"},
		},
		class: UnsupportedSchemaClass,
	}
)