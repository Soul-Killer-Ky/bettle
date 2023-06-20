// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/filesystem/internal/data/ent/predicate"
	"beetle/app/filesystem/internal/data/ent/quote"
	"beetle/app/filesystem/internal/data/ent/storage"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StorageUpdate is the builder for updating Storage entities.
type StorageUpdate struct {
	config
	hooks    []Hook
	mutation *StorageMutation
}

// Where appends a list predicates to the StorageUpdate builder.
func (su *StorageUpdate) Where(ps ...predicate.Storage) *StorageUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *StorageUpdate) SetDeletedAt(t time.Time) *StorageUpdate {
	su.mutation.SetDeletedAt(t)
	return su
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (su *StorageUpdate) SetNillableDeletedAt(t *time.Time) *StorageUpdate {
	if t != nil {
		su.SetDeletedAt(*t)
	}
	return su
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (su *StorageUpdate) ClearDeletedAt() *StorageUpdate {
	su.mutation.ClearDeletedAt()
	return su
}

// SetType sets the "type" field.
func (su *StorageUpdate) SetType(i int32) *StorageUpdate {
	su.mutation.ResetType()
	su.mutation.SetType(i)
	return su
}

// AddType adds i to the "type" field.
func (su *StorageUpdate) AddType(i int32) *StorageUpdate {
	su.mutation.AddType(i)
	return su
}

// SetPath sets the "path" field.
func (su *StorageUpdate) SetPath(s string) *StorageUpdate {
	su.mutation.SetPath(s)
	return su
}

// SetMd5 sets the "md5" field.
func (su *StorageUpdate) SetMd5(s string) *StorageUpdate {
	su.mutation.SetMd5(s)
	return su
}

// AddQuoteIDs adds the "quotes" edge to the Quote entity by IDs.
func (su *StorageUpdate) AddQuoteIDs(ids ...int) *StorageUpdate {
	su.mutation.AddQuoteIDs(ids...)
	return su
}

// AddQuotes adds the "quotes" edges to the Quote entity.
func (su *StorageUpdate) AddQuotes(q ...*Quote) *StorageUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return su.AddQuoteIDs(ids...)
}

// Mutation returns the StorageMutation object of the builder.
func (su *StorageUpdate) Mutation() *StorageMutation {
	return su.mutation
}

// ClearQuotes clears all "quotes" edges to the Quote entity.
func (su *StorageUpdate) ClearQuotes() *StorageUpdate {
	su.mutation.ClearQuotes()
	return su
}

// RemoveQuoteIDs removes the "quotes" edge to Quote entities by IDs.
func (su *StorageUpdate) RemoveQuoteIDs(ids ...int) *StorageUpdate {
	su.mutation.RemoveQuoteIDs(ids...)
	return su
}

// RemoveQuotes removes "quotes" edges to Quote entities.
func (su *StorageUpdate) RemoveQuotes(q ...*Quote) *StorageUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return su.RemoveQuoteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StorageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StorageUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StorageUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StorageUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StorageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(storage.Table, storage.Columns, sqlgraph.NewFieldSpec(storage.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.SetField(storage.FieldDeletedAt, field.TypeTime, value)
	}
	if su.mutation.DeletedAtCleared() {
		_spec.ClearField(storage.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.SetField(storage.FieldType, field.TypeInt32, value)
	}
	if value, ok := su.mutation.AddedType(); ok {
		_spec.AddField(storage.FieldType, field.TypeInt32, value)
	}
	if value, ok := su.mutation.Path(); ok {
		_spec.SetField(storage.FieldPath, field.TypeString, value)
	}
	if value, ok := su.mutation.Md5(); ok {
		_spec.SetField(storage.FieldMd5, field.TypeString, value)
	}
	if su.mutation.QuotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   storage.QuotesTable,
			Columns: []string{storage.QuotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedQuotesIDs(); len(nodes) > 0 && !su.mutation.QuotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   storage.QuotesTable,
			Columns: []string{storage.QuotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.QuotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   storage.QuotesTable,
			Columns: []string{storage.QuotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StorageUpdateOne is the builder for updating a single Storage entity.
type StorageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StorageMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *StorageUpdateOne) SetDeletedAt(t time.Time) *StorageUpdateOne {
	suo.mutation.SetDeletedAt(t)
	return suo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suo *StorageUpdateOne) SetNillableDeletedAt(t *time.Time) *StorageUpdateOne {
	if t != nil {
		suo.SetDeletedAt(*t)
	}
	return suo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suo *StorageUpdateOne) ClearDeletedAt() *StorageUpdateOne {
	suo.mutation.ClearDeletedAt()
	return suo
}

// SetType sets the "type" field.
func (suo *StorageUpdateOne) SetType(i int32) *StorageUpdateOne {
	suo.mutation.ResetType()
	suo.mutation.SetType(i)
	return suo
}

// AddType adds i to the "type" field.
func (suo *StorageUpdateOne) AddType(i int32) *StorageUpdateOne {
	suo.mutation.AddType(i)
	return suo
}

// SetPath sets the "path" field.
func (suo *StorageUpdateOne) SetPath(s string) *StorageUpdateOne {
	suo.mutation.SetPath(s)
	return suo
}

// SetMd5 sets the "md5" field.
func (suo *StorageUpdateOne) SetMd5(s string) *StorageUpdateOne {
	suo.mutation.SetMd5(s)
	return suo
}

// AddQuoteIDs adds the "quotes" edge to the Quote entity by IDs.
func (suo *StorageUpdateOne) AddQuoteIDs(ids ...int) *StorageUpdateOne {
	suo.mutation.AddQuoteIDs(ids...)
	return suo
}

// AddQuotes adds the "quotes" edges to the Quote entity.
func (suo *StorageUpdateOne) AddQuotes(q ...*Quote) *StorageUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return suo.AddQuoteIDs(ids...)
}

// Mutation returns the StorageMutation object of the builder.
func (suo *StorageUpdateOne) Mutation() *StorageMutation {
	return suo.mutation
}

// ClearQuotes clears all "quotes" edges to the Quote entity.
func (suo *StorageUpdateOne) ClearQuotes() *StorageUpdateOne {
	suo.mutation.ClearQuotes()
	return suo
}

// RemoveQuoteIDs removes the "quotes" edge to Quote entities by IDs.
func (suo *StorageUpdateOne) RemoveQuoteIDs(ids ...int) *StorageUpdateOne {
	suo.mutation.RemoveQuoteIDs(ids...)
	return suo
}

// RemoveQuotes removes "quotes" edges to Quote entities.
func (suo *StorageUpdateOne) RemoveQuotes(q ...*Quote) *StorageUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return suo.RemoveQuoteIDs(ids...)
}

// Where appends a list predicates to the StorageUpdate builder.
func (suo *StorageUpdateOne) Where(ps ...predicate.Storage) *StorageUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StorageUpdateOne) Select(field string, fields ...string) *StorageUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Storage entity.
func (suo *StorageUpdateOne) Save(ctx context.Context) (*Storage, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StorageUpdateOne) SaveX(ctx context.Context) *Storage {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StorageUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StorageUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StorageUpdateOne) sqlSave(ctx context.Context) (_node *Storage, err error) {
	_spec := sqlgraph.NewUpdateSpec(storage.Table, storage.Columns, sqlgraph.NewFieldSpec(storage.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Storage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storage.FieldID)
		for _, f := range fields {
			if !storage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != storage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.SetField(storage.FieldDeletedAt, field.TypeTime, value)
	}
	if suo.mutation.DeletedAtCleared() {
		_spec.ClearField(storage.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.SetField(storage.FieldType, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.AddedType(); ok {
		_spec.AddField(storage.FieldType, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.Path(); ok {
		_spec.SetField(storage.FieldPath, field.TypeString, value)
	}
	if value, ok := suo.mutation.Md5(); ok {
		_spec.SetField(storage.FieldMd5, field.TypeString, value)
	}
	if suo.mutation.QuotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   storage.QuotesTable,
			Columns: []string{storage.QuotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedQuotesIDs(); len(nodes) > 0 && !suo.mutation.QuotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   storage.QuotesTable,
			Columns: []string{storage.QuotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.QuotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   storage.QuotesTable,
			Columns: []string{storage.QuotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Storage{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}