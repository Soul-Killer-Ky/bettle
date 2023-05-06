package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			MaxLen(30).Unique().Comment("唯一用户标识"),
		field.String("password").
			MaxLen(100),
		field.String("nickname").
			MaxLen(20).Comment("昵称"),
		field.String("avatar").
			MaxLen(50).Comment("头像"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		UpdateTimeMixin{},
		SoftDeleteMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", Friend.Type),
		edge.To("friends", Friend.Type),
	}
}
