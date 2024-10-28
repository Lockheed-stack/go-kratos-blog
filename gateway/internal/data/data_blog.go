package data

import (
	"context"
	"gateway/internal/biz"
	"strconv"
	"sync"
	"time"

	"gateway/api/articles"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayBlogRepo struct {
	data          *Data
	log           *log.Helper
	lock          sync.Mutex
	ch            chan struct{}
	statistics_pv map[uint32]uint32
	hits_trigger  *trigger
}
type trigger struct {
	hits         uint64
	last_hits    uint64
	trigger_time int64
	ticker       *time.Ticker
}

func NewGatewayBlogRepo(data *Data, logger log.Logger) biz.GatewayBlogRepo {
	tri := &trigger{
		trigger_time: time.Now().Unix(),
		ticker:       time.NewTicker(time.Hour),
	}

	r := &gatewayBlogRepo{
		data:          data,
		log:           log.NewHelper(logger),
		lock:          sync.Mutex{},
		ch:            make(chan struct{}, 100),
		statistics_pv: make(map[uint32]uint32),
		hits_trigger:  tri,
	}
	go func() {
		for {
			<-tri.ticker.C
			// If it has been less than 5 minutes since the last trigger, the save operation will not be executed.
			if time.Now().Unix()-tri.trigger_time > 300 {
				increment := r.hits_trigger.hits - r.hits_trigger.last_hits
				if increment >= 1 { // At least 1 request in an hour
					r.savePageviewToDB()
				}
			}
		}
	}()
	go func() {
		for {
			select {
			case <-r.ch:
				{
					r.hits_trigger.hits += 1
					now := time.Now().Unix()
					increment := r.hits_trigger.hits - r.hits_trigger.last_hits
					time_span := now - r.hits_trigger.trigger_time
					r.log.Debugf("time span:%v, incr:%v, pv:%v\n", time_span, increment, r.statistics_pv)

					// At least 10000 requests in 60s or at least 100 requests in 300s
					if (time_span <= 60 && increment >= 10000) || (time_span <= 300 && increment >= 100) {
						r.hits_trigger.last_hits = r.hits_trigger.hits
						r.hits_trigger.trigger_time = now
						r.savePageviewToDB()
					}
				}
			case <-r.data.Cancel_CTX.Done():
				{
					close(r.ch)
					r.log.Info("channel closed!")
					return
				}
			}

		}
	}()

	return r
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
	// key := "random" + strconv.Itoa(int(req.Count))
	// selected_articles, err := GetBlogsListRedis(r.data.Redis_cli, key)
	// if err == nil {
	// 	result := &articles.GetRandomArticlesReply{
	// 		SelectedArticles: selected_articles,
	// 		Code:             200,
	// 	}
	// 	return result, nil
	// }

	// slow path
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.GetRandomArticles(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	// set redis key
	// err = SetBlogsListRedis(r.data.Redis_cli, key, result.SelectedArticles)
	// if err != nil {
	// 	r.log.Error(err)
	// }

	return result, nil
}
func (r *gatewayBlogRepo) GRPC_GetSingleBlog(req *articles.GetSingleArticleRequest) (*articles.GetSingleArticleReply, error) {

	/* ------------------- fast path -------------------- */
	article_id_str := strconv.Itoa(int(req.ArticleID))
	info, err := GetOneBlogRedis(r.data.Redis_cli, article_id_str)
	if err == nil { // redis cache matched

		// NOTE: map is non-concurrent safety, it cannot concurrent write
		r.lock.Lock()
		info.PageView = r.statistics_pv[uint32(req.ArticleID)]
		r.statistics_pv[uint32(req.ArticleID)] += 1
		r.lock.Unlock()

		r.ch <- struct{}{}
		result := &articles.GetSingleArticleReply{
			Article: info,
			Code:    200,
			Msg:     "ok",
		}
		return result, nil
	}

	/* -------------- slow path --------------- */
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	result, err := client.GetSingleArticle(context.Background(), req)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	r.lock.Lock()
	// concurrent: if there are many requests enter slow path, we need to precisely count the number of visits.
	if val, ok := r.statistics_pv[uint32(req.ArticleID)]; ok {
		result.Article.PageView = val
		r.statistics_pv[uint32(req.ArticleID)] += 1
		r.hits_trigger.hits += 1
		r.log.Info("pv: ", r.statistics_pv[uint32(req.ArticleID)])
	} else {
		r.statistics_pv[uint32(req.ArticleID)] = result.Article.PageView
	}
	r.lock.Unlock()
	r.ch <- struct{}{}

	// set redis key
	go func() {
		err := SetOneBlogRedis(r.data.Redis_cli, article_id_str, result.Article)
		if err != nil {
			r.log.Error(err)
		}
	}()

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
	key := strconv.Itoa(int(req.ArticleInfo.ID))
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
func (r *gatewayBlogRepo) savePageviewToDB() {

	// pv_map := make(map[uint32]uint32)
	// for k, v := range r.statistics_pv {
	// 	pv_map[k] = v
	// }
	client := articles.NewArticlesClient(r.data.ConnGRPC_blog)
	req := &articles.UpdateArticlesPageviewRequest{
		// Pageview: pv_map,
		Pageview: r.statistics_pv,
	}
	go func() {
		result, err := client.UpdateArticlesPageview(context.Background(), req)
		if err != nil {
			r.log.Error(err)
		} else {
			r.log.Infof("Save to DB:%v", result.Msg)
		}
	}()
}
