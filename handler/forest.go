package handler

import (
	"fmt"
	"net/http"
	"rdm/def"
	"rdm/tpl"
)

func forest(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P0) {
		Fprint(w, "没有使用Forest的权限")
		return
	}
	if r.Method == "GET" {
		if err := tpl.Forest.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}
