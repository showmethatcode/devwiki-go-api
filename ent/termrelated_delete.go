// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"devwiki/ent/predicate"
	"devwiki/ent/termrelated"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TermRelatedDelete is the builder for deleting a TermRelated entity.
type TermRelatedDelete struct {
	config
	hooks    []Hook
	mutation *TermRelatedMutation
}

// Where appends a list predicates to the TermRelatedDelete builder.
func (trd *TermRelatedDelete) Where(ps ...predicate.TermRelated) *TermRelatedDelete {
	trd.mutation.Where(ps...)
	return trd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (trd *TermRelatedDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(trd.hooks) == 0 {
		affected, err = trd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TermRelatedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			trd.mutation = mutation
			affected, err = trd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(trd.hooks) - 1; i >= 0; i-- {
			if trd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = trd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, trd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (trd *TermRelatedDelete) ExecX(ctx context.Context) int {
	n, err := trd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (trd *TermRelatedDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: termrelated.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: termrelated.FieldID,
			},
		},
	}
	if ps := trd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, trd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TermRelatedDeleteOne is the builder for deleting a single TermRelated entity.
type TermRelatedDeleteOne struct {
	trd *TermRelatedDelete
}

// Exec executes the deletion query.
func (trdo *TermRelatedDeleteOne) Exec(ctx context.Context) error {
	n, err := trdo.trd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{termrelated.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (trdo *TermRelatedDeleteOne) ExecX(ctx context.Context) {
	trdo.trd.ExecX(ctx)
}
