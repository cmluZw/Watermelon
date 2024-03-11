package scan

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"watermelon/common"
	"watermelon/common/bin"
)

var SimilarityContentList map[string]string //记录长度

// 程序运行时间：4m15.0272312s
func SimpleScanByDict(dict_path string, url string) {
	var wg sync.WaitGroup
	paths := common.OpenFileByPath(dict_path)
	count := len(paths) // 总数为字典的长度
	flag := 0
	if !strings.Contains(strings.Replace(url, "//", "", -1), "/") {
		url = url + "/"
	}
	//日志文件
	url_domain := common.Check_url_invalid(url)
	// 获取当前时间
	currentTime := time.Now()

	// 根据当前时间生成文件名
	fileName := fmt.Sprintf("%d-%02d-%02d_%02d-%02d-%02d-%s.txt",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second(), url_domain)
	log_name := "./log/" + fileName

	infoChan := make(chan common.Info)
	for _, path := range paths {
		wg.Add(1)
		defer wg.Done()
		flag = flag + 1
		fmt.Printf("%d/%d", flag, count)
		urlpath := url + path
		go func() {
			requestinfo := common.Request(urlpath, &wg)
			infoChan <- requestinfo
		}()

		// 主 goroutine 从 resultChan 中接收返回值
		info := <-infoChan

		switch {
		case info.StatusCode >= 200 && info.StatusCode < 300: //200-300 200
			reslut_info := common.Print_statuscode_20x(info.StatusCode, info.PageLength, path)
			common.Record_log(reslut_info, log_name)
		case info.StatusCode >= 300 && info.StatusCode < 400: //300-400 302
			reslut_info := common.Print_statuscode_30x(info.StatusCode, info.PageLength, path, info.Redirect)
			common.Record_log(reslut_info, log_name)
		case info.StatusCode >= 400 && info.StatusCode < 500 && info.StatusCode != 404: // 400-500 not 404
			reslut_info := common.Print_statuscode_40x_not_404(info.StatusCode, info.PageLength, path)
			common.Record_log(reslut_info, log_name)
		case info.StatusCode > 500 && info.StatusCode < 600 && info.StatusCode != 500: //500-600 not 500
			reslut_info := common.Print_statuscode_50x_not_500(info.StatusCode, info.PageLength, path)
			common.Record_log(reslut_info, log_name)
		case info.StatusCode == 500: //500
			reslut_info := common.Print_statuscode_500(info.StatusCode, info.PageLength, path)
			common.Record_log(reslut_info, log_name)
		case info.StatusCode == 404:
			post_info := common.POST_Request(urlpath, "application/json", "{'username'':'password'}", &wg)
			if post_info.StatusCode != 404 {
				reslut_info := common.Print_post(info.StatusCode, info.PageLength, path)
				common.Record_log(reslut_info, log_name)
			}
		}

		time.Sleep(time.Millisecond)
		fmt.Print("\033[2K\r")

	}

	endTime := time.Now()

	// 计算时间间隔
	duration := endTime.Sub(currentTime)

	fmt.Printf("程序运行时间：%v\n", duration)
	wg.Wait() // 等待所有goroutine完成

}

func ScanSimilarity(dict_path string, url string) {
	bin.Extract_404_page(url)
	var wg sync.WaitGroup
	paths := common.OpenFileByPath(dict_path)
	count := len(paths) // 总数为字典的长度
	flag := 0
	if !strings.Contains(strings.Replace(url, "//", "", -1), "/") {
		url = url + "/"
	}
	//日志文件
	url_domain := common.Check_url_invalid(url)
	// 获取当前时间
	currentTime := time.Now()

	// 根据当前时间生成文件名
	fileName := fmt.Sprintf("%d-%02d-%02d_%02d-%02d-%02d-%s.txt",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second(), url_domain)
	log_name := "./log/" + fileName

	infoChan := make(chan common.Info)
	for _, path := range paths {
		wg.Add(1)
		defer wg.Done()
		flag = flag + 1
		fmt.Printf("%d/%d", flag, count)
		urlpath := url + path
		go func() {
			requestinfo := common.Request(urlpath, &wg)
			infoChan <- requestinfo
		}()

		// 主 goroutine 从 resultChan 中接收返回值
		info := <-infoChan
		switch {
		case info.StatusCode >= 200 && info.StatusCode < 300: //200-300 200
			check := bin.Status200(info)
			if check == 1 {
				reslut_info := common.Print_statuscode_20x(info.StatusCode, info.PageLength, path)
				common.Record_log(reslut_info, log_name)
			}
		case info.StatusCode >= 300 && info.StatusCode < 400: //300-400 302
			reslut_info := common.Print_statuscode_30x(info.StatusCode, info.PageLength, path, info.Redirect)
			common.Record_log(reslut_info, log_name)
		case info.StatusCode >= 400 && info.StatusCode < 500 && info.StatusCode != 404: // 400-500 not 404
			reslut_info := common.Print_statuscode_40x_not_404(info.StatusCode, info.PageLength, path)
			common.Record_log(reslut_info, log_name)
		case info.StatusCode > 500 && info.StatusCode < 600 && info.StatusCode != 500: //500-600 not 500
			reslut_info := common.Print_statuscode_50x_not_500(info.StatusCode, info.PageLength, path)
			common.Record_log(reslut_info, log_name)
		case info.StatusCode == 500: //500
			reslut_info := common.Print_statuscode_500(info.StatusCode, info.PageLength, path)
			common.Record_log(reslut_info, log_name)
		case info.StatusCode == 404:
			post_info := common.POST_Request(urlpath, "application/json", "{'username'':'password'}", &wg)
			if post_info.StatusCode != 404 {
				reslut_info := common.Print_post(info.StatusCode, info.PageLength, path)
				common.Record_log(reslut_info, log_name)
			}
		}

		time.Sleep(time.Millisecond)
		fmt.Print("\033[2K\r")

	}

	endTime := time.Now()

	// 计算时间间隔
	duration := endTime.Sub(currentTime)

	fmt.Printf("程序运行时间：%v\n", duration)
	wg.Wait() // 等待所有goroutine完成

}

