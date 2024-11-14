package data

import (
	pb "statistics_user/api/stat_user"
	"statistics_user/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
)

type statUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewStatUserRepo(data *Data, logger log.Logger) biz.StatUserRepo {
	return &statUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *statUserRepo) GetUserSevenDaysData(uid uint64) ([]*pb.DayStatistics, error) {
	result := []*pb.DayStatistics{}

	sqlRes := r.data.db.Model(&biz.StatUser{}).Select("created_at", "pv", "uv").Where("uid=? and created_at >= CURDATE() - INTERVAL ? DAY", uid, 7).Scan(&result)
	if err := sqlRes.Error; err != nil {
		r.log.Error(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			default:
				return nil, errors.New(400, "INVALID_REQUEST", "")
			}
		}
	}
	return result, nil
}

func (r *statUserRepo) SetUserTodayStatData(data []biz.StatUser) error {
	sqlRes := r.data.db.Create(&data)
	if err := sqlRes.Error; err != nil {
		r.log.Error(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1452: // a foreign key constraint fails
				return errors.New(400, "INVALID_USER_ID", "")
			default:
				return errors.New(400, "INVALID_REQUEST", "")
			}
		}
	}
	return nil
}
