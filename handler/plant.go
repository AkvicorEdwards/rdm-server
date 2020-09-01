package handler

import (
	"net/http"
	"net/url"
	"rdm/def"
	"rdm/operator"
	"rdm/session"
	"strconv"
	"time"
)

// status:
//  0  Error
// 	-1 有未完成的任务
//  1  计时 结束任务成功
//  2  计时 开始任务成功
//  3  次数 完成成功
//  4  充值 充值或提现成功

func plant(w http.ResponseWriter, r *http.Request) {
	UpdateCookieSession(w, r)
	if !CheckUserPermission(r, def.P0) {
		Fprint(w, "没有使用Forest的权限")
		return
	}
	if r.Method == "GET" {
		decodeTag, _ := url.QueryUnescape(r.FormValue("tag"))
		tag := StringSplitBySpace(decodeTag)
		uuid := session.GetUUID(r)
		wuid, err := strconv.Atoi(r.FormValue("wuid"))
		if err != nil {
			Fprint(w, `{"status":"0"}`)
			return
		}
		if !operator.CheckWork(wuid, uuid) {
			Fprint(w, `{"status":"0"}`)
			return
		}

		work := operator.GetWork(wuid, uuid)

		switch work.Type {
		case def.Type1: // 充值/提现
			rmb, err := strconv.ParseInt(r.FormValue("quantity"), 10, 64)
			if err != nil {
				Fprint(w, `{"status":"0"}`)
				return
			}
			rmbIn, rmbOut, coin, res := operator.FinishBonus(uuid, wuid, rmb)
			if !res {
				Fprint(w, `{"status":"0"}`)
				return
			}
			UpdateCookieSession(w, r)
			Fprintf(w, `{"rmb_in":"%d","rmb_out":"%d","coin":"%d","status":"4"}`, rmbIn, rmbOut, coin)
			return
		case def.Type2: // 以次数为单位
			coin, ok := operator.FinishOnce(uuid, wuid, time.Now().Unix(), tag)
			if !ok {
				Fprint(w, `{"status":"0"}`)
				return
			}
			UpdateCookieSession(w, r)
			Fprintf(w, `{"coin":"%d","status":"3"}`, coin)
			return
		case def.Type3: // 以分钟为单位
			// 查询是否有未完成的任务
			lastRecord, ok := operator.HasUnfinishedRecord(uuid)
			if ok {
				if wuid == lastRecord.Wuid {
					//fmt.Println("plant1")
					timeStart, timeEnd, coin, ok := operator.FinishRecord(lastRecord.Urid, tag, uuid)
					if ok {
						UpdateCookieSession(w, r)
						// 结束任务成功
						Fprintf(w, `{"time_start":"%d","time_end":"%d","coin":"%d","status":"1"}`, timeStart, timeEnd, coin)
						return
					} else {
						// 结束任务失败
						Fprint(w, `{"status":"0"}`)
						return
					}
				} else {
					// 有未完成的任务，无法开始新任务 返回未完成任务的id
					Fprintf(w, `{"status":"-1","urid":"%d"}`, lastRecord.Urid)
					return
				}
			}
			// 开始这个任务，（返回任务开始的时间）
			startTime, ok := operator.Plant(uuid, wuid, tag)
			if !ok {
				Fprint(w, `{"status":"0"}`)
				return
			}

			UpdateCookieSession(w, r)
			Fprintf(w, `{"status":"2","time_start":"%d"}`, startTime)
		default:
			Fprint(w, `{"status":"0"}`)
			return
		}

	}
}
