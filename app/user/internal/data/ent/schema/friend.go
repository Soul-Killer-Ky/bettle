package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type Friend struct {
	ent.Schema
}

func (Friend) Fields() []ent.Field {
	return []ent.Field{}
}

func (Friend) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		UpdateTimeMixin{},
		SoftDeleteMixin{},
	}
}

func (Friend) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("users").
			Unique(),
		edge.From("friend", User.Type).
			Ref("friends").
			Unique(),
	}
}
