package tool

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"errors"
	"io"
	"reflect"
)

// 判断obj是否在target中，target支持的类型arrary,slice,map
func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

//进行zlib解压缩 + base64Decode, 輸入string 輸出 string
func DoZlibUnCompressGetString(Str string) string {
	compressSrc, err := base64.StdEncoding.DecodeString(Str)
	if err != nil {
		return ""
	}
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	if err != nil {
		return ""
	}
	io.Copy(&out, r)

	// 輸出
	OutStr := string(out.Bytes())
	return OutStr
}
