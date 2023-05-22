package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ChatMessage struct {
	ent.Schema
}

func (ChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("from").Comment("发送者"),
		field.Int("sender").Comment("接收者"),
		field.Text("message").Comment("消息内容"),
	}
}

func (ChatMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		SoftDeleteMixin{},
	}
}

func (ChatMessage) Edges() []ent.Edge {
	return nil
}
