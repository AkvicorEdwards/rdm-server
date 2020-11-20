package handler

import (
	"fmt"
	"net/http"
	"rdm/def"
	"rdm/tpl"
)

func todo(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P24) {
		Fprint(w, "没有使用ToDo的权限")
		return
	}
	if r.Method == "GET" {
		if err := tpl.ToDo.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}