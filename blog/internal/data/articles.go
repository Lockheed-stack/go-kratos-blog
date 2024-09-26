package data

import (
	pb "blog/api/articles"
	"blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// create
func (ap *articleRepo) CreateAnArticle(article *biz.Article) error {

	result := ap.data.db.Create(article)
	if result.Error != nil {
		ap.log.Error(result.Error)
		return result.Error
	}
	ap.log.Infof("Create a new blog:%s; Row affected:%v\n", article.Title, result.RowsAffected)
	return nil
}

// select
func (ap *articleRepo) GetArticlesInSameCategory_Pagination(pageSize uint32, offset uint32, cid uint32) ([]*pb.DetailArticleInfo, error) {

	var result = []*pb.DetailArticleInfo{}
	sqlRes := ap.data.db.Table("article").Where("cid=?", cid).Limit(int(pageSize)).Offset(int(offset)).Scan(&result)

	if sqlRes.Error != nil {
		ap.log.Error(sqlRes.Error)
		return nil, sqlRes.Error
	}

	return result, nil
}
func (ap *articleRepo) GetArticlesByCidAndUid_Pagination(pageSize uint32, offset uint32, cid uint32, uid uint32) ([]*pb.DetailArticleInfo, error) {
	var result = []*pb.DetailArticleInfo{}
	sqlRes := ap.data.db.Table("article").Where("cid=? and uid=?", cid, uid).Limit(int(pageSize)).Offset(int(offset)).Scan(&result)

	if sqlRes.Error != nil {
		ap.log.Error(sqlRes.Error)
		return nil, sqlRes.Error
	}

	return result, nil
}
func (ap *articleRepo) GetOneArticle(id uint32) (*biz.Article, error) {
	var article = &biz.Article{}

	sqlRes := ap.data.db.Where("id=?", id).First(article)
	if err := sqlRes.Error; err != nil {
		ap.log.Error(err)
		return article, err
	}

	return article, nil
}

// update
func (ap *articleRepo) UpdateOneArticle(article *biz.Article) (uint32, error) {
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img
	sqlRes := ap.data.db.Model(article).Updates(maps)

	if err := sqlRes.Error; err != nil {
		return uint32(sqlRes.RowsAffected), err
	}
	if sqlRes.RowsAffected == 0 {
		ap.log.Errorf("Try to update article '%v', but zero row affected\n", article.ID)
	}
	return uint32(sqlRes.RowsAffected), nil
}

// delete
func (ap *articleRepo) RemoveOneArticle(id uint32) (uint32, error) {

	sqlRes := ap.data.db.Where("id=?", id).Delete(&biz.Article{})

	if err := sqlRes.Error; err != nil {
		ap.log.Error(err)
		return uint32(sqlRes.RowsAffected), err
	}
	if sqlRes.RowsAffected == 0 {
		ap.log.Errorf("Try to delete article '%v', but zero row affected\n", id)
	}
	return uint32(sqlRes.RowsAffected), nil
}
