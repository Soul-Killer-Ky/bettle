// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/im/internal/data/ent/loadrecord"
	"beetle/app/im/internal/data/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoadRecordQuery is the builder for querying LoadRecord entities.
type LoadRecordQuery struct {
	config
	ctx        *QueryContext
	order      []loadrecord.OrderOption
	inters     []Interceptor
	predicates []predicate.LoadRecord
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LoadRecordQuery builder.
func (lrq *LoadRecordQuery) Where(ps ...predicate.LoadRecord) *LoadRecordQuery {
	lrq.predicates = append(lrq.predicates, ps...)
	return lrq
}

// Limit the number of records to be returned by this query.
func (lrq *LoadRecordQuery) Limit(limit int) *LoadRecordQuery {
	lrq.ctx.Limit = &limit
	return lrq
}

// Offset to start from.
func (lrq *LoadRecordQuery) Offset(offset int) *LoadRecordQuery {
	lrq.ctx.Offset = &offset
	return lrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lrq *LoadRecordQuery) Unique(unique bool) *LoadRecordQuery {
	lrq.ctx.Unique = &unique
	return lrq
}

// Order specifies how the records should be ordered.
func (lrq *LoadRecordQuery) Order(o ...loadrecord.OrderOption) *LoadRecordQuery {
	lrq.order = append(lrq.order, o...)
	return lrq
}

// First returns the first LoadRecord entity from the query.
// Returns a *NotFoundError when no LoadRecord was found.
func (lrq *LoadRecordQuery) First(ctx context.Context) (*LoadRecord, error) {
	nodes, err := lrq.Limit(1).All(setContextOp(ctx, lrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loadrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lrq *LoadRecordQuery) FirstX(ctx context.Context) *LoadRecord {
	node, err := lrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LoadRecord ID from the query.
// Returns a *NotFoundError when no LoadRecord ID was found.
func (lrq *LoadRecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lrq.Limit(1).IDs(setContextOp(ctx, lrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loadrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lrq *LoadRecordQuery) FirstIDX(ctx context.Context) int {
	id, err := lrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LoadRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LoadRecord entity is found.
// Returns a *NotFoundError when no LoadRecord entities are found.
func (lrq *LoadRecordQuery) Only(ctx context.Context) (*LoadRecord, error) {
	nodes, err := lrq.Limit(2).All(setContextOp(ctx, lrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loadrecord.Label}
	default:
		return nil, &NotSingularError{loadrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lrq *LoadRecordQuery) OnlyX(ctx context.Context) *LoadRecord {
	node, err := lrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LoadRecord ID in the query.
// Returns a *NotSingularError when more than one LoadRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (lrq *LoadRecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lrq.Limit(2).IDs(setContextOp(ctx, lrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loadrecord.Label}
	default:
		err = &NotSingularError{loadrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lrq *LoadRecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := lrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LoadRecords.
func (lrq *LoadRecordQuery) All(ctx context.Context) ([]*LoadRecord, error) {
	ctx = setContextOp(ctx, lrq.ctx, "All")
	if err := lrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LoadRecord, *LoadRecordQuery]()
	return withInterceptors[[]*LoadRecord](ctx, lrq, qr, lrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lrq *LoadRecordQuery) AllX(ctx context.Context) []*LoadRecord {
	nodes, err := lrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LoadRecord IDs.
func (lrq *LoadRecordQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lrq.ctx.Unique == nil && lrq.path != nil {
		lrq.Unique(true)
	}
	ctx = setContextOp(ctx, lrq.ctx, "IDs")
	if err = lrq.Select(loadrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lrq *LoadRecordQuery) IDsX(ctx context.Context) []int {
	ids, err := lrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lrq *LoadRecordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lrq.ctx, "Count")
	if err := lrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lrq, querierCount[*LoadRecordQuery](), lrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lrq *LoadRecordQuery) CountX(ctx context.Context) int {
	count, err := lrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lrq *LoadRecordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lrq.ctx, "Exist")
	switch _, err := lrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lrq *LoadRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := lrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LoadRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lrq *LoadRecordQuery) Clone() *LoadRecordQuery {
	if lrq == nil {
		return nil
	}
	return &LoadRecordQuery{
		config:     lrq.config,
		ctx:        lrq.ctx.Clone(),
		order:      append([]loadrecord.OrderOption{}, lrq.order...),
		inters:     append([]Interceptor{}, lrq.inters...),
		predicates: append([]predicate.LoadRecord{}, lrq.predicates...),
		// clone intermediate query.
		sql:  lrq.sql.Clone(),
		path: lrq.path,
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
//	client.LoadRecord.Query().
//		GroupBy(loadrecord.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (lrq *LoadRecordQuery) GroupBy(field string, fields ...string) *LoadRecordGroupBy {
	lrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LoadRecordGroupBy{build: lrq}
	grbuild.flds = &lrq.ctx.Fields
	grbuild.label = loadrecord.Label
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
//	client.LoadRecord.Query().
//		Select(loadrecord.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (lrq *LoadRecordQuery) Select(fields ...string) *LoadRecordSelect {
	lrq.ctx.Fields = append(lrq.ctx.Fields, fields...)
	sbuild := &LoadRecordSelect{LoadRecordQuery: lrq}
	sbuild.label = loadrecord.Label
	sbuild.flds, sbuild.scan = &lrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LoadRecordSelect configured with the given aggregations.
func (lrq *LoadRecordQuery) Aggregate(fns ...AggregateFunc) *LoadRecordSelect {
	return lrq.Select().Aggregate(fns...)
}

func (lrq *LoadRecordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lrq); err != nil {
				return err
			}
		}
	}
	for _, f := range lrq.ctx.Fields {
		if !loadrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lrq.path != nil {
		prev, err := lrq.path(ctx)
		if err != nil {
			return err
		}
		lrq.sql = prev
	}
	return nil
}

func (lrq *LoadRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LoadRecord, error) {
	var (
		nodes = []*LoadRecord{}
		_spec = lrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LoadRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LoadRecord{config: lrq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lrq *LoadRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lrq.querySpec()
	_spec.Node.Columns = lrq.ctx.Fields
	if len(lrq.ctx.Fields) > 0 {
		_spec.Unique = lrq.ctx.Unique != nil && *lrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lrq.driver, _spec)
}

func (lrq *LoadRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(loadrecord.Table, loadrecord.Columns, sqlgraph.NewFieldSpec(loadrecord.FieldID, field.TypeInt))
	_spec.From = lrq.sql
	if unique := lrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lrq.path != nil {
		_spec.Unique = true
	}
	if fields := lrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loadrecord.FieldID)
		for i := range fields {
			if fields[i] != loadrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lrq *LoadRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lrq.driver.Dialect())
	t1 := builder.Table(loadrecord.Table)
	columns := lrq.ctx.Fields
	if len(columns) == 0 {
		columns = loadrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lrq.sql != nil {
		selector = lrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lrq.ctx.Unique != nil && *lrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lrq.predicates {
		p(selector)
	}
	for _, p := range lrq.order {
		p(selector)
	}
	if offset := lrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LoadRecordGroupBy is the group-by builder for LoadRecord entities.
type LoadRecordGroupBy struct {
	selector
	build *LoadRecordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lrgb *LoadRecordGroupBy) Aggregate(fns ...AggregateFunc) *LoadRecordGroupBy {
	lrgb.fns = append(lrgb.fns, fns...)
	return lrgb
}

// Scan applies the selector query and scans the result into the given value.
func (lrgb *LoadRecordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lrgb.build.ctx, "GroupBy")
	if err := lrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoadRecordQuery, *LoadRecordGroupBy](ctx, lrgb.build, lrgb, lrgb.build.inters, v)
}

func (lrgb *LoadRecordGroupBy) sqlScan(ctx context.Context, root *LoadRecordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lrgb.fns))
	for _, fn := range lrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lrgb.flds)+len(lrgb.fns))
		for _, f := range *lrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LoadRecordSelect is the builder for selecting fields of LoadRecord entities.
type LoadRecordSelect struct {
	*LoadRecordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lrs *LoadRecordSelect) Aggregate(fns ...AggregateFunc) *LoadRecordSelect {
	lrs.fns = append(lrs.fns, fns...)
	return lrs
}

// Scan applies the selector query and scans the result into the given value.
func (lrs *LoadRecordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lrs.ctx, "Select")
	if err := lrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoadRecordQuery, *LoadRecordSelect](ctx, lrs.LoadRecordQuery, lrs, lrs.inters, v)
}

func (lrs *LoadRecordSelect) sqlScan(ctx context.Context, root *LoadRecordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lrs.fns))
	for _, fn := range lrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
