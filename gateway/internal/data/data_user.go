package data

import (
	"context"
	"gateway/api/users"
	"gateway/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayUserRepo(data *Data, logger log.Logger) biz.GatewayUserRepo {
	return &gatewayUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *gatewayUserRepo) GRPC_CreateUser(req *users.CreateUsersRequest) (*users.CreateUsersReply, error) {
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	result, err := client.CreateUsers(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayUserRepo) GRPC_DeleteUser(req *users.DeleteUsersRequest) (*users.DeleteUsersReply, error) {
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	result, err := client.DeleteUsers(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayUserRepo) GRPC_AuthUser(req *users.AuthUsersRequest) (*users.AuthUsersReply, error) {
	client := users.NewUsersClient(r.data.ConnGRPC_user)
	result, err := client.AuthUsers(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
