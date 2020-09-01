package operator

import (
	"time"
)

func abs64(n int64) int64 {
	if n > 0 {
		return n
	}
	return -n
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func DayUnix(year, month, day int) (start int64, end int64) {
	start = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Unix()
	end = time.Date(year, time.Month(month), day, 23, 59, 59, 0, time.Local).Unix()
	return start, end
}

func MonthUnix(year, month int) (start int64, end int64) {
	nextMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	nextMonth = nextMonth.AddDate(0, 1, -1)
	start = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local).Unix()
	end = time.Date(nextMonth.Year(), nextMonth.Month(), nextMonth.Day(), 23, 59, 59, 0, time.Local).Unix()
	return start, end
}

func YearUnix(year int) (start int64, end int64) {
	nextYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
	nextYear = nextYear.AddDate(1, 0, -1)
	start = time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local).Unix()
	end = time.Date(nextYear.Year(), nextYear.Month(), nextYear.Day(), 23, 59, 59, 0, time.Local).Unix()
	return start, end
}

func CustomUnix(yearStart, monthStart, dayStart, yearEnd, monthEnd, dayEnd int) (start int64, end int64) {
	if yearStart == 0 {
		start = 0
	} else if monthStart == 0 {
		start, _ = YearUnix(yearStart)
	} else if dayStart == 0 {
		start, _ = MonthUnix(yearStart, monthStart)
	} else {
		start, _ = DayUnix(yearStart, monthStart, dayStart)
	}
	if yearEnd == 0 {
		_, end = YearUnix(2100)
	} else if monthEnd == 0 {
		_, end = YearUnix(yearEnd)
	} else if dayEnd == 0 {
		_, end = MonthUnix(yearEnd, monthEnd)
	} else {
		_, end = DayUnix(yearEnd, monthEnd, dayEnd)
	}
	return
}
