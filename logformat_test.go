package go_util

import (
	"fmt"
	"testing"
)

func TestLogField_ToFormatString(t *testing.T) {
	formatString := NewLogFormat().
		SetFormat("测试f1 %v ,测试 f2 %v").
		Set("f1", "value1").
		Set("f2", "value2").
		ToJsonString()
	fmt.Println(formatString)
}

type User struct {
	Name string
	Age  int
}

func TestLogField_ToFormatString2(t *testing.T) {
	user := &User{Name: "gene", Age: 30}
	formatString := NewLogFormat().
		SetText("init text").
		//Set("f1", "value1").
		//Set("f2", "value2").
		Set("user", user).
		ToJsonString()
	fmt.Println(formatString)

}

func TestLogField_ToFormatString3(t *testing.T) {
	formatString := NewLogFormat().
		SetFormat("用户ID: %v").
		//Set("f1", "value1").
		//Set("f2", "value2").
		Set("user_id", 55).
		ToFormatJsonString()
	fmt.Println(formatString)
}

func TestLogField_ToFormatString4(t *testing.T) {
	formatString := NewLogFormat().
		SetFormat("----------- %v -----------").
		//Set("f1", "value1").
		//Set("f2", "value2").
		Set("start", "start").
		ToFormatJsonString()
	fmt.Println(formatString)
}