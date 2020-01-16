package go_util

import (
	"fmt"
	"strconv"
	"time"
)

func GetTimeFromStrDate(date string) (year, month, day int) {
	const shortForm = "2006-01-02"
	d, err := time.Parse(shortForm, date)
	if err != nil {
		Debug("出生日期解析错误！")
		return 0, 0, 0
	}
	year = d.Year()
	month = int(d.Month())
	day = d.Day()
	return
}

func GetUnixNowTime() int64{
	timeLayout := "2006-01-02"
	date := time.Now().Format(timeLayout)
	fmt.Println(date)
	loc, _ := time.LoadLocation("Local")                      //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, date, loc) //使用模板在对应时区转化为time.time类型
	return theTime.Unix()
}

//计算下一个月份1号
func GetNextMonthUnixTime() int64 {
	date1 := time.Time.AddDate(time.Now(),0, 1, 0)
	next_month := date1.Format("01")
	year:=strconv.Itoa(time.Now().Year())
	dateStr := fmt.Sprintf("%s-%s-%s %s", year , next_month , "01","00:00:00")
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 00:00:00", dateStr,loc)

	return theTime.Unix()
}


//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day() + 1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取下个月的第一天
func GetFirstDateOfNextMonth(d time.Time) time.Time{
	d = d.AddDate(0,1,-d.Day()+1)
	return GetZeroTime(d)
}

//获取下个月的最后一天
func GetLastDateOfNextMonth(d time.Time) time.Time{
	return GetFirstDateOfNextMonth(d).AddDate(0,1,-1)
}

//获取下下个月的第一天
func GetFirstDateOfNextNextMonth(d time.Time)time.Time{
	return GetLastDateOfNextMonth(d).AddDate(0,0,1)
}

//获取下下个月的第一天（返回Long)
func GetFirstDateOfNextNextMonthUnixTime(d time.Time) int64{
	return GetLastDateOfNextMonth(d).AddDate(0,0,1).Unix()
}

//计算N天之后
func GetAfterNDay(day int) int64{
	k:= time.Now()
	ad, _ := time.ParseDuration("24h")
	d:= time.Duration(day)
	return k.Add(ad * d).Unix()
}