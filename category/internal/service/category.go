package service

import (
	"context"

	pb "category/api/category"
	"category/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

type CategoryService struct {
	pb.UnimplementedCategoryServer
	uc *biz.CategoryUsecase
}

func NewCategoryService(uc *biz.CategoryUsecase) *CategoryService {
	return &CategoryService{
		uc: uc,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryReply, error) {
	resp := &pb.CreateCategoryReply{}
	err := s.uc.CreateOneCategory(req.CategoryName)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}
func (s *CategoryService) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	resp := &pb.ListCategoryReply{}
	categories, err := s.uc.GetSelectedCategory(req.PageSize, req.PageNum)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
	} else {
		resp.CategoryArray = categories
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryReply, error) {
	resp := &pb.DeleteCategoryReply{}
	err := s.uc.DeleteOneCategory(req.ID)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}
func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryReply, error) {
	resp := &pb.UpdateCategoryReply{}
	category := &biz.Category{
		Model: gorm.Model{ID: uint(req.CategoryID)},
		Name:  req.CategoryName,
	}
	err := s.uc.UpdateOneCategory(req.CategoryID, category)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}
