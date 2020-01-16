package go_util

import (
	"go/types"
	"regexp"
	"strings"
)

const (
	REG_RETRYTIMES  = 50
	UUID_MAX_LENGTH = 50
	UUID_MIN_LENGTH = 10
)



func CheckDeviceId(deviceId string) bool {
	if len(deviceId) < UUID_MIN_LENGTH || len(deviceId) > UUID_MAX_LENGTH {
		return false
	}
	return true
}

func CheckEmail(email string) bool {
	Re := regexp.MustCompile(`^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$`)
	return Re.MatchString(email)
}

func CheckPassWord(passwd string) bool {
	pw := strings.ToLower(passwd)
	pattern := "^[a-zA-Z0-9_.]{8,20}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(pw)
}

func CheckUserName(userName string) bool {
	userName = strings.ToLower(userName)
	pattern := "^[a-zA-Z][a-zA-Z0-9_]{5,17}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(userName)
}

//是否在数组
func InArray(needle interface{}, hystack_array interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack_array.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack_array.([]int) {
			if key == item {
				return true
			}
		}

	case int32:
		for _, item := range hystack_array.([]int32) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack_array.([]int64) {
			if key == item {
				return true
			}
		}

	default:
		return false
	}

	return false
}

//(三元运算)
func If(conditions bool, trueVal, falseVal interface{}) interface{} {
	if conditions {
		return trueVal
	}
	return falseVal
}


//判断是否为空
func IS_Empty(arg interface{}) bool {

	switch arg.(type) {

	case int:
		return If(arg.(int) == int(0), true, false).(bool)

	case int32:
		return If(arg.(int32) == int32(0), true, false).(bool)
	case int64:
		return If(arg.(int64) == int64(0), true, false).(bool)

	case float64:
		return If(arg.(float64) == float64(0.00), true, false).(bool)

	case []byte:

		return If(len(arg.([]byte))  == 0, true, false).(bool)

	case string:
		return If(arg.(string) == " " || arg.(string) == "" || arg.(string) == "0" || arg.(string) == "NULL", true, false).(bool)

	case map[string]interface{}:

		return If(len(arg.(map[string]interface{})) == 0, true, false).(bool)

	case []interface{}:

		return If(len(arg.([]interface{})) == 0, true, false).(bool)

	case []string:

		return If(len(arg.([]string)) == 0, true, false).(bool)


	case []int64:

		return If(len(arg.([]int64)) == 0, true, false).(bool)

	case []float64:

		return If(len(arg.([]float64)) == 0, true, false).(bool)

	case []int:

		return If(len(arg.([]int)) == 0, true, false).(bool)

	case []map[string]interface{}:

		return If(len(arg.([]map[string]interface{})) == 0, true, false).(bool)

	case types.Nil:

		return If(arg == nil, true, false).(bool)

	case bool:

		return If(!arg.(bool), true, false).(bool)

	default:

		return true
	}
}
