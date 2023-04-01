// #!/bin/bash

// # 指定Go编译器路径和应用程序名称
// export GOPATH=$HOME/go
// export GOROOT=/usr/local/go
// export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
// export APP_NAME=iosapp

// # 设置编译环境变量
// export CGO_ENABLED=1
// export CC=$(xcrun --sdk iphoneos --find clang)
// export CXX=$(xcrun --sdk iphoneos --find clang++)
// export GOOS=darwin
// export GOARCH=arm64

// # 创建一个新的Go模块并获取所需的依赖项
// go mod init $APP_NAME
// go get github.com/DHowett/go-plist

// # 编译Go程序
// go build -o $APP_NAME main.go

// # 将编译的可执行文件传输到iOS设备上
// scp $APP_NAME root@${IP_ADDRESS}:/var/root/

// # 在iOS设备上运行编译的程序，并将输出存储到本地文件
// ssh root@${IP_ADDRESS} "chmod +x /var/root/$APP_NAME && /var/root/$APP_NAME > /var/root/app-list.txt"

// # 从iOS设备中下载输出文件
// scp root@${IP_ADDRESS}:/var/root/app-list.txt app-list.txt

// # 删除编译的可执行文件和Go模块
// rm $APP_NAME
// rm go.mod
// rm go.sum

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func runCommand(command string, args ...string) (string, error) {
    cmd := exec.Command(command, args...)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        return "", err
    }
    return stdout.String(), nil
}


func main() {
	// 设置环境变量
    ccPath, err := runCommand("xcrun", "--sdk", "iphoneos", "--find", "clang")
    if err != nil {
        fmt.Println("Error finding clang:", err)
        return
    }
    cxxPath, err := runCommand("xcrun", "--sdk", "iphoneos", "--find", "clang++")
    if err != nil {
        fmt.Println("Error finding clang++:", err)
        return
    }
    os.Setenv("CC", ccPath)
    os.Setenv("CXX", cxxPath)
    fmt.Println("CC:", os.Getenv("CC"))
    fmt.Println("CXX:", os.Getenv("CXX"))
	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", "arm64")

	// 从os.Args中获取目标文件名和命令行参数
	args := os.Args
	// 创建 `exec.Cmd` 结构体
	cmd := exec.Command("go", args[1:]...)
	// 执行命令，并获取输出
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command: %v\nOutput: %s\n", err, out)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
