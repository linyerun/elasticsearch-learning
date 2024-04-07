package common

import (
	es "github.com/elastic/go-elasticsearch/v7"
	"sync"
)

var once sync.Once
var t *es.Client

func NewClient() *es.Client {
	once.Do(func() {
		var err error
		t, err = es.NewClient(es.Config{Addresses: []string{"http://127.0.0.1:9200"}}) // 不可以少了http://这个前缀
		if err != nil {
			panic(err)
		}
	})
	return t
}
