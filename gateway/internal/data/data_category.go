package data

import (
	"context"
	"gateway/api/category"
	"gateway/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayCategoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayCategoryRepo(data *Data, logger log.Logger) biz.GatewayCategoryRepo {
	return &gatewayCategoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *gatewayCategoryRepo) GRPC_CreateOneCategory(req *category.CreateCategoryRequest) (*category.CreateCategoryReply, error) {
	client := category.NewCategoryClient(r.data.ConnGRPC_category)
	result, err := client.CreateCategory(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayCategoryRepo) GRPC_ListCategory(req *category.ListCategoryRequest) (*category.ListCategoryReply, error) {

	/* ------------- fast path ---------------------*/
	key := "category_list"
	categoryList, err := GetCategoryRedis(r.data.Redis_cli, key)
	if err == nil { // redis cache matched
		result := &category.ListCategoryReply{
			CategoryArray: categoryList,
			Code:          200,
		}
		return result, nil
	}

	/*------------- slow path ----------------------*/
	client := category.NewCategoryClient(r.data.ConnGRPC_category)
	result, err := client.ListCategory(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayCategoryRepo) GRPC_DeleteOneCategory(req *category.DeleteCategoryRequest) (*category.DeleteCategoryReply, error) {
	client := category.NewCategoryClient(r.data.ConnGRPC_category)
	result, err := client.DeleteCategory(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayCategoryRepo) GRPC_UpdateOneCategory(req *category.UpdateCategoryRequest) (*category.UpdateCategoryReply, error) {
	client := category.NewCategoryClient(r.data.ConnGRPC_category)
	result, err := client.UpdateCategory(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
