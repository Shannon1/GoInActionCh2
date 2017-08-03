package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

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

