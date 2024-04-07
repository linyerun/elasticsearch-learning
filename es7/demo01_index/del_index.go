package main

import "elasticsearch-learning/es7/common"

func main() {
	response, err := common.NewClient().Indices.Delete([]string{"student01", "shopping"})
	if err != nil {
		panic(err)
	}

	common.PrintJson(response)
}
