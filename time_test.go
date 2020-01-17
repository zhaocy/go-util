package go_util

import (
	"fmt"
	"testing"
	"time"
)

func TestGetNextMonthUnixTime(t *testing.T) {
	fmt.Println(GetNextMonthUnixTime())
	//now := time.Now()                                                                    //获取当前时间，放到now里面，要给next用  
	//next := now.Add(time.Hour * 24)                                                      //通过now偏移24小时
	//next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
	//fmt.Println(next)
}

func TestGetUnixNowTime2(t *testing.T) {
	//fmt.Println(GetUnixNowTime())
	fmt.Println(GetFirstDateOfNextNextMonth(time.Now()))
}

func TestGetAfterNDay(t *testing.T) {
	day := GetAfterNDay(7)
	fmt.Println(day-time.Now().Unix())
}
