package routers

import (
	_ "blog-service/docs"
	"blog-service/internal/middleware"
	v1 "blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()      // 创建一个默认的路由引擎
	r.Use(gin.Logger()) // 使用日志中间件和恢复中间件
	r.Use(gin.Recovery())

	r.Use(middleware.Translations()) // 注册国际化翻译中间件
	// url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle() // 创建一个新的文章控制器
	tag := v1.NewTag()         // 创建一个新的标签控制器

	apiv1 := r.Group("/api/v1") // 创建一个新的路由组
	{
		// 注册标签路由
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		// 注册文章路由
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
