echo "开始构建"
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GO111MODULE=on
export GOPROXY=https://goproxy.cn

# Print Go version
go version

echo "current: ${USER}"

go build -o user_srv_main user_srv/main.go
echo "构建结束"