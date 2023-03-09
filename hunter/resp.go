package hunter

import (
	"encoding/json"
	"fmt"
	"os"
)

type Reqe struct {
	Is_risk   string `json:"is_risk"`
	Url       string `json:"url"`
	Ip        string `jsom:"ip"`
	Port      string `jsom:"port"`
	Web_Title string `jsom:"web_title"`
	Domain    string `jsom:"domain"`
	Component string `jsom:"component"`
	Os        string `jsom:"os"`
	Company   string `jsom:"company"`
	Banner    string `jsom:"banner"`
}

func Files(body []byte) {
	var Data Reqe
	err := json.Unmarshal(body, &Data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 将JSON数据写入文件
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(Data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Data saved to file.")
}

//func Excel(bodyString string) {
//	// 创建一个新的 Excel 文件
//	file := xlsx.NewFile()
//	// 创建一个名为 "Sheet1" 的工作表
//	sheet, err := file.AddSheet("Sheet1")
//	if err != nil {
//		fmt.Println(err)
//		return
//	} // 在工作表中添加一些数据
//	row := sheet.AddRow()
//	cell := row.AddCell()
//	cell.Value = "is_risk"
//	cell = row.AddCell()
//	cell.Value = "url"
//	cell = row.AddCell()
//	cell.Value = "ip"
//	cell = row.AddCell()
//	cell.Value = "port"
//	cell = row.AddCell()
//	cell.Value = "web_title"
//	cell = row.AddCell()
//	cell.Value = "domain"
//	cell = row.AddCell()
//	cell.Value = "component"
//	cell = row.AddCell()
//	cell.Value = "os"
//	cell = row.AddCell()
//	cell.Value = "company"
//	cell = row.AddCell()
//	cell.Value = "banner"
//
//	var data []map[string]interface{}
//	bodyReader := bytes.NewBuffer([]byte(bodyString))
//	err = json.NewDecoder(bodyReader).Decode(&data)if err != nil {
//		panic(err)
//	}
//	for row, item := range data { // 遍历JSON数据中的每个对象
//		rowNum := row + 2 // 行号从2开始，因为第一行已经写入了表头
//		for col, header := range 9 { // 遍历表头中的每一列
//			cell := fmt.Sprintf("%s%d", string('A'+col), rowNum) // 将列号与行号组合成单元格地址
//			value, ok := item[header] // 获取该列的值
//			if ok { // 如果该列的值存在
//				f.SetCellValue(sheetName, cell, value) // 在单元格中写入该列的值
//			}
//		}
//	}
//	//保存xlsx文件
//	err = file.Save("example.xlsx")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//var excel Reqe
//
//}
