package spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"watermelon/common"
	"watermelon/config"
)

func Spider_Dict_Generate(url string) {
	file, err := os.Create(config.Spider_dict)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	path_list := common.Get_path(common.FindLinks(string(content)))
	for _, path := range path_list {
		file.WriteString(path + "\n")
	}
}
