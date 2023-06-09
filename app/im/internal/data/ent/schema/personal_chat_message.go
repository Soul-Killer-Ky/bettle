package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type PersonalChatMessage struct {
	ent.Schema
}

func (PersonalChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("from").Comment("发送者"),
		field.Int("sender").Comment("接收者"),
		field.Text("message").Comment("消息内容"),
	}
}

func (PersonalChatMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		SoftDeleteMixin{},
	}
}

func (PersonalChatMessage) Edges() []ent.Edge {
	return nil
}
