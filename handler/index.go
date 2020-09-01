package handler

import (
	"fmt"
	"net/http"
	"rdm/cookie"
	"rdm/operator"
	"rdm/session"
	"rdm/tpl"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Index.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Login.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := operator.Login(username, password)

	if !ok {
		_, _ = fmt.Fprintln(w, "用户不存在或密码错误")
		return
	}

	session.SetUserInfo(w, r, user)
	cookie.SetUserInfo(w, user)

	http.Redirect(w, r, "/", 302)

}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Register.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}
	username := r.FormValue("username")
	nickname := r.FormValue("nickname")
	password := r.FormValue("password")
	permissionCode := r.FormValue("p_code")

	ok := operator.AddUser(username, nickname, password, permissionCode)
	if !ok {
		_, _ = fmt.Fprintln(w, "ERROR")
		return
	}

	http.Redirect(w, r, "/", 302)
}
