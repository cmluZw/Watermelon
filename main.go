package main

import (
	"os"
	"watermelon/application"
)

func main() {
	// 获取命令行参数
	args_all := os.Args
	args_tmp := args_all[1:] //去头
	info := application.Parse(args_tmp)
	application.Start(info)
}
