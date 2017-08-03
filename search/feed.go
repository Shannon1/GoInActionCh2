package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// 定义一个类型用来表示(存储)数据源信息
// 最后 ` 反引号里的部分被称作标记（tag）。这个标记里描述了 JSON 解码的元数据，
// 用于创建 Feed 类型值的切片。每个标记将结构类型里字段对应到 JSON 文档里指定名字的字段。
type Feed struct {
	Name 	string 	`json:"site"`
	URI 	string 	`json:"link"`
	Type 	string 	`json:"type"`
}


// 读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 将文件解码到一个存储Feed指针类型的切片（数组）中
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}

