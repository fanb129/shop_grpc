package router

import (
	"github.com/gin-gonic/gin"

	"shop_api/order_web/api/order"
	"shop_api/order_web/api/pay"
	"shop_api/order_web/middlewares"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders").Use(middlewares.JWTAuth()).Use(middlewares.Trace())
	{
		OrderRouter.GET("", order.List)       // 订单列表
		OrderRouter.POST("", order.New)       // 新建订单
		OrderRouter.GET("/:id", order.Detail) // 订单详情
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}
}
