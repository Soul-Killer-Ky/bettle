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
		field.String("phone").
			MaxLen(20).Unique().Comment("手机号"),
		field.String("username").
			MaxLen(30).Nillable().Unique().Comment("唯一用户标识"),
		field.String("password").
			MaxLen(100).Comment("密码"),
		field.String("nickname").
			MaxLen(20).Comment("昵称"),
		field.String("avatar").
			MaxLen(50).Comment("头像"),
		field.String("memo").
			MaxLen(50).Comment("个人简介"),
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
		// friends
		edge.To("friends", User.Type),

		// groups
		edge.From("joined_groups", Group.Type).Ref("users").
			Through("group_members", GroupMember.Type),
		edge.To("created_groups", Group.Type),
	}
}
