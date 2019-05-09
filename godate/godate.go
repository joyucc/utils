package godate

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type goDate struct {
	Time     time.Time
	TimeZone *time.Location
	FirstDayOfWeek time.Weekday
}

//IsBefore checks if the goDate is before the passed goDate
func (d *goDate) IsBefore(compare *goDate) bool {
	return d.Time.Before(compare.Time)
}

//IsAfter checks if the goDate is before the passed goDate
func (d *goDate) IsAfter(compare *goDate) bool {
	return d.Time.After(compare.Time)
}

//Sub subtracts the 'count' from the goDate using the unit passed
func (d goDate) Sub(count int, unit Unit) *goDate {
	return d.Add(-count, unit)
}

//Add adds the 'count' from the goDate using the unit passed
func (d goDate) Add(count int, unit Unit) *goDate {
	d.Time = d.Time.Add(time.Duration(unit * Unit(count)))
	return &d
}

//Get the difference as a duration type
func (d *goDate) DifferenceAsDuration(compare *goDate) time.Duration {
	return d.Time.Sub(compare.Time)
}

//Difference Returns the difference between the Godate and another in the specified unit
//If the difference is negative then the 'compare' date occurs after the date
//Else it occurs before the the date
func (d goDate) Difference(compare *goDate, unit Unit) int {
	difference := d.DifferenceAsFloat(compare, unit)
	return int(difference)
}

//Get the difference as a float
func (d goDate) DifferenceAsFloat(compare *goDate, unit Unit) float64 {
	duration := d.DifferenceAsDuration(compare)
	return float64(duration) / float64(time.Duration(unit))
}

//Gets the difference between the relative to the date value in the form of
//1 month before
//1 month after
func (d goDate) DifferenceForHumans(compare *goDate) string {
	differenceString, differenceInt := d.AbsDifferenceForHumans(compare)
	if differenceInt > 0 {
		return differenceString + " before"
	} else {
		return differenceString + " after"
	}
}

//Gets the difference between the relative to current time value in the form of
//1 month ago
//1 month from now
func (d goDate) DifferenceFromNowForHumans() string {
	now := Now(d.TimeZone)
	differenceString, differenceInt := now.AbsDifferenceForHumans(&d)
	if differenceInt > 0 {
		return differenceString + " ago"
	} else {
		return differenceString + " from now"
	}
}

//Get the abs difference relative to compare time in the form
//1 month
//2 days
func (d goDate) AbsDifferenceForHumans(compare *goDate) (string, int) {
	sentence := make([]string, 2, 2)
	duration := Unit(math.Abs(float64(d.DifferenceAsDuration(compare))))
	var unit Unit
	if duration >= YEAR {
		unit = YEAR
	} else if duration < YEAR && duration >= MONTH {
		unit = MONTH
	} else if duration < MONTH && duration >= WEEK {
		unit = WEEK
	} else if duration < WEEK && duration >= DAY {
		unit = DAY
	} else if duration < DAY && duration >= HOUR {
		unit = HOUR
	} else if duration < HOUR && duration >= MINUTE {
		unit = MINUTE
	} else {
		unit = SECOND
	}
	difference := d.Difference(compare, unit)
	sentence[0] = strconv.Itoa(int(math.Abs(float64(difference))))
	sentence[1] = unit.String()
	if difference == 1 || difference == -1 {
		sentence[1] = strings.TrimSuffix(sentence[1], "s")
	}
	return strings.Join(sentence, " "), difference
}

//MidDay gets the midday time usually 12:00 PM of the current day
func (d *goDate) MidDay() *goDate {
	y, m, day := d.Time.Date()
	return &goDate{time.Date(y, m, day, 12, 0, 0, 0, d.TimeZone), d.TimeZone,0}
}

//ToDateTimeString Formats and returns the goDate in the form 2006-01-02 15:04:05
func (d *goDate) ToDateTimeString() string{
	return d.Format("2006-01-02 15:04:05")
}

//ToDateString Formats and returns the goDate in the form 2006-01-02
func (d *goDate) ToDateString() string{
	return d.Format("2006-01-02")
}

//ToFormattedDateString Formats and returns the goDate in the form Jan 02, 2006
func (d *goDate) ToFormattedDateString() string{
	return d.Format("Jan 02, 2006")
}

//ToTimeString Formats and returns the goDate in the form 15:04:05
func (d *goDate) ToTimeString() string{
	return d.Format("15:04:05")
}

//ToDayTimeString Formats and returns the goDate in the form Mon, Jan 2, 2006 03:04 PM
func (d *goDate) ToDayTimeString() string{
	return d.Format("Mon, Jan 2, 2006 03:04 PM")
}

func (d *goDate) Format(format string) string{
	return d.Time.Format(format)
}

func (d *goDate) SetFirstDay(weekday time.Weekday){
	d.FirstDayOfWeek = weekday
}

func (d goDate) String() string{
	return d.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
}

func (d *goDate) IsLeapYear() bool {
	return (d.Time.Year()%100 != 0 && d.Time.Year()%4 == 0) || (d.Time.Year()%400 == 0)
}

//Timestamp 同 Unix 获取时间戳
func (d *goDate) Timestamp() int64 {
	return d.Time.Unix()
}


func parseWithFormat(str string) (t time.Time, err error) {
	for _, format := range TimeFormats {
		t, err = time.Parse(format, str)
		if err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}

var TimeFormats = []string{"1/2/2006", "1/2/2006 15:4:5", "2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "1-2", "15:4:5", "15:4", "15", "15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02T15:04:05-07:00"}
var hasTimeRegexp = regexp.MustCompile(`(\s+|^\s*)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`) // match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, etc
var onlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)      // match 15:04:05, 15, 15:04:05.000, 15:04:05.000000, etc

// Parse parse string to time
func (d *goDate) Parse(strs ...string) (t time.Time, err error) {
	var (
		setCurrentTime  bool
		parseTime       []int
		currentTime     = []int{d.Time.Nanosecond(), d.Time.Second(), d.Time.Minute(), d.Time.Hour(), d.Time.Day(), int(d.Time.Month()), d.Time.Year()}
		currentLocation = d.Time.Location()
		onlyTimeInStr   = true
	)

	for _, str := range strs {
		hasTimeInStr := hasTimeRegexp.MatchString(str) // match 15:04:05, 15
		onlyTimeInStr = hasTimeInStr && onlyTimeInStr && onlyTimeRegexp.MatchString(str)
		if t, err = parseWithFormat(str); err == nil {
			location := t.Location()
			if location.String() == "UTC" {
				location = currentLocation
			}

			parseTime = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}

			for i, v := range parseTime {
				// Don't reset hour, minute, second if current time str including time
				if hasTimeInStr && i <= 3 {
					continue
				}

				// If value is zero, replace it with current time
				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				// if current time only includes time, should change day, month to current time
				if onlyTimeInStr {
					if i == 4 || i == 5 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}

			t = time.Date(parseTime[6], time.Month(parseTime[5]), parseTime[4], parseTime[3], parseTime[2], parseTime[1], parseTime[0], location)
			currentTime = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
		}
	}
	return
}

// MustParse must parse string to time or it will panic
func (d *goDate) MustParse(strs ...string) (t time.Time) {
	t, err := d.Parse(strs...)
	if err != nil {
		panic(err)
	}
	return t
}