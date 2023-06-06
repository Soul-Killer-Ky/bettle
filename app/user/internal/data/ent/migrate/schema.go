// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "type", Type: field.TypeInt32},
		{Name: "name", Type: field.TypeString, Size: 30},
		{Name: "icon", Type: field.TypeString, Size: 50},
		{Name: "memo", Type: field.TypeString, Size: 50},
		{Name: "user_created_groups", Type: field.TypeInt, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_users_created_groups",
				Columns:    []*schema.Column{GroupsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupMembersColumns holds the columns for the "group_members" table.
	GroupMembersColumns = []*schema.Column{
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "role", Type: field.TypeInt32},
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupMembersTable holds the schema information for the "group_members" table.
	GroupMembersTable = &schema.Table{
		Name:       "group_members",
		Columns:    GroupMembersColumns,
		PrimaryKey: []*schema.Column{GroupMembersColumns[3], GroupMembersColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_members_groups_group",
				Columns:    []*schema.Column{GroupMembersColumns[3]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "group_members_users_user",
				Columns:    []*schema.Column{GroupMembersColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "phone", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 30},
		{Name: "password", Type: field.TypeString, Size: 100},
		{Name: "nickname", Type: field.TypeString, Size: 20},
		{Name: "avatar", Type: field.TypeString, Size: 50},
		{Name: "memo", Type: field.TypeString, Size: 50},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserFriendsColumns holds the columns for the "user_friends" table.
	UserFriendsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// UserFriendsTable holds the schema information for the "user_friends" table.
	UserFriendsTable = &schema.Table{
		Name:       "user_friends",
		Columns:    UserFriendsColumns,
		PrimaryKey: []*schema.Column{UserFriendsColumns[0], UserFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_friends_user_id",
				Columns:    []*schema.Column{UserFriendsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_friends_friend_id",
				Columns:    []*schema.Column{UserFriendsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsTable,
		GroupMembersTable,
		UsersTable,
		UserFriendsTable,
	}
)

func init() {
	GroupsTable.ForeignKeys[0].RefTable = UsersTable
	GroupMembersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupMembersTable.ForeignKeys[1].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[1].RefTable = UsersTable
}
