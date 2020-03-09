package go_util

import (
	"fmt"
	"testing"
)

func TestLogField_ToFormatString(t *testing.T) {
	formatString := NewLogFormat().
		SetText("init text").
		//SetFormat("测试f1 %v ,测试 f2 %v").
		//Set("f1", "value1").
		//Set("f2", "value2").
		ToJsonString()
	fmt.Println(formatString)
}
