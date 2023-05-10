package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"

	pb2 "beetle/api/beetle/v1"
	pb "beetle/api/im/service/v1"
	ws2 "beetle/app/im/internal/pkg/ws"
	"beetle/internal/pkg/jwt"
	"beetle/internal/pkg/ws"

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

// ImRepo is a Greater repo.
type ImRepo interface {
	Connect(ctx context.Context) error
	ListByUserID(context.Context, int) ([]*Group, error)
}

// ImUseCase is a Im usecase.
type ImUseCase struct {
	log *log.Helper

	repo ImRepo
}

// NewImUseCase new a im usecase.
func NewImUseCase(repo ImRepo, logger log.Logger) *ImUseCase {
	return &ImUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (ic *ImUseCase) Connect(ctx context.Context, req *http.Request, resp *http.ResponseWriter) error {
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

func (ic *ImUseCase) Save() {

}

func (ic *ImUseCase) GetGroups(ctx context.Context, userID int) ([]*Group, error) {
	groups, err := ic.repo.ListByUserID(ctx, userID)
	if err != nil {
		ic.log.WithContext(ctx).Errorf("list groups error: %s", err)
		return nil, errors.BadRequest("QUERY_ERROR", "list groups error")
	}

	return groups, nil
}
