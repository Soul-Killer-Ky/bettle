// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"beetle/app/im/internal/data/ent"
	"beetle/app/im/internal/data/ent/groupchatmessage"
	"beetle/app/im/internal/data/ent/personalchatmessage"
	"beetle/app/im/internal/data/ent/predicate"
	"beetle/app/im/internal/data/ent/synchronizerecord"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The GroupChatMessageFunc type is an adapter to allow the use of ordinary function as a Querier.
type GroupChatMessageFunc func(context.Context, *ent.GroupChatMessageQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f GroupChatMessageFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.GroupChatMessageQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.GroupChatMessageQuery", q)
}

// The TraverseGroupChatMessage type is an adapter to allow the use of ordinary function as Traverser.
type TraverseGroupChatMessage func(context.Context, *ent.GroupChatMessageQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseGroupChatMessage) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseGroupChatMessage) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GroupChatMessageQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.GroupChatMessageQuery", q)
}

// The PersonalChatMessageFunc type is an adapter to allow the use of ordinary function as a Querier.
type PersonalChatMessageFunc func(context.Context, *ent.PersonalChatMessageQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PersonalChatMessageFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PersonalChatMessageQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PersonalChatMessageQuery", q)
}

// The TraversePersonalChatMessage type is an adapter to allow the use of ordinary function as Traverser.
type TraversePersonalChatMessage func(context.Context, *ent.PersonalChatMessageQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePersonalChatMessage) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePersonalChatMessage) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PersonalChatMessageQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PersonalChatMessageQuery", q)
}

// The SynchronizeRecordFunc type is an adapter to allow the use of ordinary function as a Querier.
type SynchronizeRecordFunc func(context.Context, *ent.SynchronizeRecordQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f SynchronizeRecordFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.SynchronizeRecordQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.SynchronizeRecordQuery", q)
}

// The TraverseSynchronizeRecord type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSynchronizeRecord func(context.Context, *ent.SynchronizeRecordQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSynchronizeRecord) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSynchronizeRecord) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SynchronizeRecordQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.SynchronizeRecordQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.GroupChatMessageQuery:
		return &query[*ent.GroupChatMessageQuery, predicate.GroupChatMessage, groupchatmessage.OrderOption]{typ: ent.TypeGroupChatMessage, tq: q}, nil
	case *ent.PersonalChatMessageQuery:
		return &query[*ent.PersonalChatMessageQuery, predicate.PersonalChatMessage, personalchatmessage.OrderOption]{typ: ent.TypePersonalChatMessage, tq: q}, nil
	case *ent.SynchronizeRecordQuery:
		return &query[*ent.SynchronizeRecordQuery, predicate.SynchronizeRecord, synchronizerecord.OrderOption]{typ: ent.TypeSynchronizeRecord, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
