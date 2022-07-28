package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TermRevision holds the schema definition for the TermRevision entity.
type TermRevision struct {
	ent.Schema
}

// Fields of the TermRevision.
func (TermRevision) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").
			Optional(),

		field.Int("term_id").
			Optional(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the TermRevision.
func (TermRevision) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pointers", TermPointer.Type),

		edge.From("term", Term.Type).
			Ref("revisions").
			Unique().
			Field("term_id"),
	}
}
