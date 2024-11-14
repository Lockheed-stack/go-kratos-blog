package service

import (
	"context"

	pb "statistics_user/api/stat_user"
	"statistics_user/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
)

type StatUserService struct {
	pb.UnimplementedStatUserServer
	uc *biz.StatUserUsecase
}

func NewStatUserService(uc *biz.StatUserUsecase) *StatUserService {
	return &StatUserService{
		uc: uc,
	}
}

func (s *StatUserService) GetUserSevenDaysStat(ctx context.Context, req *pb.GetUserSevenDaysStatRequest) (*pb.GetUserSevenDaysStatReply, error) {
	resp := &pb.GetUserSevenDaysStatReply{}
	result, err := s.uc.GetUserSevenDaysStatistics(req.Uid)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Message
	} else {
		resp.SevenDaysData = result
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}
func (s *StatUserService) SetUserStatInfo(ctx context.Context, req *pb.SetUserStatInfoRequest) (*pb.SetUserStatInfoReply, error) {
	resp := &pb.SetUserStatInfoReply{}

	err := s.uc.SetUserTodayStatistics(req.Data)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Message
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}

	return resp, nil
}
