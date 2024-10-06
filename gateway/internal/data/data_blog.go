package data

import (
	"context"
	"gateway/internal/biz"
	"strconv"

	"gateway/api/articles"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayBlogRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayBlogRepo(data *Data, logger log.Logger) biz.GatewayBlogRepo {
	return &gatewayBlogRepo{
		data: data,
		log:  log.NewHelper(logger),
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
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.GetRecommendArticles(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, err
}
func (r *gatewayBlogRepo) GRPC_GetBlogsByRandom(req *articles.GetRandomArticlesRequest) (*articles.GetRandomArticlesReply, error) {
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.GetRandomArticles(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return result, err
}
func (r *gatewayBlogRepo) GRPC_GetSingleBlog(req *articles.GetSingleArticleRequest) (*articles.GetSingleArticleReply, error) {

	// fast path
	article_id_str := strconv.Itoa(int(req.ArticleID))
	info, err := GetOneBlogRedis(r.data.Redis_cli, article_id_str)
	if err == nil {
		// r.log.Info("using fast path")
		result := &articles.GetSingleArticleReply{
			Article: info,
			Code:    200,
			Msg:     "ok",
		}
		return result, nil
	}

	// slow path
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.GetSingleArticle(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
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
