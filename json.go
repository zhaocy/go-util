package go_util

import jsoniter "github.com/json-iterator/go"

func ToJSON(data interface{}) string {
	jsonStr, err := jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(data)
	if err != nil {
		Error(err)
		return ""
	}
	return jsonStr
}
