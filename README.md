# goIosForMacCompileSh

```
#!/bin/bash

# 指定Go编译器路径和应用程序名称
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
export APP_NAME=iosapp

# 设置编译环境变量
export CGO_ENABLED=1
export CC=$(xcrun --sdk iphoneos --find clang)
export CXX=$(xcrun --sdk iphoneos --find clang++)
export GOOS=darwin
export GOARCH=arm64

# 创建一个新的Go模块并获取所需的依赖项
go mod init $APP_NAME
go get github.com/DHowett/go-plist

# 编译Go程序
go build -o $APP_NAME main.go

# 将编译的可执行文件传输到iOS设备上
scp $APP_NAME root@${IP_ADDRESS}:/var/root/

# 在iOS设备上运行编译的程序，并将输出存储到本地文件
ssh root@${IP_ADDRESS} "chmod +x /var/root/$APP_NAME && /var/root/$APP_NAME > /var/root/app-list.txt"

# 从iOS设备中下载输出文件
scp root@${IP_ADDRESS}:/var/root/app-list.txt app-list.txt

# 删除编译的可执行文件和Go模块
rm $APP_NAME
rm go.mod
rm go.sum
