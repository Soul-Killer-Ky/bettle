// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"beetle/app/user/internal/data/ent/schema"
	"beetle/app/user/internal/data/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	userMixinHooks2 := userMixin[2].Hooks()
	user.Hooks[0] = userMixinHooks2[0]
	userMixinInters2 := userMixin[2].Interceptors()
	user.Interceptors[0] = userMixinInters2[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields1[0].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[2].Descriptor()
	// user.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	user.NicknameValidator = userDescNickname.Validators[0].(func(string) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[3].Descriptor()
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = userDescAvatar.Validators[0].(func(string) error)
}

const (
	Version = "v0.11.9"                                         // Version of ent codegen.
	Sum     = "h1:dbbCkAiPVTRBIJwoZctiSYjB7zxQIBOzVSU5H9VYIQI=" // Sum of ent codegen.
)
