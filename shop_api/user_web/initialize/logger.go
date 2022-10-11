package initialize

import "go.uber.org/zap"

func InitLogger() {
	dev := zap.NewDevelopmentConfig()
	//dev.OutputPaths = []string{"E:\\go\\shop_grpc\\shop_api\\user_web\\dev.log"}
	logger, _ := dev.Build()
	zap.ReplaceGlobals(logger)
}
