package service

import (
	"context"

	pb "blog/api/articles"
	"blog/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

type ArticlesService struct {
	pb.UnimplementedArticlesServer
	ac *biz.ArticleUsecase
}

func NewArticlesService(ac *biz.ArticleUsecase) *ArticlesService {
	return &ArticlesService{
		ac: ac,
	}
}

func (s *ArticlesService) CreateArticles(ctx context.Context, req *pb.CreateArticlesRequest) (*pb.CreateArticlesReply, error) {

	article := &biz.Article{
		Title:    req.Title,
		Cid:      uint(req.Cid),
		Uid:      uint(req.Uid),
		Desc:     req.Desc,
		Content:  req.Content,
		Img:      req.Img,
		PageView: 0,
	}
	resp := &pb.CreateArticlesReply{}
	err := s.ac.CreateArticle(article)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
		return resp, err
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}

func (s *ArticlesService) GetArticlesInSameCategory(ctx context.Context, req *pb.GetArticlesInSameCategoryRequest) (*pb.GetArticlesInSameCategoryReply, error) {

	result, count, err := s.ac.GetSelectedArticlesByCid(req.PageSize, req.PageNum, req.CID)
	resp := &pb.GetArticlesInSameCategoryReply{}
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		return resp, err
	} else {
		resp.SelectedArticles = result
		resp.Total = count
		resp.Code = 200
	}
	return resp, nil
}

func (s *ArticlesService) GetArticlesByCidAndUid(ctx context.Context, req *pb.GetArticlesByCidAndUidRequest) (*pb.GetArticlesByCidAndUidReply, error) {
	resp := &pb.GetArticlesByCidAndUidReply{}
	result, count, err := s.ac.GetSelectedArticlesByCidAndUid(req.PageSize, req.PageNum, req.CID, req.UID)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		return resp, err
	} else {
		resp.SelectedArticles = result
		resp.Total = count
		resp.Code = 200
	}
	return resp, nil
}
func (s *ArticlesService) GetRecommendArticles(ctx context.Context, req *pb.GetRecommendArticlesRequest) (*pb.GetRecommendArticlesReply, error) {
	resp := &pb.GetRecommendArticlesReply{}
	result, err := s.ac.GetArticlesForRecommend(req.PageSize, req.PageNum)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		return resp, err
	} else {
		resp.SelectedArticles = result
		resp.Code = 200
	}
	return resp, nil
}
func (s *ArticlesService) GetRandomArticles(ctx context.Context, req *pb.GetRandomArticlesRequest) (*pb.GetRandomArticlesReply, error) {
	resp := &pb.GetRandomArticlesReply{}
	result, err := s.ac.GetArticlesByRandom(req.Count)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		return resp, err
	} else {
		resp.SelectedArticles = result
		resp.Code = 200
	}
	return resp, nil
}
func (s *ArticlesService) GetSingleArticle(ctx context.Context, req *pb.GetSingleArticleRequest) (*pb.GetSingleArticleReply, error) {

	result, err := s.ac.GetArticleByID(req.ArticleID)
	resp := &pb.GetSingleArticleReply{}
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Msg = kratos_err.Reason
		resp.Code = uint32(kratos_err.Code)
		return resp, err
	} else {
		resp.Article = &pb.DetailArticleInfo{
			CreatedAt: result.CreatedAt.String(),
			UpdatedAt: result.CreatedAt.String(),
			Title:     result.Title,
			Desc:      result.Desc,
			Content:   result.Content,
			Img:       result.Img,
			PageView:  uint32(result.PageView),
			ID:        uint64(result.ID),
			Uid:       uint64(result.Uid),
			Cid:       uint64(result.Cid),
		}
		resp.Msg = "OK"
		resp.Code = 200
	}

	return resp, nil
}

func (s *ArticlesService) UpdateArticles(ctx context.Context, req *pb.UpdateArticlesRequest) (*pb.UpdateArticlesReply, error) {
	article := &biz.Article{
		Model:   gorm.Model{ID: uint(req.ArticleInfo.ID)},
		Cid:     uint(req.ArticleInfo.Cid),
		Desc:    req.ArticleInfo.Desc,
		Content: req.ArticleInfo.Content,
		Img:     req.ArticleInfo.Img,
		Title:   req.ArticleInfo.Title,
		Uid:     uint(req.ArticleInfo.Uid),
	}

	resp := &pb.UpdateArticlesReply{}
	err := s.ac.UpdateArticle(article)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
		return resp, err
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}

func (s *ArticlesService) UpdateArticlesPageview(ctx context.Context, req *pb.UpdateArticlesPageviewRequest) (*pb.UpdateArticlesPageviewReply, error) {
	resp := &pb.UpdateArticlesPageviewReply{}
	err := s.ac.UpdateArticlesPageview(req.Pageview)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Message
		return resp, err
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}

func (s *ArticlesService) DeleteArticles(ctx context.Context, req *pb.DeleteArticlesRequest) (*pb.DeleteArticlesReply, error) {

	resp := &pb.DeleteArticlesReply{}
	err := s.ac.DeleteArticleByID(req.ArticleID)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
		return resp, err
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}

func (s *ArticlesService) CheckExistenceOfBlog(ctx context.Context, req *pb.CheckExistenceOfBlogRequest) (*pb.CheckExistenceOfBlogReply, error) {
	resp := &pb.CheckExistenceOfBlogReply{}
	val, err := s.ac.CheckWhetherBlogIDExistence(uint64(req.ArticleID))
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	resp.Existence = val
	return resp, nil
}
