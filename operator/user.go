package operator

import (
	"encoding/json"
	"rdm/dam"
	"rdm/def"
	"strconv"
)

// 添加一个用户
func AddUser(username, nickname, password, permissionCode string) bool {

	if !(def.CheckUsername(username) && def.CheckPassword(password)) {
		return false
	}

	pc := dam.GetPermissionCode(permissionCode)

	if pc.Pcid == 0 || pc.Deleted != 0 || pc.UsedBy != 0 {
		return false
	}

	if uuid, ok := dam.AddUser(username, nickname, password, pc.Permission); ok {
		dam.UsedPermissionCode(pc.Pcid, uuid)
	}

	return true
}

// 删除一组用户
func DeleteUsers(uuids []string) bool {
	uuid := make([]int, len(uuids))
	var err error
	for k, v := range uuids {
		uuid[k], err = strconv.Atoi(v)
		if err != nil {
			return false
		}
	}
	return dam.DeleteUsers(uuid)
}

// 更新一条用户记录
func UpdateUser(uuid int, nickname, password string, permission, rmbIn, rmbOut, coin, deleted int64) bool {
	if !def.CheckPassword(password) {
		return false
	}
	return dam.UpdateUser(dam.User{
		Uuid:     uuid,
		Nickname: nickname,
		Password: password,
		Permission: permission,
		RmbIn:    rmbIn,
		RmbOut:   rmbOut,
		Coin:     coin,
		Deleted:  deleted,
	})
}

// 获取所有用户记录的json格式
func GetAllUserJson() string {
	users := dam.GetAllUser()
	data, err := json.Marshal(users)
	if err != nil {
		return ""
	}
	return string(data)
}

// 登录
func Login(username, password string) (*dam.User, bool) {
	user := dam.GetUser(username)
	if user.Deleted != 0 || user.Uuid <= 0 || user.Password != password {
		return &dam.User{}, false
	}
	return user, true
}

// 获取一条用户记录
func GetUser(uuid int) (*dam.User, bool) {
	CheckAutoWork(uuid)
	user := dam.GetUser(uuid)
	if user.Deleted != 0 || user.Uuid <= 0 {
		return &dam.User{}, false
	}
	return user, true
}
