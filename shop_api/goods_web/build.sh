echo "开始构建"
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GO111MODULE=on
export GOPROXY=https://goproxy.cn

# Print Go version
go version

echo "current: ${USER}"
#拷贝配置文件到target下
mkdir -vp target/goods_web
cp config-pro.yaml target/goods_web/config-pro.yaml
cp start.sh target/

go mod tidy
go build -o target/goods_web_main main.go
echo "构建结束"