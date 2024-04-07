package main

import (
	"bytes"
	"elasticsearch-learning/es7/common"
	"encoding/json"
	"fmt"
)

// 0887736364
// 9737441111
func main() {
	//data, err := json.Marshal(common.Book{Author: "随风", Age: 2, Name: "设计模式"})
	//if err != nil {
	//	panic(err)
	//}
	//
	//resp, err := common.NewClient().Create("book", common.GetId(), bytes.NewReader(data))
	//if err != nil {
	//	panic(err)
	//}
	//
	//common.PrintJson(resp)

	createUsers()
}

func createUsers() {
	client := common.NewClient()
	names := []string{"华为", "小米", "小鹏", "理想", "蔚来"}
	for i := 0; i < 100; i++ {
		data, _ := json.Marshal(common.User{Age: i, Name: names[i%len(names)], Class: fmt.Sprintf("高三%d班", i+1), Introduce: "自我介绍..."})
		_, err := client.Create("user", common.GetId(), bytes.NewReader(data))
		if err != nil {
			panic(err)
		}
		fmt.Println(i)
	}
}
