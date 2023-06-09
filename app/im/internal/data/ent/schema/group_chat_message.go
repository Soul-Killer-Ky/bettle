package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type GroupChatMessage struct {
	ent.Schema
}

func (GroupChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("from").Comment("发送者"),
		field.Int("group_id").Comment("群ID"),
		field.Text("message").Comment("消息内容"),
	}
}

func (GroupChatMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		SoftDeleteMixin{},
	}
}

func (GroupChatMessage) Edges() []ent.Edge {
	return nil
}
