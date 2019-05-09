package godate

import (
	"time"
)

type Unit time.Duration

//UNIT CONSTANTS
const (
	SECOND Unit = Unit(time.Second)
	MINUTE = 60 * SECOND
	HOUR = 60 * MINUTE
	DAY   = HOUR * 24
	WEEK  = DAY * 7
	MONTH = DAY * 30
	YEAR  = DAY * 365
)

const (
	//January 一月
	January = 1 + iota
	//February 二月
	February
	//March 三月
	March
	//April 四月
	April
	//May 五月
	May
	//June 六月
	June
	//July 七月
	July
	//August 八月
	August
	//September 九月
	September
	//October 十月
	October
	//November 十一月
	November
	//December 十二月
	December
)

var UnitStrings = map[Unit]string{
	SECOND: "seconds",
	MINUTE: "minutes",
	HOUR:   "hours",
	DAY:    "days",
	WEEK:   "weeks",
	MONTH:  "months",
	YEAR:   "years",
}

const (
	//Sunday 周日
	Sunday = "Sunday"
	//Monday 周一
	Monday = "Monday"
	//Tuesday 周二
	Tuesday = "Tuesday"
	//Wednesday 三
	Wednesday = "Wednesday"
	//Thursday 四
	Thursday = "Thursday"
	//Friday 五
	Friday = "Friday"
	//Saturday 六
	Saturday = "Saturday"
)

func (u Unit) String() string{
	return UnitStrings[u]
}

func Create(time time.Time) *goDate {
	return &goDate{time,time.Location(),0}
}

func Now(location *time.Location) *goDate {
	return &goDate{time.Now().In(location),location,0}
}

func Tomorrow(location *time.Location) *goDate {
	tomorrow := Now(location).Add(1,DAY)
	return tomorrow
}

func Yesterday(location *time.Location) *goDate {
	yesterday := Now(location).Sub(1,DAY)
	return yesterday
}

func Parse(layout, value string) (*goDate,error){
	parsedTime, err := time.Parse(layout,value)
	if err != nil{
		return nil, err
	}
	return &goDate{Time: parsedTime, TimeZone: parsedTime.Location()},nil
}
