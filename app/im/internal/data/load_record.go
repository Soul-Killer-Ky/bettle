package data

import (
	"context"

	pbUser "beetle/api/user/service/v1"
	"beetle/app/im/internal/biz"
	"beetle/app/im/internal/data/ent"
	"beetle/app/im/internal/data/ent/loadrecord"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type loadRecordRepo struct {
	data *Data
	log  *log.Helper
}

// NewLoadRecordRepo .
func NewLoadRecordRepo(data *Data, logger log.Logger) biz.LoadRecordRepo {
	return &loadRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *loadRecordRepo) SaveLoadRecord(ctx context.Context, uid int, deviceID uuid.UUID) error {
	_, err := r.data.db.LoadRecord.Create().SetUserID(uid).SetDeviceID(deviceID).Save(ctx)
	return err
}

func (r *loadRecordRepo) GetLastLoadRecord(ctx context.Context, uid int, deviceID uuid.UUID) (*ent.LoadRecord, error) {
	rs, err := r.data.db.LoadRecord.Query().Where(loadrecord.UserID(uid), loadrecord.DeviceID(deviceID)).
		Order(loadrecord.ByCreatedAt(sql.OrderDesc())).First(ctx)
	return rs, err
}

func (r *loadRecordRepo) GetUser(ctx context.Context, uid int) (*pbUser.GetUserReply, error) {
	svc, err := r.data.registry.GetService(context.Background(), "user-svc")
	log.Infof("svc: %v, err: %v", svc, err)
	svc, err = r.data.registry.GetService(context.Background(), "im-svc")
	log.Infof("svc: %v, err: %v", svc, err)
	reply, err := r.data.uc.GetUser(ctx, &pbUser.GetUserRequest{Id: int64(uid)})
	return reply, err
}
