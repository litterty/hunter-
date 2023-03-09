package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"go_hunter2/hunter"
	"os"
)

func main() {
	var search string
	api_key, page, page_size := hunter.Read()
	flag.StringVar(&search, "s", "", "搜索规则")
	//flag.StringVar(, "o", "", "保存到excel文件")
	flag.Parse()
	// 判断是否需要显示帮助信息
	if flag.NFlag() == 0 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// 判断search是否为空
	if search == "" {
		fmt.Println("Error: search keyword is required")
		os.Exit(1)
	}
	search = base64.URLEncoding.EncodeToString([]byte(search))
	//fmt.Println(api_key, page, page_size, search)
	body, err := hunter.Req(api_key, page, page_size, search)
	if err != nil {
		os.Exit(1)
	}
	hunter.Files(body)
	bodyString := string(body) //byte转换成string

	fmt.Println(bodyString, err)
	//hunter.Excel(bodyString)
}
