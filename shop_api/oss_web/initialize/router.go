package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_api/oss_web/middlewares"
	"shop_api/oss_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	Router.LoadHTMLFiles(fmt.Sprintf("oss_web/templates/index.html"))
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	Router.StaticFS("/static", http.Dir(fmt.Sprintf("oss_web/static")))
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	Router.GET("", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "posts/index",
		})
	})

	//配置跨域
	Router.Use(middlewares.Cors())

	//ApiGroup := Router.Group("/oss/v1")
	ApiGroup := Router.Group("/v1")
	router.InitOssRouter(ApiGroup)

	return Router
}
