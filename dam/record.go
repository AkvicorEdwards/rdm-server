package dam

import (
	"encoding/json"
	"log"
)

// 向 record 表添加记录
func AddRecord(r Record) (int, bool) {
	if !Connected {
		Connect()
	}
	lockRecord.Lock()
	defer lockRecord.Unlock()

	r.Urid = GetInc("record") + 1

	if err := db.Table("record").Create(r.Transfer()).Error; err != nil {
		log.Println(err)
		return r.Urid, false
	}

	UpdateInc("record", r.Urid)

	return r.Urid, true
}

// 从 record 表中删除一条记录
func DeleteRecord(urid int) bool {
	if !Connected {
		Connect()
	}
	lockRecord.Lock()
	defer lockRecord.Unlock()

	if err := db.Table("record").Where("urid=?", urid).Delete(&RecordDam{Urid: urid}); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 从 record 表中删除一组记录
func DeleteRecords(urid []int) bool {
	if !Connected {
		Connect()
	}
	lockRecord.Lock()
	defer lockRecord.Unlock()

	if err := db.Table("record").Delete(RecordDam{}, "urid IN (?)", urid).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 更新 record 表中的一条记录
// urid作为索引，同时 urid，uuid，wuid 不会被更新
func UpdateRecord(urid, wuid int, timeStart, timeEnd, coin int64, tag []string) bool {
	if !Connected {
		Connect()
	}
	lockRecord.Lock()
	defer lockRecord.Unlock()

	ragStr, _ := json.Marshal(tag)

	res := db.Table("record").Where("urid=?", urid).Updates(map[string]interface{}{
		"wuid":       wuid,
		"time_start": timeStart,
		"time_end":   timeEnd,
		"coin":       coin,
		"tag":        string(ragStr),
	})

	if res.Error != nil {
		log.Println(res.Error)
		return false
	}

	if res.RowsAffected == 0 {
		return false
	}

	return true
}

// 完成 record 表中的一条记录
func FinishRecord(urid int, timeEnd, coin int64, tag []string) bool {
	if !Connected {
		Connect()
	}
	lockRecord.Lock()
	defer lockRecord.Unlock()

	tagStr, err := json.Marshal(tag)
	if err != nil {
		return false
	}

	res := db.Table("record").Where("urid=?", urid).Updates(map[string]interface{}{
		"time_end": timeEnd,
		"coin":     coin,
		"tag": 		string(tagStr),
	})

	if res.Error != nil {
		log.Println(res.Error)
		return false
	}

	if res.RowsAffected == 0 {
		return false
	}

	return true
}

// 从 record 表中获取一条记录
func GetRecord(urid int) *Record {
	if !Connected {
		Connect()
	}

	work := &RecordDam{}
	db.Table("record").Where("urid=?", urid).Limit(1).Find(work)

	return work.Transfer()
}

// 从 record 表中获取用户某个work的最后一条记录
func GetLastRecord(uuid, wuid int) *Record {
	if !Connected {
		Connect()
	}

	work := &RecordDam{}
	db.Table("record").Where("uuid=? AND wuid=?", uuid, wuid).Order("urid desc").Limit(1).Find(work)

	return work.Transfer()
}

// 从 record 表中获取用户某个work的最后一条记录及对应的work
func GetLastRecordWithWork(uuid, wuid int) *RecordWithWork {
	if !Connected {
		Connect()
	}

	record := GetLastRecord(uuid, wuid)
	if record.Wuid == 0 {
		return &RecordWithWork{}
	}

	work := GetWork(wuid, uuid, true)
	if work.Wuid == 0 {
		return &RecordWithWork{}
	}

	return &RecordWithWork{
		Record: *record,
		Work:   *work,
	}
}

// 从 record 表中获取用户的最后一条记录
func GetLastUserRecord(uuid int) *Record {
	if !Connected {
		Connect()
	}

	work := &RecordDam{}
	db.Table("record").Where("uuid=? AND time_end=0", uuid).Last(work)

	return work.Transfer()
}

// 从 record 表中获取用户的所有记录
func GetRecords(uuid int) []Record {
	if !Connected {
		Connect()
	}

	recordDams := make([]RecordDam, 0)
	db.Table("record").Where("uuid=?", uuid).Find(&recordDams)

	records := make([]Record, 0)
	for _, v := range recordDams {
		records = append(records, *v.Transfer())
	}

	return records
}

// 从 record 表中获取所有记录
func GetAllRecord() []Record {
	if !Connected {
		Connect()
	}

	recordDams := make([]RecordDam, 0)
	db.Table("record").Find(&recordDams)

	records := make([]Record, 0)
	for _, v := range recordDams {
		records = append(records, *v.Transfer())
	}

	return records
}

// 从 record 表中获取所有记录
func GetUserRecord(uuid int) []Record {
	if !Connected {
		Connect()
	}

	recordDams := make([]RecordDam, 0)
	db.Table("record").Where("uuid=?", uuid).Find(&recordDams)

	records := make([]Record, 0)
	for _, v := range recordDams {
		records = append(records, *v.Transfer())
	}

	return records
}

func GetUserRecordByTime(uuid int, start, end int64) []Record {
	if !Connected {
		Connect()
	}

	recordDams := make([]RecordDam, 0)
	records := make([]Record, 0)
	db.Table("record").Where("uuid=? AND ((time_start>=? AND time_start<?) "+
		"OR (time_end>? AND time_end<=?))", uuid, start, end, start, end).Find(&recordDams)

	for k, v := range recordDams {
		if v.TimeStart < start {
			recordDams[k].Coin = v.Coin * (v.TimeEnd - start + 1) / (v.TimeEnd - v.TimeStart + 1)
			recordDams[k].TimeStart = start
		}
		if v.TimeEnd > end {
			recordDams[k].Coin = v.Coin * (end - v.TimeStart + 1) / (v.TimeEnd - v.TimeStart + 1)
			recordDams[k].TimeEnd = end
		}
		records = append(records, *v.Transfer())
	}

	return records
}

func GetRecordByTime(start, end int64) []Record {
	if !Connected {
		Connect()
	}

	recordDams := make([]RecordDam, 0)
	records := make([]Record, 0)
	db.Table("record").Where(" (time_start>=? AND time_start<?) "+
		"OR (time_end>? AND time_end<=?)", start, end, start, end).Find(&recordDams)

	for k, v := range recordDams {
		if v.TimeStart < start {
			recordDams[k].Coin = v.Coin * (v.TimeEnd - start + 1) / (v.TimeEnd - v.TimeStart + 1)
			recordDams[k].TimeStart = start
		}
		if v.TimeEnd > end {
			recordDams[k].Coin = v.Coin * (end - v.TimeStart + 1) / (v.TimeEnd - v.TimeStart + 1)
			recordDams[k].TimeEnd = end
		}
		records = append(records, *v.Transfer())
	}

	return records
}
