package main

import "elasticsearch-learning/es7/common"

func main() {
	resp, err := common.NewClient().Delete("book", "9737441111")
	if err != nil {
		panic(err)
	}

	common.PrintJson(resp)
}
