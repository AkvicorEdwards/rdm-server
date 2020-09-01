package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rdm/def"
	"rdm/operator"
	"rdm/tpl"
	"strconv"
)

func user(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P12) {
		Fprint(w, "没有查看用户列表的权限")
		return
	}
	if r.Method == "GET" {
		if err := tpl.User.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P11) {
		Fprint(w, "没有创建用户的权限")
		return
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		nickname := r.FormValue("nickname")
		password := r.FormValue("password")
		permissionCode := r.FormValue("p_code")

		if operator.AddUser(username, nickname, password, permissionCode) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P13) {
		Fprint(w, "没有删除用户的权限")
		return
	}
	if r.Method == "POST" {
		uuids := r.FormValue("uuid")
		uuid := make([]string, 0)
		err := json.Unmarshal([]byte(uuids), &uuid)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if operator.DeleteUsers(uuid) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P13) {
		Fprint(w, "没有修改用户的权限")
		return
	}
	if r.Method == "POST" {
		uuid, err := strconv.Atoi(r.FormValue("uuid"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		nickname := r.FormValue("nickname")
		password := r.FormValue("password")

		permission, err := strconv.ParseInt(r.FormValue("permission"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		rmbIn, err := strconv.ParseInt(r.FormValue("rmb_in"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		rmbOut, err := strconv.ParseInt(r.FormValue("rmb_out"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		coin, err := strconv.ParseInt(r.FormValue("coin"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		deleted, err := strconv.ParseInt(r.FormValue("deleted"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if operator.UpdateUser(uuid, nickname, password, permission, rmbIn, rmbOut, coin, deleted) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func getUserList(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P12) {
		Fprint(w, "没有获得用户列表的权限")
		return
	}
	if r.Method == "GET" {
		Fprint(w, operator.GetAllUserJson())
	}
}
