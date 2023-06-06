package service

import (
	pb "beetle/api/user/service/v1"
	"beetle/app/user/internal/biz"
	"beetle/app/user/internal/conf"
	"beetle/internal/pkg/jwt"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
)

type UserService struct {
	pb.UnimplementedUserServer

	authc *conf.Auth
	uc    *biz.UserUseCase
	fc    *biz.FriendUseCase
	gc    *biz.GroupUseCase
}

func NewUserService(c *conf.Auth, uc *biz.UserUseCase, fc *biz.FriendUseCase, gc *biz.GroupUseCase) *UserService {
	return &UserService{authc: c, uc: uc, fc: fc, gc: gc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserReply, error) {
	user, err := s.uc.VerifyUser(ctx, req)
	if err != nil {
		return nil, err
	}
	token, err := s.uc.CreateToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.LoginUserReply{
		Token:       token,
		ExpiredTime: s.authc.JwtExp,
		Id:          int64(user.ID),
		Phone:       user.Phone,
		Username:    *user.Username,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	u, err := s.uc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		UserId:   int64(u.ID),
		Phone:    u.Phone,
		Username: *u.Username,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Memo:     u.Memo,
	}, nil
}

func (s *UserService) ListFriend(ctx context.Context, _ *pb.ListFriendRequest) (*pb.ListFriendReply, error) {
	userID, err := jwt.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	ps, err := s.fc.ListFriend(ctx, userID)
	friends := make([]*pb.Friend, 0)
	for _, p := range ps {
		friends = append(friends, &pb.Friend{
			UserId:   int64(p.UserID),
			Phone:    p.Phone,
			Username: *p.Username,
			Nickname: p.Nickname,
			Avatar:   p.Avatar,
			Memo:     p.Memo,
		})
	}

	return &pb.ListFriendReply{
		Friends: friends,
	}, nil
}

func (s *UserService) AddFriend(ctx context.Context, req *pb.AddFriendRequest) (*pb.AddFriendReply, error) {
	userID, err := jwt.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	u, err := s.uc.GetUserByPhone(ctx, req.GetPhone())
	if err != nil {
		return nil, err
	}
	if u.ID == userID {
		return nil, errors.Forbidden("FORBIDDEN", "Prohibit adding yourself")
	}
	friendUserID := u.ID
	p, err := s.fc.AddFriend(ctx, userID, friendUserID)
	if err != nil {
		return nil, err
	}

	return &pb.AddFriendReply{
		Friend: &pb.Friend{
			UserId:   int64(p.UserID),
			Phone:    p.Phone,
			Username: *p.Username,
			Nickname: p.Nickname,
			Avatar:   p.Avatar,
			Memo:     p.Memo,
		},
	}, nil
}

func (s *UserService) ListGroup(ctx context.Context, req *pb.ListGroupRequest) (*pb.ListGroupReply, error) {
	userID, err := jwt.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	ps, err := s.gc.ListGroup(ctx, userID)
	groups := make([]*pb.Group, 0)
	for _, p := range ps {
		groups = append(groups, &pb.Group{
			Id:   int64(p.ID),
			Type: p.Type,
			Name: p.Name,
			Icon: p.Icon,
			Memo: p.Memo,
		})
	}
	return &pb.ListGroupReply{
		Groups: groups,
	}, nil
}

func (s *UserService) JoinGroup(ctx context.Context, req *pb.JoinGroupRequest) (*pb.JoinGroupReply, error) {
	userID, err := jwt.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = s.gc.JoinGroup(ctx, userID, req.GroupName)
	if err != nil {
		return nil, err
	}
	return &pb.JoinGroupReply{}, nil
}

func (s *UserService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupReply, error) {
	userID, err := jwt.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	gp, err := s.gc.CreateGroup(ctx, userID, req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateGroupReply{
		Group: &pb.Group{
			Id:   int64(gp.ID),
			Type: gp.Type,
			Name: gp.Name,
			Icon: gp.Icon,
			Memo: gp.Memo,
		},
	}, nil
}
