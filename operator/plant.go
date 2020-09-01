package operator

import (
	"rdm/dam"
	"rdm/def"
	"time"
)

// 检查用户是否有未完成的任务，如果有则返回该任务的id
func HasUnfinishedRecord(uuid int) (*dam.Record, bool) {
	record := dam.GetLastUserRecord(uuid)
	if record.Wuid == 0 {
		return record, false
	}
	return record, true
}

// 完成一个计时任务
func FinishRecord(urid int, tag []string, uuid int) (timeStart, timeEnd, coin int64, res bool) {
	record := dam.GetRecord(urid)
	if !def.CheckPermission(dam.GetUser(record.Uuid).Permission, def.P0) {
		return
	}
	work := dam.GetWork(record.Wuid, uuid, true)
	timeEnd = time.Now().Unix()
	coin = (timeEnd - record.TimeStart) / 60 * work.Coin
	if dam.FinishRecord(urid, timeEnd, coin, tag) {
		dam.UpdatePlant(record.Uuid, 0)
		dam.UpdateUserCoin(record.Uuid, 0, 0, coin)
		return record.TimeStart, timeEnd, coin, true
	}
	return 0, 0, 0, false
}

// 完成一个记次任务
func FinishOnce(uuid, wuid int, t int64, tag []string) (coin int64, res bool) {
	if !def.CheckPermission(dam.GetUser(uuid).Permission, def.P0) {
		return
	}
	work := dam.GetWork(wuid, uuid, true)
	if _, ok := dam.AddRecord(dam.Record{
		Uuid:      uuid,
		Wuid:      wuid,
		TimeStart: t,
		TimeEnd:   t,
		Coin:      work.Coin,
		Tag:       tag,
	}); ok {
		dam.UpdateUserCoin(uuid, 0, 0, work.Coin)
		return work.Coin, true
	}
	return 0, false
}

// 完成一次充值或提现
func FinishBonus(uuid, wuid int, rmb int64) (rmbIn, rmbOut, coin int64, res bool) {
	if !def.CheckPermission(dam.GetUser(uuid).Permission, def.P1) {
		return
	}
	t := time.Now().Unix()
	work := dam.GetWork(wuid, uuid, true)
	if work.Coin < 0 { // 提现
		rmbIn = 0
		rmbOut = abs64(rmb)
		coin = abs64(rmb)* work.Coin
	} else { // 充值
		rmbIn = abs64(rmb)
		rmbOut = 0
		coin = abs64(rmb)* work.Coin
		rmb = -rmb
	}

	if dam.AddTransaction(dam.Transaction{
		Uuid:     uuid,
		UnixTime: t,
		Rmb:	  rmb,
		Coin:     coin,
	}) {
		dam.UpdateUserCoin(uuid, rmbIn, rmbOut, coin)
		return rmbIn, rmbOut, coin, true
	}
	return 0, 0, 0, false
}

// 开始一项计时任务
func Plant(uuid, wuid int, tag []string) (startTime int64, res bool) {
	if !def.CheckPermission(dam.GetUser(uuid).Permission, def.P0) {
		return
	}
	startTime = time.Now().Unix()
	if urid, ok := dam.AddRecord(dam.Record{
		Uuid:      uuid,
		Wuid:      wuid,
		TimeStart: startTime,
		TimeEnd:   0,
		Coin:      0,
		Tag:	   tag,
	}); ok {
		dam.UpdatePlant(uuid, urid)
		return startTime, true
	}
	return 0, false
}

func CheckAutoWork(uuid int) {
	if !def.CheckPermission(dam.GetUser(uuid).Permission, def.P0) {
		return
	}
	works := dam.GetAutoWorks(uuid, false)
	for _, i := range works {
		if i.Associate == 0 {
			i.Associate = i.Wuid
		}
		lastRecordListen := dam.GetLastRecord(uuid, i.Associate)
		lastRecordThis := dam.GetLastRecord(uuid, i.Wuid)
		if lastRecordListen.Wuid == 0 && lastRecordThis.Wuid == 0 {
			FinishOnce(uuid, i.Wuid, time.Now().Unix(), []string{"auto"})
			continue
		}

		for {
			lastRecordListen = dam.GetLastRecord(uuid, i.Associate)
			lastRecordThis = dam.GetLastRecord(uuid, i.Wuid)
			lastTime := max(lastRecordListen.TimeEnd, lastRecordThis.TimeEnd) // 最后一次完成的时间
			seg := int64(i.Unit)
			if lastTime + seg - time.Now().Unix() <= 0 {
				FinishOnce(uuid, i.Wuid, lastTime + seg, []string{"auto"})
			} else {
				break
			}
		}
	}
}
