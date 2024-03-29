// Code generated by ent, DO NOT EDIT.

package ent

import (
	"beetle/app/im/internal/data/ent/synchronizerecord"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// SynchronizeRecord is the model entity for the SynchronizeRecord schema.
type SynchronizeRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 用户ID
	UserID int `json:"user_id,omitempty"`
	// 设备号
	DeviceID uuid.UUID `json:"device_id,omitempty"`
	// 消息ID
	LastMessageID int64 `json:"last_message_id,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SynchronizeRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case synchronizerecord.FieldID, synchronizerecord.FieldUserID, synchronizerecord.FieldLastMessageID:
			values[i] = new(sql.NullInt64)
		case synchronizerecord.FieldCreatedAt, synchronizerecord.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case synchronizerecord.FieldDeviceID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SynchronizeRecord fields.
func (sr *SynchronizeRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case synchronizerecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sr.ID = int(value.Int64)
		case synchronizerecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sr.CreatedAt = value.Time
			}
		case synchronizerecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sr.UpdatedAt = value.Time
			}
		case synchronizerecord.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				sr.UserID = int(value.Int64)
			}
		case synchronizerecord.FieldDeviceID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value != nil {
				sr.DeviceID = *value
			}
		case synchronizerecord.FieldLastMessageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_message_id", values[i])
			} else if value.Valid {
				sr.LastMessageID = value.Int64
			}
		default:
			sr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SynchronizeRecord.
// This includes values selected through modifiers, order, etc.
func (sr *SynchronizeRecord) Value(name string) (ent.Value, error) {
	return sr.selectValues.Get(name)
}

// Update returns a builder for updating this SynchronizeRecord.
// Note that you need to call SynchronizeRecord.Unwrap() before calling this method if this SynchronizeRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (sr *SynchronizeRecord) Update() *SynchronizeRecordUpdateOne {
	return NewSynchronizeRecordClient(sr.config).UpdateOne(sr)
}

// Unwrap unwraps the SynchronizeRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sr *SynchronizeRecord) Unwrap() *SynchronizeRecord {
	_tx, ok := sr.config.driver.(*txDriver)
	if !ok {
		panic("ent: SynchronizeRecord is not a transactional entity")
	}
	sr.config.driver = _tx.drv
	return sr
}

// String implements the fmt.Stringer.
func (sr *SynchronizeRecord) String() string {
	var builder strings.Builder
	builder.WriteString("SynchronizeRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", sr.UserID))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", sr.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("last_message_id=")
	builder.WriteString(fmt.Sprintf("%v", sr.LastMessageID))
	builder.WriteByte(')')
	return builder.String()
}

// SynchronizeRecords is a parsable slice of SynchronizeRecord.
type SynchronizeRecords []*SynchronizeRecord
