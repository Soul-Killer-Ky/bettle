package biz

import (
	"context"

	pb "beetle/api/user/service/v1"
	"beetle/app/user/internal/conf"
	"beetle/internal/pkg/jwt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(pb.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type Consume struct {
	ID         int64
	UserID     int64
	OrderID    string
	OrderPrice int64
}

type ConsumeRepo interface {
	CreateConsume(ctx context.Context, a *Consume) (int64, error)
}

// User is a User model.
type User struct {
	gorm.Model
	Username string `gorm:"size:30;unique;<-:create"`
	Password string `gorm:"size:100"`
	Nickname string `gorm:"size:20"`
	Avatar   string `gorm:"size:50"`
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

// UserUseCase is a User usecase.
type UserUseCase struct {
	authc       *conf.Auth
	consumeRepo ConsumeRepo
	userRepo    UserRepo
	tx          Transaction
	log         *log.Helper
}

// NewUserUseCase new a User usecase.
func NewUserUseCase(authc *conf.Auth, repo ConsumeRepo, userRepo UserRepo, tx Transaction, logger log.Logger) *UserUseCase {
	return &UserUseCase{authc: authc, consumeRepo: repo, userRepo: userRepo, tx: tx, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Consume(ctx context.Context, c *Consume) (int64, error) {
	var (
		err error
		id  int64
	)
	// 调用事务实例
	err = uc.tx.ExecTx(ctx, func(ctx context.Context) error {
		id, err = uc.consumeRepo.CreateConsume(ctx, c)
		if err != nil {
			return err
		}

		_, err = uc.userRepo.Save(ctx, &User{
			Username: "",
		})
		return err
	})
	return id, err
}

func (uc *UserUseCase) VerifyUser(ctx context.Context, req *pb.LoginUserRequest) (*User, error) {
	uc.log.WithContext(ctx).Infof("VerifyUser: %v", req.Username)
	if req.Username != "root" {
		//return nil, pb.ErrorUserAccountFormatError("user account must be %s", "root")
	}
	if len(req.Password) != 6 {
		return nil, pb.ErrorUserPasswordFormatError("user password's length must be %d", 6)
	}
	u, err := uc.userRepo.FindByName(ctx, req.Username)
	if err != nil {
		return nil, pb.ErrorUserNotFound("user not found")
	}
	if u.Password != req.Password {
		return nil, pb.ErrorUserPasswordError("user password error")
	}
	return u, nil
}

func (uc *UserUseCase) CreateToken(ctx context.Context, u *User) (string, error) {
	signedToken, err := jwt.CreateToken(uc.authc.JwtKey, jwt.WithCustomUserClaims(u.ID, uc.authc.JwtExp))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
