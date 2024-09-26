package biz

import (
	pb "blog/api/articles"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string `gorm:"type:varchar(100);not null" json:"title"`
	Cid      uint32 `gorm:"type:int;not null;UNSIGNED" json:"cid"`
	Uid      uint32 `gorm:"type:int;not null;UNSIGNED" json:"uid"`
	Desc     string `gorm:"type:varchar(200)" json:"desc"`
	Content  string `gorm:"type:longtext" json:"content"`
	Img      string `gorm:"type:longtext" json:"img"`
	PageView uint   `gorm:"type:uint;defualt:0" json:"pv"`
}

type ArticleRepo interface {
	CreateAnArticle(*Article) error
	GetArticlesInSameCategory_Pagination(pagesize uint32, pagenum uint32, cid uint32) ([]*pb.DetailArticleInfo, error)
	GetArticlesByCidAndUid_Pagination(pagesize uint32, pagenum uint32, cid uint32, uid uint32) ([]*pb.DetailArticleInfo, error)
	GetOneArticle(uint32) (*Article, error)

	UpdateOneArticle(*Article) (uint32, error)
	RemoveOneArticle(uint32) (uint32, error)
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

// create
func (uc *ArticleUsecase) CreateArticle(a *Article) error {
	err := uc.repo.CreateAnArticle(a)
	if err != nil {
		return pb.ErrorErrArticleInvalidIdOrTitle("please check id or title\n")
	}
	return err
}

// select
func (uc *ArticleUsecase) GetSelectedArticlesByCid(pageSize uint32, pageNum uint32, cid uint32) ([]*pb.DetailArticleInfo, error) {
	if pageSize > 50 {
		pageSize = 50
	}
	var offset uint32 = (pageNum - 1) * pageSize

	return uc.repo.GetArticlesInSameCategory_Pagination(pageSize, offset, cid)
}
func (uc *ArticleUsecase) GetSelectedArticlesByCidAndUid(pageSize uint32, pageNum uint32, cid uint32, uid uint32) ([]*pb.DetailArticleInfo, error) {
	if pageSize > 50 {
		pageSize = 50
	}

	var offset uint32 = (pageNum - 1) * pageSize
	return uc.repo.GetArticlesByCidAndUid_Pagination(pageSize, offset, cid, uid)
}
func (uc *ArticleUsecase) GetArticleByID(id uint32) (*Article, error) {
	article, err := uc.repo.GetOneArticle(id)
	if err != nil {
		return article, pb.ErrorErrArticleNotExist("Article '%v' doesn't exist\n", id)
	}
	return article, nil
}

// update
func (uc *ArticleUsecase) UpdateArticle(a *Article) error {

	rows, err := uc.repo.UpdateOneArticle(a)
	if err != nil || rows == 0 {
		return pb.ErrorErrArticleNotExist("Article '%v' doesn't exist\n", a.ID)
	}
	return nil
}

// delete
func (uc *ArticleUsecase) DeleteArticleByID(id uint32) error {
	rows, err := uc.repo.RemoveOneArticle(id)
	if err != nil || rows == 0 {
		return pb.ErrorErrArticleNotExist("Article '%v' doesn't exist\n", id)
	}
	return nil
}
