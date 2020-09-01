package cookie

import (
	"fmt"
	"net/http"
	"rdm/dam"
	"time"
)

// 设置客户端Cookie
// 有效期1年
// 包含 [uuid], [username], [nickname]
//     [plant], [rmb_in], [rmb_out], [coin]
func SetUserInfo(w http.ResponseWriter, user *dam.User) {
	// Cookie
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)

	http.SetCookie(w, &http.Cookie{
		Name:    "uuid",
		Value:   fmt.Sprint(user.Uuid),
		Expires: expiration,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   user.Username,
		Expires: expiration,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "nickname",
		Value:   user.Nickname,
		Expires: expiration,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "plant",
		Value:   fmt.Sprint(user.Plant),
		Expires: expiration,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "rmb_in",
		Value:   fmt.Sprint(user.RmbIn),
		Expires: expiration,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "rmb_out",
		Value:   fmt.Sprint(user.RmbOut),
		Expires: expiration,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "coin",
		Value:   fmt.Sprint(user.Coin),
		Expires: expiration,
	})
}
