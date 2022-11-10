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
mkdir -vp target/userop_srv
cp config-pro.yaml target/userop_srv/config-pro.yaml

cd
go build -o target/userop_srv_main main.go
echo "构建结束"