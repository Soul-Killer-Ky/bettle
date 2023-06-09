// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/im/internal/data/ent/groupchatmessage"
	"beetle/app/im/internal/data/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupChatMessageUpdate is the builder for updating GroupChatMessage entities.
type GroupChatMessageUpdate struct {
	config
	hooks    []Hook
	mutation *GroupChatMessageMutation
}

// Where appends a list predicates to the GroupChatMessageUpdate builder.
func (gcmu *GroupChatMessageUpdate) Where(ps ...predicate.GroupChatMessage) *GroupChatMessageUpdate {
	gcmu.mutation.Where(ps...)
	return gcmu
}

// SetDeletedAt sets the "deleted_at" field.
func (gcmu *GroupChatMessageUpdate) SetDeletedAt(t time.Time) *GroupChatMessageUpdate {
	gcmu.mutation.SetDeletedAt(t)
	return gcmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gcmu *GroupChatMessageUpdate) SetNillableDeletedAt(t *time.Time) *GroupChatMessageUpdate {
	if t != nil {
		gcmu.SetDeletedAt(*t)
	}
	return gcmu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gcmu *GroupChatMessageUpdate) ClearDeletedAt() *GroupChatMessageUpdate {
	gcmu.mutation.ClearDeletedAt()
	return gcmu
}

// SetFrom sets the "from" field.
func (gcmu *GroupChatMessageUpdate) SetFrom(i int) *GroupChatMessageUpdate {
	gcmu.mutation.ResetFrom()
	gcmu.mutation.SetFrom(i)
	return gcmu
}

// AddFrom adds i to the "from" field.
func (gcmu *GroupChatMessageUpdate) AddFrom(i int) *GroupChatMessageUpdate {
	gcmu.mutation.AddFrom(i)
	return gcmu
}

// SetGroupID sets the "group_id" field.
func (gcmu *GroupChatMessageUpdate) SetGroupID(i int) *GroupChatMessageUpdate {
	gcmu.mutation.ResetGroupID()
	gcmu.mutation.SetGroupID(i)
	return gcmu
}

// AddGroupID adds i to the "group_id" field.
func (gcmu *GroupChatMessageUpdate) AddGroupID(i int) *GroupChatMessageUpdate {
	gcmu.mutation.AddGroupID(i)
	return gcmu
}

// SetMessage sets the "message" field.
func (gcmu *GroupChatMessageUpdate) SetMessage(s string) *GroupChatMessageUpdate {
	gcmu.mutation.SetMessage(s)
	return gcmu
}

