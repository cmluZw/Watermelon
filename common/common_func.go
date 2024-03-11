package common

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"watermelon/config"
)

var client = &http.Client{}

var mu sync.Mutex

type Info struct {
	URL        string
	StatusCode int
	PageLength int
	Content    string
	Err        error
	Redirect   string
	Method     string
}

func Check_url_invalid(url string) string {
	// 定义正则表达式，匹配域名部分
	re := regexp.MustCompile(`^http[s]?://([^/]+)`)
	// 执行正则匹配
	match := re.FindStringSubmatch(url)
	if len(match) >= 2 {
		// 获取匹配到的域名部分
		domain := match[1]
		return domain
	} else {
		return ""
	}
	return ""
}

func check_domain_invalid(domain string) bool {
	// 正则表达式匹配域名的模式
	pattern := `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z]{2,})+$`

	// 编译正则表达式
	re := regexp.MustCompile(pattern)

	// 判断字符串是否匹配域名模式
	if re.MatchString(domain) {
		return true
	} else {
		return false
	}
	return false
}

func Deal_url(url string) string {
	if strings.HasSuffix(url, "/") {
		return url
	} else {
		return url + "/"
	}
}

func Request(url string, wg *sync.WaitGroup) Info {
	//defer wg.Done() // 当函数执行完成时，减少 WaitGroup 的计数器
	resp, err := client.Get(url)
	if err != nil {
		return Info{url, -1, 0, "", err, "", "GET"}
	}
	status_code := resp.StatusCode
	content, _ := ioutil.ReadAll(resp.Body)
	page_length := len(content)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusFound {
		// 获取重定向后的链接
		redirectURL := resp.Header.Get("Location")
		return Info{url, status_code, int(page_length), string(content), nil, redirectURL, "GET"}
	}
	return Info{url, status_code, int(page_length), string(content), nil, "", "GET"}
}

func POST_Request(url string, content_type string, data string, wg *sync.WaitGroup) Info {
	//defer wg.Done() // 当函数执行完成时，减少 WaitGroup 的计数器
	payload := strings.NewReader(data)

	resp, err := client.Post(url, content_type, payload)
	if err != nil {
		return Info{url, -1, 0, "", err, "", "GET"}
	}
	status_code := resp.StatusCode
	content, _ := ioutil.ReadAll(resp.Body)
	page_length := len(content)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusFound {
		// 获取重定向后的链接
		redirectURL := resp.Header.Get("Location")
		return Info{url, status_code, int(page_length), string(content), nil, redirectURL, "GET"}
	}
	return Info{url, status_code, int(page_length), string(content), nil, "", "GET"}
}

func Clear_info(url string, content string, page_length int) Info {
	status_code := 000
	return Info{url, status_code, page_length, string(content), nil, "", "GET"}
}

// 查找链接,其他正则后续参考findsomethings，暂放
func FindLinks(content string) []string {
	// 匹配以 .js 结尾的链接
	//re := regexp.MustCompile(`(?i)(?U)(https?://\S+?\.js\S*)`)
	re := regexp.MustCompile(`href=["']?([^"'>]+)["']?`)
	matches := re.FindAllStringSubmatch(content, -1)
	var links []string
	for _, match := range matches {
		match_tmp := match[1]
		if strings.Contains(match_tmp, " ") {
			match_tmp = strings.Split(match_tmp, " ")[0]
		}

		links = append(links, match_tmp)
	}
	return links
}

func Get_path(links []string) []string {
	var path_list []string
	for _, link := range links {
		if strings.Contains(link, "://") {
			link_tmp := strings.Replace(link, "://", "", -1)
			path_tmp_list := strings.Split(link_tmp, "/")
			for _, path_tmp := range path_tmp_list[1:] {
				flag := 0
				if path_tmp != "" {
					for _, while_item := range config.White_list {
						if strings.Contains(path_tmp, while_item) {
							flag = 1
						}
					}
					if flag == 0 {
						if !check_domain_invalid(path_tmp) {
							path_list = append(path_list, path_tmp)
						}
					} else {
						continue
					}

				}
			}
		} else {
			path_tmp_list := strings.Split(link, "/")
			flag := 0
			for _, path_tmp := range path_tmp_list[1:] {
				if path_tmp != "" {

					for _, while_item := range config.White_list {
						if strings.Contains(path_tmp, while_item) {
							flag = 1
						}
					}
					if flag == 0 {
						if !check_domain_invalid(path_tmp) {
							path_list = append(path_list, path_tmp)
						}

					} else {
						continue
					}
				}
			}
		}

	}

	return set(path_list)
}

func set(list []string) []string {
	// 使用 map 来去重
	uniqueMap := make(map[string]bool)
	for _, item := range list {
		uniqueMap[item] = true
	}

	// 构建去重后的列表
	uniqueList := make([]string, 0, len(uniqueMap))
	for item := range uniqueMap {
		uniqueList = append(uniqueList, item)
	}
	return uniqueList
}
