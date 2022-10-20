package initialize

import "go.uber.org/zap"

func InitLogger() {
	dev := zap.NewDevelopmentConfig()
	dev.OutputPaths = []string{"./tmp/log/inventory_srv/dev.log", "stderr"}
	logger, _ := dev.Build()
	zap.ReplaceGlobals(logger)
}
