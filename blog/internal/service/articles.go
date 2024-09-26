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
		return resp, err
	} else {
		resp.Code = 200
		resp.Msg = "OK"
	}
	return resp, nil
}

func (s *ArticlesService) GetArticlesInSameCategory(ctx context.Context, req *pb.GetArticlesInSameCategoryRequest) (*pb.GetArticlesInSameCategoryReply, error) {

	result, err := s.ac.GetSelectedArticlesByCid(req.PageSize, req.PageNum, req.CID)
	resp := &pb.GetArticlesInSameCategoryReply{}
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

func (s *ArticlesService) GetArticlesByCidAndUid(ctx context.Context, req *pb.GetArticlesByCidAndUidRequest) (*pb.GetArticlesByCidAndUidReply, error) {
	resp := &pb.GetArticlesByCidAndUidReply{}
	result, err := s.ac.GetSelectedArticlesByCidAndUid(req.PageSize, req.PageNum, req.CID, req.UID)
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
			ID:        uint32(result.ID),
			UID:       uint32(result.Uid),
			CID:       uint32(result.Cid),
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
