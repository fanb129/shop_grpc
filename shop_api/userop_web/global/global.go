package global

import (
	ut "github.com/go-playground/universal-translator"
	"shop_api/userop_web/config"
	"shop_api/userop_web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig

	NacosConfig *config.NacosConfig

	GoodsSrvClient proto.GoodsClient
	MessageClient  proto.MessageClient
	AddressClient  proto.AddressClient
	UserFavClient  proto.UserFavClient
)
