package biz

import (
	"context"

	pb "beetle/api/user/service/v1"
	"beetle/app/user/internal/conf"
	"beetle/internal/pkg/jwt"
	"beetle/internal/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

// User is a User model.
type User struct {
	ID       int
	Username string
	Password string
	Nickname string
	Avatar   string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	FindByName(context.Context, string) (*User, error)
	ListByHello(context.Context, string) ([]*User, error)
	ListAll(context.Context) ([]*User, error)
}

// UserUseCase is a User use case.
type UserUseCase struct {
	authConf *conf.Auth
	userRepo UserRepo
	log      *log.Helper
}

// NewUserUseCase new a User use case.
func NewUserUseCase(authConf *conf.Auth, userRepo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{authConf: authConf, userRepo: userRepo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) VerifyUser(ctx context.Context, req *pb.LoginUserRequest) (*User, error) {
	uc.log.WithContext(ctx).Infof("VerifyUser: %v", req.Username)
	u, err := uc.userRepo.FindByName(ctx, req.Username)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("user not found: %s", err)
		return nil, pb.ErrorUserNotFound("user not found")
	}
	if !util.Md5Equal(u.Password, req.Password) {
		return nil, pb.ErrorUserPasswordError("user password error")
	}

	return u, nil
}

func (uc *UserUseCase) CreateToken(ctx context.Context, u *User) (string, error) {
	signedToken, err := jwt.CreateToken(uc.authConf.JwtKey, jwt.WithCustomUserClaims(u.ID, uc.authConf.JwtExp))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
