package bin

import (
	"fmt"
	"os"
	"watermelon/config"
)

func Dict_Generate() {
	file, err := os.Create("./dict/Blaste/len5.txt")
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()

	// 定义字符集合
	characters := config.Blaste_str

	// 循环生成所有可能的字符串，并写入文件
	for length := config.Blaste_max_length; length <= config.Blaste_max_length; length++ {
		generateStrings(file, characters, "", length)
	}

	fmt.Println("生成完毕")
}

// 递归生成所有可能的字符串，并写入文件
func generateStrings(file *os.File, characters, currentString string, length int) {
	if length == 0 {
		// 写入字符串到文件中
		file.WriteString(currentString + "\n")
		return
	}

	// 递归生成所有可能的字符串
	for _, char := range characters {
		generateStrings(file, characters, currentString+string(char), length-1)
	}
}
