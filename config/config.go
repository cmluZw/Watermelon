package config

// 全局配置
var (
	AppName    = "WaterMelon"
	AppVersion = "1.0.0"
	LogLevel   = "debug"
	proxy      = ""
	userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
		"Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.181 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 11; Pixel 3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.181 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/88.0.4324.152 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPad; CPU OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/88.0.4324.152 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:86.0) Gecko/20100101 Firefox/86.0",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:86.0) Gecko/20100101 Firefox/86.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:86.0) Gecko/20100101 Firefox/86.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 11.2; rv:86.0) Gecko/20100101 Firefox/86.0",
		"Mozilla/5.0 (Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0",
		"Mozilla/5.0 (Android 10; Mobile; rv:86.0) Gecko/86.0 Firefox/86.0",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/32.0 Mobile/15E148 Safari/605.1.15",
		"Mozilla/5.0 (iPad; CPU OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/32.0 Mobile/15E148 Safari/605.1.15",
		"Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.181 Mobile Safari/537.36 EdgA/46.03.4.5155",
		"Mozilla/5.0 (Linux; Android 11; Pixel 3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.181 Mobile Safari/537.36 EdgA/46.03.4.5155",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36 Edg/88.0.705.81",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36 Edg/88.0.705.81",
	}
	references          = ""
	delay_time          = ""              //时延
	deepth              = 2               //扫描深度
	exclude_status_code = []string{"404"} //排除状态码
	Similarity          = 0.95            //是否开启相似度检验
	similar_length      = 0               //相似的页面长度
	similar_content     = ""              //相似的内容
	Is_String_Blaste    = 0               //是否开启爆破
	Blaste_min_length   = 1               //爆破的字符串最小长度
	Blaste_max_length   = 5               //爆破的字符串最大长度
	Blaste_str          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"

	Is_Spider   = 0         //是否打开页面爬取
	Spider_type = []string{ //爬取后缀类型
		"html", "js",
	}
	customize_http_header = make(map[string]string) //自定义请求头
	session               = ""                      //session设置
	dict_path             = ""                      //字典路径
	String_404            = []string{"404", "not found", "404 not found", "error"}
	Error_path            = "/Watermelon_project"
	Request_method        = []string{"POST", "PUT", "MOVE", "PATCH", "OPTIONS", "HEAD"}
	White_list            = []string{".css", ".jpg", ".png", ".gif", ".ico"}
	Spider_dict           = "./dict/Spider/dict.txt"
)
