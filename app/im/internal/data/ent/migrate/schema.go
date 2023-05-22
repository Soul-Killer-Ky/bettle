// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatMessagesColumns holds the columns for the "chat_messages" table.
	ChatMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "from", Type: field.TypeInt},
		{Name: "sender", Type: field.TypeInt},
		{Name: "message", Type: field.TypeString, Size: 2147483647},
	}
	// ChatMessagesTable holds the schema information for the "chat_messages" table.
	ChatMessagesTable = &schema.Table{
		Name:       "chat_messages",
		Columns:    ChatMessagesColumns,
		PrimaryKey: []*schema.Column{ChatMessagesColumns[0]},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 30},
		{Name: "icon", Type: field.TypeString, Size: 50},
		{Name: "memo", Type: field.TypeString, Size: 50},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatMessagesTable,
		GroupsTable,
	}
)

func init() {
}
