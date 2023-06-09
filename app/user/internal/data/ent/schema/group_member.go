package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type GroupMember struct {
	ent.Schema
}

func (GroupMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("group_id", "user_id"),
	}
}

func (GroupMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("role").Default(0).Comment("群身份"),

		field.Int("group_id"),
		field.Int("user_id"),
	}
}

func (GroupMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		SoftDeleteMixin{},
	}
}

func (GroupMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", Group.Type).
			Unique().
			Required().
			Field("group_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}
