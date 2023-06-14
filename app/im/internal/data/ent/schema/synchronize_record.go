package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type SynchronizeRecord struct {
	ent.Schema
}

func (SynchronizeRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户ID"),
		field.UUID("device_id", uuid.New()).Comment("设备号"),
		field.Int64("last_message_id").Comment("消息ID"),
	}
}

func (SynchronizeRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		UpdateTimeMixin{},
	}
}

func (SynchronizeRecord) Edges() []ent.Edge {
	return nil
}

func (SynchronizeRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "device_id").
			Unique(),
	}
}
