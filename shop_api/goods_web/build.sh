echo "开始构建"
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GO111MODULE=on
export GOPROXY=https://goproxy.cn

# Print Go version
go version

echo "current: ${USER}"
#拷贝配置文件到target下
mkdir -vp goods_web/target/goods_web
cp goods_web/config-pro.yaml goods_web/target/goods_web/config-pro.yaml
cp goods_web/start.sh goods_web/target/start.sh

go build -o goods_web/target/goods_web_main goods_web/main.go
echo "构建结束"