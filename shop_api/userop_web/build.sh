echo "开始构建"
#export GOPATH=$WORKSPACE/..
#export PATH=$PATH:$GOROOT/bin
#
## Print Go version
#go version
#
#export GO111MODULE=on
#export GOPROXY=https://goproxy.cn
#export ENV=local

echo "current: ${USER}"
#拷贝配置文件到target下
mkdir -vp target/userop_web
cp config-pro.yaml target/userop_web/config-pro.yaml

go build -o target/userop_web_main main.go
echo "构建结束"