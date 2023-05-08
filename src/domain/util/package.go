package util

import (
	"encoding/json"
	"fmt"
)

func ParseStruct(objName string, obj interface{}) string {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("[%s]%s", objName, jsonData)
}
