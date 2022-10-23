package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/userop_web/api/message"
	"shop_api/userop_web/middlewares"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middlewares.JWTAuth())
	{
		MessageRouter.GET("", message.List) // 消息留言列表
		MessageRouter.POST("", message.New) //新建消息
	}
}
