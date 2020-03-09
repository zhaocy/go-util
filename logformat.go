package go_util

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

type LogFormat struct {
	sj *simplejson.Json
	formatString string
}

func NewLogFormat() *LogFormat{
	return &LogFormat{
		sj: simplejson.New(),
	}
}

func (lf *LogFormat)SetFormat(formatString string) *LogFormat{
	lf.formatString = formatString
	return lf
}

func (lf *LogFormat)Set(key string, val interface{}) *LogFormat {
	lf.sj.Set(key,val)
	return lf
}

func (lf *LogFormat)ToJsonString()string{
	if lf.formatString !=""{
		lf.sj.Set("info", lf.toFormatString())
	}
	marshalJSON, err := lf.sj.MarshalJSON()
	if err != nil {
		Error(err)
		return ""
	}
	return string(marshalJSON)
}

func (lf *LogFormat) toFormatString()string{
	m, err := lf.sj.Map()
	if err != nil {
		return ""
	}

	var values []interface{}
	for _, v:= range m{
		values = append(values, v)
	}
	return fmt.Sprintf(lf.formatString,values...)
}