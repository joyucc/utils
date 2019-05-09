package godate

import "strconv"

//CountDayForYear 返回一年的天数，如果是闰年则返回366天
func (d *goDate) CountDayForYear() int {
	if d.IsLeapYear() {
		return 366
	}
	return 365
}

//CountDayForMonth 返回每个月的天数
func (d *goDate) CountDayForMonth() int {
	month := d.Format("1")
	m, _ := strconv.Atoi(month)
	switch m {
	case January, March, May, July, August, October, December:
		return 31
	case February:
		if c.IsLeapYear() {
			return 29
		}
		return 28
	case April, June, September, November:
		return 30
	default:
		return 30

	}
}
