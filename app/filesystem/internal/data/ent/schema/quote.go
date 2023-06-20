package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Quote struct {
	ent.Schema
}

func (Quote) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.Int("created_by").Comment("上传者"),
	}
}

func (Quote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		SoftDeleteMixin{},
	}
}

func (Quote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("storage", Storage.Type).Ref("quotes").Unique(),
	}
}
