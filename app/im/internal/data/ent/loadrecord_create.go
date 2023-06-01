// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/im/internal/data/ent/loadrecord"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LoadRecordCreate is the builder for creating a LoadRecord entity.
type LoadRecordCreate struct {
	config
	mutation *LoadRecordMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lrc *LoadRecordCreate) SetCreatedAt(t time.Time) *LoadRecordCreate {
	lrc.mutation.SetCreatedAt(t)
	return lrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lrc *LoadRecordCreate) SetNillableCreatedAt(t *time.Time) *LoadRecordCreate {
	if t != nil {
		lrc.SetCreatedAt(*t)
	}
	return lrc
}

// SetUserID sets the "user_id" field.
func (lrc *LoadRecordCreate) SetUserID(i int) *LoadRecordCreate {
	lrc.mutation.SetUserID(i)
	return lrc
}

// SetDeviceID sets the "device_id" field.
func (lrc *LoadRecordCreate) SetDeviceID(u uuid.UUID) *LoadRecordCreate {
	lrc.mutation.SetDeviceID(u)
	return lrc
}

// Mutation returns the LoadRecordMutation object of the builder.
func (lrc *LoadRecordCreate) Mutation() *LoadRecordMutation {
	return lrc.mutation
}

// Save creates the LoadRecord in the database.
func (lrc *LoadRecordCreate) Save(ctx context.Context) (*LoadRecord, error) {
	lrc.defaults()
	return withHooks(ctx, lrc.sqlSave, lrc.mutation, lrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lrc *LoadRecordCreate) SaveX(ctx context.Context) *LoadRecord {
	v, err := lrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lrc *LoadRecordCreate) Exec(ctx context.Context) error {
	_, err := lrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lrc *LoadRecordCreate) ExecX(ctx context.Context) {
	if err := lrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lrc *LoadRecordCreate) defaults() {
	if _, ok := lrc.mutation.CreatedAt(); !ok {
		v := loadrecord.DefaultCreatedAt()
		lrc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lrc *LoadRecordCreate) check() error {
	if _, ok := lrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LoadRecord.created_at"`)}
	}
	if _, ok := lrc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "LoadRecord.user_id"`)}
	}
	if _, ok := lrc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device_id", err: errors.New(`ent: missing required field "LoadRecord.device_id"`)}
	}
	return nil
}

func (lrc *LoadRecordCreate) sqlSave(ctx context.Context) (*LoadRecord, error) {
	if err := lrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lrc.mutation.id = &_node.ID
	lrc.mutation.done = true
	return _node, nil
}

func (lrc *LoadRecordCreate) createSpec() (*LoadRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &LoadRecord{config: lrc.config}
		_spec = sqlgraph.NewCreateSpec(loadrecord.Table, sqlgraph.NewFieldSpec(loadrecord.FieldID, field.TypeInt))
	)
	if value, ok := lrc.mutation.CreatedAt(); ok {
		_spec.SetField(loadrecord.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lrc.mutation.UserID(); ok {
		_spec.SetField(loadrecord.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := lrc.mutation.DeviceID(); ok {
		_spec.SetField(loadrecord.FieldDeviceID, field.TypeUUID, value)
		_node.DeviceID = value
	}
	return _node, _spec
}

// LoadRecordCreateBulk is the builder for creating many LoadRecord entities in bulk.
type LoadRecordCreateBulk struct {
	config
	builders []*LoadRecordCreate
}

// Save creates the LoadRecord entities in the database.
func (lrcb *LoadRecordCreateBulk) Save(ctx context.Context) ([]*LoadRecord, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lrcb.builders))
	nodes := make([]*LoadRecord, len(lrcb.builders))
	mutators := make([]Mutator, len(lrcb.builders))
	for i := range lrcb.builders {
		func(i int, root context.Context) {
			builder := lrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LoadRecordMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lrcb *LoadRecordCreateBulk) SaveX(ctx context.Context) []*LoadRecord {
	v, err := lrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lrcb *LoadRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := lrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lrcb *LoadRecordCreateBulk) ExecX(ctx context.Context) {
	if err := lrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
