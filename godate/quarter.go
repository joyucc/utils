package godate

import "time"

//Quarter 季度类型
type Quarter int8

//Next 返回下一季度
func (q Quarter) Next() Quarter {
	if q == 4 {
		return 1
	}
	return q + 1
}

//Last 返回上一个季度
func (q Quarter) Last() Quarter {
	if q == 1 {
		return 4
	}
	return q - 1
}



//CurrentQuarter 返回当前季度
func (d *goDate) CurrentQuarter() Quarter {
	switch {
	case 1 <= d.Time.Month() && d.Time.Month() <= 3:
		return 1
	case 4 <= d.Time.Month() && d.Time.Month() <= 6:
		return 2
	case 7 <= d.Time.Month() && d.Time.Month() <= 9:
		return 3
	case 10 <= d.Time.Month() && d.Time.Month() <= 12:
		return 4
	default:
		return 0
	}
}

//IsCurrentQuarter 判断是否是当前季度
func (d *goDate) IsCurrentQuarter() bool {
	return d.CurrentQuarter() == Now(time.Local).CurrentQuarter()
}

//IsNextQuarter 判断是否是下一季度
func (d *goDate) IsNextQuarter() bool {
	return d.CurrentQuarter() == Now(time.Local).CurrentQuarter().Next()
}

//IsLastQuarter 判断是否是上一季度
func (d *goDate) IsLastQuarter() bool {
	return d.CurrentQuarter() == Now(time.Local).CurrentQuarter().Last()
}
