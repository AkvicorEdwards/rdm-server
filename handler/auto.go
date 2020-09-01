package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rdm/dam"
	"rdm/def"
	"rdm/session"
	"rdm/tpl"
	"strconv"
	"time"
)

func auto(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P14, def.P16)) {
		Fprint(w, "没有查看自动任务剩余时间的权限")
		return
	}
	if r.Method == "GET" {
		if err := tpl.Auto.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func autoRemain(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.CombinePermission(def.P14, def.P16)) {
		Fprint(w, "没有查看自动任务剩余时间的权限")
		return
	}
	if r.Method == "GET" {
		uuid := session.GetUUID(r)
		if len(r.FormValue("uuid")) != 0 && CheckUserPermission(r, def.P16) {
			t, err := strconv.Atoi(r.FormValue("uuid"))
			if err != nil {
				uuid = dam.GetUser(r.FormValue("uuid")).Uuid
			} else {
				uuid = t
			}
		}
		works := dam.GetAutoWorks(uuid, false)
		var res []struct {
			Name   string `json:"name"`
			Coin   int64  `json:"coin"`
			Remain int64  `json:"remain"`
		}
		for _, i := range works {
			lastRecordListen := dam.GetLastRecord(uuid, i.Associate)
			lastRecordThis := dam.GetLastRecord(uuid, i.Wuid)
			realTime := max(lastRecordListen.TimeEnd, lastRecordThis.TimeEnd)
			realTime = max((realTime+int64(i.Unit))-time.Now().Unix(), 0)
			res = append(res, struct {
				Name   string `json:"name"`
				Coin   int64  `json:"coin"`
				Remain int64  `json:"remain"`
			}{
				Name:   i.Name,
				Coin:   i.Coin,
				Remain: realTime,
			})
		}
		resp, err := json.Marshal(res)
		if err != nil {
			Fprint(w, `{"status":"0"}`)
			return
		}
		Fprint(w, `{"status":"5","data":`+string(resp)+`}`)
	}
}
