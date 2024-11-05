package service

import (
	"context"

	pb "user/api/users"
	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
)

type UsersService struct {
	pb.UnimplementedUsersServer
	uc *biz.UserUsecase
}

func NewUsersService(uc *biz.UserUsecase) *UsersService {
	return &UsersService{
		uc: uc,
	}
}

func (s *UsersService) CreateUsers(ctx context.Context, req *pb.CreateUsersRequest) (*pb.CreateUsersReply, error) {
	resp := &pb.CreateUsersReply{}
	err := s.uc.CreateOneUser(req.UserName, req.Password)
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
func (s *UsersService) UpdateUsers(ctx context.Context, req *pb.UpdateUsersRequest) (*pb.UpdateUsersReply, error) {
	return &pb.UpdateUsersReply{}, nil
}
func (s *UsersService) DeleteUsers(ctx context.Context, req *pb.DeleteUsersRequest) (*pb.DeleteUsersReply, error) {
	resp := &pb.DeleteUsersReply{}
	err := s.uc.RemoveOneUser(req.ID)
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
func (s *UsersService) GetSelectedUsers(ctx context.Context, req *pb.GetSelectedUsersRequest) (*pb.GetSelectedUsersReply, error) {
	resp := &pb.GetSelectedUsersReply{}
	result, err := s.uc.GetUsersByIDs(req.UsersID)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
	} else {
		resp.SelectedUsers = result
		resp.Code = 200
	}
	return resp, nil
}
func (s *UsersService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersReply, error) {
	return &pb.ListUsersReply{}, nil
}
func (s *UsersService) AuthUsers(ctx context.Context, req *pb.AuthUsersRequest) (*pb.AuthUsersReply, error) {
	resp := &pb.AuthUsersReply{}
	result, err := s.uc.UserLoginAuth(req.UserName, req.Password)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
	} else {
		resp.Code = 200
		resp.Msg = "OK"
		resp.SelectedUser = result
	}
	return resp, nil
}

func (s *UsersService) GetUserStatisticsInfo(ctx context.Context, req *pb.GetStatisticsRequest) (*pb.GetStatisticsReply, error) {
	resp := &pb.GetStatisticsReply{}
	result, err := s.uc.GetStatisticsInfoByID(req.ID)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
	} else {
		resp.Code = 200
		resp.Info = result
		resp.Msg = "OK"
	}
	return resp, nil
}
