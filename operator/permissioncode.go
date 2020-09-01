package operator

import (
	"encoding/json"
	"rdm/dam"
	"strconv"
)

// 添加一个注册码
func AddPermissionCode(code string, permission, deadline int64, generatedBy int) bool {
	return dam.AddPermissionCode(dam.PermissionCode{
		Code:        code,
		Permission:  permission,
		Deadline:    deadline,
		GeneratedBy: generatedBy,
	})
}

// 删除一组注册码
func DeletePermissionCodes(pcids []string) bool {
	pcid := make([]int, len(pcids))
	var err error
	for k, v := range pcids {
		pcid[k], err = strconv.Atoi(v)
		if err != nil {
			return false
		}
	}
	return dam.DeletePermissionCodes(pcid)
}

// 更新一条注册码
func UpdatePermissionCode(pcid, generatedBy, usedBy int, code string, permission, deadline, deleted int64) bool {
	return dam.UpdatePermissionCode(dam.PermissionCode{
		Pcid:        pcid,
		Code:        code,
		Permission:  permission,
		Deadline:    deadline,
		GeneratedBy: generatedBy,
		UsedBy:      usedBy,
		Deleted:     deleted,
	})
}

// 获取所有注册码的json格式
func GetAllPermissionCodeJson() string {
	permissionCodes := dam.GetAllPermissionCode()
	data, err := json.Marshal(permissionCodes)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取一个注册码
func GetPermissionCode(pcid int) (*dam.PermissionCode, bool) {
	permissionCode := dam.GetPermissionCode(pcid)
	if permissionCode.Deleted != 0 || permissionCode.Deadline != 0 || permissionCode.UsedBy != 0 {
		return &dam.PermissionCode{}, false
	}
	return permissionCode, true
}
