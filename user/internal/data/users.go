package data

import (
	pb "user/api/users"
	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CheckDuplicateUsername(name string) bool {
	sqlRes := r.data.db.Model(&biz.User{}).Select("name").Where("name=?", name)
	if sqlRes.Error != nil {
		r.log.Error(sqlRes.Error)
		return true
	}
	if sqlRes.RowsAffected > 0 {
		return true
	}
	return false
}

func (r *userRepo) CreateUser(username string, passwd string) error {
	user := &biz.User{
		Username: username,
		Password: passwd,
	}
	sqlRes := r.data.db.Create(user)
	if err := sqlRes.Error; err != nil {
		r.log.Error(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062: // duplicate entry
				return errors.New(400, "ERR_USER_PRE_EXISTING", "")
			default:
				return errors.New(400, "ERR_USER_INVALID_REQUEST", "")
			}
		}
	}
	return nil
}

func (r *userRepo) RemoveUser(id uint64) error {

	user := &biz.User{}

	sqlRes := r.data.db.Where("id=?", id).Delete(user)
	if err := sqlRes.Error; err != nil {
		r.log.Error(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			default:
				return errors.New(400, "ERR_USER_INVALID_REQUEST", "")
			}
		}
	}

	if sqlRes.RowsAffected == 0 {
		return errors.New(404, "ERR_USER_NOT_FOUND", "")
	}

	return nil
}

func (r *userRepo) AuthLogin(name string, pwd string) (*pb.UserPublicInfo, error) {

	user := &biz.User{}
	r.data.db.Where("username=?", name).First(user)

	if user.ID == 0 || user.Password != pwd {
		return nil, errors.New(400, "ERR_USER_USERNAME_PASSWORD_WRONG", "")
	}
	result := &pb.UserPublicInfo{}
	result.ID = uint64(user.ID)
	result.Avatar = user.Avatar
	result.SelfDesc = user.SelfDesc
	result.Username = user.Username

	return result, nil
}

func (r *userRepo) GetSelectedUsers(selectedFields []string, IDs []uint64) ([]*pb.UserPublicInfo, error) {
	result := []*pb.UserPublicInfo{}
	sqlRes := r.data.db.Model(&biz.User{}).Select(selectedFields).Where("id IN ?", IDs).Scan(&result)
	if err := sqlRes.Error; err != nil {
		r.log.Error(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			default:
				return nil, errors.New(400, "ERR_USER_INVALID_REQUEST", "")
			}
		}
	}
	return result, nil
}

func (r *userRepo) GetUserStatitics(id uint64) (*pb.StatisticsInfo, error) {
	result := &pb.StatisticsInfo{}
	// user := &biz.User{}
	selectedFidlds := []string{
		"total_login_days",
		"total_blogs",
		"total_pageviews",
		"total_uniqueviews",
	}
	sqlRes := r.data.db.Model(&biz.User{}).Select(selectedFidlds).Where("id = ?", id).Scan(result)
	if err := sqlRes.Error; err != nil {
		r.log.Error(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			default:
				return nil, errors.New(400, "ERR_USER_INVALID_REQUEST", "")
			}
		}
	}

	// result.TotalBlogs = uint64(user.TotalBlogs)
	// result.TotalLoginDays = uint64(user.TotalLoginDays)
	// result.TotalPageviews = uint64(user.TotalPageviews)
	// result.TotalUniqueviews = uint64(user.TotalUniqueviews)
	// fmt.Println(user, sqlRes.RowsAffected)
	return result, nil
}

func (r *userRepo) UpdateUserStatisticsInfo(infos []*pb.StatisticsInfo) (int64, error) {
	user := &biz.User{}
	update_fields := make(map[string]interface{})
	var rowsAffected int64
	for _, v := range infos {
		user.ID = uint(v.ID)
		update_fields["total_pageviews"] = gorm.Expr("total_pageviews + ?", v.TotalPageviews)
		update_fields["total_uniqueviews"] = gorm.Expr("total_uniqueviews + ?", v.TotalUniqueviews)
		update_fields["total_blogs"] = gorm.Expr("total_blogs + ?", v.TotalBlogs)
		sqlRes := r.data.db.Model(user).Updates(update_fields)
		if err := sqlRes.Error; err != nil {
			return rowsAffected, err
		}
		rowsAffected += 1
	}
	return rowsAffected, nil
}
func (r *userRepo) UpdateUserPublicInfo(info *pb.UserPublicInfo) error {
	user := &biz.User{}
	user.ID = uint(info.ID)
	sqlRes := r.data.db.Model(user).Updates(biz.User{Username: info.Username, Avatar: info.Avatar, SelfDesc: info.SelfDesc})
	if err := sqlRes.Error; err != nil {
		r.log.Error(err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062: //duplicate entry
				return errors.New(400, "ERR_USER_PRE_EXISTING", "")
			default:
				return errors.New(400, "ERR_USER_INVALID_REQUEST", "")

			}
		}
	}
	return nil
}
