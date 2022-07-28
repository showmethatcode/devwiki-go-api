package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// TermRelated holds the schema definition for the TermRelated entity.
type TermRelated struct {
	ent.Schema
}

// Fields of the TermRelated.
func (TermRelated) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the TermRelated.
func (TermRelated) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subject", Term.Type).
			Ref("subject_id").
			Unique(),

		edge.From("related", Term.Type).
			Ref("related_id").
			Unique(),
	}
}
