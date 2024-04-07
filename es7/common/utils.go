package common

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"math/rand"
	"os"
)

func PrintJson(resp *esapi.Response) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	json.Indent(&out, data, "", "  ")
	out.WriteTo(os.Stdout)
}

func GetId() string {
	var res []byte
	for i := 0; i < 10; i++ {
		res = append(res, '0'+byte(rand.Int()%10))
	}
	return string(res)
}
