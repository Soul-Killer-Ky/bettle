// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeJoinedGroups holds the string denoting the joined_groups edge name in mutations.
	EdgeJoinedGroups = "joined_groups"
	// EdgeCreatedGroups holds the string denoting the created_groups edge name in mutations.
	EdgeCreatedGroups = "created_groups"
	// EdgeGroupMembers holds the string denoting the group_members edge name in mutations.
	EdgeGroupMembers = "group_members"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "user_friends"
	// JoinedGroupsTable is the table that holds the joined_groups relation/edge. The primary key declared below.
	JoinedGroupsTable = "group_members"
	// JoinedGroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	JoinedGroupsInverseTable = "groups"
	// CreatedGroupsTable is the table that holds the created_groups relation/edge.
	CreatedGroupsTable = "groups"
	// CreatedGroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	CreatedGroupsInverseTable = "groups"
	// CreatedGroupsColumn is the table column denoting the created_groups relation/edge.
	CreatedGroupsColumn = "user_created_groups"
	// GroupMembersTable is the table that holds the group_members relation/edge.
	GroupMembersTable = "group_members"
	// GroupMembersInverseTable is the table name for the GroupMember entity.
	// It exists in this package in order to avoid circular dependency with the "groupmember" package.
	GroupMembersInverseTable = "group_members"
	// GroupMembersColumn is the table column denoting the group_members relation/edge.
	GroupMembersColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldPhone,
	FieldUsername,
	FieldPassword,
	FieldNickname,
	FieldAvatar,
	FieldMemo,
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
	// JoinedGroupsPrimaryKey and JoinedGroupsColumn2 are the table columns denoting the
	// primary key for the joined_groups relation (M2M).
	JoinedGroupsPrimaryKey = []string{"group_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "beetle/app/user/internal/data/ent/runtime"
//
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	AvatarValidator func(string) error
	// MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	MemoValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// ByFriendsCount orders the results by friends count.
func ByFriendsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFriendsStep(), opts...)
	}
}

// ByFriends orders the results by friends terms.
func ByFriends(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByJoinedGroupsCount orders the results by joined_groups count.
func ByJoinedGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newJoinedGroupsStep(), opts...)
	}
}

// ByJoinedGroups orders the results by joined_groups terms.
func ByJoinedGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newJoinedGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedGroupsCount orders the results by created_groups count.
func ByCreatedGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedGroupsStep(), opts...)
	}
}

// ByCreatedGroups orders the results by created_groups terms.
func ByCreatedGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupMembersCount orders the results by group_members count.
func ByGroupMembersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupMembersStep(), opts...)
	}
}

// ByGroupMembers orders the results by group_members terms.
func ByGroupMembers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupMembersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFriendsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, FriendsTable, FriendsPrimaryKey...),
	)
}
func newJoinedGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(JoinedGroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, JoinedGroupsTable, JoinedGroupsPrimaryKey...),
	)
}
func newCreatedGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedGroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedGroupsTable, CreatedGroupsColumn),
	)
}
func newGroupMembersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupMembersInverseTable, GroupMembersColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, GroupMembersTable, GroupMembersColumn),
	)
}
