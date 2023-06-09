package data

import (
	"beetle/app/im/internal/biz"
	"beetle/app/im/internal/data/ent/synchronizerecord"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type synchronizeRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewSynchronizeRecordRepo(data *Data, logger log.Logger) biz.SynchronizeRecordRepo {
	return &synchronizeRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *synchronizeRecordRepo) FindByUserDeviceID(ctx context.Context, uid int, deviceID uuid.UUID) (*biz.SynchronizedRecord, error) {
	p, err := r.data.db.SynchronizeRecord.Query().Where(synchronizerecord.UserID(uid), synchronizerecord.DeviceID(deviceID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.SynchronizedRecord{
		UserID:    p.UserID,
		DeviceID:  p.DeviceID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, err
}

func (r *synchronizeRecordRepo) Create(ctx context.Context, uid int, deviceID uuid.UUID) (*biz.SynchronizedRecord, error) {
	p, err := r.data.db.SynchronizeRecord.Create().SetUserID(uid).SetDeviceID(deviceID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.SynchronizedRecord{
		UserID:    p.UserID,
		DeviceID:  p.DeviceID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (r *synchronizeRecordRepo) Update(ctx context.Context, uid int, deviceID uuid.UUID) (*biz.SynchronizedRecord, error) {
	p, err := r.data.db.SynchronizeRecord.Query().Where(synchronizerecord.UserID(uid), synchronizerecord.DeviceID(deviceID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	p, err = p.Update().Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.SynchronizedRecord{
		UserID:    p.UserID,
		DeviceID:  p.DeviceID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}
