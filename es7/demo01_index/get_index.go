package main

import "elasticsearch-learning/es7/common"

func main() {
	client := common.NewClient()

	resp, err := client.Indices.Get([]string{"book"})
	if err != nil {
		panic(err)
	}

	common.PrintJson(resp)
}
