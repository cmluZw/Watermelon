package application

import (
	"fmt"
	"strings"
	"watermelon/common"
)

type parse_result struct {
	URL       string
	scan_mode string
	filepath  string
	result    int
}

func Parse(args []string) parse_result {
	scan_mode, url, filename := "default", "", "./dict/dict.txt"
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], "-h") {
			pring_usage()
			return parse_result{"", "", filename, 0}
		} else if strings.Contains(args[i], "-m") {
			scan_mode = args[i+1]
		} else if strings.Contains(args[i], "-u") {
			url_tmp := args[i+1]
			if common.Check_url_invalid(url_tmp) != "" {
				url = url_tmp
			} else {
				pring_usage()
				return parse_result{"", "", filename, 0}
			}
		} else if strings.Contains(args[i], "-f") {
			filename_tmp := args[i+1]
			if common.Check_file_exist(filename_tmp) {
				filename = filename_tmp
			}

		}
	}
	if url != "" && scan_mode != "" && filename != "" {
		return parse_result{url, scan_mode, filename, 1}
	}
	return parse_result{url, scan_mode, filename, 0}
}

func pring_usage() {
	Watermelon_logo := `
 __      __         __                              .__                 
/  \    /  \_____ _/  |_  ___________  _____   ____ |  |   ____   ____  
\   \/\/   /\__  \\   __\/ __ \_  __ \/     \_/ __ \|  |  /  _ \ /    \ 
 \        /  / __ \|  | \  ___/|  | \/  Y Y  \  ___/|  |_(  <_> )   |  \
  \__/\  /  (____  /__|  \___  >__|  |__|_|  /\___  >____/\____/|___|  /
       \/        \/          \/            \/     \/                 \/ 
`
	default_usage_help := "watermelon is a dirsearch tool by cmluZw"
	default_usage_info := "watermelon example:"
	default_usage_example := `
[!] watermelon -u http://www.example.com/       default scan,just spider current page and use default dict to dirsearch
[!] watermelon -u http://www.example.com/ -m s  open Similar scan
`
	fmt.Print(Watermelon_logo)
	fmt.Println(default_usage_help)
	fmt.Print(default_usage_info)
	fmt.Print(default_usage_example)
}
