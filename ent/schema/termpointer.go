package schema

import "entgo.io/ent"

// TermPointer holds the schema definition for the TermPointer entity.
type TermPointer struct {
	ent.Schema
}

// Fields of the TermPointer.
func (TermPointer) Fields() []ent.Field {
	return nil
}

// Edges of the TermPointer.
func (TermPointer) Edges() []ent.Edge {
	return nil
}
