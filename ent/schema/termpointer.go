package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TermPointer holds the schema definition for the TermPointer entity.
type TermPointer struct {
	ent.Schema
}

// Fields of the TermPointer.
func (TermPointer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("term_id").
			Optional(),

		field.Int("revision_id").
			Optional(),
	}
}

// Edges of the TermPointer.
func (TermPointer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("term", Term.Type).
			Ref("pointers").
			Unique().
			Field("term_id"),

		edge.From("revision", TermRevision.Type).
			Ref("pointers").
			Unique().
			Field("revision_id"),
	}
}
