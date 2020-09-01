package handler

import (
	"fmt"
	"net/http"
	"rdm/dam"
	"rdm/def"
	"rdm/operator"
	"rdm/session"
	"rdm/tpl"
	"strconv"
)

func analyze(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P22, def.P23)) {
		Fprint(w, "没有查看分析的权限")
		return
	}
	if r.Method == "GET" {
		if err := tpl.Analyze.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func getAnalyze(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		UpdateCookieSession(w, r)
		if !CheckUserPermission(r, def.CombinePermission(def.P22, def.P23)) {
			Fprint(w, "没有查看分析的权限")
			return
		}
		uuid := session.GetUUID(r)
		if len(r.FormValue("uuid")) != 0 && CheckUserPermission(r, def.P23) {
			t, err := strconv.Atoi(r.FormValue("uuid"))
			if err != nil {
				uuid = dam.GetUser(r.FormValue("uuid")).Uuid
			} else {
				uuid = t
			}
		}
		switch r.FormValue("c") {
		case "day":
			year, err := strconv.Atoi(r.FormValue("year"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			month, err := strconv.Atoi(r.FormValue("month"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			day, err := strconv.Atoi(r.FormValue("day"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			Fprint(w, operator.AnalyzeDayJson(uuid, year, month, day))
		case "month":
			year, err := strconv.Atoi(r.FormValue("year"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			month, err := strconv.Atoi(r.FormValue("month"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			Fprint(w, operator.AnalyzeMonthJson(uuid, year, month))
		case "year":
			year, err := strconv.Atoi(r.FormValue("year"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			Fprint(w, operator.AnalyzeYearJson(uuid, year))
		case "custom":
			year, err := strconv.Atoi(r.FormValue("year"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			month, err := strconv.Atoi(r.FormValue("month"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			day, err := strconv.Atoi(r.FormValue("day"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			yearEnd, err := strconv.Atoi(r.FormValue("year_end"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			monthEnd, err := strconv.Atoi(r.FormValue("month_end"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			dayEnd, err := strconv.Atoi(r.FormValue("day_end"))
			if err != nil {
				Fprint(w, "Failure")
				return
			}
			Fprint(w, operator.AnalyzeCustomJson(uuid, year, month, day, yearEnd, monthEnd, dayEnd))
		default:
			Fprint(w, "Failure")
		}

	}
}
