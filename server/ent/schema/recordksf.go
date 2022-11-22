package schema

import (
	"time"
    "entgo.io/ent"
	
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
	
)

// RecordKsf holds the schema definition for the RecordKsf entity.
type RecordKsf struct {
	ent.Schema
}

// Annotations of the RecordKsf.
func (RecordKsf) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "records_ksf"},
    }
}

// Fields of the RecordKsf.
func (RecordKsf) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Time("timestamp").Default(time.Now),
		field.String("player_name"),
		field.String("map_name"),
		field.String("time"),
		field.String("improvement"),
		field.String("server"),
	}
}

// Edges of the RecordKsf.
func (RecordKsf) Edges() []ent.Edge {
	return nil
}
