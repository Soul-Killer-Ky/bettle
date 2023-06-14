package biz

import (
	"context"
	"time"

	"beetle/app/im/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type SynchronizedRecord struct {
	UserID        int
	DeviceID      uuid.UUID
	LastMessageID int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type SynchronizeRecordRepo interface {
	FindByUserDeviceID(context.Context, int, uuid.UUID) (*SynchronizedRecord, error)
	Create(context.Context, int, uuid.UUID, int64) (*SynchronizedRecord, error)
	Update(context.Context, int, uuid.UUID, int64) (*SynchronizedRecord, error)
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

func (c *SynchronizeRecordUseCase) SaveSynchronizeRecord(ctx context.Context, uid int, deviceID uuid.UUID, messageID int64) (*SynchronizedRecord, error) {
	p, err := c.repo.FindByUserDeviceID(ctx, uid, deviceID)
	if err != nil {
		if ent.IsNotFound(err) {
			p, err = c.repo.Create(ctx, uid, deviceID, messageID)
		} else {
			return nil, err
		}
	} else {
		p, err = c.repo.Update(ctx, uid, deviceID, messageID)
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

func (c *SynchronizeRecordUseCase) GetLastMessageID(ctx context.Context, uid int, deviceID uuid.UUID) (int64, error) {
	p, err := c.repo.FindByUserDeviceID(ctx, uid, deviceID)
	if err != nil {
		if ent.IsNotFound(err) {
			return 0, nil
		}
		return 0, err
	}
	return p.LastMessageID, nil
}
