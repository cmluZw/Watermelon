package common

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"time"
)

func Print_statuscode_20x(status_code int, length int, web_path string) string {

	currentTime := time.Now()
	// 使用Format方法将时间格式化为 "15:04:05" 的格式
	formattedTime := currentTime.Format("15:04:05")
	info_tmp := "\r[%s]  %d - %d -   %s \n"
	info := fmt.Sprintf(info_tmp, formattedTime, status_code, length, web_path)
	color.Green(info)
	return info
}

func Print_statuscode_30x(status_code int, length int, web_path string, redirect_url string) string {
	currentTime := time.Now()
	// 使用Format方法将时间格式化为 "15:04:05" 的格式
	formattedTime := currentTime.Format("15:04:05")
	info_tmp := "\r[%s]  %d - %d -   %s   --> %s \n"
	info := fmt.Sprintf(info_tmp, formattedTime, status_code, length, web_path, redirect_url)
	color.Yellow(info)
	return info
}

func Print_statuscode_40x_not_404(status_code int, length int, web_path string) string {
	currentTime := time.Now()
	// 使用Format方法将时间格式化为 "15:04:05" 的格式
	formattedTime := currentTime.Format("15:04:05")
	info_tmp := "\r[%s]  %d - %d -   %s \n"
	info := fmt.Sprintf(info_tmp, formattedTime, status_code, length, web_path)
	color.Blue(info)
	return info
}

func Print_statuscode_50x_not_500(status_code int, length int, web_path string) string {
	currentTime := time.Now()
	// 使用Format方法将时间格式化为 "15:04:05" 的格式
	formattedTime := currentTime.Format("15:04:05")
	info_tmp := "\r[%s]  %d - %d -   %s \n"
	info := fmt.Sprintf(info_tmp, formattedTime, status_code, length, web_path)
	color.Magenta(info)
	return info
}

func Print_statuscode_500(status_code int, length int, web_path string) string {
	currentTime := time.Now()
	// 使用Format方法将时间格式化为 "15:04:05" 的格式
	formattedTime := currentTime.Format("15:04:05")
	info_tmp := "\r[%s]  %d - %d -   %s \n"
	info := fmt.Sprintf(info_tmp, formattedTime, status_code, length, web_path)
	color.Red(info)
	return info
}

func Print_post(status_code int, length int, web_path string) string {
	currentTime := time.Now()
	// 使用Format方法将时间格式化为 "15:04:05" 的格式
	formattedTime := currentTime.Format("15:04:05")
	info_tmp := "\r[%s]  %d - %d -  POST - %s \n"
	info := fmt.Sprintf(info_tmp, formattedTime, status_code, length, web_path)
	color.Red(info)
	return info
}

func Record_log(info string, log_name string) {

	file, err := os.OpenFile(log_name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建一个带缓冲的写入器
	writer := bufio.NewWriter(file)

	// 将内容写入文件
	_, err = writer.WriteString(strings.ReplaceAll(info, "\r", ""))
	if err != nil {
		return
	}

	// 刷新缓冲区，确保所有数据都被写入文件
	err = writer.Flush()
	if err != nil {
		//fmt.Println("无法刷新缓冲区:", err)
		return
	}
}
