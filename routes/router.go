package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	routerV1 := r.Group("api/v1")

	{
		//用户模块路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
	}

	{
		//用户模块路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.POST("login", v1.Login)

		//分类模块的路由接口
		routerV1.GET("category", v1.GetCate)
		routerV1.GET("category/articleList", v1.GetCategoryArticle)

		//文章模块的路由接口
		routerV1.GET("article", v1.GetArticle)
		routerV1.GET("article/info/:id", v1.GetArticleInfo)
	}
	r.Run(utils.HttpPort)
}
