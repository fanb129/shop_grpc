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
mkdir -vp target/inventory_srv
cp config-pro.yaml target/inventory_srv/config-pro.yaml

go build -o target/inventory_srv_main main.go
echo "构建结束"