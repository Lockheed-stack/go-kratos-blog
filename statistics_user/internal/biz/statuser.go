package biz

import (
	pb "statistics_user/api/stat_user"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

type StatUser struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Pv        uint `gorm:"DEFAULT:0"`
	Uv        uint `gorm:"DEFAULT:0"`
	Uid       uint
	User      User `gorm:"foreignKey:Uid"`
}
type User struct {
	gorm.Model
}

type StatUserUsecase struct {
	repo StatUserRepo
}

type StatUserRepo interface {
	GetUserSevenDaysData(uid uint64) ([]*pb.DayStatistics, error)
	SetUserTodayStatData(data []StatUser) error
}

func NewStatUserUsecase(repo StatUserRepo) *StatUserUsecase {
	return &StatUserUsecase{
		repo: repo,
	}
}

func (uc *StatUserUsecase) GetUserSevenDaysStatistics(uid uint64) ([]*pb.DayStatistics, error) {

	result, err := uc.repo.GetUserSevenDaysData(uid)
	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		default:
			return nil, pb.ErrorInvalidRequest("")
		}
	}
	return result, err
}

func (uc *StatUserUsecase) SetUserTodayStatistics(req []*pb.DayStatistics) error {
	users_data := make([]StatUser, len(req))
	now := time.Now()
	yesterday := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 0, 0, now.Location()).AddDate(0, 0, -1)
	for i, v := range req {
		users_data[i].Uid = uint(v.Uid)
		users_data[i].Pv = uint(v.Pv)
		users_data[i].Uv = uint(v.Uv)
		users_data[i].CreatedAt = yesterday
	}
	err := uc.repo.SetUserTodayStatData(users_data)
	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		case "INVALID_USER_ID":
			return pb.ErrorInvalidUserId("request contain invalid user id")
		default:
			return pb.ErrorInvalidRequest("bad request")
		}
	}
	return nil
}
