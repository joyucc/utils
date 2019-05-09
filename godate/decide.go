package godate

import (
	"strconv"
	"time"
)

//IsSunday 判断是不是周日
func (d *goDate) IsSunday() bool {
	return d.Time.Weekday().String() == Sunday
}

//IsMonday 判断是不是周一
func (d *goDate) IsMonday() bool {
	return d.Time.Weekday().String() == Monday
}

//IsTuesday 判断是不是周二
func (d *goDate) IsTuesday() bool {
	return d.Time.Weekday().String() == Tuesday
}

//IsWednesday 判断是不是周三
func (d *goDate) IsWednesday() bool {
	return d.Time.Weekday().String() == Wednesday
}

//IsThursday 判断是不是周四
func (d *goDate) IsThursday() bool {
	return d.Time.Weekday().String() == Thursday
}

//IsFriday 判断是不是周五
func (d *goDate) IsFriday() bool {
	return d.Time.Weekday().String() == Friday
}

//IsSaturday 判断是不是周六
func (d *goDate) IsSaturday() bool {
	return d.Time.Weekday().String() == Saturday
}

//IsWeekend 判断是不是周末
func (d *goDate) IsWeekend() bool {
	return d.IsSunday() || d.IsSaturday()
}

//IsWeekday 是否是工作日
func (d *goDate) IsWorkday() bool {
	return !d.IsWeekend()
}

//IsCurrentYear 判断是不是今年
func (d *goDate) IsCurrentYear() bool {
	curYear := Now(time.Local).Format("01")
	return curYear == strconv.Itoa(d.Time.Year())
}

//IsNextYear 判断是不是明天
func (d *goDate) IsNextYear() bool {
	curYear := time.Now().Year()
	return d.Time.Year()-curYear == 1
}

//IsLastYear 判断是不是去年
func (d *goDate) IsLastYear() bool {
	curYear := time.Now().Year()
	return curYear-d.Time.Year() == 1
}

//IsCurrentDay 判断是不是今天
func (d *goDate) IsCurrentDay() bool {
	curDay := time.Now().Day()
	return curDay == d.Time.Day()
}

//IsNextDay 判断是不是明天
func (d *goDate) IsNextDay() bool {
	curDay := time.Now().Day()
	return d.Time.Day()-curDay == 1
}

//IsLastDay 判断是不是昨天
func (d *goDate) IsLastDay() bool {
	curDay := time.Now().Day()
	return curDay-d.Time.Day() == 1
}

//IsCurrentHour 判断是否是当前小时
func (d *goDate) IsCurrentHour() bool {
	curHour := time.Now().Hour()
	return curHour == d.Time.Hour()
}

//IsNextHour 判断是否是下一小时
func (d *goDate) IsNextHour() bool {
	curHour := time.Now().Hour()

	return d.Time.Hour()-curHour == 1
}

//IsLastHour 判断是否是上一小时
func (d *goDate) IsLastHour() bool {
	curHour := time.Now().Hour()

	return curHour-d.Time.Hour() == 1
}

//IsCurrentWeek 判断是不是当前周
func (d *goDate) IsCurrentWeek() bool {
	curWeek := time.Now().Weekday()

	return curWeek == d.Time.Weekday()
}

//IsNextWeek 判断是不是下周
func (d *goDate) IsNextWeek() bool {
	curWeek := time.Now().Weekday()
	return d.Time.Weekday()-curWeek == 1
}

//IsLastWeek 判断是否是上周
func (d *goDate) IsLastWeek() bool {
	curWeek := time.Now().Weekday()

	return curWeek-d.Time.Weekday()== 1
}

//IsCurrentMinute 判断是否是当前分钟
func (d *goDate) IsCurrentMinute() bool {
	curMinute := time.Now().Minute()
	return curMinute == d.Time.Minute()
}

//IsNextMinute 判断是否是下一分钟
func (d *goDate) IsNextMinute() bool {
	curMinute := time.Now().Minute()

	return d.Time.Minute()-curMinute == 1
}

//IsLastMinute 判断是否是上一分钟
func (d *goDate) IsLastMinute() bool {
	curMinute := time.Now().Minute()

	return curMinute-d.Time.Minute() == 1
}

//IsCurrentSecond 判断是否是当前秒
func (d *goDate) IsCurrentSecond() bool {
	curSec := time.Now().Second()
	return curSec == d.Time.Second()
}

//IsNextSecond 判断是否是下一秒
func (d *goDate) IsNextSecond() bool {
	curSec := time.Now().Second()

	return d.Time.Second()-curSec == 1
}

//IsLastSecond 判断是否是上一秒
func (d *goDate) IsLastSecond() bool {
	curSec := time.Now().Second()

	return curSec-d.Time.Second() == 1
}

//IsCurrentMonth 判断是否是当前月
func (d *goDate) IsCurrentMonth() bool {
	curMonth := time.Now().Month()
	return curMonth == d.Time.Month()
}

//IsNextMonth 判断是否是下一个月
func (d *goDate) IsNextMonth() bool {
	curMonth := time.Now().Month()
	return d.Time.Month()-curMonth == 1
}

//IsLastMonth 判断是否是上一个月
func (d *goDate) IsLastMonth() bool {
	curMonth := time.Now().Month()

	return curMonth-d.Time.Month() == 1
}
