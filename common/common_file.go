package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 读取文件，从传入的path里
func OpenFileByPath(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		//fmt.Print(err)
		fmt.Print(path+"读取错误", err)
		return nil
	}
	defer file.Close()
	var content []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != " " {
			content = append(content, scanner.Text())
		}
	}
	return content
}

func Clear_file(path string) int {
	// 打开文件
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return -1
	}
	defer file.Close()

	// 清空文件
	err = file.Truncate(0)
	if err != nil {
		fmt.Println("无法清空文件:", err)
		return -1
	}

	return 1
}

func Check_file_exist(filename string) bool {
	// 获取文件信息
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("文件 %s 不存在,将使用默认字典\n", filename)
		return false
	} else if info.Size() == 0 {
		fmt.Printf("文件 %s 存在但大小为0，将使用默认字典\n", filename)
		return false
	} else {
		fmt.Printf("文件 %s 存在且大小不为0\n", filename)
		return true
	}
	return false
}
