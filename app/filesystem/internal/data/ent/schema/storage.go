package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Storage struct {
	ent.Schema
}

func (Storage) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("type").Comment("类型"),
		field.String("path").Comment("存储路径"),
		field.String("md5").Comment("md5编码").Unique(),
	}
}

func (Storage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CreateTimeMixin{},
		SoftDeleteMixin{},
	}
}

func (Storage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("quotes", Quote.Type),
	}
}
