package hunter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const BaseURL = "https://hunter.qianxin.com/openApi/search?"

type MyConfig struct {
	Api_key   string `json:"api_key"`
	Page      int    `json:"page"`
	Page_size int    `jsom:"page_size"`
}

func Read() (string, int, int) {
	//读取配置文件
	file, err := os.Open("B:\\代码\\go_hunter2\\hunter\\config.json")
	if err != nil {
		// 处理错误
	}
	defer file.Close()

	var config MyConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		// 处理错误
	}

	// 打印参数
	//fmt.Println(config.Api_key, config.Page, config.Page_size)
	return config.Api_key, config.Page, config.Page_size
}

func Req(api_key string, page int, page_size int, search string) ([]byte, error) {
	//api请求
	api_url := fmt.Sprintf("%sapi-key=%s&search=%s&page=%d&page_size=%d", BaseURL, api_key, search, page, page_size)
	fmt.Println(api_url)
	response, err := http.Get(api_url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
