package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("type").Comment("群类型"),
		field.String("name").
			MaxLen(30).Comment("名称"),
		field.String("icon").
			MaxLen(50).Comment("图标"),
		field.String("memo").
			MaxLen(50).Comment("简介"),
	}
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		UpdateTimeMixin{},
		SoftDeleteMixin{},
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).Through("group_members", GroupMember.Type),

		edge.From("created_by", User.Type).
			Ref("created_groups").
			Unique(),
	}
}
