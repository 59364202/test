// Copyright 2017 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package datetime provides a function to convert value type time.Time to format.
package datetime

import (
	"haii.or.th/api/server/model/setting"
	"strings"
	"time"
)

const (
	DATE     = "date"
	DATETIME = "datetime"
	RFC3339  = "rfc3339"
)

// DatetimeFormat return format datetime.
//
//	Parameters:
//		t time value type time.Time
//		format output format type string
// 	Return:
//		format datetime
//
func DatetimeFormat(t time.Time, format string) string {
	format = strings.ToUpper(format)
	switch format {
	case DATE:
		return t.Format(setting.GetSystemSetting("setting.Default.DateFormat"))
	case DATETIME:
		return t.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
	case RFC3339:
		return t.Format(time.RFC3339)
	default:
		return t.Format(time.RFC3339)
	}
	return t.Format(time.RFC3339)
}

var weekDayTH = [...]string{
	"อาทิตย์",
	"จันทร์",
	"อังคาร",
	"พุธ",
	"พฤหัสบดี",
	"ศุกร์",
	"เสาร์",
}

//	convert weekday from eng to thai
//	Parameter:
//		wd
//			weekday from time.WeekDay
//	Return:
//		weekday in thai language
func WeekDayTH(wd time.Weekday) string {
	return weekDayTH[int(wd)]
}

var monthsTH = [...]string{
	"มกราคม",
	"กุมภาพันธ์",
	"มีนาคม",
	"เมษายน",
	"พฤษภาคม",
	"มิถุนายน",
	"กรกฎาคม",
	"สิงหาคม",
	"กันยายน",
	"ตุลาคม",
	"พฤศจิกายน",
	"ธันวาคม",
}
var monthsTH_short = [...]string{
	"ม.ค.",
	"ก.พ.",
	"มี.ค.",
	"เม.ย.",
	"พ.ค.",
	"มิ.ย.",
	"ก.ค.",
	"ส.ค.",
	"ก.ย.",
	"ต.ค.",
	"พ.ย.",
	"ธ.ค.",
}

//	month name in th
//	Parameter:
//		m
//			month from time.Month
//	Return:
//		month in thai language
func MonthTH(m time.Month) string { return monthsTH[int(m)-1] }

// 	short month name in th
//	Parameter:
//		m
//			month from time.Month
//	Return:
//		short month in thai language
func MonthTHShort(m time.Month) string { return monthsTH_short[int(m)-1] }

// get first date of week
// example:
//	tn := time.Now().UTC()
//	fmt.Println(tn)
//	year, week := tn.ISOWeek()
//	fmt.Println(year, week)
//
//	date := FirstDayOfISOWeek(year, week, time.UTC)
// sanity check
//	isoYear, isoWeek := date.ISOWeek()
//	if year != isoYear || week != isoWeek {
//		panic(fmt.Sprintf("Input: year %v, week %v. Result: %v, year %v, week %v\n", year, week, date, isoYear, isoWeek))
//	}
//	fmt.Printf("Input: year %v, week %v. Result: %v (%v)\n", year, week, date, date.Weekday())

//	Return:
//		first date of week
func FirstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	return date
}
