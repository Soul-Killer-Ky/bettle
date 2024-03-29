// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/im/internal/data/ent/predicate"
	"beetle/app/im/internal/data/ent/synchronizerecord"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SynchronizeRecordQuery is the builder for querying SynchronizeRecord entities.
type SynchronizeRecordQuery struct {
	config
	ctx        *QueryContext
	order      []synchronizerecord.OrderOption
	inters     []Interceptor
	predicates []predicate.SynchronizeRecord
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SynchronizeRecordQuery builder.
func (srq *SynchronizeRecordQuery) Where(ps ...predicate.SynchronizeRecord) *SynchronizeRecordQuery {
	srq.predicates = append(srq.predicates, ps...)
	return srq
}

// Limit the number of records to be returned by this query.
func (srq *SynchronizeRecordQuery) Limit(limit int) *SynchronizeRecordQuery {
	srq.ctx.Limit = &limit
	return srq
}

// Offset to start from.
func (srq *SynchronizeRecordQuery) Offset(offset int) *SynchronizeRecordQuery {
	srq.ctx.Offset = &offset
	return srq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (srq *SynchronizeRecordQuery) Unique(unique bool) *SynchronizeRecordQuery {
	srq.ctx.Unique = &unique
	return srq
}

// Order specifies how the records should be ordered.
func (srq *SynchronizeRecordQuery) Order(o ...synchronizerecord.OrderOption) *SynchronizeRecordQuery {
	srq.order = append(srq.order, o...)
	return srq
}

// First returns the first SynchronizeRecord entity from the query.
// Returns a *NotFoundError when no SynchronizeRecord was found.
func (srq *SynchronizeRecordQuery) First(ctx context.Context) (*SynchronizeRecord, error) {
	nodes, err := srq.Limit(1).All(setContextOp(ctx, srq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{synchronizerecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (srq *SynchronizeRecordQuery) FirstX(ctx context.Context) *SynchronizeRecord {
	node, err := srq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SynchronizeRecord ID from the query.
// Returns a *NotFoundError when no SynchronizeRecord ID was found.
func (srq *SynchronizeRecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = srq.Limit(1).IDs(setContextOp(ctx, srq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{synchronizerecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (srq *SynchronizeRecordQuery) FirstIDX(ctx context.Context) int {
	id, err := srq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SynchronizeRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SynchronizeRecord entity is found.
// Returns a *NotFoundError when no SynchronizeRecord entities are found.
func (srq *SynchronizeRecordQuery) Only(ctx context.Context) (*SynchronizeRecord, error) {
	nodes, err := srq.Limit(2).All(setContextOp(ctx, srq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{synchronizerecord.Label}
	default:
		return nil, &NotSingularError{synchronizerecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (srq *SynchronizeRecordQuery) OnlyX(ctx context.Context) *SynchronizeRecord {
	node, err := srq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SynchronizeRecord ID in the query.
// Returns a *NotSingularError when more than one SynchronizeRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (srq *SynchronizeRecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = srq.Limit(2).IDs(setContextOp(ctx, srq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{synchronizerecord.Label}
	default:
		err = &NotSingularError{synchronizerecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (srq *SynchronizeRecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := srq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SynchronizeRecords.
func (srq *SynchronizeRecordQuery) All(ctx context.Context) ([]*SynchronizeRecord, error) {
	ctx = setContextOp(ctx, srq.ctx, "All")
	if err := srq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SynchronizeRecord, *SynchronizeRecordQuery]()
	return withInterceptors[[]*SynchronizeRecord](ctx, srq, qr, srq.inters)
}

// AllX is like All, but panics if an error occurs.
func (srq *SynchronizeRecordQuery) AllX(ctx context.Context) []*SynchronizeRecord {
	nodes, err := srq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SynchronizeRecord IDs.
func (srq *SynchronizeRecordQuery) IDs(ctx context.Context) (ids []int, err error) {
	if srq.ctx.Unique == nil && srq.path != nil {
		srq.Unique(true)
	}
	ctx = setContextOp(ctx, srq.ctx, "IDs")
	if err = srq.Select(synchronizerecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (srq *SynchronizeRecordQuery) IDsX(ctx context.Context) []int {
	ids, err := srq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (srq *SynchronizeRecordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, srq.ctx, "Count")
	if err := srq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, srq, querierCount[*SynchronizeRecordQuery](), srq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (srq *SynchronizeRecordQuery) CountX(ctx context.Context) int {
	count, err := srq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (srq *SynchronizeRecordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, srq.ctx, "Exist")
	switch _, err := srq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (srq *SynchronizeRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := srq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SynchronizeRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (srq *SynchronizeRecordQuery) Clone() *SynchronizeRecordQuery {
	if srq == nil {
		return nil
	}
	return &SynchronizeRecordQuery{
		config:     srq.config,
		ctx:        srq.ctx.Clone(),
		order:      append([]synchronizerecord.OrderOption{}, srq.order...),
		inters:     append([]Interceptor{}, srq.inters...),
		predicates: append([]predicate.SynchronizeRecord{}, srq.predicates...),
		// clone intermediate query.
		sql:  srq.sql.Clone(),
		path: srq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SynchronizeRecord.Query().
//		GroupBy(synchronizerecord.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (srq *SynchronizeRecordQuery) GroupBy(field string, fields ...string) *SynchronizeRecordGroupBy {
	srq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SynchronizeRecordGroupBy{build: srq}
	grbuild.flds = &srq.ctx.Fields
	grbuild.label = synchronizerecord.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.SynchronizeRecord.Query().
//		Select(synchronizerecord.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (srq *SynchronizeRecordQuery) Select(fields ...string) *SynchronizeRecordSelect {
	srq.ctx.Fields = append(srq.ctx.Fields, fields...)
	sbuild := &SynchronizeRecordSelect{SynchronizeRecordQuery: srq}
	sbuild.label = synchronizerecord.Label
	sbuild.flds, sbuild.scan = &srq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SynchronizeRecordSelect configured with the given aggregations.
func (srq *SynchronizeRecordQuery) Aggregate(fns ...AggregateFunc) *SynchronizeRecordSelect {
	return srq.Select().Aggregate(fns...)
}

func (srq *SynchronizeRecordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range srq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, srq); err != nil {
				return err
			}
		}
	}
	for _, f := range srq.ctx.Fields {
		if !synchronizerecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if srq.path != nil {
		prev, err := srq.path(ctx)
		if err != nil {
			return err
		}
		srq.sql = prev
	}
	return nil
}

func (srq *SynchronizeRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SynchronizeRecord, error) {
	var (
		nodes = []*SynchronizeRecord{}
		_spec = srq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SynchronizeRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SynchronizeRecord{config: srq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, srq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (srq *SynchronizeRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := srq.querySpec()
	_spec.Node.Columns = srq.ctx.Fields
	if len(srq.ctx.Fields) > 0 {
		_spec.Unique = srq.ctx.Unique != nil && *srq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, srq.driver, _spec)
}

func (srq *SynchronizeRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(synchronizerecord.Table, synchronizerecord.Columns, sqlgraph.NewFieldSpec(synchronizerecord.FieldID, field.TypeInt))
	_spec.From = srq.sql
	if unique := srq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if srq.path != nil {
		_spec.Unique = true
	}
	if fields := srq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, synchronizerecord.FieldID)
		for i := range fields {
			if fields[i] != synchronizerecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := srq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := srq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := srq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := srq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (srq *SynchronizeRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(srq.driver.Dialect())
	t1 := builder.Table(synchronizerecord.Table)
	columns := srq.ctx.Fields
	if len(columns) == 0 {
		columns = synchronizerecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if srq.sql != nil {
		selector = srq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if srq.ctx.Unique != nil && *srq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range srq.predicates {
		p(selector)
	}
	for _, p := range srq.order {
		p(selector)
	}
	if offset := srq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := srq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SynchronizeRecordGroupBy is the group-by builder for SynchronizeRecord entities.
type SynchronizeRecordGroupBy struct {
	selector
	build *SynchronizeRecordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (srgb *SynchronizeRecordGroupBy) Aggregate(fns ...AggregateFunc) *SynchronizeRecordGroupBy {
	srgb.fns = append(srgb.fns, fns...)
	return srgb
}

// Scan applies the selector query and scans the result into the given value.
func (srgb *SynchronizeRecordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, srgb.build.ctx, "GroupBy")
	if err := srgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SynchronizeRecordQuery, *SynchronizeRecordGroupBy](ctx, srgb.build, srgb, srgb.build.inters, v)
}

func (srgb *SynchronizeRecordGroupBy) sqlScan(ctx context.Context, root *SynchronizeRecordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(srgb.fns))
	for _, fn := range srgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*srgb.flds)+len(srgb.fns))
		for _, f := range *srgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*srgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SynchronizeRecordSelect is the builder for selecting fields of SynchronizeRecord entities.
type SynchronizeRecordSelect struct {
	*SynchronizeRecordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (srs *SynchronizeRecordSelect) Aggregate(fns ...AggregateFunc) *SynchronizeRecordSelect {
	srs.fns = append(srs.fns, fns...)
	return srs
}

// Scan applies the selector query and scans the result into the given value.
func (srs *SynchronizeRecordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, srs.ctx, "Select")
	if err := srs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SynchronizeRecordQuery, *SynchronizeRecordSelect](ctx, srs.SynchronizeRecordQuery, srs, srs.inters, v)
}

func (srs *SynchronizeRecordSelect) sqlScan(ctx context.Context, root *SynchronizeRecordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(srs.fns))
	for _, fn := range srs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*srs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
