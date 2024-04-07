package main

import (
	"elasticsearch-learning/es7/common"
	"strings"
)

func main() {
	condition := `
		{
			"query": {
				"match": {
					"name": 2
				}
			}
		}
	`

	client := common.NewClient()
	resp, err := client.DeleteByQuery([]string{"book"}, strings.NewReader(condition))
	if err != nil {
		panic(err)
	}

	common.PrintJson(resp)
}
