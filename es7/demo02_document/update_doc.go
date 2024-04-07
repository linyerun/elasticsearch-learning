package main

import (
	"bytes"
	"elasticsearch-learning/es7/common"
	"encoding/json"
)

func main() {
	//
	data, _ := json.Marshal(map[string]any{"doc": common.Book{Name: "随心", Age: 22}})
	resp, err := common.NewClient().Update("book", "9737441111", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	common.PrintJson(resp)
}
