package bin

//通过请求肯定不存在的路径来获取页面内容，后续的使用相似度来判断404

import (
	"io/ioutil"
	"net/http"
	"strings"
	"watermelon/common"
	"watermelon/config"
)

var page_404_content string

func Extract_404_page(url string) {
	urlpath := url + config.Error_path
	resp, err := http.Get(urlpath)
	if err != nil {
		return
	}
	page_404, _ := ioutil.ReadAll(resp.Body)
	page_404_content = string(page_404)
}

func Status200(info common.Info) int {
	content_lower := strings.ToLower(info.Content)
	for _, string_404_item := range config.String_404 { //通过404 not found字符来判断状态
		if strings.Contains(content_lower, string_404_item) {
			return -1
		}
	}
	Similarity := Similarity(page_404_content, info.Content)
	if Similarity > config.Similarity {
		return -1
	}
	return 1

}

// 这个方法是如果之前的方法都结束了后还是404，那么会使用其他请求方法进行请求
func Status404(info common.Info) common.Info {
	for _, method := range config.Request_method {
		// 构建请求
		req, _ := http.NewRequest(method, info.URL, nil)
		// 发送请求
		client := &http.Client{}
		resp, _ := client.Do(req)
		defer resp.Body.Close()
		// 返回状态码
		if (resp.StatusCode >= 200 && resp.StatusCode < 400 && resp.ContentLength > 0) || resp.StatusCode == 403 {
			info.Method = method
			info.StatusCode = resp.StatusCode
			content, _ := ioutil.ReadAll(resp.Body)
			info.Content = string(content)
			info.PageLength = len(content)
			break
		}
	}
	return info
}
