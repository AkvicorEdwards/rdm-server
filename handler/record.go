package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rdm/dam"
	"rdm/def"
	"rdm/operator"
	"rdm/session"
	"rdm/tpl"
	"strconv"
)

func record(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P14, def.P16)) {
		Fprint(w, "没有查看任务记录的权限")
		return
	}
	if r.Method == "GET" {
		if err := tpl.Record.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func getRecord(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P14, def.P16)) {
		Fprint(w, "没有获取任务记录的权限")
		return
	}
	if r.Method == "GET" {
		urid, err := strconv.Atoi(r.FormValue("urid"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		Fprint(w, operator.GetRecordJson(urid))
	}
}

func getRecordList(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P14, def.P16)) {
		Fprint(w, "没有获取任务记录的权限")
		return
	}
	if r.Method == "GET" {
		uuid := session.GetUUID(r)
		if len(r.FormValue("uuid")) != 0 && CheckUserPermission(r, def.P16) {
			t, err := strconv.Atoi(r.FormValue("uuid"))
			if err != nil {
				uuid = dam.GetUser(r.FormValue("uuid")).Uuid
			} else {
				uuid = t
			}
		}
		Fprint(w, operator.GetUserRecordWithWorkJson(uuid))
	}
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P15, def.P17)) {
		Fprint(w, "没有修改任务记录的权限")
		return
	}
	if r.Method == "POST" {
		urid, err := strconv.Atoi(r.FormValue("urid"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		wuid, err := strconv.Atoi(r.FormValue("wuid"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		startTime, err := strconv.ParseInt(r.FormValue("start_time"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		endTime, err := strconv.ParseInt(r.FormValue("end_time"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		tag := StringSplitBySpace(r.FormValue("tag"))

		if operator.UpdateRecord(urid, wuid, startTime, endTime, tag) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func deleteRecords(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P15, def.P17)) {
		Fprint(w, "没有删除任务记录的权限")
		return
	}
	if r.Method == "POST" {
		urids := r.FormValue("urid")
		urid := make([]string, 0)
		err := json.Unmarshal([]byte(urids), &urid)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if operator.DeleteRecords(urid) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}
