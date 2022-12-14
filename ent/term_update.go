// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"devwiki/ent/predicate"
	"devwiki/ent/term"
	"devwiki/ent/termpointer"
	"devwiki/ent/termrelated"
	"devwiki/ent/termrevision"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TermUpdate is the builder for updating Term entities.
type TermUpdate struct {
	config
	hooks    []Hook
	mutation *TermMutation
}

// Where appends a list predicates to the TermUpdate builder.
func (tu *TermUpdate) Where(ps ...predicate.Term) *TermUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TermUpdate) SetName(s string) *TermUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TermUpdate) SetUpdatedAt(t time.Time) *TermUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// AddRevisionIDs adds the "revisions" edge to the TermRevision entity by IDs.
func (tu *TermUpdate) AddRevisionIDs(ids ...int) *TermUpdate {
	tu.mutation.AddRevisionIDs(ids...)
	return tu
}

// AddRevisions adds the "revisions" edges to the TermRevision entity.
func (tu *TermUpdate) AddRevisions(t ...*TermRevision) *TermUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddRevisionIDs(ids...)
}

// AddPointerIDs adds the "pointers" edge to the TermPointer entity by IDs.
func (tu *TermUpdate) AddPointerIDs(ids ...int) *TermUpdate {
	tu.mutation.AddPointerIDs(ids...)
	return tu
}

// AddPointers adds the "pointers" edges to the TermPointer entity.
func (tu *TermUpdate) AddPointers(t ...*TermPointer) *TermUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddPointerIDs(ids...)
}

// AddSubjectIDIDs adds the "subject_id" edge to the TermRelated entity by IDs.
func (tu *TermUpdate) AddSubjectIDIDs(ids ...int) *TermUpdate {
	tu.mutation.AddSubjectIDIDs(ids...)
	return tu
}

// AddSubjectID adds the "subject_id" edges to the TermRelated entity.
func (tu *TermUpdate) AddSubjectID(t ...*TermRelated) *TermUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddSubjectIDIDs(ids...)
}

// AddRelatedIDIDs adds the "related_id" edge to the TermRelated entity by IDs.
func (tu *TermUpdate) AddRelatedIDIDs(ids ...int) *TermUpdate {
	tu.mutation.AddRelatedIDIDs(ids...)
	return tu
}

// AddRelatedID adds the "related_id" edges to the TermRelated entity.
func (tu *TermUpdate) AddRelatedID(t ...*TermRelated) *TermUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddRelatedIDIDs(ids...)
}

// Mutation returns the TermMutation object of the builder.
func (tu *TermUpdate) Mutation() *TermMutation {
	return tu.mutation
}

// ClearRevisions clears all "revisions" edges to the TermRevision entity.
func (tu *TermUpdate) ClearRevisions() *TermUpdate {
	tu.mutation.ClearRevisions()
	return tu
}

// RemoveRevisionIDs removes the "revisions" edge to TermRevision entities by IDs.
func (tu *TermUpdate) RemoveRevisionIDs(ids ...int) *TermUpdate {
	tu.mutation.RemoveRevisionIDs(ids...)
	return tu
}

// RemoveRevisions removes "revisions" edges to TermRevision entities.
func (tu *TermUpdate) RemoveRevisions(t ...*TermRevision) *TermUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveRevisionIDs(ids...)
}

// ClearPointers clears all "pointers" edges to the TermPointer entity.
func (tu *TermUpdate) ClearPointers() *TermUpdate {
	tu.mutation.ClearPointers()
	return tu
}

// RemovePointerIDs removes the "pointers" edge to TermPointer entities by IDs.
func (tu *TermUpdate) RemovePointerIDs(ids ...int) *TermUpdate {
	tu.mutation.RemovePointerIDs(ids...)
	return tu
}

// RemovePointers removes "pointers" edges to TermPointer entities.
func (tu *TermUpdate) RemovePointers(t ...*TermPointer) *TermUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemovePointerIDs(ids...)
}

// ClearSubjectID clears all "subject_id" edges to the TermRelated entity.
func (tu *TermUpdate) ClearSubjectID() *TermUpdate {
	tu.mutation.ClearSubjectID()
	return tu
}

// RemoveSubjectIDIDs removes the "subject_id" edge to TermRelated entities by IDs.
func (tu *TermUpdate) RemoveSubjectIDIDs(ids ...int) *TermUpdate {
	tu.mutation.RemoveSubjectIDIDs(ids...)
	return tu
}

// RemoveSubjectID removes "subject_id" edges to TermRelated entities.
func (tu *TermUpdate) RemoveSubjectID(t ...*TermRelated) *TermUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveSubjectIDIDs(ids...)
}

