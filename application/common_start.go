package application

import (
	"watermelon/common"
	"watermelon/scan"
	"watermelon/spider"
)

func Start(info parse_result) {
	if info.result == 0 {
		return
	}
	url := common.Deal_url(info.URL)
	spider.Spider_Dict_Generate(url)
	scan_mode := info.scan_mode
	if scan_mode == "default" {
		scan.SimpleScanByDict(info.filepath, url)
	} else if scan_mode == "s" {
		scan.ScanSimilarity(info.filepath, url)
	}
}
