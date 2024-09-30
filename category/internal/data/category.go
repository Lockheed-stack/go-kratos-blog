package data

import (
	pb "category/api/category"
	"category/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *categoryRepo) CreateCategory(name string) error {
	// the category name is unique
	sqlRes := r.data.db.Create(&biz.Category{Name: name})
	if sqlRes.Error != nil {
		if mysqlErr, ok := sqlRes.Error.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062: //Duplicate entry
				{
					r.log.Error(mysqlErr.Message)
					return errors.New(400, "ERR_CATEGORY_PRE_EXISTING", "")
				}
			}
		}
	}

	return nil
}
func (r *categoryRepo) GetCategory_Pagination(pageSize uint32, offset uint32) ([]*pb.CategoryInfo, error) {
	var result = []*pb.CategoryInfo{}

	sqlRes := r.data.db.Model(&biz.Category{}).Limit(int(pageSize)).Offset(int(offset)).Scan(&result)
	if err := sqlRes.Error; err != nil {
		r.log.Error(err)
		return result, err
	}
	return result, nil
}
func (r *categoryRepo) UpdateCategoryByID(id uint32, category *biz.Category) error {
	sqlRes := r.data.db.Updates(category)

	if sqlRes.Error != nil {
		if mysqlErr, ok := sqlRes.Error.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1406: // Data too long for column
				{
					r.log.Error(mysqlErr.Message)
					return errors.New(400, "ERR_CATEGORY_INVALID_TITLE", "title length should be less than 30")
				}
			case 1451: // foreign key constraint fails
				{
					r.log.Error(mysqlErr.Message)
					return errors.New(400, "ERR_CATEGORY_FOREIGN_KEY_CONSTRAINT", "")
				}
			}
		}
	}

	if sqlRes.RowsAffected == 0 {
		r.log.Infof("Try to update category '%v', but zero row affected\n", id)
		return errors.New(400, "ERR_CATEGORY_INVALID_ID", "please check the category id")
	}

	return nil
}
func (r *categoryRepo) DeleteCategoryByID(id uint32) error {
	sqlRes := r.data.db.Unscoped().Where("id=?", id).Delete(&biz.Category{})

	if err := sqlRes.Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1451: // foreign key constraint fails
				{
					r.log.Error(mysqlErr.Message)
					return errors.New(400, "ERR_CATEGORY_FOREIGN_KEY_CONSTRAINT", "")
				}
			}
		}
	}

	if sqlRes.RowsAffected == 0 {
		r.log.Infof("Try to delete category '%v', but zero row affected\n", id)
		return errors.New(400, "ERR_CATEGORY_INVALID_ID", "please check the category id")
	}
	return nil
}
