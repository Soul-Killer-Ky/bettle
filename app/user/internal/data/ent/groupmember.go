// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/user/internal/data/ent/group"
	"beetle/app/user/internal/data/ent/groupmember"
	"beetle/app/user/internal/data/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// GroupMember is the model entity for the GroupMember schema.
type GroupMember struct {
	config `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 群身份
	Role int32 `json:"role,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID int `json:"group_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupMemberQuery when eager-loading is set.
	Edges        GroupMemberEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupMemberEdges holds the relations/edges for other nodes in the graph.
type GroupMemberEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMemberEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMemberEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupMember) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupmember.FieldRole, groupmember.FieldGroupID, groupmember.FieldUserID:
			values[i] = new(sql.NullInt64)
		case groupmember.FieldCreatedAt, groupmember.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupMember fields.
func (gm *GroupMember) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupmember.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gm.CreatedAt = value.Time
			}
		case groupmember.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gm.DeletedAt = new(time.Time)
				*gm.DeletedAt = value.Time
			}
		case groupmember.FieldRole:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				gm.Role = int32(value.Int64)
			}
		case groupmember.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				gm.GroupID = int(value.Int64)
			}
		case groupmember.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				gm.UserID = int(value.Int64)
			}
		default:
			gm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupMember.
// This includes values selected through modifiers, order, etc.
func (gm *GroupMember) Value(name string) (ent.Value, error) {
	return gm.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the GroupMember entity.
func (gm *GroupMember) QueryGroup() *GroupQuery {
	return NewGroupMemberClient(gm.config).QueryGroup(gm)
}

// QueryUser queries the "user" edge of the GroupMember entity.
func (gm *GroupMember) QueryUser() *UserQuery {
	return NewGroupMemberClient(gm.config).QueryUser(gm)
}

// Update returns a builder for updating this GroupMember.
// Note that you need to call GroupMember.Unwrap() before calling this method if this GroupMember
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GroupMember) Update() *GroupMemberUpdateOne {
	return NewGroupMemberClient(gm.config).UpdateOne(gm)
}

// Unwrap unwraps the GroupMember entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GroupMember) Unwrap() *GroupMember {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupMember is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GroupMember) String() string {
	var builder strings.Builder
	builder.WriteString("GroupMember(")
	builder.WriteString("created_at=")
	builder.WriteString(gm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := gm.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", gm.Role))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", gm.GroupID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", gm.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// GroupMembers is a parsable slice of GroupMember.
type GroupMembers []*GroupMember