// Mutation returns the GroupChatMessageMutation object of the builder.
func (gcmu *GroupChatMessageUpdate) Mutation() *GroupChatMessageMutation {
	return gcmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gcmu *GroupChatMessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gcmu.sqlSave, gcmu.mutation, gcmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gcmu *GroupChatMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := gcmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gcmu *GroupChatMessageUpdate) Exec(ctx context.Context) error {
	_, err := gcmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcmu *GroupChatMessageUpdate) ExecX(ctx context.Context) {
	if err := gcmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gcmu *GroupChatMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(groupchatmessage.Table, groupchatmessage.Columns, sqlgraph.NewFieldSpec(groupchatmessage.FieldID, field.TypeInt))
	if ps := gcmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcmu.mutation.DeletedAt(); ok {
		_spec.SetField(groupchatmessage.FieldDeletedAt, field.TypeTime, value)
	}
	if gcmu.mutation.DeletedAtCleared() {
		_spec.ClearField(groupchatmessage.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gcmu.mutation.From(); ok {
		_spec.SetField(groupchatmessage.FieldFrom, field.TypeInt, value)
	}
	if value, ok := gcmu.mutation.AddedFrom(); ok {
		_spec.AddField(groupchatmessage.FieldFrom, field.TypeInt, value)
	}
	if value, ok := gcmu.mutation.GroupID(); ok {
		_spec.SetField(groupchatmessage.FieldGroupID, field.TypeInt, value)
	}
	if value, ok := gcmu.mutation.AddedGroupID(); ok {
		_spec.AddField(groupchatmessage.FieldGroupID, field.TypeInt, value)
	}
	if value, ok := gcmu.mutation.Message(); ok {
		_spec.SetField(groupchatmessage.FieldMessage, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gcmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupchatmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gcmu.mutation.done = true
	return n, nil
}

// GroupChatMessageUpdateOne is the builder for updating a single GroupChatMessage entity.
type GroupChatMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupChatMessageMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (gcmuo *GroupChatMessageUpdateOne) SetDeletedAt(t time.Time) *GroupChatMessageUpdateOne {
	gcmuo.mutation.SetDeletedAt(t)
	return gcmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gcmuo *GroupChatMessageUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupChatMessageUpdateOne {
	if t != nil {
		gcmuo.SetDeletedAt(*t)
	}
	return gcmuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gcmuo *GroupChatMessageUpdateOne) ClearDeletedAt() *GroupChatMessageUpdateOne {
	gcmuo.mutation.ClearDeletedAt()
	return gcmuo
}

// SetFrom sets the "from" field.
func (gcmuo *GroupChatMessageUpdateOne) SetFrom(i int) *GroupChatMessageUpdateOne {
	gcmuo.mutation.ResetFrom()
	gcmuo.mutation.SetFrom(i)
	return gcmuo
}

// AddFrom adds i to the "from" field.
func (gcmuo *GroupChatMessageUpdateOne) AddFrom(i int) *GroupChatMessageUpdateOne {
	gcmuo.mutation.AddFrom(i)
	return gcmuo
}

// SetGroupID sets the "group_id" field.
func (gcmuo *GroupChatMessageUpdateOne) SetGroupID(i int) *GroupChatMessageUpdateOne {
	gcmuo.mutation.ResetGroupID()
	gcmuo.mutation.SetGroupID(i)
	return gcmuo
}

// AddGroupID adds i to the "group_id" field.
func (gcmuo *GroupChatMessageUpdateOne) AddGroupID(i int) *GroupChatMessageUpdateOne {
	gcmuo.mutation.AddGroupID(i)
	return gcmuo
}

// SetMessage sets the "message" field.
func (gcmuo *GroupChatMessageUpdateOne) SetMessage(s string) *GroupChatMessageUpdateOne {
	gcmuo.mutation.SetMessage(s)
	return gcmuo
}

// Mutation returns the GroupChatMessageMutation object of the builder.
func (gcmuo *GroupChatMessageUpdateOne) Mutation() *GroupChatMessageMutation {
	return gcmuo.mutation
}

// Where appends a list predicates to the GroupChatMessageUpdate builder.
func (gcmuo *GroupChatMessageUpdateOne) Where(ps ...predicate.GroupChatMessage) *GroupChatMessageUpdateOne {
	gcmuo.mutation.Where(ps...)
	return gcmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gcmuo *GroupChatMessageUpdateOne) Select(field string, fields ...string) *GroupChatMessageUpdateOne {
	gcmuo.fields = append([]string{field}, fields...)
	return gcmuo
}

// Save executes the query and returns the updated GroupChatMessage entity.
func (gcmuo *GroupChatMessageUpdateOne) Save(ctx context.Context) (*GroupChatMessage, error) {
	return withHooks(ctx, gcmuo.sqlSave, gcmuo.mutation, gcmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gcmuo *GroupChatMessageUpdateOne) SaveX(ctx context.Context) *GroupChatMessage {
	node, err := gcmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gcmuo *GroupChatMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := gcmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcmuo *GroupChatMessageUpdateOne) ExecX(ctx context.Context) {
	if err := gcmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gcmuo *GroupChatMessageUpdateOne) sqlSave(ctx context.Context) (_node *GroupChatMessage, err error) {
	_spec := sqlgraph.NewUpdateSpec(groupchatmessage.Table, groupchatmessage.Columns, sqlgraph.NewFieldSpec(groupchatmessage.FieldID, field.TypeInt))
	id, ok := gcmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GroupChatMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gcmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupchatmessage.FieldID)
		for _, f := range fields {
			if !groupchatmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != groupchatmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gcmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcmuo.mutation.DeletedAt(); ok {
		_spec.SetField(groupchatmessage.FieldDeletedAt, field.TypeTime, value)
	}
	if gcmuo.mutation.DeletedAtCleared() {
		_spec.ClearField(groupchatmessage.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gcmuo.mutation.From(); ok {
		_spec.SetField(groupchatmessage.FieldFrom, field.TypeInt, value)
	}
	if value, ok := gcmuo.mutation.AddedFrom(); ok {
		_spec.AddField(groupchatmessage.FieldFrom, field.TypeInt, value)
	}
	if value, ok := gcmuo.mutation.GroupID(); ok {
		_spec.SetField(groupchatmessage.FieldGroupID, field.TypeInt, value)
	}
	if value, ok := gcmuo.mutation.AddedGroupID(); ok {
		_spec.AddField(groupchatmessage.FieldGroupID, field.TypeInt, value)
	}
	if value, ok := gcmuo.mutation.Message(); ok {
		_spec.SetField(groupchatmessage.FieldMessage, field.TypeString, value)
	}
	_node = &GroupChatMessage{config: gcmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gcmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupchatmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gcmuo.mutation.done = true
	return _node, nil
}
