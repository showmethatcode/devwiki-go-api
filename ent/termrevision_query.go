// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"devwiki/ent/predicate"
	"devwiki/ent/term"
	"devwiki/ent/termrevision"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TermRevisionQuery is the builder for querying TermRevision entities.
type TermRevisionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TermRevision
	// eager-loading edges.
	withTerm *TermQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TermRevisionQuery builder.
func (trq *TermRevisionQuery) Where(ps ...predicate.TermRevision) *TermRevisionQuery {
	trq.predicates = append(trq.predicates, ps...)
	return trq
}

// Limit adds a limit step to the query.
func (trq *TermRevisionQuery) Limit(limit int) *TermRevisionQuery {
	trq.limit = &limit
	return trq
}

// Offset adds an offset step to the query.
func (trq *TermRevisionQuery) Offset(offset int) *TermRevisionQuery {
	trq.offset = &offset
	return trq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (trq *TermRevisionQuery) Unique(unique bool) *TermRevisionQuery {
	trq.unique = &unique
	return trq
}

// Order adds an order step to the query.
func (trq *TermRevisionQuery) Order(o ...OrderFunc) *TermRevisionQuery {
	trq.order = append(trq.order, o...)
	return trq
}

// QueryTerm chains the current query on the "term" edge.
func (trq *TermRevisionQuery) QueryTerm() *TermQuery {
	query := &TermQuery{config: trq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := trq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := trq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(termrevision.Table, termrevision.FieldID, selector),
			sqlgraph.To(term.Table, term.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, termrevision.TermTable, termrevision.TermColumn),
		)
		fromU = sqlgraph.SetNeighbors(trq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TermRevision entity from the query.
// Returns a *NotFoundError when no TermRevision was found.
func (trq *TermRevisionQuery) First(ctx context.Context) (*TermRevision, error) {
	nodes, err := trq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{termrevision.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (trq *TermRevisionQuery) FirstX(ctx context.Context) *TermRevision {
	node, err := trq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TermRevision ID from the query.
// Returns a *NotFoundError when no TermRevision ID was found.
func (trq *TermRevisionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{termrevision.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (trq *TermRevisionQuery) FirstIDX(ctx context.Context) int {
	id, err := trq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TermRevision entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TermRevision entity is found.
// Returns a *NotFoundError when no TermRevision entities are found.
func (trq *TermRevisionQuery) Only(ctx context.Context) (*TermRevision, error) {
	nodes, err := trq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{termrevision.Label}
	default:
		return nil, &NotSingularError{termrevision.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (trq *TermRevisionQuery) OnlyX(ctx context.Context) *TermRevision {
	node, err := trq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TermRevision ID in the query.
// Returns a *NotSingularError when more than one TermRevision ID is found.
// Returns a *NotFoundError when no entities are found.
func (trq *TermRevisionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{termrevision.Label}
	default:
		err = &NotSingularError{termrevision.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (trq *TermRevisionQuery) OnlyIDX(ctx context.Context) int {
	id, err := trq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TermRevisions.
func (trq *TermRevisionQuery) All(ctx context.Context) ([]*TermRevision, error) {
	if err := trq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return trq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (trq *TermRevisionQuery) AllX(ctx context.Context) []*TermRevision {
	nodes, err := trq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TermRevision IDs.
func (trq *TermRevisionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := trq.Select(termrevision.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (trq *TermRevisionQuery) IDsX(ctx context.Context) []int {
	ids, err := trq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (trq *TermRevisionQuery) Count(ctx context.Context) (int, error) {
	if err := trq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return trq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (trq *TermRevisionQuery) CountX(ctx context.Context) int {
	count, err := trq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (trq *TermRevisionQuery) Exist(ctx context.Context) (bool, error) {
	if err := trq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return trq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (trq *TermRevisionQuery) ExistX(ctx context.Context) bool {
	exist, err := trq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TermRevisionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (trq *TermRevisionQuery) Clone() *TermRevisionQuery {
	if trq == nil {
		return nil
	}
	return &TermRevisionQuery{
		config:     trq.config,
		limit:      trq.limit,
		offset:     trq.offset,
		order:      append([]OrderFunc{}, trq.order...),
		predicates: append([]predicate.TermRevision{}, trq.predicates...),
		withTerm:   trq.withTerm.Clone(),
		// clone intermediate query.
		sql:    trq.sql.Clone(),
		path:   trq.path,
		unique: trq.unique,
	}
}

// WithTerm tells the query-builder to eager-load the nodes that are connected to
// the "term" edge. The optional arguments are used to configure the query builder of the edge.
func (trq *TermRevisionQuery) WithTerm(opts ...func(*TermQuery)) *TermRevisionQuery {
	query := &TermQuery{config: trq.config}
	for _, opt := range opts {
		opt(query)
	}
	trq.withTerm = query
	return trq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TermRevision.Query().
//		GroupBy(termrevision.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (trq *TermRevisionQuery) GroupBy(field string, fields ...string) *TermRevisionGroupBy {
	grbuild := &TermRevisionGroupBy{config: trq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := trq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return trq.sqlQuery(ctx), nil
	}
	grbuild.label = termrevision.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	client.TermRevision.Query().
//		Select(termrevision.FieldDescription).
//		Scan(ctx, &v)
//
func (trq *TermRevisionQuery) Select(fields ...string) *TermRevisionSelect {
	trq.fields = append(trq.fields, fields...)
	selbuild := &TermRevisionSelect{TermRevisionQuery: trq}
	selbuild.label = termrevision.Label
	selbuild.flds, selbuild.scan = &trq.fields, selbuild.Scan
	return selbuild
}

func (trq *TermRevisionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range trq.fields {
		if !termrevision.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if trq.path != nil {
		prev, err := trq.path(ctx)
		if err != nil {
			return err
		}
		trq.sql = prev
	}
	return nil
}

func (trq *TermRevisionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TermRevision, error) {
	var (
		nodes       = []*TermRevision{}
		_spec       = trq.querySpec()
		loadedTypes = [1]bool{
			trq.withTerm != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TermRevision).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TermRevision{config: trq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, trq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := trq.withTerm; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TermRevision)
		for i := range nodes {
			fk := nodes[i].TermID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(term.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "term_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Term = n
			}
		}
	}

	return nodes, nil
}

func (trq *TermRevisionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := trq.querySpec()
	_spec.Node.Columns = trq.fields
	if len(trq.fields) > 0 {
		_spec.Unique = trq.unique != nil && *trq.unique
	}
	return sqlgraph.CountNodes(ctx, trq.driver, _spec)
}

func (trq *TermRevisionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := trq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (trq *TermRevisionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   termrevision.Table,
			Columns: termrevision.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: termrevision.FieldID,
			},
		},
		From:   trq.sql,
		Unique: true,
	}
	if unique := trq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := trq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, termrevision.FieldID)
		for i := range fields {
			if fields[i] != termrevision.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := trq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := trq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := trq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := trq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (trq *TermRevisionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(trq.driver.Dialect())
	t1 := builder.Table(termrevision.Table)
	columns := trq.fields
	if len(columns) == 0 {
		columns = termrevision.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if trq.sql != nil {
		selector = trq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if trq.unique != nil && *trq.unique {
		selector.Distinct()
	}
	for _, p := range trq.predicates {
		p(selector)
	}
	for _, p := range trq.order {
		p(selector)
	}
	if offset := trq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := trq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TermRevisionGroupBy is the group-by builder for TermRevision entities.
type TermRevisionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (trgb *TermRevisionGroupBy) Aggregate(fns ...AggregateFunc) *TermRevisionGroupBy {
	trgb.fns = append(trgb.fns, fns...)
	return trgb
}

// Scan applies the group-by query and scans the result into the given value.
func (trgb *TermRevisionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := trgb.path(ctx)
	if err != nil {
		return err
	}
	trgb.sql = query
	return trgb.sqlScan(ctx, v)
}

func (trgb *TermRevisionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range trgb.fields {
		if !termrevision.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := trgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (trgb *TermRevisionGroupBy) sqlQuery() *sql.Selector {
	selector := trgb.sql.Select()
	aggregation := make([]string, 0, len(trgb.fns))
	for _, fn := range trgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(trgb.fields)+len(trgb.fns))
		for _, f := range trgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(trgb.fields...)...)
}

// TermRevisionSelect is the builder for selecting fields of TermRevision entities.
type TermRevisionSelect struct {
	*TermRevisionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (trs *TermRevisionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := trs.prepareQuery(ctx); err != nil {
		return err
	}
	trs.sql = trs.TermRevisionQuery.sqlQuery(ctx)
	return trs.sqlScan(ctx, v)
}

func (trs *TermRevisionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := trs.sql.Query()
	if err := trs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}