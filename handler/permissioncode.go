package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rdm/def"
	"rdm/operator"
	"rdm/session"
	"rdm/tpl"
	"strconv"
)

func permissionCode(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P9) {
		Fprint(w, "没有查看注册码的权限")
		return
	}
	if r.Method == "GET" {
		if err := tpl.PermissionCode.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func addPermissionCode(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P8) {
		Fprint(w, "没有创建注册码的权限")
		return
	}
	if r.Method == "POST" {
		uuid := session.GetUUID(r)
		permissionCode := r.FormValue("p_code")
		permission, err := strconv.ParseInt(r.FormValue("permission"), 10, 64)
		if err != nil {
			Fprint(w, `Failure`)
			return
		}
		deadline, err := strconv.ParseInt(r.FormValue("deadline"), 10, 64)
		if err != nil {
			Fprint(w, `Failure`)
			return
		}

		if operator.AddPermissionCode(permissionCode, permission, deadline, uuid) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func deletePermissionCodes(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P10) {
		Fprint(w, "没有删除注册码的权限")
		return
	}
	if r.Method == "POST" {
		pcids := r.FormValue("pcid")
		pcid := make([]string, 0)
		err := json.Unmarshal([]byte(pcids), &pcid)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if operator.DeletePermissionCodes(pcid) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func updatePermissionCode(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P10) {
		Fprint(w, "没有修改注册码的权限")
		return
	}
	if r.Method == "POST" {
		pcid, err := strconv.Atoi(r.FormValue("pcid"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		code := r.FormValue("code")

		permission, err := strconv.ParseInt(r.FormValue("permission"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		deadline, err := strconv.ParseInt(r.FormValue("deadline"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		generatedBy, err := strconv.Atoi(r.FormValue("generated_by"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		usedBy, err := strconv.Atoi(r.FormValue("used_by"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		deleted, err := strconv.ParseInt(r.FormValue("deleted"), 10, 64)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if operator.UpdatePermissionCode(pcid, generatedBy, usedBy, code, permission, deadline, deleted) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func getPermissionCodeList(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P9) {
		Fprint(w, "没有查看注册码的权限")
		return
	}
	if r.Method == "GET" {
		Fprint(w, operator.GetAllPermissionCodeJson())
	}
}
