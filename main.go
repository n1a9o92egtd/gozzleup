package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 设置环境变量
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", "$(xcrun --sdk iphoneos --find clang)")
	os.Setenv("CXX", "$(xcrun --sdk iphoneos --find clang++)")
	os.Setenv("GOOS", "ios")
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
