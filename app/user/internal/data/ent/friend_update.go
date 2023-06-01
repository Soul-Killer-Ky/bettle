// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/user/internal/data/ent/friend"
	"beetle/app/user/internal/data/ent/predicate"
	"beetle/app/user/internal/data/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendUpdate is the builder for updating Friend entities.
type FriendUpdate struct {
	config
	hooks    []Hook
	mutation *FriendMutation
}

// Where appends a list predicates to the FriendUpdate builder.
func (fu *FriendUpdate) Where(ps ...predicate.Friend) *FriendUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FriendUpdate) SetUpdatedAt(t time.Time) *FriendUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetDeletedAt sets the "deleted_at" field.
func (fu *FriendUpdate) SetDeletedAt(t time.Time) *FriendUpdate {
	fu.mutation.SetDeletedAt(t)
	return fu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fu *FriendUpdate) SetNillableDeletedAt(t *time.Time) *FriendUpdate {
	if t != nil {
		fu.SetDeletedAt(*t)
	}
	return fu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fu *FriendUpdate) ClearDeletedAt() *FriendUpdate {
	fu.mutation.ClearDeletedAt()
	return fu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (fu *FriendUpdate) SetUserID(id int) *FriendUpdate {
	fu.mutation.SetUserID(id)
	return fu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (fu *FriendUpdate) SetNillableUserID(id *int) *FriendUpdate {
	if id != nil {
		fu = fu.SetUserID(*id)
	}
	return fu
}

// SetUser sets the "user" edge to the User entity.
func (fu *FriendUpdate) SetUser(u *User) *FriendUpdate {
	return fu.SetUserID(u.ID)
}

// SetFriendID sets the "friend" edge to the User entity by ID.
func (fu *FriendUpdate) SetFriendID(id int) *FriendUpdate {
	fu.mutation.SetFriendID(id)
	return fu
}

// SetNillableFriendID sets the "friend" edge to the User entity by ID if the given value is not nil.
func (fu *FriendUpdate) SetNillableFriendID(id *int) *FriendUpdate {
	if id != nil {
		fu = fu.SetFriendID(*id)
	}
	return fu
}

// SetFriend sets the "friend" edge to the User entity.
func (fu *FriendUpdate) SetFriend(u *User) *FriendUpdate {
	return fu.SetFriendID(u.ID)
}

// Mutation returns the FriendMutation object of the builder.
func (fu *FriendUpdate) Mutation() *FriendMutation {
	return fu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fu *FriendUpdate) ClearUser() *FriendUpdate {
	fu.mutation.ClearUser()
	return fu
}

// ClearFriend clears the "friend" edge to the User entity.
func (fu *FriendUpdate) ClearFriend() *FriendUpdate {
	fu.mutation.ClearFriend()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FriendUpdate) Save(ctx context.Context) (int, error) {
	if err := fu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FriendUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FriendUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FriendUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FriendUpdate) defaults() error {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		if friend.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized friend.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := friend.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (fu *FriendUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(friend.Table, friend.Columns, sqlgraph.NewFieldSpec(friend.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(friend.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fu.mutation.DeletedAt(); ok {
		_spec.SetField(friend.FieldDeletedAt, field.TypeTime, value)
	}
	if fu.mutation.DeletedAtCleared() {
		_spec.ClearField(friend.FieldDeletedAt, field.TypeTime)
	}
	if fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.UserTable,
			Columns: []string{friend.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.UserTable,
			Columns: []string{friend.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.FriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.FriendTable,
			Columns: []string{friend.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.FriendTable,
			Columns: []string{friend.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FriendUpdateOne is the builder for updating a single Friend entity.
type FriendUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FriendUpdateOne) SetUpdatedAt(t time.Time) *FriendUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fuo *FriendUpdateOne) SetDeletedAt(t time.Time) *FriendUpdateOne {
	fuo.mutation.SetDeletedAt(t)
	return fuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableDeletedAt(t *time.Time) *FriendUpdateOne {
	if t != nil {
		fuo.SetDeletedAt(*t)
	}
	return fuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fuo *FriendUpdateOne) ClearDeletedAt() *FriendUpdateOne {
	fuo.mutation.ClearDeletedAt()
	return fuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (fuo *FriendUpdateOne) SetUserID(id int) *FriendUpdateOne {
	fuo.mutation.SetUserID(id)
	return fuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableUserID(id *int) *FriendUpdateOne {
	if id != nil {
		fuo = fuo.SetUserID(*id)
	}
	return fuo
}

// SetUser sets the "user" edge to the User entity.
func (fuo *FriendUpdateOne) SetUser(u *User) *FriendUpdateOne {
	return fuo.SetUserID(u.ID)
}

// SetFriendID sets the "friend" edge to the User entity by ID.
func (fuo *FriendUpdateOne) SetFriendID(id int) *FriendUpdateOne {
	fuo.mutation.SetFriendID(id)
	return fuo
}

// SetNillableFriendID sets the "friend" edge to the User entity by ID if the given value is not nil.
func (fuo *FriendUpdateOne) SetNillableFriendID(id *int) *FriendUpdateOne {
	if id != nil {
		fuo = fuo.SetFriendID(*id)
	}
	return fuo
}

// SetFriend sets the "friend" edge to the User entity.
func (fuo *FriendUpdateOne) SetFriend(u *User) *FriendUpdateOne {
	return fuo.SetFriendID(u.ID)
}

// Mutation returns the FriendMutation object of the builder.
func (fuo *FriendUpdateOne) Mutation() *FriendMutation {
	return fuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fuo *FriendUpdateOne) ClearUser() *FriendUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// ClearFriend clears the "friend" edge to the User entity.
func (fuo *FriendUpdateOne) ClearFriend() *FriendUpdateOne {
	fuo.mutation.ClearFriend()
	return fuo
}

// Where appends a list predicates to the FriendUpdate builder.
func (fuo *FriendUpdateOne) Where(ps ...predicate.Friend) *FriendUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FriendUpdateOne) Select(field string, fields ...string) *FriendUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Friend entity.
func (fuo *FriendUpdateOne) Save(ctx context.Context) (*Friend, error) {
	if err := fuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FriendUpdateOne) SaveX(ctx context.Context) *Friend {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FriendUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FriendUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FriendUpdateOne) defaults() error {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		if friend.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized friend.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := friend.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (fuo *FriendUpdateOne) sqlSave(ctx context.Context) (_node *Friend, err error) {
	_spec := sqlgraph.NewUpdateSpec(friend.Table, friend.Columns, sqlgraph.NewFieldSpec(friend.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Friend.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friend.FieldID)
		for _, f := range fields {
			if !friend.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(friend.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.DeletedAt(); ok {
		_spec.SetField(friend.FieldDeletedAt, field.TypeTime, value)
	}
	if fuo.mutation.DeletedAtCleared() {
		_spec.ClearField(friend.FieldDeletedAt, field.TypeTime)
	}
	if fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.UserTable,
			Columns: []string{friend.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.UserTable,
			Columns: []string{friend.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.FriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.FriendTable,
			Columns: []string{friend.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friend.FriendTable,
			Columns: []string{friend.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Friend{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
