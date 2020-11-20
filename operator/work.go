package operator

import (
	"encoding/json"
	"rdm/dam"
	"strconv"
)

// 添加Work
func AddWork(uuid int, name string, wType, wUnit int, coin int64, associate int) bool {
	return dam.AddWork(dam.Work{
		Uuid:      uuid,
		Name:      name,
		Type:      wType,
		Unit:      wUnit,
		Coin:      coin,
		Associate: associate,
	})
}

// 检查work是否有效
func CheckWork(key interface{}, uuid int) bool {
	work := dam.GetWork(key, uuid, true)
	return work.Wuid > 0
}

// 获取work
func GetWork(key interface{}, uuid int) *dam.Work {
	return dam.GetWork(key, uuid, true)
}

// 获取json格式的Work
func GetWorkJson(wuid int, uuid int) string {
	work := dam.GetWork(wuid, uuid, true)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}

// 批量删除Work
func DeleteWorks(wuids []string) bool {
	tWuid := make([]int, len(wuids))
	var err error
	for k, v := range wuids {
		tWuid[k], err = strconv.Atoi(v)
		if err != nil {
			return false
		}
	}
	return dam.DeleteWorks(tWuid)
}

// 更新Work
func UpdateWork(wuid int, name string, wType, wUnit int, coin int64, associate int, deleted int64) bool {
	return dam.UpdateWork(dam.Work{
		Wuid:      wuid,
		Name:      name,
		Type:      wType,
		Unit:      wUnit,
		Coin:      coin,
		Associate: associate,
		Deleted:   deleted,
	})
}

// 获取json格式的所有Work
func GetAllWorkJson(uuid int) string {
	work := dam.GetAllWorks(uuid, true)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}
func GetAllWorkJsonNoDeleted(uuid int) string {
	work := dam.GetAllWorks(uuid, false)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}
func GetAllWorkJsonNoDeletedNoToDo(uuid int) string {
	work := dam.GetAllWorksNoToDo(uuid, false)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取json格式的计时类型的Work
func GetTimeWorkJson(uuid int) string {
	work := dam.GetTimeWorks(uuid, true)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取json格式的记次类型的Work
func GetOnceWorkJson(uuid int) string {
	work := dam.GetOnceWorks(uuid, true)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取json格式的自定义类型的Work
func GetCustomWorkJson(uuid int) string {
	work := dam.GetCustomWorks(uuid, true)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取json格式的自动类型的Work
func GetAutoWorkJson(uuid int) string {
	work := dam.GetAutoWorks(uuid, true)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}

func GetAllToDoWorkJson(uuid int) string {
	work := dam.GetToDoWorks(uuid, true)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}

func GetToDoWorkJson(uuid int) string {
	work := dam.GetToDoWorks(uuid, false)
	data, err := json.Marshal(work)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取json格式用户自动类型的Work
func GetUserAutoWorkJson(uuid int) string {
	work := dam.GetAutoWorks(uuid, true)
	res := make([]*dam.RecordWithWork, 0)
	for _, v := range work {
		lr := dam.GetLastRecordWithWork(uuid, v.Wuid)
		if lr.Work.Wuid == 0 {
			continue
		}
		res = append(res, lr)
	}
	data, err := json.Marshal(res)
	if err != nil {
		return ""
	}
	return string(data)
}
