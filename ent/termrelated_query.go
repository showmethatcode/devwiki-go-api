// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"devwiki/ent/predicate"
	"devwiki/ent/term"
	"devwiki/ent/termrelated"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TermRelatedQuery is the builder for querying TermRelated entities.
type TermRelatedQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TermRelated
	// eager-loading edges.
	withSubject *TermQuery
	withRelated *TermQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TermRelatedQuery builder.
func (trq *TermRelatedQuery) Where(ps ...predicate.TermRelated) *TermRelatedQuery {
	trq.predicates = append(trq.predicates, ps...)
	return trq
}

// Limit adds a limit step to the query.
func (trq *TermRelatedQuery) Limit(limit int) *TermRelatedQuery {
	trq.limit = &limit
	return trq
}

// Offset adds an offset step to the query.
func (trq *TermRelatedQuery) Offset(offset int) *TermRelatedQuery {
	trq.offset = &offset
	return trq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (trq *TermRelatedQuery) Unique(unique bool) *TermRelatedQuery {
	trq.unique = &unique
	return trq
}

// Order adds an order step to the query.
func (trq *TermRelatedQuery) Order(o ...OrderFunc) *TermRelatedQuery {
	trq.order = append(trq.order, o...)
	return trq
}

// QuerySubject chains the current query on the "subject" edge.
func (trq *TermRelatedQuery) QuerySubject() *TermQuery {
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
			sqlgraph.From(termrelated.Table, termrelated.FieldID, selector),
			sqlgraph.To(term.Table, term.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, termrelated.SubjectTable, termrelated.SubjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(trq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRelated chains the current query on the "related" edge.
func (trq *TermRelatedQuery) QueryRelated() *TermQuery {
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
			sqlgraph.From(termrelated.Table, termrelated.FieldID, selector),
			sqlgraph.To(term.Table, term.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, termrelated.RelatedTable, termrelated.RelatedColumn),
		)
		fromU = sqlgraph.SetNeighbors(trq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TermRelated entity from the query.
// Returns a *NotFoundError when no TermRelated was found.
func (trq *TermRelatedQuery) First(ctx context.Context) (*TermRelated, error) {
	nodes, err := trq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{termrelated.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (trq *TermRelatedQuery) FirstX(ctx context.Context) *TermRelated {
	node, err := trq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TermRelated ID from the query.
// Returns a *NotFoundError when no TermRelated ID was found.
func (trq *TermRelatedQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{termrelated.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (trq *TermRelatedQuery) FirstIDX(ctx context.Context) int {
	id, err := trq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TermRelated entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TermRelated entity is found.
// Returns a *NotFoundError when no TermRelated entities are found.
func (trq *TermRelatedQuery) Only(ctx context.Context) (*TermRelated, error) {
	nodes, err := trq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{termrelated.Label}
	default:
		return nil, &NotSingularError{termrelated.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (trq *TermRelatedQuery) OnlyX(ctx context.Context) *TermRelated {
	node, err := trq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TermRelated ID in the query.
// Returns a *NotSingularError when more than one TermRelated ID is found.
// Returns a *NotFoundError when no entities are found.
func (trq *TermRelatedQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{termrelated.Label}
	default:
		err = &NotSingularError{termrelated.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (trq *TermRelatedQuery) OnlyIDX(ctx context.Context) int {
	id, err := trq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TermRelateds.
func (trq *TermRelatedQuery) All(ctx context.Context) ([]*TermRelated, error) {
	if err := trq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return trq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (trq *TermRelatedQuery) AllX(ctx context.Context) []*TermRelated {
	nodes, err := trq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TermRelated IDs.
func (trq *TermRelatedQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := trq.Select(termrelated.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (trq *TermRelatedQuery) IDsX(ctx context.Context) []int {
	ids, err := trq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (trq *TermRelatedQuery) Count(ctx context.Context) (int, error) {
	if err := trq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return trq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (trq *TermRelatedQuery) CountX(ctx context.Context) int {
	count, err := trq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (trq *TermRelatedQuery) Exist(ctx context.Context) (bool, error) {
	if err := trq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return trq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (trq *TermRelatedQuery) ExistX(ctx context.Context) bool {
	exist, err := trq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TermRelatedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (trq *TermRelatedQuery) Clone() *TermRelatedQuery {
	if trq == nil {
		return nil
	}
	return &TermRelatedQuery{
		config:      trq.config,
		limit:       trq.limit,
		offset:      trq.offset,
		order:       append([]OrderFunc{}, trq.order...),
		predicates:  append([]predicate.TermRelated{}, trq.predicates...),
		withSubject: trq.withSubject.Clone(),
		withRelated: trq.withRelated.Clone(),
		// clone intermediate query.
		sql:    trq.sql.Clone(),
		path:   trq.path,
		unique: trq.unique,
	}
}

// WithSubject tells the query-builder to eager-load the nodes that are connected to
// the "subject" edge. The optional arguments are used to configure the query builder of the edge.
func (trq *TermRelatedQuery) WithSubject(opts ...func(*TermQuery)) *TermRelatedQuery {
	query := &TermQuery{config: trq.config}
	for _, opt := range opts {
		opt(query)
	}
	trq.withSubject = query
	return trq
}

// WithRelated tells the query-builder to eager-load the nodes that are connected to
// the "related" edge. The optional arguments are used to configure the query builder of the edge.
func (trq *TermRelatedQuery) WithRelated(opts ...func(*TermQuery)) *TermRelatedQuery {
	query := &TermQuery{config: trq.config}
	for _, opt := range opts {
		opt(query)
	}
	trq.withRelated = query
	return trq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (trq *TermRelatedQuery) GroupBy(field string, fields ...string) *TermRelatedGroupBy {
	grbuild := &TermRelatedGroupBy{config: trq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := trq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return trq.sqlQuery(ctx), nil
	}
	grbuild.label = termrelated.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (trq *TermRelatedQuery) Select(fields ...string) *TermRelatedSelect {
	trq.fields = append(trq.fields, fields...)
	selbuild := &TermRelatedSelect{TermRelatedQuery: trq}
	selbuild.label = termrelated.Label
	selbuild.flds, selbuild.scan = &trq.fields, selbuild.Scan
	return selbuild
}

func (trq *TermRelatedQuery) prepareQuery(ctx context.Context) error {
	for _, f := range trq.fields {
		if !termrelated.ValidColumn(f) {
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

func (trq *TermRelatedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TermRelated, error) {
	var (
		nodes       = []*TermRelated{}
		withFKs     = trq.withFKs
		_spec       = trq.querySpec()
		loadedTypes = [2]bool{
			trq.withSubject != nil,
			trq.withRelated != nil,
		}
	)
	if trq.withSubject != nil || trq.withRelated != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, termrelated.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TermRelated).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TermRelated{config: trq.config}
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

	if query := trq.withSubject; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TermRelated)
		for i := range nodes {
			if nodes[i].term_subject_id == nil {
				continue
			}
			fk := *nodes[i].term_subject_id
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
				return nil, fmt.Errorf(`unexpected foreign-key "term_subject_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Subject = n
			}
		}
	}

	if query := trq.withRelated; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TermRelated)
		for i := range nodes {
			if nodes[i].term_related_id == nil {
				continue
			}
			fk := *nodes[i].term_related_id
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
				return nil, fmt.Errorf(`unexpected foreign-key "term_related_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Related = n
			}
		}
	}

	return nodes, nil
}

func (trq *TermRelatedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := trq.querySpec()
	_spec.Node.Columns = trq.fields
	if len(trq.fields) > 0 {
		_spec.Unique = trq.unique != nil && *trq.unique
	}
	return sqlgraph.CountNodes(ctx, trq.driver, _spec)
}

func (trq *TermRelatedQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := trq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (trq *TermRelatedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   termrelated.Table,
			Columns: termrelated.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: termrelated.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, termrelated.FieldID)
		for i := range fields {
			if fields[i] != termrelated.FieldID {
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

func (trq *TermRelatedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(trq.driver.Dialect())
	t1 := builder.Table(termrelated.Table)
	columns := trq.fields
	if len(columns) == 0 {
		columns = termrelated.Columns
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

// TermRelatedGroupBy is the group-by builder for TermRelated entities.
type TermRelatedGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (trgb *TermRelatedGroupBy) Aggregate(fns ...AggregateFunc) *TermRelatedGroupBy {
	trgb.fns = append(trgb.fns, fns...)
	return trgb
}

// Scan applies the group-by query and scans the result into the given value.
func (trgb *TermRelatedGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := trgb.path(ctx)
	if err != nil {
		return err
	}
	trgb.sql = query
	return trgb.sqlScan(ctx, v)
}

func (trgb *TermRelatedGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range trgb.fields {
		if !termrelated.ValidColumn(f) {
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

func (trgb *TermRelatedGroupBy) sqlQuery() *sql.Selector {
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

// TermRelatedSelect is the builder for selecting fields of TermRelated entities.
type TermRelatedSelect struct {
	*TermRelatedQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (trs *TermRelatedSelect) Scan(ctx context.Context, v interface{}) error {
	if err := trs.prepareQuery(ctx); err != nil {
		return err
	}
	trs.sql = trs.TermRelatedQuery.sqlQuery(ctx)
	return trs.sqlScan(ctx, v)
}

func (trs *TermRelatedSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := trs.sql.Query()
	if err := trs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}