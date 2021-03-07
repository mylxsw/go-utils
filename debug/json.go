package debug

import (
	"encoding/json"
	"fmt"
)

// MustMarshalJSON 将对象编码为 json， 出错时返回字符串 对象解析失败
// 注意：仅在调试时使用
func MustMarshalJSON(obj interface{}) string {
	data, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return fmt.Sprintf("parse object to json failed: %v", err)
	}

	return string(data)
}
