package global

import (
	ut "github.com/go-playground/universal-translator"
	"shop_api/user_web/config"
	"shop_api/user_web/proto"
) // 多语言翻译验证

var (
	Trans         ut.Translator
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	NacosConfig   *config.NacosConfig  = &config.NacosConfig{}
	UserSrvClient proto.UserClient
)
