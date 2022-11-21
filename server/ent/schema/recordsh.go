package schema

import (
	"time"
    "entgo.io/ent"
	
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
	
)

// RecordSh holds the schema definition for the RecordSh entity.
type RecordSh struct {
	ent.Schema
}

// Annotations of the RecordSh.
func (RecordSh) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "records_sh"},
    }
}


// Fields of the RecordSh.
func (RecordSh) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Time("timestamp").Default(time.Now),
		field.String("player_name"),
		field.String("player_id"),
		field.String("type"),
		field.Int("track"),
		field.String("map_name"),
		field.String("time"),
		field.String("improvement"),
		field.String("server"),
	}
}


// Edges of the RecordSh.
func (RecordSh) Edges() []ent.Edge {
	return nil
}
