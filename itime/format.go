package itime

import (
	"strconv"
	"time"
)

func GetCurentYearMonthString() string {
	timeNow := time.Now()
	month := timeNow.Month()
	yearMonth := strconv.Itoa(timeNow.Year())
	if month < 10 {
		yearMonth += "0"
	}

	yearMonth += strconv.Itoa(int(month))
	return yearMonth
}

func GetLastYearMonth() (int, int) {
	timeNow := time.Now()
	year := timeNow.Year()
	month := timeNow.Month()
	if month == 1 {
		return year - 1, 12
	} else {
		return year, int(month - 1)
	}
}

func GetLastYearMonthString() string {
	lastYear, lastMonth := GetLastYearMonth()
	yearMonth := strconv.Itoa(lastYear)
	if lastMonth < 10 {
		yearMonth += "0"
	}
	yearMonth += strconv.Itoa(lastMonth)
	return yearMonth
}

/*
* 函数名
*   GetCurentWeekFirstUnixTime
*
* 说明
*       获取本周一的凌晨零点时间戳
*
* 参数说明
*
* RETURNS
*   UNIX时间戳
 */
func GetCurentWeekFirstUnixTime() int64 {
	stdtime := time.Now()

	t := stdtime
	t2 := stdtime
	days := stdtime.Weekday()
	switch days {
	case 0:
		t = stdtime.AddDate(0, 0, -6)
	case 1:
		t = stdtime.AddDate(0, 0, 0)
	case 2:
		t = stdtime.AddDate(0, 0, -1)
	case 3:
		t = stdtime.AddDate(0, 0, -2)
	case 4:
		t = stdtime.AddDate(0, 0, -3)
	case 5:
		t = stdtime.AddDate(0, 0, -4)
	case 6:
		t = stdtime.AddDate(0, 0, 5)
	}
	t2 = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)

	return t2.Unix()
}

// 按照特定格式（"2006-01-02"）字符串返还本周一的日期
func GetCurentWeekFirstDate() string {
	stdtime := time.Now()

	t := stdtime
	t2 := stdtime
	days := stdtime.Weekday()
	switch days {
	case 0:
		t = stdtime.AddDate(0, 0, -6)
	case 1:
		t = stdtime.AddDate(0, 0, 0)
	case 2:
		t = stdtime.AddDate(0, 0, -1)
	case 3:
		t = stdtime.AddDate(0, 0, -2)
	case 4:
		t = stdtime.AddDate(0, 0, -3)
	case 5:
		t = stdtime.AddDate(0, 0, -4)
	case 6:
		t = stdtime.AddDate(0, 0, -5)
	}
	t2 = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)

	return t2.Format("2006-01-02")
}
