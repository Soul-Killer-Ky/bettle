// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 30},
		{Name: "password", Type: field.TypeString, Size: 100},
		{Name: "nickname", Type: field.TypeString, Size: 20},
		{Name: "avatar", Type: field.TypeString, Size: 50},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}
