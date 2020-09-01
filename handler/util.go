package handler

import (
	"net/http"
	"rdm/cookie"
	"rdm/operator"
	"rdm/session"
	"regexp"
	"strings"
)

// 需要在所有对w的操作之前进行
func UpdateCookieSession(w http.ResponseWriter, r *http.Request) (uuid int) {
	uuid = session.GetUUID(r)
	userInfo, _ := operator.GetUser(uuid)
	cookie.SetUserInfo(w, userInfo)
	session.SetUserInfo(w, r, userInfo)
	return
}

func StringSplitBySpace(str string) []string {
	return regexp.MustCompile(`(\s+)|(\n+\s+\n+)`).Split(strings.TrimSpace(str), -1)
}