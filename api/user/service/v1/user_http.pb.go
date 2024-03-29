// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.6.1
// source: user/service/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserAddFriend = "/api.user.service.v1.User/AddFriend"
const OperationUserCreateGroup = "/api.user.service.v1.User/CreateGroup"
const OperationUserCreateUser = "/api.user.service.v1.User/CreateUser"
const OperationUserDeleteUser = "/api.user.service.v1.User/DeleteUser"
const OperationUserGetUser = "/api.user.service.v1.User/GetUser"
const OperationUserJoinGroup = "/api.user.service.v1.User/JoinGroup"
const OperationUserListFriend = "/api.user.service.v1.User/ListFriend"
const OperationUserListGroup = "/api.user.service.v1.User/ListGroup"
const OperationUserListUser = "/api.user.service.v1.User/ListUser"
const OperationUserListUserByGroup = "/api.user.service.v1.User/ListUserByGroup"
const OperationUserLoginUser = "/api.user.service.v1.User/LoginUser"
const OperationUserUpdateUser = "/api.user.service.v1.User/UpdateUser"

type UserHTTPServer interface {
	AddFriend(context.Context, *AddFriendRequest) (*AddFriendReply, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupReply, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupReply, error)
	ListFriend(context.Context, *ListFriendRequest) (*ListFriendReply, error)
	ListGroup(context.Context, *ListGroupRequest) (*ListGroupReply, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	ListUserByGroup(context.Context, *ListUserByGroupRequest) (*ListUserByGroupReply, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/user/v1/create", _User_CreateUser0_HTTP_Handler(srv))
	r.POST("/user/v1/update", _User_UpdateUser0_HTTP_Handler(srv))
	r.POST("/user/v1/delete", _User_DeleteUser0_HTTP_Handler(srv))
	r.POST("/user/v1/list", _User_ListUser0_HTTP_Handler(srv))
	r.POST("/user/v1/login", _User_LoginUser0_HTTP_Handler(srv))
	r.POST("/user/v1/get", _User_GetUser0_HTTP_Handler(srv))
	r.POST("/user/v1/friend/list", _User_ListFriend0_HTTP_Handler(srv))
	r.POST("/user/v1/friend/add", _User_AddFriend0_HTTP_Handler(srv))
	r.POST("/user/v1/group/list", _User_ListGroup0_HTTP_Handler(srv))
	r.POST("/user/v1/group/join", _User_JoinGroup0_HTTP_Handler(srv))
	r.POST("/user/v1/group/create", _User_CreateGroup0_HTTP_Handler(srv))
	r.POST("/user/v1/group/list_user", _User_ListUserByGroup0_HTTP_Handler(srv))
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_LoginUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLoginUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginUser(ctx, req.(*LoginUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListFriend0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListFriendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListFriend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListFriend(ctx, req.(*ListFriendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListFriendReply)
		return ctx.Result(200, reply)
	}
}

func _User_AddFriend0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddFriendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAddFriend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddFriend(ctx, req.(*AddFriendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddFriendReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListGroup0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListGroup(ctx, req.(*ListGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListGroupReply)
		return ctx.Result(200, reply)
	}
}

func _User_JoinGroup0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in JoinGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserJoinGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.JoinGroup(ctx, req.(*JoinGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*JoinGroupReply)
		return ctx.Result(200, reply)
	}
}

func _User_CreateGroup0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateGroup(ctx, req.(*CreateGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateGroupReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListUserByGroup0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserByGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListUserByGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserByGroup(ctx, req.(*ListUserByGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserByGroupReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	AddFriend(ctx context.Context, req *AddFriendRequest, opts ...http.CallOption) (rsp *AddFriendReply, err error)
	CreateGroup(ctx context.Context, req *CreateGroupRequest, opts ...http.CallOption) (rsp *CreateGroupReply, err error)
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	JoinGroup(ctx context.Context, req *JoinGroupRequest, opts ...http.CallOption) (rsp *JoinGroupReply, err error)
	ListFriend(ctx context.Context, req *ListFriendRequest, opts ...http.CallOption) (rsp *ListFriendReply, err error)
	ListGroup(ctx context.Context, req *ListGroupRequest, opts ...http.CallOption) (rsp *ListGroupReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	ListUserByGroup(ctx context.Context, req *ListUserByGroupRequest, opts ...http.CallOption) (rsp *ListUserByGroupReply, err error)
	LoginUser(ctx context.Context, req *LoginUserRequest, opts ...http.CallOption) (rsp *LoginUserReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...http.CallOption) (*AddFriendReply, error) {
	var out AddFriendReply
	pattern := "/user/v1/friend/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAddFriend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...http.CallOption) (*CreateGroupReply, error) {
	var out CreateGroupReply
	pattern := "/user/v1/group/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/user/v1/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserReply, error) {
	var out DeleteUserReply
	pattern := "/user/v1/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/user/v1/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...http.CallOption) (*JoinGroupReply, error) {
	var out JoinGroupReply
	pattern := "/user/v1/group/join"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserJoinGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListFriend(ctx context.Context, in *ListFriendRequest, opts ...http.CallOption) (*ListFriendReply, error) {
	var out ListFriendReply
	pattern := "/user/v1/friend/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserListFriend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListGroup(ctx context.Context, in *ListGroupRequest, opts ...http.CallOption) (*ListGroupReply, error) {
	var out ListGroupReply
	pattern := "/user/v1/group/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserListGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/user/v1/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListUserByGroup(ctx context.Context, in *ListUserByGroupRequest, opts ...http.CallOption) (*ListUserByGroupReply, error) {
	var out ListUserByGroupReply
	pattern := "/user/v1/group/list_user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserListUserByGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...http.CallOption) (*LoginUserReply, error) {
	var out LoginUserReply
	pattern := "/user/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLoginUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserReply, error) {
	var out UpdateUserReply
	pattern := "/user/v1/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
