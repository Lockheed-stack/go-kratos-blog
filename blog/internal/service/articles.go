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
		Cid:      req.Cid,
		Uid:      req.Uid,
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
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}

func (s *ArticlesService) GetArticles(ctx context.Context, req *pb.GetArticlesRequest) (*pb.GetArticlesReply, error) {

	result, err := s.ac.GetSelectedArticles(req.PageSize, req.PageNum)
	resp := &pb.GetArticlesReply{}
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
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
	} else {
		resp.Article = &pb.GetSingleArticleReply_RespondMsg{
			CreatedAt: result.CreatedAt.String(),
			UpdatedAt: result.CreatedAt.String(),
			Title:     result.Title,
			Desc:      result.Desc,
			Content:   result.Content,
			PageView:  uint32(result.PageView),
		}
		resp.Msg = "OK"
		resp.Code = 200
	}

	return resp, nil
}

func (s *ArticlesService) UpdateArticles(ctx context.Context, req *pb.UpdateArticlesRequest) (*pb.UpdateArticlesReply, error) {
	article := &biz.Article{
		Model:   gorm.Model{ID: uint(req.ArticleID)},
		Cid:     req.Cid,
		Desc:    req.Desc,
		Content: req.Content,
		Img:     req.Thumbnail,
		Title:   req.Title,
	}

	resp := &pb.UpdateArticlesReply{}
	err := s.ac.UpdateArticle(article)
	if err != nil {
		kratos_err := err.(*errors.Error)
		resp.Code = uint32(kratos_err.Code)
		resp.Msg = kratos_err.Reason
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
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}
