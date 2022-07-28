// Code generated by ent, DO NOT EDIT.

package ent

import (
	"devwiki/ent/term"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Term is the model entity for the Term schema.
type Term struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TermQuery when eager-loading is set.
	Edges TermEdges `json:"edges"`
}

// TermEdges holds the relations/edges for other nodes in the graph.
type TermEdges struct {
	// Revisions holds the value of the revisions edge.
	Revisions []*TermRevision `json:"revisions,omitempty"`
	// Pointers holds the value of the pointers edge.
	Pointers []*TermPointer `json:"pointers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RevisionsOrErr returns the Revisions value or an error if the edge
// was not loaded in eager-loading.
func (e TermEdges) RevisionsOrErr() ([]*TermRevision, error) {
	if e.loadedTypes[0] {
		return e.Revisions, nil
	}
	return nil, &NotLoadedError{edge: "revisions"}
}

// PointersOrErr returns the Pointers value or an error if the edge
// was not loaded in eager-loading.
func (e TermEdges) PointersOrErr() ([]*TermPointer, error) {
	if e.loadedTypes[1] {
		return e.Pointers, nil
	}
	return nil, &NotLoadedError{edge: "pointers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Term) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case term.FieldID:
			values[i] = new(sql.NullInt64)
		case term.FieldName:
			values[i] = new(sql.NullString)
		case term.FieldCreatedAt, term.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Term", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Term fields.
func (t *Term) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case term.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case term.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case term.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case term.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryRevisions queries the "revisions" edge of the Term entity.
func (t *Term) QueryRevisions() *TermRevisionQuery {
	return (&TermClient{config: t.config}).QueryRevisions(t)
}

// QueryPointers queries the "pointers" edge of the Term entity.
func (t *Term) QueryPointers() *TermPointerQuery {
	return (&TermClient{config: t.config}).QueryPointers(t)
}

// Update returns a builder for updating this Term.
// Note that you need to call Term.Unwrap() before calling this method if this Term
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Term) Update() *TermUpdateOne {
	return (&TermClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Term entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Term) Unwrap() *Term {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Term is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Term) String() string {
	var builder strings.Builder
	builder.WriteString("Term(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Terms is a parsable slice of Term.
type Terms []*Term

func (t Terms) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
