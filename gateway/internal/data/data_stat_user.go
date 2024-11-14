package data

import (
	"context"
	"gateway/api/stat_user"
	"gateway/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayStatUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayStatUserRepo(data *Data, logger log.Logger) biz.GatewayStatUserRepo {
	return &gatewayStatUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *gatewayStatUserRepo) GRPC_GetUserSevenDayStat(req *stat_user.GetUserSevenDaysStatRequest) (*stat_user.GetUserSevenDaysStatReply, error) {

	client := stat_user.NewStatUserClient(r.data.ConnGRPC_stat_user)
	result, err := client.GetUserSevenDaysStat(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}

func (r *gatewayStatUserRepo) GRPC_SetUserTodayStatData(req *stat_user.SetUserStatInfoRequest) (*stat_user.SetUserStatInfoReply, error) {
	client := stat_user.NewStatUserClient(r.data.ConnGRPC_stat_user)
	result, err := client.SetUserStatInfo(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
