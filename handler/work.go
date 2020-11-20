package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rdm/dam"
	"rdm/def"
	"rdm/operator"
	"rdm/session"
	"rdm/tpl"
	"strconv"
)

func work(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P3, def.P6)) {
		Fprint(w, "没有查看任务列表的权限")
		return
	}
	if r.Method == "GET" {
		if err := tpl.Work.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func addWork(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P2, def.P5)) {
		Fprint(w, "没有创建任务的权限")
		return
	}
	if r.Method == "POST" {
		uuid := session.GetUUID(r)
		name := r.FormValue("name")
		wType, err := strconv.Atoi(r.FormValue("type"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		wUnit, err := strconv.Atoi(r.FormValue("unit"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		coin, err := strconv.ParseInt(r.FormValue("coin"), 10, 64)
		associate, err := strconv.Atoi(r.FormValue("associate"))
		if err != nil {
			associate = 0
		}
		if operator.AddWork(uuid, name, wType, wUnit, coin, associate) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func deleteWorks(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P4, def.P7)) {
		Fprint(w, "没有删除任务的权限")
		return
	}
	if r.Method == "POST" {
		wuids := r.FormValue("wuid")
		wuid := make([]string, 0)
		err := json.Unmarshal([]byte(wuids), &wuid)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if operator.DeleteWorks(wuid) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func updateWork(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P4, def.P7)) {
		Fprint(w, "没有修改任务的权限")
		return
	}
	if r.Method == "POST" {
		wuid, err := strconv.Atoi(r.FormValue("wuid"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		name := r.FormValue("name")
		iType, err := strconv.Atoi(r.FormValue("type"))
		if err != nil {
			log.Println(err)
			return
		}
		iUnit, err := strconv.Atoi(r.FormValue("unit"))
		if err != nil {
			log.Println(err)
			return
		}
		coin, err := strconv.ParseInt(r.FormValue("coin"), 10, 64)
		if err != nil {
			log.Println(err)
			return
		}
		associatedAuto, err := strconv.Atoi(r.FormValue("associate"))
		if err != nil {
			log.Println(err)
			return
		}
		deleted, err := strconv.ParseInt(r.FormValue("deleted"), 10, 64)
		if err != nil {
			log.Println(err)
			Fprint(w, "Failure")
			return
		}

		if operator.UpdateWork(wuid, name, iType, iUnit, coin, associatedAuto, deleted) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func getWork(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P3, def.P6)) {
		Fprint(w, "没有获取任务信息的权限")
		return
	}
	if r.Method == "GET" {
		wuid, err := strconv.Atoi(r.FormValue("wuid"))
		if err != nil {
			return
		}
		Fprint(w, operator.GetWorkJson(wuid, session.GetUUID(r)))
	}
}

func getWorkList(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P3, def.P6)) {
		Fprint(w, "没有获取任务信息的权限")
		return
	}
	if r.Method == "GET" {
		uuid := session.GetUUID(r)
		if len(r.FormValue("uuid")) != 0 && CheckUserPermission(r, def.P6) {
			t, err := strconv.Atoi(r.FormValue("uuid"))
			if err != nil {
				uuid = dam.GetUser(r.FormValue("uuid")).Uuid
			} else {
				uuid = t
			}
		}
		Fprint(w, operator.GetAllWorkJsonNoDeletedNoToDo(uuid))
	}
}

func getWorkTime(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P3, def.P6)) {
		Fprint(w, "没有获取任务信息的权限")
		return
	}
	if r.Method == "GET" {
		Fprint(w, operator.GetTimeWorkJson(session.GetUUID(r)))
	}
}

func getWorkOnce(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P3, def.P6)) {
		Fprint(w, "没有获取任务信息的权限")
		return
	}
	if r.Method == "GET" {
		Fprint(w, operator.GetOnceWorkJson(session.GetUUID(r)))
	}
}

func getWorkCustom(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P3, def.P6)) {
		Fprint(w, "没有获取任务信息的权限")
		return
	}
	if r.Method == "GET" {
		Fprint(w, operator.GetCustomWorkJson(session.GetUUID(r)))
	}
}

func getWorkAuto(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P3, def.P6)) {
		Fprint(w, "没有获取任务信息的权限")
		return
	}
	if r.Method == "GET" {
		Fprint(w, operator.GetAutoWorkJson(session.GetUUID(r)))
	}
}

func getWorkToDo(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P24) {
		Fprint(w, "没有获取任务信息的权限")
		return
	}
	if r.Method == "GET" {
		Fprint(w, operator.GetToDoWorkJson(session.GetUUID(r)))
	}
}

func getUserWorkAuto(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P3, def.P6)) {
		Fprint(w, "没有获取任务信息的权限")
		return
	}
	if r.Method == "GET" {
		uuid := session.GetUUID(r)
		Fprint(w, operator.GetUserAutoWorkJson(uuid))
	}
}
