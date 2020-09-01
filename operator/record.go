package operator

import (
	"encoding/json"
	"rdm/dam"
	"rdm/def"
	"strconv"
)

// 获取记录的json格式
func GetRecordJson(urid int) string {
	record := dam.GetRecord(urid)
	data, err := json.Marshal(record)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取所有记录的json格式
func GetAllRecordJson() string {
	records := dam.GetAllRecord()
	data, err := json.Marshal(records)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取所有记录的json格式
func GetAllRecordWithWorkJson() string {
	records := recordWithWork(dam.GetAllRecord())
	data, err := json.Marshal(records)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取所有记录的json格式
func GetUserRecordWithWorkJson(uuid int) string {
	records := recordWithWork(dam.GetUserRecord(uuid))
	data, err := json.Marshal(records)
	if err != nil {
		return ""
	}
	return string(data)
}

// 更新一条记录
func UpdateRecord(urid, wuid int, startTime, endTime int64, tag []string) bool {
	record := dam.GetRecord(urid)
	if !dam.UpdateUserCoin(record.Uuid, 0, 0, -1*record.Coin) {
		return false
	}
	work := dam.GetWork(wuid, record.Uuid, true)
	coin := func() int64 {
		if endTime == 0 {
			return 0
		}
		switch work.Type {
		case def.Type0, def.Type1, def.Type2:
			return work.Coin
		case def.Type3:
			return (endTime - startTime) / 60 * work.Coin
		default:
			return (endTime - startTime) / 60 * work.Coin
		}
	}()
	if !dam.UpdateUserCoin(record.Uuid, 0, 0, coin) {
		dam.UpdateUserCoin(record.Uuid, 0, 0, record.Coin)
		return false
	}

	return dam.UpdateRecord(urid, wuid, startTime, endTime, coin, tag)
}

// 删除一条记录
func DeleteRecords(urids []string) bool {
	urid := make([]int, len(urids))
	var err error
	for k, v := range urids {
		urid[k], err = strconv.Atoi(v)
		if err != nil {
			return false
		}
		record := dam.GetRecord(urid[k])
		dam.UpdateUserCoin(record.Uuid, 0, 0, -1*record.Coin)
	}
	return dam.DeleteRecords(urid)
}
