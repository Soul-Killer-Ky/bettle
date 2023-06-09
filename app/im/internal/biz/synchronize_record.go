package biz

import (
	"context"
	"time"

	"beetle/app/im/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type SynchronizedRecord struct {
	UserID    int
	DeviceID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SynchronizeRecordRepo interface {
	FindByUserDeviceID(context.Context, int, uuid.UUID) (*SynchronizedRecord, error)
	Create(context.Context, int, uuid.UUID) (*SynchronizedRecord, error)
	Update(context.Context, int, uuid.UUID) (*SynchronizedRecord, error)
}

type SynchronizeRecordUseCase struct {
	log *log.Helper

	repo SynchronizeRecordRepo
}

func NewLoadRecordUseCase(repo SynchronizeRecordRepo, logger log.Logger) *SynchronizeRecordUseCase {
	return &SynchronizeRecordUseCase{
		repo: repo,
		log:  log.NewHelper(log.NewFilter(logger, log.FilterLevel(log.LevelDebug))),
	}
}

func (c *SynchronizeRecordUseCase) SaveSynchronizeRecord(ctx context.Context, uid int, deviceID uuid.UUID) (*SynchronizedRecord, error) {
	p, err := c.repo.FindByUserDeviceID(ctx, uid, deviceID)
	if err != nil {
		if ent.IsNotFound(err) {
			p, err = c.repo.Create(ctx, uid, deviceID)
		} else {
			return nil, err
		}
	} else {
		p, err = c.repo.Update(ctx, uid, deviceID)
	}
	return p, err
}

func (c *SynchronizeRecordUseCase) GetLastTime(ctx context.Context, uid int, deviceID uuid.UUID) (*time.Time, error) {
	p, err := c.repo.FindByUserDeviceID(ctx, uid, deviceID)
	if err != nil {
		if ent.IsNotFound(err) {
			defaultTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
			return &defaultTime, nil
		}
		return nil, err
	}
	return &p.UpdatedAt, nil
}
