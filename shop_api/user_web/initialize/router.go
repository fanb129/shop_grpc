package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_api/user_web/middlewares"
	"shop_api/user_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 配置跨域
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/u/v1")
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	return Router
}
