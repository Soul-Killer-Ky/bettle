package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(30).Comment("名称"),
		field.String("icon").
			MaxLen(50).Comment("图标"),
		field.String("memo").
			MaxLen(50).Comment("个人简介"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}
