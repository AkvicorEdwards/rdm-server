package session

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"rdm/dam"
)

// 获取全局session
func GetUser(r *http.Request) (*sessions.Session, error) {
	return Get(r, "user")
}

// 设置用户的Session
func SetUserInfo(w http.ResponseWriter, r *http.Request, user *dam.User) {
	// Session
	ses, _ := GetUser(r)
	ses.Values["uuid"] = user.Uuid
	ses.Values["username"] = user.Username
	ses.Values["nickname"] = user.Nickname
	ses.Values["password"] = user.Password
	ses.Values["permission"] = user.Permission
	ses.Values["plant"] = user.Plant
	ses.Values["rmb_in"] = user.RmbIn
	ses.Values["rmb_out"] = user.RmbOut
	ses.Values["coin"] = user.Coin
	ses.Options.MaxAge = 60 * 60 * 24
	err := ses.Save(r, w)
	if err != nil {
		_, _ = fmt.Fprintln(w, "ERROR session SetUserInfo")
		return
	}
}

func GetUUID(r *http.Request) int {
	ses, err := Get(r, "user")
	if err != nil {
		return 0
	}
	uuid, ok := ses.Values["uuid"].(int)
	if !ok {
		return 0
	}
	return uuid
}

func GetPermission(r *http.Request) int64 {
	ses, err := Get(r, "user")
	if err != nil {
		return 0
	}
	permission, ok := ses.Values["permission"].(int64)
	if !ok {
		return 0
	}
	return permission
}
