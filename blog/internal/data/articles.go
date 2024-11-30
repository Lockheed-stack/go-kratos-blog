package data

import (
	pb "blog/api/articles"
	"blog/internal/biz"
	"math/rand"
	"time"

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
func (ap *articleRepo) GetArticlesInSameCategory_Pagination(pageSize uint32, offset uint32, cid uint64) ([]*pb.DetailArticleInfo, uint32, error) {

	var result = []*pb.DetailArticleInfo{}
	var count int64 = 0

	sqlRes := ap.data.db.Model(&biz.Article{}).Where("cid=?", cid).Limit(int(pageSize)).Offset(int(offset)).Scan(&result)
	if sqlRes.Error != nil {
		ap.log.Error(sqlRes.Error)
		return nil, uint32(count), sqlRes.Error
	}
	ap.data.db.Model(&biz.Article{}).Where("cid=?", cid).Count(&count)
	return result, uint32(count), nil
}
func (ap *articleRepo) GetArticlesByCidAndUid_Pagination(pageSize uint32, offset uint32, cid uint64, uid uint64) ([]*pb.DetailArticleInfo, uint32, error) {

	var result = []*pb.DetailArticleInfo{}
	var count int64 = 0

	sqlRes := ap.data.db.Model(&biz.Article{}).Where("cid=? and uid=?", cid, uid).Limit(int(pageSize)).Offset(int(offset)).Scan(&result)
	if sqlRes.Error != nil {
		ap.log.Error(sqlRes.Error)
		return nil, uint32(count), sqlRes.Error
	}
	ap.data.db.Model(&biz.Article{}).Where("cid=? and uid=?", cid, uid).Count(&count)
	return result, uint32(count), nil
}
func (ap *articleRepo) GetArticlesForRecommend_Pagination(pageSize uint32, offset uint32) ([]*pb.DetailArticleInfo, error) {
	var result = []*pb.DetailArticleInfo{}
	sqlRes := ap.data.db.Model(&biz.Article{}).Order("page_view desc").Limit(int(pageSize)).Offset(int(offset)).Scan(&result)
	if sqlRes.Error != nil {
		ap.log.Error(sqlRes.Error)
		return nil, sqlRes.Error
	}
	return result, nil
}
func (ap *articleRepo) GetArticlesByRandomSelect(count uint32) ([]*pb.DetailArticleInfo, error) {
	var result = []*pb.DetailArticleInfo{}
	var total int64
	ap.data.db.Model(&biz.Article{}).Count(&total)
	if total == 0 {
		return result, nil
	}

	max_start_index := total - int64(count) + 1
	var random_start_index int64
	if max_start_index > 0 {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		random_start_index = r.Int63n(max_start_index)
	}

	sqlRes := ap.data.db.Model(&biz.Article{}).Where("id > ?", random_start_index).Limit(int(count)).Scan(&result)

	if sqlRes.Error != nil {
		ap.log.Error(sqlRes.Error)
		return nil, sqlRes.Error
	}
	return result, nil
}
func (ap *articleRepo) GetOneArticle(id uint64) (*biz.Article, error) {
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
	// check the blog which wants to be updated belongs to its owner
	var original_article = &biz.Article{}
	sqlRes := ap.data.db.Select("uid").Where("id=?", article.ID).Find(original_article)
	if err := sqlRes.Error; err != nil {
		return 0, err
	}
	if original_article.Uid != article.Uid {
		return 0, nil
	}

	// update blog
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img
	sqlRes = ap.data.db.Model(article).Updates(maps)

	if err := sqlRes.Error; err != nil {
		return uint32(sqlRes.RowsAffected), err
	}
	if sqlRes.RowsAffected == 0 {
		ap.log.Errorf("Try to update article '%v', but zero row affected\n", article.ID)
	}
	return uint32(sqlRes.RowsAffected), nil
}
func (ap *articleRepo) UpdateArticlesPageview(id_pageview map[uint32]uint32) (uint32, error) {
	article := &biz.Article{}
	var rowsAffected int64
	for k, v := range id_pageview {
		article.ID = uint(k)
		sqlRes := ap.data.db.Model(article).Update("page_view", v)
		if err := sqlRes.Error; err != nil {
			return uint32(rowsAffected), err
		}
		rowsAffected += sqlRes.RowsAffected
	}

	return uint32(rowsAffected), nil
}

// delete
func (ap *articleRepo) RemoveOneArticle(id uint64) (uint32, error) {

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
