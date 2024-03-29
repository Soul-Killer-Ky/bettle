// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/filesystem/internal/data/ent/predicate"
	"beetle/app/filesystem/internal/data/ent/quote"
	"beetle/app/filesystem/internal/data/ent/storage"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeQuote   = "Quote"
	TypeStorage = "Storage"
)

// QuoteMutation represents an operation that mutates the Quote nodes in the graph.
type QuoteMutation struct {
	config
	op             Op
	typ            string
	id             *int
	created_at     *time.Time
	deleted_at     *time.Time
	name           *string
	created_by     *int
	addcreated_by  *int
	clearedFields  map[string]struct{}
	storage        *int
	clearedstorage bool
	done           bool
	oldValue       func(context.Context) (*Quote, error)
	predicates     []predicate.Quote
}

var _ ent.Mutation = (*QuoteMutation)(nil)

// quoteOption allows management of the mutation configuration using functional options.
type quoteOption func(*QuoteMutation)

// newQuoteMutation creates new mutation for the Quote entity.
func newQuoteMutation(c config, op Op, opts ...quoteOption) *QuoteMutation {
	m := &QuoteMutation{
		config:        c,
		op:            op,
		typ:           TypeQuote,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withQuoteID sets the ID field of the mutation.
func withQuoteID(id int) quoteOption {
	return func(m *QuoteMutation) {
		var (
			err   error
			once  sync.Once
			value *Quote
		)
		m.oldValue = func(ctx context.Context) (*Quote, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Quote.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withQuote sets the old Quote of the mutation.
func withQuote(node *Quote) quoteOption {
	return func(m *QuoteMutation) {
		m.oldValue = func(context.Context) (*Quote, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m QuoteMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m QuoteMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *QuoteMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *QuoteMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Quote.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *QuoteMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *QuoteMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Quote entity.
// If the Quote object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *QuoteMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *QuoteMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetDeletedAt sets the "deleted_at" field.
func (m *QuoteMutation) SetDeletedAt(t time.Time) {
	m.deleted_at = &t
}

// DeletedAt returns the value of the "deleted_at" field in the mutation.
func (m *QuoteMutation) DeletedAt() (r time.Time, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old "deleted_at" field's value of the Quote entity.
// If the Quote object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *QuoteMutation) OldDeletedAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (m *QuoteMutation) ClearDeletedAt() {
	m.deleted_at = nil
	m.clearedFields[quote.FieldDeletedAt] = struct{}{}
}

// DeletedAtCleared returns if the "deleted_at" field was cleared in this mutation.
func (m *QuoteMutation) DeletedAtCleared() bool {
	_, ok := m.clearedFields[quote.FieldDeletedAt]
	return ok
}

// ResetDeletedAt resets all changes to the "deleted_at" field.
func (m *QuoteMutation) ResetDeletedAt() {
	m.deleted_at = nil
	delete(m.clearedFields, quote.FieldDeletedAt)
}

// SetName sets the "name" field.
func (m *QuoteMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *QuoteMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Quote entity.
// If the Quote object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *QuoteMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *QuoteMutation) ResetName() {
	m.name = nil
}

// SetCreatedBy sets the "created_by" field.
func (m *QuoteMutation) SetCreatedBy(i int) {
	m.created_by = &i
	m.addcreated_by = nil
}

// CreatedBy returns the value of the "created_by" field in the mutation.
func (m *QuoteMutation) CreatedBy() (r int, exists bool) {
	v := m.created_by
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedBy returns the old "created_by" field's value of the Quote entity.
// If the Quote object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *QuoteMutation) OldCreatedBy(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedBy is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedBy requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedBy: %w", err)
	}
	return oldValue.CreatedBy, nil
}

// AddCreatedBy adds i to the "created_by" field.
func (m *QuoteMutation) AddCreatedBy(i int) {
	if m.addcreated_by != nil {
		*m.addcreated_by += i
	} else {
		m.addcreated_by = &i
	}
}

// AddedCreatedBy returns the value that was added to the "created_by" field in this mutation.
func (m *QuoteMutation) AddedCreatedBy() (r int, exists bool) {
	v := m.addcreated_by
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedBy resets all changes to the "created_by" field.
func (m *QuoteMutation) ResetCreatedBy() {
	m.created_by = nil
	m.addcreated_by = nil
}

// SetStorageID sets the "storage" edge to the Storage entity by id.
func (m *QuoteMutation) SetStorageID(id int) {
	m.storage = &id
}

// ClearStorage clears the "storage" edge to the Storage entity.
func (m *QuoteMutation) ClearStorage() {
	m.clearedstorage = true
}

// StorageCleared reports if the "storage" edge to the Storage entity was cleared.
func (m *QuoteMutation) StorageCleared() bool {
	return m.clearedstorage
}

// StorageID returns the "storage" edge ID in the mutation.
func (m *QuoteMutation) StorageID() (id int, exists bool) {
	if m.storage != nil {
		return *m.storage, true
	}
	return
}

// StorageIDs returns the "storage" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// StorageID instead. It exists only for internal usage by the builders.
func (m *QuoteMutation) StorageIDs() (ids []int) {
	if id := m.storage; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetStorage resets all changes to the "storage" edge.
func (m *QuoteMutation) ResetStorage() {
	m.storage = nil
	m.clearedstorage = false
}

// Where appends a list predicates to the QuoteMutation builder.
func (m *QuoteMutation) Where(ps ...predicate.Quote) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the QuoteMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *QuoteMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Quote, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *QuoteMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *QuoteMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Quote).
func (m *QuoteMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *QuoteMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, quote.FieldCreatedAt)
	}
	if m.deleted_at != nil {
		fields = append(fields, quote.FieldDeletedAt)
	}
	if m.name != nil {
		fields = append(fields, quote.FieldName)
	}
	if m.created_by != nil {
		fields = append(fields, quote.FieldCreatedBy)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *QuoteMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case quote.FieldCreatedAt:
		return m.CreatedAt()
	case quote.FieldDeletedAt:
		return m.DeletedAt()
	case quote.FieldName:
		return m.Name()
	case quote.FieldCreatedBy:
		return m.CreatedBy()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *QuoteMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case quote.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case quote.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	case quote.FieldName:
		return m.OldName(ctx)
	case quote.FieldCreatedBy:
		return m.OldCreatedBy(ctx)
	}
	return nil, fmt.Errorf("unknown Quote field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *QuoteMutation) SetField(name string, value ent.Value) error {
	switch name {
	case quote.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case quote.FieldDeletedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	case quote.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case quote.FieldCreatedBy:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedBy(v)
		return nil
	}
	return fmt.Errorf("unknown Quote field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *QuoteMutation) AddedFields() []string {
	var fields []string
	if m.addcreated_by != nil {
		fields = append(fields, quote.FieldCreatedBy)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *QuoteMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case quote.FieldCreatedBy:
		return m.AddedCreatedBy()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *QuoteMutation) AddField(name string, value ent.Value) error {
	switch name {
	case quote.FieldCreatedBy:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreatedBy(v)
		return nil
	}
	return fmt.Errorf("unknown Quote numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *QuoteMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(quote.FieldDeletedAt) {
		fields = append(fields, quote.FieldDeletedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *QuoteMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *QuoteMutation) ClearField(name string) error {
	switch name {
	case quote.FieldDeletedAt:
		m.ClearDeletedAt()
		return nil
	}
	return fmt.Errorf("unknown Quote nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *QuoteMutation) ResetField(name string) error {
	switch name {
	case quote.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case quote.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	case quote.FieldName:
		m.ResetName()
		return nil
	case quote.FieldCreatedBy:
		m.ResetCreatedBy()
		return nil
	}
	return fmt.Errorf("unknown Quote field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *QuoteMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.storage != nil {
		edges = append(edges, quote.EdgeStorage)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *QuoteMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case quote.EdgeStorage:
		if id := m.storage; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *QuoteMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *QuoteMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *QuoteMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedstorage {
		edges = append(edges, quote.EdgeStorage)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *QuoteMutation) EdgeCleared(name string) bool {
	switch name {
	case quote.EdgeStorage:
		return m.clearedstorage
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *QuoteMutation) ClearEdge(name string) error {
	switch name {
	case quote.EdgeStorage:
		m.ClearStorage()
		return nil
	}
	return fmt.Errorf("unknown Quote unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *QuoteMutation) ResetEdge(name string) error {
	switch name {
	case quote.EdgeStorage:
		m.ResetStorage()
		return nil
	}
	return fmt.Errorf("unknown Quote edge %s", name)
}

// StorageMutation represents an operation that mutates the Storage nodes in the graph.
type StorageMutation struct {
	config
	op            Op
	typ           string
	id            *int
	created_at    *time.Time
	deleted_at    *time.Time
	_type         *int32
	add_type      *int32
	_path         *string
	md5           *string
	clearedFields map[string]struct{}
	quotes        map[int]struct{}
	removedquotes map[int]struct{}
	clearedquotes bool
	done          bool
	oldValue      func(context.Context) (*Storage, error)
	predicates    []predicate.Storage
}

var _ ent.Mutation = (*StorageMutation)(nil)

// storageOption allows management of the mutation configuration using functional options.
type storageOption func(*StorageMutation)

// newStorageMutation creates new mutation for the Storage entity.
func newStorageMutation(c config, op Op, opts ...storageOption) *StorageMutation {
	m := &StorageMutation{
		config:        c,
		op:            op,
		typ:           TypeStorage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStorageID sets the ID field of the mutation.
func withStorageID(id int) storageOption {
	return func(m *StorageMutation) {
		var (
			err   error
			once  sync.Once
			value *Storage
		)
		m.oldValue = func(ctx context.Context) (*Storage, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Storage.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStorage sets the old Storage of the mutation.
func withStorage(node *Storage) storageOption {
	return func(m *StorageMutation) {
		m.oldValue = func(context.Context) (*Storage, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m StorageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StorageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *StorageMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *StorageMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Storage.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *StorageMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *StorageMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Storage entity.
// If the Storage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StorageMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *StorageMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetDeletedAt sets the "deleted_at" field.
func (m *StorageMutation) SetDeletedAt(t time.Time) {
	m.deleted_at = &t
}

// DeletedAt returns the value of the "deleted_at" field in the mutation.
func (m *StorageMutation) DeletedAt() (r time.Time, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old "deleted_at" field's value of the Storage entity.
// If the Storage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StorageMutation) OldDeletedAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (m *StorageMutation) ClearDeletedAt() {
	m.deleted_at = nil
	m.clearedFields[storage.FieldDeletedAt] = struct{}{}
}

// DeletedAtCleared returns if the "deleted_at" field was cleared in this mutation.
func (m *StorageMutation) DeletedAtCleared() bool {
	_, ok := m.clearedFields[storage.FieldDeletedAt]
	return ok
}

// ResetDeletedAt resets all changes to the "deleted_at" field.
func (m *StorageMutation) ResetDeletedAt() {
	m.deleted_at = nil
	delete(m.clearedFields, storage.FieldDeletedAt)
}

// SetType sets the "type" field.
func (m *StorageMutation) SetType(i int32) {
	m._type = &i
	m.add_type = nil
}

// GetType returns the value of the "type" field in the mutation.
func (m *StorageMutation) GetType() (r int32, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the Storage entity.
// If the Storage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StorageMutation) OldType(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// AddType adds i to the "type" field.
func (m *StorageMutation) AddType(i int32) {
	if m.add_type != nil {
		*m.add_type += i
	} else {
		m.add_type = &i
	}
}

// AddedType returns the value that was added to the "type" field in this mutation.
func (m *StorageMutation) AddedType() (r int32, exists bool) {
	v := m.add_type
	if v == nil {
		return
	}
	return *v, true
}

// ResetType resets all changes to the "type" field.
func (m *StorageMutation) ResetType() {
	m._type = nil
	m.add_type = nil
}

// SetPath sets the "path" field.
func (m *StorageMutation) SetPath(s string) {
	m._path = &s
}

// Path returns the value of the "path" field in the mutation.
func (m *StorageMutation) Path() (r string, exists bool) {
	v := m._path
	if v == nil {
		return
	}
	return *v, true
}

// OldPath returns the old "path" field's value of the Storage entity.
// If the Storage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StorageMutation) OldPath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPath: %w", err)
	}
	return oldValue.Path, nil
}

// ResetPath resets all changes to the "path" field.
func (m *StorageMutation) ResetPath() {
	m._path = nil
}

// SetMd5 sets the "md5" field.
func (m *StorageMutation) SetMd5(s string) {
	m.md5 = &s
}

// Md5 returns the value of the "md5" field in the mutation.
func (m *StorageMutation) Md5() (r string, exists bool) {
	v := m.md5
	if v == nil {
		return
	}
	return *v, true
}

// OldMd5 returns the old "md5" field's value of the Storage entity.
// If the Storage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StorageMutation) OldMd5(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMd5 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMd5 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMd5: %w", err)
	}
	return oldValue.Md5, nil
}

// ResetMd5 resets all changes to the "md5" field.
func (m *StorageMutation) ResetMd5() {
	m.md5 = nil
}

// AddQuoteIDs adds the "quotes" edge to the Quote entity by ids.
func (m *StorageMutation) AddQuoteIDs(ids ...int) {
	if m.quotes == nil {
		m.quotes = make(map[int]struct{})
	}
	for i := range ids {
		m.quotes[ids[i]] = struct{}{}
	}
}

// ClearQuotes clears the "quotes" edge to the Quote entity.
func (m *StorageMutation) ClearQuotes() {
	m.clearedquotes = true
}

// QuotesCleared reports if the "quotes" edge to the Quote entity was cleared.
func (m *StorageMutation) QuotesCleared() bool {
	return m.clearedquotes
}

// RemoveQuoteIDs removes the "quotes" edge to the Quote entity by IDs.
func (m *StorageMutation) RemoveQuoteIDs(ids ...int) {
	if m.removedquotes == nil {
		m.removedquotes = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.quotes, ids[i])
		m.removedquotes[ids[i]] = struct{}{}
	}
}

// RemovedQuotes returns the removed IDs of the "quotes" edge to the Quote entity.
func (m *StorageMutation) RemovedQuotesIDs() (ids []int) {
	for id := range m.removedquotes {
		ids = append(ids, id)
	}
	return
}

// QuotesIDs returns the "quotes" edge IDs in the mutation.
func (m *StorageMutation) QuotesIDs() (ids []int) {
	for id := range m.quotes {
		ids = append(ids, id)
	}
	return
}

// ResetQuotes resets all changes to the "quotes" edge.
func (m *StorageMutation) ResetQuotes() {
	m.quotes = nil
	m.clearedquotes = false
	m.removedquotes = nil
}

// Where appends a list predicates to the StorageMutation builder.
func (m *StorageMutation) Where(ps ...predicate.Storage) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the StorageMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *StorageMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Storage, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *StorageMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *StorageMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Storage).
func (m *StorageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *StorageMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, storage.FieldCreatedAt)
	}
	if m.deleted_at != nil {
		fields = append(fields, storage.FieldDeletedAt)
	}
	if m._type != nil {
		fields = append(fields, storage.FieldType)
	}
	if m._path != nil {
		fields = append(fields, storage.FieldPath)
	}
	if m.md5 != nil {
		fields = append(fields, storage.FieldMd5)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *StorageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case storage.FieldCreatedAt:
		return m.CreatedAt()
	case storage.FieldDeletedAt:
		return m.DeletedAt()
	case storage.FieldType:
		return m.GetType()
	case storage.FieldPath:
		return m.Path()
	case storage.FieldMd5:
		return m.Md5()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *StorageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case storage.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case storage.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	case storage.FieldType:
		return m.OldType(ctx)
	case storage.FieldPath:
		return m.OldPath(ctx)
	case storage.FieldMd5:
		return m.OldMd5(ctx)
	}
	return nil, fmt.Errorf("unknown Storage field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StorageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case storage.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case storage.FieldDeletedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	case storage.FieldType:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case storage.FieldPath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPath(v)
		return nil
	case storage.FieldMd5:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMd5(v)
		return nil
	}
	return fmt.Errorf("unknown Storage field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *StorageMutation) AddedFields() []string {
	var fields []string
	if m.add_type != nil {
		fields = append(fields, storage.FieldType)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *StorageMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case storage.FieldType:
		return m.AddedType()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StorageMutation) AddField(name string, value ent.Value) error {
	switch name {
	case storage.FieldType:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddType(v)
		return nil
	}
	return fmt.Errorf("unknown Storage numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *StorageMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(storage.FieldDeletedAt) {
		fields = append(fields, storage.FieldDeletedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *StorageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *StorageMutation) ClearField(name string) error {
	switch name {
	case storage.FieldDeletedAt:
		m.ClearDeletedAt()
		return nil
	}
	return fmt.Errorf("unknown Storage nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *StorageMutation) ResetField(name string) error {
	switch name {
	case storage.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case storage.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	case storage.FieldType:
		m.ResetType()
		return nil
	case storage.FieldPath:
		m.ResetPath()
		return nil
	case storage.FieldMd5:
		m.ResetMd5()
		return nil
	}
	return fmt.Errorf("unknown Storage field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *StorageMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.quotes != nil {
		edges = append(edges, storage.EdgeQuotes)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *StorageMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case storage.EdgeQuotes:
		ids := make([]ent.Value, 0, len(m.quotes))
		for id := range m.quotes {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *StorageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedquotes != nil {
		edges = append(edges, storage.EdgeQuotes)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *StorageMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case storage.EdgeQuotes:
		ids := make([]ent.Value, 0, len(m.removedquotes))
		for id := range m.removedquotes {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *StorageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedquotes {
		edges = append(edges, storage.EdgeQuotes)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *StorageMutation) EdgeCleared(name string) bool {
	switch name {
	case storage.EdgeQuotes:
		return m.clearedquotes
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *StorageMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Storage unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *StorageMutation) ResetEdge(name string) error {
	switch name {
	case storage.EdgeQuotes:
		m.ResetQuotes()
		return nil
	}
	return fmt.Errorf("unknown Storage edge %s", name)
}
