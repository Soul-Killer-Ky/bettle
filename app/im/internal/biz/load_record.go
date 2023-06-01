package biz

import (
	pbUser "beetle/api/user/service/v1"
	"context"
	"time"

	"beetle/app/im/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type LoadRecordRepo interface {
	SaveLoadRecord(context.Context, int, uuid.UUID) error
	GetLastLoadRecord(context.Context, int, uuid.UUID) (*ent.LoadRecord, error)
	GetUser(context.Context, int) (*pbUser.GetUserReply, error)
}

type LoadRecordUseCase struct {
	log *log.Helper

	repo LoadRecordRepo
}

func NewLoadRecordUseCase(repo LoadRecordRepo, logger log.Logger) *LoadRecordUseCase {
	return &LoadRecordUseCase{
		repo: repo,
		log:  log.NewHelper(log.NewFilter(logger, log.FilterLevel(log.LevelDebug))),
	}
}

func (lc *LoadRecordUseCase) SaveLoadRecord(ctx context.Context, uid int, deviceID uuid.UUID) error {
	return lc.repo.SaveLoadRecord(ctx, uid, deviceID)
}

func (lc *LoadRecordUseCase) GetLastTime(ctx context.Context, uid int, deviceID uuid.UUID) (*time.Time, error) {
	rs, err := lc.repo.GetLastLoadRecord(ctx, uid, deviceID)
	if err != nil {
		if ent.IsNotFound(err) {
			defaultTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
			return &defaultTime, nil
		}
		return nil, err
	}
	return &rs.CreatedAt, nil
}

func (lc *LoadRecordUseCase) Test(ctx context.Context, uid int) {
	user, err := lc.repo.GetUser(ctx, uid)
	if err != nil {
		lc.log.Infof("user: %v, err: %v", user, err)
	} else {
		lc.log.Infof("user: %v", user)
	}
}
