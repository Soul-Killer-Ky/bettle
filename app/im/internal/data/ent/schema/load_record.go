package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type LoadRecord struct {
	ent.Schema
}

func (LoadRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户ID"),
		field.UUID("device_id", uuid.New()).Comment("设备号"),
	}
}

func (LoadRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
	}
}

func (LoadRecord) Edges() []ent.Edge {
	return nil
}
