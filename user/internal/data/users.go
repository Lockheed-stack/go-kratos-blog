package data

import (
	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
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

func (r *userRepo) AuthLogin(name string, pwd string) (uint64, error) {

	user := &biz.User{}
	r.data.db.Where("username=?", name).First(user)

	if user.ID == 0 || user.Password != pwd {
		return 0, errors.New(400, "ERR_USER_USERNAME_PASSWORD_WRONG", "")
	}

	return uint64(user.ID), nil
}
