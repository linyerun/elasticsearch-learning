package main

import (
	"elasticsearch-learning/es7/common"
	"strings"
)

func main() {
	client := common.NewClient()
	resp, err := client.Search(client.Search.WithIndex("book"), client.Search.WithBody(strings.NewReader(GetByManyFieldIndex())))
	if err != nil {
		panic(err)
	}
	common.PrintJson(resp)
}

func GetAll() string {
	return `
		{
			"query": {
				"match_all": {}
			},
			"from": 0,
			"size": 100
		}
	`
}

func GetByManyFieldIndex() string {
	return `
				{
					"query": {
						"multi_match": {
							"query": "高三",
							"fields": ["class", "name"]
						}
					},
					"from": 0,
					"size": 50
				}
			`
}
