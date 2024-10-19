package data

import (
	"context"
	"gateway/internal/biz"
	"strconv"
	"sync"

	"gateway/api/articles"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayBlogRepo struct {
	data          *Data
	log           *log.Helper
	lock          sync.Mutex
	statistics_pv map[uint32]uint32
}

func NewGatewayBlogRepo(data *Data, logger log.Logger) biz.GatewayBlogRepo {
	return &gatewayBlogRepo{
		data:          data,
		log:           log.NewHelper(logger),
		lock:          sync.Mutex{},
		statistics_pv: make(map[uint32]uint32),
	}
}

func (r *gatewayBlogRepo) GRPC_CreateOneBlog(req *articles.CreateArticlesRequest) (*articles.CreateArticlesReply, error) {

	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.CreateArticles(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}

func (r *gatewayBlogRepo) GRPC_GetBlogsInSameCategory(req *articles.GetArticlesInSameCategoryRequest) (*articles.GetArticlesInSameCategoryReply, error) {

	// fast path

	// slow path
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.GetArticlesInSameCategory(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayBlogRepo) GRPC_GetBlogsByCidAndUid(req *articles.GetArticlesByCidAndUidRequest) (*articles.GetArticlesByCidAndUidReply, error) {
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.GetArticlesByCidAndUid(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, nil
}
func (r *gatewayBlogRepo) GRPC_GetBlogsForRecommend(req *articles.GetRecommendArticlesRequest) (*articles.GetRecommendArticlesReply, error) {
	// fast path
	key := "recommend"
	selected_articles, err := GetBlogsListRedis(r.data.Redis_cli, key)
	if err == nil {
		result := &articles.GetRecommendArticlesReply{
			SelectedArticles: selected_articles,
			Code:             200,
		}
		return result, nil
	}

	// slow path
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.GetRecommendArticles(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	// set redis key
	err = SetBlogsListRedis(r.data.Redis_cli, key, result.SelectedArticles)
	if err != nil {
		r.log.Error(err)
	}

	return result, err
}
func (r *gatewayBlogRepo) GRPC_GetBlogsByRandom(req *articles.GetRandomArticlesRequest) (*articles.GetRandomArticlesReply, error) {

	// fast path
	key := "random" + strconv.Itoa(int(req.Count))
	selected_articles, err := GetBlogsListRedis(r.data.Redis_cli, key)
	if err == nil {
		result := &articles.GetRandomArticlesReply{
			SelectedArticles: selected_articles,
			Code:             200,
		}
		return result, nil
	}

	// slow path
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.GetRandomArticles(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	// set redis key
	err = SetBlogsListRedis(r.data.Redis_cli, key, result.SelectedArticles)
	if err != nil {
		r.log.Error(err)
	}

	return result, nil
}
func (r *gatewayBlogRepo) GRPC_GetSingleBlog(req *articles.GetSingleArticleRequest) (*articles.GetSingleArticleReply, error) {

	// fast path
	article_id_str := strconv.Itoa(int(req.ArticleID))
	info, err := GetOneBlogRedis(r.data.Redis_cli, article_id_str)
	if err == nil { // cache matched
		r.lock.Lock()
		info.PageView = r.statistics_pv[uint32(req.ArticleID)]
		r.statistics_pv[uint32(req.ArticleID)] += 1
		r.lock.Unlock()

		result := &articles.GetSingleArticleReply{
			Article: info,
			Code:    200,
			Msg:     "ok",
		}
		// update pageview
		go func() {
			r.lock.Lock()
			defer r.lock.Unlock()
			info.PageView += 1
			err := SetOneBlogRedis(r.data.Redis_cli, article_id_str, info)
			if err != nil {
				r.log.Error(err)
			}
		}()

		return result, nil
	}

	// slow path
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)

	r.lock.Lock() // lock below code until it return
	defer r.lock.Unlock()

	result, err := client.GetSingleArticle(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	// concurrent: if there are many requests enter slow path, we need to precisely count the number of visits.
	if val, ok := r.statistics_pv[uint32(req.ArticleID)]; ok {
		result.Article.PageView = val
		r.statistics_pv[uint32(req.ArticleID)] += 1
		r.log.Info("pv: ", r.statistics_pv[uint32(req.ArticleID)])
	} else {
		r.statistics_pv[uint32(req.ArticleID)] = result.Article.PageView
	}

	// set redis key
	err = SetOneBlogRedis(r.data.Redis_cli, article_id_str, result.Article)
	if err != nil {
		r.log.Error(err)
	}
	return result, nil
}

func (r *gatewayBlogRepo) GRPC_UpdateBlog(req *articles.UpdateArticlesRequest) (*articles.UpdateArticlesReply, error) {
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.UpdateArticles(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	// del redis key
	key := strconv.Itoa(int(req.ArticleID))
	err = DelOneBlogKeyRedis(r.data.Redis_cli, key)
	if err != nil {
		r.log.Error(err)
	}

	return result, nil
}

func (r *gatewayBlogRepo) GRPC_DeleteBlog(req *articles.DeleteArticlesRequest) (*articles.DeleteArticlesReply, error) {
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.DeleteArticles(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	// del redis key
	key := strconv.Itoa(int(req.ArticleID))
	err = DelOneBlogKeyRedis(r.data.Redis_cli, key)
	if err != nil {
		r.log.Error(err)
	}

	return result, nil
}

// internal functions
