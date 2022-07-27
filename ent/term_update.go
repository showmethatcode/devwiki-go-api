// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"devwiki/ent/predicate"
	"devwiki/ent/term"
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