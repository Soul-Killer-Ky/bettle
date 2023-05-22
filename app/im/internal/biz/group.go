package biz

import (
	pb2 "beetle/api/beetle/v1"
	pb "beetle/api/im/service/v1"
	ws2 "beetle/app/im/internal/pkg/ws"
	"beetle/internal/pkg/jwt"
	"beetle/internal/pkg/ws"
	"context"
	"github.com/go-kratos/kratos/v2/errors"

	"github.com/go-kratos/kratos/v2/log"
	jwt2 "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Group struct {
	ID   int
	Name string
	Icon string
	Memo string
}

// GroupRepo is a Greater repo.
type GroupRepo interface {
	Connect(ctx context.Context) error
	ListByUserID(context.Context, int) ([]*Group, error)
}

// GroupUseCase is a group usecase.
type GroupUseCase struct {
	log *log.Helper

	repo GroupRepo
}

// NewGroupUseCase new a im usecase.
func NewGroupUseCase(repo GroupRepo, logger log.Logger) *GroupUseCase {
	return &GroupUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (ic *GroupUseCase) Connect(ctx context.Context, req *http.Request, resp *http.ResponseWriter) error {
	token, ok := jwt2.FromContext(ctx)
	if !ok {
		return pb2.ErrorAuthFailure("auth fail")
	}
	claims := token.(jwt.CustomUserClaims)
	ic.log.WithContext(ctx).Infof("Connect: %v", claims.ID)
	conn, err := ws.App(*resp, req)
	if err != nil {
		return pb.ErrorImConnectFailure("Failure to establish connection")
	}
	client := ws2.NewImClient(uint(claims.ID), conn, ic.log)

	go client.Read()
	go client.Write()

	return nil
}

func (ic *GroupUseCase) GetGroups(ctx context.Context, userID int) ([]*Group, error) {
	groups, err := ic.repo.ListByUserID(ctx, userID)
	if err != nil {
		ic.log.WithContext(ctx).Errorf("list groups error: %s", err)
		return nil, errors.BadRequest("QUERY_ERROR", "list groups error")
	}

	return groups, nil
}
