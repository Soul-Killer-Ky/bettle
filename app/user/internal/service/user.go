package service

import (
	pb "beetle/api/user/service/v1"
	"beetle/app/user/internal/biz"
	"beetle/app/user/internal/conf"
	"context"
)

type UserService struct {
	pb.UnimplementedUserServer

	authc *conf.Auth
	uc    *biz.UserUseCase
}

func NewUserService(c *conf.Auth, uc *biz.UserUseCase) *UserService {
	return &UserService{authc: c, uc: uc}
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
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
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
		Username:    user.Username,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
	}, nil
}
