package router

import (
	"gateway/internal/biz"
	"gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
)

func NewGinRouter(
	blog_handler *biz.GatewayBlogUsecase,
	category_handler *biz.GatewayCategoryUsecase,
	user_handler *biz.GatewayUserUsecase,
	upload_handler *biz.GatewayUploadUsecase,
	mids *middlewares.Mids,
) *gin.Engine {

	r := gin.Default()
	r.Use(mids.CorsHandler())

	public_group := r.Group("/gateway")
	public_group.Use(kgin.Middlewares(
		recovery.Recovery(),
	))
	{
		// blogs
		public_group.GET("/blog/:blogID", blog_handler.GetOneBlog)
		public_group.GET("/list/:cid", blog_handler.GetBlogsInSameCategory)
		public_group.GET("/:uid/list/:cid", blog_handler.GetBlogsByCidAndUid)
		public_group.GET("/recommendBlogs", blog_handler.GetRecommendBlogs)
		public_group.GET("/randomBlogs", blog_handler.GetRandomBlogs)
		// category
		public_group.GET("/category", category_handler.GetCategoryList)
		// user
		public_group.POST("/login", user_handler.AuthUser)
		public_group.POST("/user/add", user_handler.CreateOneUser)
		public_group.POST("/user/publicInfo", user_handler.GetPublicUsersInfo)
	}

	auth_Required_group := r.Group("/management")
	auth_Required_group.Use(
		kgin.Middlewares(
			recovery.Recovery(),
		),
		middlewares.JwtMids(),
		middlewares.NormalUserAuth(),
	)
	{
		// blogs
		auth_Required_group.POST("/:username/add-blog", blog_handler.CreateOneBlog)
		auth_Required_group.PATCH("/:username/modify-blog", blog_handler.UpdateOneBlog)
		auth_Required_group.DELETE("/:username/rm-blog", blog_handler.DeleteOneBlog)
		// auth_Required_group.POST("/:username/upload", upload_handler.Upload)
		// category
		auth_Required_group.POST("/category/add", category_handler.CreateOneCategory)
		auth_Required_group.PATCH("/category/update", category_handler.UpdateCategory)
		auth_Required_group.DELETE("/category/rm", category_handler.DeleteCategory)
		// user
		auth_Required_group.DELETE("/user/rm", user_handler.DeleteOneUser)
		auth_Required_group.POST("/user/token-check", user_handler.TokenCheck)
	}

	return r
}
