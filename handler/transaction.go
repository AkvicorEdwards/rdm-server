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

func transaction(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P18, def.P20)) {
		Fprint(w, "没有查看交易记录的权限")
		return
	}
	if r.Method == "GET" {
		if err := tpl.Transaction.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func getTransactionList(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P18, def.P20)) {
		Fprint(w, "没有获取交易记录的权限")
		return
	}
	if r.Method == "GET" {
		uuid := session.GetUUID(r)
		if len(r.FormValue("uuid")) != 0 && CheckUserPermission(r, def.P20) {
			t, err := strconv.Atoi(r.FormValue("uuid"))
			if err != nil {
				uuid = dam.GetUser(r.FormValue("uuid")).Uuid
			} else {
				uuid = t
			}
		}
		Fprint(w, operator.GetUserTransactionJson(uuid))
	}
}

func updateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		utid, err := strconv.Atoi(r.FormValue("utid"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		UpdateCookieSession(w, r)
		if !CheckUserPermission(r, def.P21) { // 如果没有修改所有人交易记录的权限
			if CheckUserPermission(r, def.P19) { // 如果有修改自己交易记录的权限
				if operator.GetTransaction(utid).Utid != session.GetUUID(r) { // 修改的记录不属于当前用户
					Fprint(w, "没有修改交易记录的权限")
					return
				}
			} else {
				Fprint(w, "没有修改交易记录的权限")
				return
			}
		}

		wuid, err := strconv.Atoi(r.FormValue("wuid"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		unixTime, err := strconv.ParseInt(r.FormValue("unix_time"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		rmb, err := strconv.ParseInt(r.FormValue("rmb"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if operator.UpdateTransaction(utid, wuid, unixTime, rmb) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func deleteTransactions(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if r.Method == "POST" {
		utids := r.FormValue("utid")
		utid := make([]string, 0)
		err := json.Unmarshal([]byte(utids), &utid)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		for _, v := range utid {
			if !CheckUserPermission(r, def.P21) { // 如果没有修改所有人交易记录的权限
				if CheckUserPermission(r, def.P19) { // 如果有修改自己交易记录的权限
					k, err := strconv.Atoi(v)
					if err != nil {
						Fprint(w, "Failure")
						return
					}
					if operator.GetTransaction(k).Utid != session.GetUUID(r) { // 修改的记录不属于当前用户
						Fprint(w, "没有修改交易记录的权限")
						return
					}
				} else {
					Fprint(w, "没有修改交易记录的权限")
					return
				}
			}
		}

		if operator.DeleteTransactions(utid) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}
