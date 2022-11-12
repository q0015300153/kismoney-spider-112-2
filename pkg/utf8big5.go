package pkg

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// 轉換 UTF-8 至 BIG5
func Utf8Big5(str string) string {
	convertStr := bytes.NewReader([]byte(str))
	resultStr := transform.NewReader(convertStr, traditionalchinese.Big5.NewEncoder())
	data, err := ioutil.ReadAll(resultStr)
	if err != nil {
		return string(data)
	}

	return string(data)
}
