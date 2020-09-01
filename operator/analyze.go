package operator

import (
	"encoding/json"
	"rdm/dam"
)

func AnalyzeDay(uuid, year, month, day int) []dam.RecordWithWork {
	start, end := DayUnix(year, month, day)
	return recordWithWork(dam.GetUserRecordByTime(uuid, start, end))
}

func AnalyzeDayJson(uuid, year, month, day int) string {
	records := AnalyzeDay(uuid, year, month, day)
	res, err := json.Marshal(records)
	if err != nil {
		return ""
	}
	return string(res)
}

func AnalyzeMonth(uuid, year, month int) []dam.RecordWithWork {
	start, end := MonthUnix(year, month)
	return recordWithWork(dam.GetUserRecordByTime(uuid, start, end))
}

func AnalyzeMonthJson(uuid, year, month int) string {
	records := AnalyzeMonth(uuid, year, month)
	res, err := json.Marshal(records)
	if err != nil {
		return ""
	}
	return string(res)
}

func AnalyzeYear(uuid, year int) []dam.RecordWithWork {
	start, end := YearUnix(year)
	records := recordWithWork(dam.GetUserRecordByTime(uuid, start, end))
	return records
}

func AnalyzeYearJson(uuid, year int) string {
	records := AnalyzeYear(uuid, year)
	res, err := json.Marshal(records)
	if err != nil {
		return ""
	}
	return string(res)
}

func AnalyzeCustom(uuid, yearStart, monthStart, dayStart, yearEnd, monthEnd, dayEnd int) []dam.RecordWithWork {
	start, end := CustomUnix(yearStart, monthStart, dayStart, yearEnd, monthEnd, dayEnd)
	return recordWithWork(dam.GetUserRecordByTime(uuid, start, end))
}

func AnalyzeCustomJson(uuid, yearStart, monthStart, dayStart, yearEnd, monthEnd, dayEnd int) string {
	records := AnalyzeCustom(uuid, yearStart, monthStart, dayStart, yearEnd, monthEnd, dayEnd)
	res, err := json.Marshal(records)
	if err != nil {
		return ""
	}
	return string(res)
}

func recordWithWork(record []dam.Record) []dam.RecordWithWork {
	records := make([]dam.RecordWithWork, 0)
	for _, v := range record {
		records = append(records, dam.RecordWithWork{
			Record: v,
			Work:   *dam.GetWork(v.Wuid, v.Uuid, true),
		})
	}
	return records
}
