package dam

import (
	"log"
	"time"
)

// 向 permission_code 表添加记录
func AddPermissionCode(pc PermissionCode) bool {
	if !Connected {
		Connect()
	}

	lockPermissionCode.Lock()
	defer lockPermissionCode.Unlock()

	pc.Pcid = GetInc("permission_code") + 1
	pc.Deleted = 0

	if err := db.Table("permission_code").Create(&pc).Error; err != nil {
		log.Println(err)
		return false
	}

	UpdateInc("permission_code", pc.Pcid)

	return true
}

// 将 permission_code 表中目标记录的deleted置为当前时间戳
func DeletePermissionCode(pcid int) bool {
	if !Connected {
		Connect()
	}
	lockPermissionCode.Lock()
	defer lockPermissionCode.Unlock()

	res := db.Table("permission_code").Where("pcid=? AND deleted=0", pcid).Update("deleted", time.Now().Unix())
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if res.RowsAffected == 0 {
		return false
	}

	return true
}

// 将 permission_code 表中一组记录的deleted置为当前时间戳
func DeletePermissionCodes(pcid []int) bool {
	if !Connected {
		Connect()
	}

	lockPermissionCode.Lock()
	defer lockPermissionCode.Unlock()

	res := db.Table("permission_code").Where("pcid IN (?) AND deleted=0", pcid).Update("deleted", time.Now().Unix())
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
func UpdatePermissionCode(pc PermissionCode) bool {
	if !Connected {
		Connect()
	}
	lockPermissionCode.Lock()
	defer lockPermissionCode.Unlock()

	res := db.Table("permission_code").Where("pcid=?", pc.Pcid).Updates(map[string]interface{}{
		"code":         pc.Code,
		"permission":   pc.Permission,
		"deadline":     pc.Deadline,
		"generated_by": pc.GeneratedBy,
		"used_by":      pc.UsedBy,
		"deleted":      pc.Deleted,
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
func GetPermissionCode(key interface{}) *PermissionCode {
	if !Connected {
		Connect()
	}

	permissionCode := &PermissionCode{}
	query := ""
	switch key.(type) {
	case int, int64, int32, int16, int8:
		query = "pcid=?"
	case string:
		query = "code=?"
	default:
		return permissionCode
	}

	db.Table("permission_code").Where(query, key).First(permissionCode)

	return permissionCode
}

// 获取所有用户
func GetAllPermissionCode() []PermissionCode {
	if !Connected {
		Connect()
	}
	permissionCodes := make([]PermissionCode, 0)
	db.Table("permission_code").Find(&permissionCodes)
	return permissionCodes
}

func UsedPermissionCode(pcid int, usedBy int) bool {
	if !Connected {
		Connect()
	}
	lockPermissionCode.Lock()
	defer lockPermissionCode.Unlock()

	res := db.Table("permission_code").Where("pcid=?", pcid).Updates(map[string]interface{}{
		"used_by": usedBy,
		"deleted": time.Now().Unix(),
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
