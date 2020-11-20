package dam

import (
	"log"
	"rdm/def"
	"time"
)

// 向 work 表添加记录
func AddWork(work Work) bool {
	if !Connected {
		Connect()
	}
	lockWork.Lock()
	defer lockWork.Unlock()

	work.Wuid = GetInc("work") + 1
	work.Deleted = 0

	if err := db.Table("work").Create(&work).Error; err != nil {
		log.Println(err)
		return false
	}

	UpdateInc("work", work.Wuid)

	return true
}

// 将 work 表中目标记录的deleted置为当前时间戳
// key 为 wuid 或 name
func DeleteWork(key interface{}, uuid int) bool {
	if !Connected {
		Connect()
	}
	lockWork.Lock()
	defer lockWork.Unlock()

	query := ""
	switch key.(type) {
	case int, int64, int32, int16, int8:
		query = "wuid=? AND uuid=? AND deleted=0"
	case string:
		query = "name=? AND uuid=? AND deleted=0"
	default:
		return false
	}

	res := db.Table("work").Where(query, key, uuid).Update("deleted", time.Now().Unix())
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

// 将 work 表中一组记录的deleted置为当前时间戳
func DeleteWorks(wuids []int) bool {
	if !Connected {
		Connect()
	}

	lockWork.Lock()
	defer lockWork.Unlock()

	res := db.Table("work").Where("wuid IN (?) AND deleted=0", wuids).Update("deleted", time.Now().Unix())
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

// 更新 work 表中的记录
// wuid 作为 key 不会被更新
// uuid 不会被更新
func UpdateWork(work Work) bool {
	if !Connected {
		Connect()
	}
	lockWork.Lock()
	defer lockWork.Unlock()

	res := db.Table("work").Where("wuid=?", work.Wuid).Updates(map[string]interface{}{
		"name":      work.Name,
		"type":      work.Type,
		"unit":      work.Unit,
		"associate": work.Associate,
		"coin":      work.Coin,
		"deleted":   work.Deleted,
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

// 从 work 表中获取一条记录
// 如果all为true，则已删除的记录将包含在内
func GetWork(key interface{}, uuid int, all bool) *Work {
	if !Connected {
		Connect()
	}

	work := &Work{}
	query := ""
	switch key.(type) {
	case int, int64, int32, int16, int8:
		query = "wuid=? AND uuid=?"
	case string:
		query = "name=? AND uuid=?"
	default:
		return work
	}

	if all == false {
		query += " AND deleted=0"
	}

	db.Table("work").Where(query, key, uuid).First(work)

	return work
}

// 获取 work 表的所有记录
func GetAllWorks(uuid int, all bool) []Work {
	if !Connected {
		Connect()
	}
	works := make([]Work, 0)
	if all {
		db.Table("work").Where("uuid=?", uuid).Order("wuid").Find(&works)
	} else {
		db.Table("work").Where("uuid=? AND deleted=0", uuid).Order("wuid").Find(&works)
	}

	return works
}

func GetAllWorksNoToDo(uuid int, all bool) []Work {
	if !Connected {
		Connect()
	}
	works := make([]Work, 0)
	if all {
		db.Table("work").Where("uuid=? AND type=?", uuid, def.Type4).Order("wuid").Find(&works)
	} else {
		db.Table("work").Where("uuid=? AND type=? AND deleted=0", uuid, def.Type4).Order("wuid").Find(&works)
	}

	return works
}

// 获取 work 表中 按时间 计算收益的记录
func GetTimeWorks(uuid int, all bool) []Work {
	if !Connected {
		Connect()
	}
	works := make([]Work, 0)
	if all {
		db.Table("work").Where("type=? AND uuid=?", def.Type3, uuid).Order("wuid").Find(&works)
	} else {
		db.Table("work").Where("type=? AND uuid=? AND deleted=0", def.Type3, uuid).Order("wuid").Find(&works)
	}
	return works
}

// 获取 work 表中 按次数 计算收益的记录
func GetOnceWorks(uuid int, all bool) []Work {
	if !Connected {
		Connect()
	}
	works := make([]Work, 0)
	if all {
		db.Table("work").Where("type=? AND uuid=?", def.Type2, uuid).Order("wuid").Find(&works)
	} else {
		db.Table("work").Where("type=? AND uuid=? AND deleted=0", def.Type2, uuid).Order("wuid").Find(&works)
	}

	return works
}

// 获取 work 表中 自定义数量 计算收益的记录
func GetCustomWorks(uuid int, all bool) []Work {
	if !Connected {
		Connect()
	}
	works := make([]Work, 0)
	if all {
		db.Table("work").Where("type=? AND uuid=?", def.Type1, uuid).Order("wuid").Find(&works)
	} else {
		db.Table("work").Where("type=? AND uuid=? AND deleted=0", def.Type1, uuid).Order("wuid").Find(&works)
	}
	return works
}

// 获取 work 表中 自动 计算收益的记录
func GetAutoWorks(uuid int, all bool) []Work {
	if !Connected {
		Connect()
	}
	works := make([]Work, 0)
	if all {
		db.Table("work").Where("type=? AND uuid=?", def.Type0, uuid).Order("wuid").Find(&works)
	} else {
		db.Table("work").Where("type=? AND uuid=? AND deleted=0", def.Type0, uuid).Order("wuid").Find(&works)
	}
	return works
}

func GetToDoWorks(uuid int, all bool) []Work {
	if !Connected {
		Connect()
	}
	works := make([]Work, 0)
	if all {
		db.Table("work").Where("type=? AND uuid=?", def.Type4, uuid).Order("wuid").Find(&works)
	} else {
		db.Table("work").Where("type=? AND uuid=? AND deleted=0", def.Type4, uuid).Order("wuid").Find(&works)
	}
	return works
}