// ClearRelatedID clears all "related_id" edges to the TermRelated entity.
func (tu *TermUpdate) ClearRelatedID() *TermUpdate {
	tu.mutation.ClearRelatedID()
	return tu
}

// RemoveRelatedIDIDs removes the "related_id" edge to TermRelated entities by IDs.
func (tu *TermUpdate) RemoveRelatedIDIDs(ids ...int) *TermUpdate {
	tu.mutation.RemoveRelatedIDIDs(ids...)
	return tu
}

// RemoveRelatedID removes "related_id" edges to TermRelated entities.
func (tu *TermUpdate) RemoveRelatedID(t ...*TermRelated) *TermUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveRelatedIDIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TermUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TermMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TermUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TermUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TermUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TermUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := term.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

func (tu *TermUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   term.Table,
			Columns: term.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: term.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: term.FieldName,
		})
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: term.FieldUpdatedAt,
		})
	}
	if tu.mutation.RevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RevisionsTable,
			Columns: []string{term.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrevision.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedRevisionsIDs(); len(nodes) > 0 && !tu.mutation.RevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RevisionsTable,
			Columns: []string{term.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrevision.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RevisionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RevisionsTable,
			Columns: []string{term.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrevision.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.PointersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.PointersTable,
			Columns: []string{term.PointersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termpointer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedPointersIDs(); len(nodes) > 0 && !tu.mutation.PointersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.PointersTable,
			Columns: []string{term.PointersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termpointer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.PointersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.PointersTable,
			Columns: []string{term.PointersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termpointer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.SubjectIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.SubjectIDTable,
			Columns: []string{term.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedSubjectIDIDs(); len(nodes) > 0 && !tu.mutation.SubjectIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.SubjectIDTable,
			Columns: []string{term.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SubjectIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.SubjectIDTable,
			Columns: []string{term.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.RelatedIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RelatedIDTable,
			Columns: []string{term.RelatedIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedRelatedIDIDs(); len(nodes) > 0 && !tu.mutation.RelatedIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RelatedIDTable,
			Columns: []string{term.RelatedIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RelatedIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RelatedIDTable,
			Columns: []string{term.RelatedIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{term.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TermUpdateOne is the builder for updating a single Term entity.
type TermUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TermMutation
}

// SetName sets the "name" field.
func (tuo *TermUpdateOne) SetName(s string) *TermUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TermUpdateOne) SetUpdatedAt(t time.Time) *TermUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// AddRevisionIDs adds the "revisions" edge to the TermRevision entity by IDs.
func (tuo *TermUpdateOne) AddRevisionIDs(ids ...int) *TermUpdateOne {
	tuo.mutation.AddRevisionIDs(ids...)
	return tuo
}

// AddRevisions adds the "revisions" edges to the TermRevision entity.
func (tuo *TermUpdateOne) AddRevisions(t ...*TermRevision) *TermUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddRevisionIDs(ids...)
}

// AddPointerIDs adds the "pointers" edge to the TermPointer entity by IDs.
func (tuo *TermUpdateOne) AddPointerIDs(ids ...int) *TermUpdateOne {
	tuo.mutation.AddPointerIDs(ids...)
	return tuo
}

// AddPointers adds the "pointers" edges to the TermPointer entity.
func (tuo *TermUpdateOne) AddPointers(t ...*TermPointer) *TermUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddPointerIDs(ids...)
}

// AddSubjectIDIDs adds the "subject_id" edge to the TermRelated entity by IDs.
func (tuo *TermUpdateOne) AddSubjectIDIDs(ids ...int) *TermUpdateOne {
	tuo.mutation.AddSubjectIDIDs(ids...)
	return tuo
}

// AddSubjectID adds the "subject_id" edges to the TermRelated entity.
func (tuo *TermUpdateOne) AddSubjectID(t ...*TermRelated) *TermUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddSubjectIDIDs(ids...)
}

// AddRelatedIDIDs adds the "related_id" edge to the TermRelated entity by IDs.
func (tuo *TermUpdateOne) AddRelatedIDIDs(ids ...int) *TermUpdateOne {
	tuo.mutation.AddRelatedIDIDs(ids...)
	return tuo
}

// AddRelatedID adds the "related_id" edges to the TermRelated entity.
func (tuo *TermUpdateOne) AddRelatedID(t ...*TermRelated) *TermUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddRelatedIDIDs(ids...)
}

// Mutation returns the TermMutation object of the builder.
func (tuo *TermUpdateOne) Mutation() *TermMutation {
	return tuo.mutation
}

// ClearRevisions clears all "revisions" edges to the TermRevision entity.
func (tuo *TermUpdateOne) ClearRevisions() *TermUpdateOne {
	tuo.mutation.ClearRevisions()
	return tuo
}

// RemoveRevisionIDs removes the "revisions" edge to TermRevision entities by IDs.
func (tuo *TermUpdateOne) RemoveRevisionIDs(ids ...int) *TermUpdateOne {
	tuo.mutation.RemoveRevisionIDs(ids...)
	return tuo
}

// RemoveRevisions removes "revisions" edges to TermRevision entities.
func (tuo *TermUpdateOne) RemoveRevisions(t ...*TermRevision) *TermUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveRevisionIDs(ids...)
}

// ClearPointers clears all "pointers" edges to the TermPointer entity.
func (tuo *TermUpdateOne) ClearPointers() *TermUpdateOne {
	tuo.mutation.ClearPointers()
	return tuo
}

// RemovePointerIDs removes the "pointers" edge to TermPointer entities by IDs.
func (tuo *TermUpdateOne) RemovePointerIDs(ids ...int) *TermUpdateOne {
	tuo.mutation.RemovePointerIDs(ids...)
	return tuo
}

// RemovePointers removes "pointers" edges to TermPointer entities.
func (tuo *TermUpdateOne) RemovePointers(t ...*TermPointer) *TermUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemovePointerIDs(ids...)
}

// ClearSubjectID clears all "subject_id" edges to the TermRelated entity.
func (tuo *TermUpdateOne) ClearSubjectID() *TermUpdateOne {
	tuo.mutation.ClearSubjectID()
	return tuo
}

// RemoveSubjectIDIDs removes the "subject_id" edge to TermRelated entities by IDs.
func (tuo *TermUpdateOne) RemoveSubjectIDIDs(ids ...int) *TermUpdateOne {
	tuo.mutation.RemoveSubjectIDIDs(ids...)
	return tuo
}

// RemoveSubjectID removes "subject_id" edges to TermRelated entities.
func (tuo *TermUpdateOne) RemoveSubjectID(t ...*TermRelated) *TermUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveSubjectIDIDs(ids...)
}

// ClearRelatedID clears all "related_id" edges to the TermRelated entity.
func (tuo *TermUpdateOne) ClearRelatedID() *TermUpdateOne {
	tuo.mutation.ClearRelatedID()
	return tuo
}

// RemoveRelatedIDIDs removes the "related_id" edge to TermRelated entities by IDs.
func (tuo *TermUpdateOne) RemoveRelatedIDIDs(ids ...int) *TermUpdateOne {
	tuo.mutation.RemoveRelatedIDIDs(ids...)
	return tuo
}

// RemoveRelatedID removes "related_id" edges to TermRelated entities.
func (tuo *TermUpdateOne) RemoveRelatedID(t ...*TermRelated) *TermUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveRelatedIDIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TermUpdateOne) Select(field string, fields ...string) *TermUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Term entity.
func (tuo *TermUpdateOne) Save(ctx context.Context) (*Term, error) {
	var (
		err  error
		node *Term
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TermMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Term)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TermMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TermUpdateOne) SaveX(ctx context.Context) *Term {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TermUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TermUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TermUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := term.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

func (tuo *TermUpdateOne) sqlSave(ctx context.Context) (_node *Term, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   term.Table,
			Columns: term.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: term.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Term.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, term.FieldID)
		for _, f := range fields {
			if !term.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != term.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: term.FieldName,
		})
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: term.FieldUpdatedAt,
		})
	}
	if tuo.mutation.RevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RevisionsTable,
			Columns: []string{term.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrevision.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedRevisionsIDs(); len(nodes) > 0 && !tuo.mutation.RevisionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RevisionsTable,
			Columns: []string{term.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrevision.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RevisionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RevisionsTable,
			Columns: []string{term.RevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrevision.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.PointersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.PointersTable,
			Columns: []string{term.PointersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termpointer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedPointersIDs(); len(nodes) > 0 && !tuo.mutation.PointersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.PointersTable,
			Columns: []string{term.PointersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termpointer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.PointersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.PointersTable,
			Columns: []string{term.PointersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termpointer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.SubjectIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.SubjectIDTable,
			Columns: []string{term.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedSubjectIDIDs(); len(nodes) > 0 && !tuo.mutation.SubjectIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.SubjectIDTable,
			Columns: []string{term.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SubjectIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.SubjectIDTable,
			Columns: []string{term.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.RelatedIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RelatedIDTable,
			Columns: []string{term.RelatedIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedRelatedIDIDs(); len(nodes) > 0 && !tuo.mutation.RelatedIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RelatedIDTable,
			Columns: []string{term.RelatedIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RelatedIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.RelatedIDTable,
			Columns: []string{term.RelatedIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: termrelated.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Term{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{term.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
