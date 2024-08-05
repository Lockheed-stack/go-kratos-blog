package biz

import (
	pb "category/api/category"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(30);uniqueIndex;not null"`
}

type CategoryRepo interface {
	CreateCategory(string) error
	GetCategory_Pagination(uint32, uint32) ([]*pb.ListCategoryReply_CategoryInfo, error)
	UpdateCategoryByID(uint32, *Category) error
	DeleteCategoryByID(uint32) error
}

type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *CategoryUsecase) CreateOneCategory(name string) error {
	err := uc.repo.CreateCategory(name)
	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		case "ERR_CATEGORY_PRE_EXISTING":
			return pb.ErrorErrCategoryPreExisting("")
		}
	}
	return nil
}

func (uc *CategoryUsecase) GetSelectedCategory(pageSize uint32, pageNum uint32) ([]*pb.ListCategoryReply_CategoryInfo, error) {
	if pageSize > 50 {
		pageSize = 50
	}
	var offset uint32 = (pageNum - 1) * pageSize

	return uc.repo.GetCategory_Pagination(pageSize, offset)
}

func (uc *CategoryUsecase) DeleteOneCategory(id uint32) error {
	err := uc.repo.DeleteCategoryByID(id)
	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		case "ERR_CATEGORY_INVALID_ID":
			return pb.ErrorErrCategoryInvalidId("")
		case "ERR_CATEGORY_FOREIGN_KEY_CONSTRAINT":
			return pb.ErrorErrCategoryForeignKeyConstraint("")
		}
	}
	return nil
}

func (uc *CategoryUsecase) UpdateOneCategory(id uint32, category *Category) error {
	err := uc.repo.UpdateCategoryByID(id, category)

	if err != nil {
		e := errors.FromError(err)
		switch e.Reason {
		case "ERR_CATEGORY_INVALID_ID":
			return pb.ErrorErrCategoryInvalidId("")
		case "ERR_CATEGORY_INVALID_TITLE":
			return pb.ErrorErrCategoryInvalidTitle("")
		case "ERR_CATEGORY_FOREIGN_KEY_CONSTRAINT":
			return pb.ErrorErrCategoryForeignKeyConstraint("")
		}
	}
	return nil
}
