package main

import (
	"elasticsearch-learning/es7/common"
	"strings"
)

func main() {
	indexMapping := `
		{
		  "mappings": {
			"properties": {
			  "author": {
				"index": true,
				"type": "text"
			  },
			  "age": {
				"index": false,
				"type": "integer"
			  },
			  "name": {
				"index": true,
				"type": "text"
			  }
			}
		  }
		}
	`

	client := common.NewClient()
	resp, err := client.Indices.Create("book", client.Indices.Create.WithBody(strings.NewReader(indexMapping)))
	if err != nil {
		panic(err)
	}

	common.PrintJson(resp)
}
