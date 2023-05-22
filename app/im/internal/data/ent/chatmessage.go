// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/im/internal/data/ent/chatmessage"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ChatMessage is the model entity for the ChatMessage schema.
type ChatMessage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 发送者
	From int `json:"from,omitempty"`
	// 接收者
	Sender int `json:"sender,omitempty"`
	// 消息内容
	Message string `json:"message,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChatMessage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chatmessage.FieldID, chatmessage.FieldFrom, chatmessage.FieldSender:
			values[i] = new(sql.NullInt64)
		case chatmessage.FieldMessage:
			values[i] = new(sql.NullString)
		case chatmessage.FieldCreatedAt, chatmessage.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ChatMessage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChatMessage fields.
func (cm *ChatMessage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chatmessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cm.ID = int(value.Int64)
		case chatmessage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cm.CreatedAt = value.Time
			}
		case chatmessage.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cm.DeletedAt = new(time.Time)
				*cm.DeletedAt = value.Time
			}
		case chatmessage.FieldFrom:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				cm.From = int(value.Int64)
			}
		case chatmessage.FieldSender:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value.Valid {
				cm.Sender = int(value.Int64)
			}
		case chatmessage.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				cm.Message = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ChatMessage.
// Note that you need to call ChatMessage.Unwrap() before calling this method if this ChatMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (cm *ChatMessage) Update() *ChatMessageUpdateOne {
	return NewChatMessageClient(cm.config).UpdateOne(cm)
}

// Unwrap unwraps the ChatMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cm *ChatMessage) Unwrap() *ChatMessage {
	_tx, ok := cm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChatMessage is not a transactional entity")
	}
	cm.config.driver = _tx.drv
	return cm
}

// String implements the fmt.Stringer.
func (cm *ChatMessage) String() string {
	var builder strings.Builder
	builder.WriteString("ChatMessage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(cm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := cm.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(fmt.Sprintf("%v", cm.From))
	builder.WriteString(", ")
	builder.WriteString("sender=")
	builder.WriteString(fmt.Sprintf("%v", cm.Sender))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(cm.Message)
	builder.WriteByte(')')
	return builder.String()
}

// ChatMessages is a parsable slice of ChatMessage.
type ChatMessages []*ChatMessage