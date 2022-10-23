package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/userop_web/api/address"
	"shop_api/userop_web/middlewares"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middlewares.JWTAuth(), address.List)          // 地址列表
		AddressRouter.DELETE("/:id", middlewares.JWTAuth(), address.Delete) // 删除地址
		AddressRouter.POST("", middlewares.JWTAuth(), address.New)          //添加地址
		AddressRouter.PUT("/:id", middlewares.JWTAuth(), address.Update)    //修改地址信息
	}
}
