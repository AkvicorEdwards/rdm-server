package handler

import (
	"fmt"
	"net/http"
	"rdm/def"
	"rdm/tpl"
)

func topUp(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P1) {
		Fprint(w, "没有使用充值的权限")
		return
	}

	if r.Method == "GET" {
		UpdateCookieSession(w, r)
		if err := tpl.TopUp.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}