//func ScanByDict_Similar(dict_path string, url string) {
//
//	var wg sync.WaitGroup
//	SimilarityContentList = make(map[string]string) //初始化键值对
//	SimilarityContentList["content"] = ""
//	SimilarityContentList["frequency"] = "0"
//	SimilarityContentList["lock"] = "0" //未上锁
//
//	paths := common.OpenFileByPath(dict_path)
//	count := len(paths) // 总数为字典的长度
//	flag := 0
//	if !strings.Contains(strings.Replace(url, "//", "", -1), "/") {
//		url = url + "/"
//	}
//	////日志文件
//	//url_domain := common.Check_url_invalid(url)
//	//// 获取当前时间
//	//currentTime := time.Now()
//
//	//// 根据当前时间生成文件名
//	//fileName := fmt.Sprintf("%d-%02d-%02d_%02d-%02d-%02d-%s.txt",
//	//	currentTime.Year(), currentTime.Month(), currentTime.Day(),
//	//	currentTime.Hour(), currentTime.Minute(), currentTime.Second(), url_domain)
//	//log_name := "./dict/" + fileName
//	for _, path := range paths {
//		wg.Add(1)
//		defer wg.Done()
//		flag = flag + 1
//		fmt.Printf("%d/%d", flag, count)
//		urlpath := url + path
//		info := common.Request(urlpath, &wg)
//		Similarity := bin.Similarity(SimilarityContentList["content"], info.Content)
//		if Similarity > 0.99 && SimilarityContentList["lock"] == "0" {
//			frequencyInt, _ := strconv.Atoi(SimilarityContentList["frequency"])
//			SimilarityContentList["frequency"] = strconv.Itoa(frequencyInt + 1)
//			SimilarityContentList["content"] = info.Content
//		}
//		current, _ := strconv.Atoi(SimilarityContentList["frequency"])
//		if current < 20 && SimilarityContentList["lock"] == "0" {
//			SimilarityContentList["content"] = info.Content
//		} else {
//			SimilarityContentList["lock"] = "1" //上锁
//		}
//		fmt.Println(urlpath)
//		fmt.Println(info.StatusCode)
//		fmt.Println("++++++++++++++++++")
//
//		if SimilarityContentList["lock"] == "1" {
//			info = common.Clear_info(url, info.Content, len(info.Content))
//		}
//		fmt.Println(info.StatusCode)
//
//		//	switch {
//		//	case info.StatusCode >= 200 && info.StatusCode < 300: //200-300 200
//		//		reslut_info := common.Print_statuscode_20x(info.StatusCode, info.PageLength, path)
//		//		common.Record_log(reslut_info, log_name)
//		//	case info.StatusCode >= 300 && info.StatusCode < 400: //300-400 302
//		//		reslut_info := common.Print_statuscode_30x(info.StatusCode, info.PageLength, path, info.Redirect)
//		//		common.Record_log(reslut_info, log_name)
//		//	case info.StatusCode >= 400 && info.StatusCode < 500 && info.StatusCode != 404: // 400-500 not 404
//		//		reslut_info := common.Print_statuscode_40x_not_404(info.StatusCode, info.PageLength, path)
//		//		common.Record_log(reslut_info, log_name)
//		//	case info.StatusCode > 500 && info.StatusCode < 600 && info.StatusCode != 500: //500-600 not 500
//		//		reslut_info := common.Print_statuscode_50x_not_500(info.StatusCode, info.PageLength, path)
//		//		common.Record_log(reslut_info, log_name)
//		//	case info.StatusCode == 500: //500
//		//		reslut_info := common.Print_statuscode_500(info.StatusCode, info.PageLength, path)
//		//		common.Record_log(reslut_info, log_name)
//		//	case info.StatusCode == 404:
//		//		post_info := common.POST_Request(urlpath, "application/json", "{'username'':'password'}", &wg)
//		//		if post_info.StatusCode != 404 {
//		//			reslut_info := common.Print_post(info.StatusCode, info.PageLength, path)
//		//			common.Record_log(reslut_info, log_name)
//		//		}
//		//
//		//	}
//		//
//		//	time.Sleep(time.Millisecond)
//		//	fmt.Print("\033[2K\r")
//		//
//	}
//	wg.Wait() // 等待所有goroutine完成
//}
