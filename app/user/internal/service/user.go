package service

import (
	pb "beetle/api/user/service/v1"
	"beetle/app/user/internal/biz"
	"beetle/app/user/internal/conf"
	"beetle/internal/pkg/jwt"
	"context"
)

type UserService struct {
	pb.UnimplementedUserServer

	authc *conf.Auth
	uc    *biz.UserUseCase
	fc    *biz.FriendUseCase
}

func NewUserService(c *conf.Auth, uc *biz.UserUseCase, fc *biz.FriendUseCase) *UserService {
	return &UserService{authc: c, uc: uc, fc: fc}
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
		Username:    user.Username,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
	}, nil
}

func (s *UserService) GetFriend(ctx context.Context, req *pb.GetFriendRequest) (*pb.GetFriendReply, error) {
	userID, err := jwt.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	ps, err := s.fc.GetFriends(ctx, userID)
	friends := make([]*pb.GetFriendReply_Friend, 0)
	for _, p := range ps {
		friends = append(friends, &pb.GetFriendReply_Friend{
			Id:       int64(p.ID),
			UserId:   int64(p.UserID),
			Username: p.Username,
			Nickname: p.Nickname,
			Avatar:   p.Avatar,
			Memo:     p.Memo,
		})
	}
	return &pb.GetFriendReply{
		Friends: friends,
	}, nil
}
