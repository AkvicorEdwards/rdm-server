package dam

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// 向 user 表添加记录
func AddUser(username, nickname, password string, permission int64) (int, bool) {
	if !Connected {
		Connect()
	}
	lockUser.Lock()
	defer lockUser.Unlock()

	user := User{
		Uuid:       GetInc("user") + 1,
		Username:   username,
		Nickname:   nickname,
		Password:   password,
		Permission: permission,
	}

	if err := db.Table("user").Create(&user).Error; err != nil {
		log.Println(err)
		return 0, false
	}

	UpdateInc("user", user.Uuid)

	return user.Uuid, true
}

// 将 user 表中目标记录的deleted置为当前时间戳
func DeleteUser(u interface{}) bool {
	if !Connected {
		Connect()
	}
	lockUser.Lock()
	defer lockUser.Unlock()

	query := ""
	switch u.(type) {
	case int, int64, int32, int16, int8:
		query = "uuid=? AND deleted=0"
	case string:
		query = "username=? AND deleted=0"
	default:
		return false
	}

	res := db.Table("user").Where(query, u).Update("deleted", time.Now().Unix())
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

// 删除 user 表中的一条记录
func DeleteUsers(uuid []int) bool {
	if !Connected {
		Connect()
	}

	lockUser.Lock()
	defer lockUser.Unlock()

	res := db.Table("user").Where("uuid IN (?) AND deleted=0", uuid).Update("deleted", time.Now().Unix())
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if res.RowsAffected == 0 {
		return false
	}
	return true

}

// 更新 user 表中的一条记录
// uuid和username不会被更新
func UpdateUser(user User) bool {
	if !Connected {
		Connect()
	}
	lockUser.Lock()
	defer lockUser.Unlock()

	res := db.Table("user").Where("uuid=?", user.Uuid).Updates(map[string]interface{}{
		"nickname":   user.Nickname,
		"password":   user.Password,
		"permission": user.Permission,
		"plant":      user.Plant,
		"rmb_in":     user.RmbIn,
		"rmb_out":    user.RmbOut,
		"coin":       user.Coin,
		"deleted":    user.Deleted,
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

// 更新用户的密码
func UpdatePassword(u interface{}, password, newPassword string) bool {
	if !Connected {
		Connect()
	}
	lockUser.Lock()
	defer lockUser.Unlock()

	query := ""
	switch u.(type) {
	case int, int64, int32, int16, int8:
		query = "uuid=? AND password=?"
	case string:
		query = "username=? AND password=?"
	default:
		return false
	}

	res := db.Table("user").Where(query, u, password).Update("password", newPassword)

	if res.Error != nil {
		log.Println(res.Error)
		return false
	}

	if res.RowsAffected == 0 {
		return false
	}

	return true
}

// 从 user 表中获取一条记录
func GetUser(u interface{}) *User {
	if !Connected {
		Connect()
	}
	user := &User{}
	query := ""
	switch u.(type) {
	case int, int64, int32, int16, int8:
		query = "uuid=?"
	case string:
		query = "username=?"
	default:
		return user
	}
	db.Table("user").Where(query, u).First(user)

	return user
}

// 更新用户的rmb和coin
func UpdateUserCoin(uuid int, rmbIn, rmbOut, coin int64) bool {
	if !Connected {
		Connect()
	}
	lockUser.Lock()
	defer lockUser.Unlock()

	res := db.Table("user").Where("uuid=?", uuid).UpdateColumns(map[string]interface{}{
		"rmb_in":  gorm.Expr("rmb_in+?", rmbIn),
		"rmb_out": gorm.Expr("rmb_out+?", rmbOut),
		"coin":    gorm.Expr("coin+?", coin),
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

// 获取所有用户
func GetAllUser() []User {
	if !Connected {
		Connect()
	}
	users := make([]User, 0)
	db.Table("user").Find(&users)
	return users
}

// 更新用户的plant
func UpdatePlant(uuid, wuid int) bool {
	if !Connected {
		Connect()
	}

	lockUser.Lock()
	defer lockUser.Unlock()

	res := db.Table("user").Where("uuid=?", uuid).Update("plant", wuid)

	if res.Error != nil {
		log.Println(res.Error)
		return false
	}

	if res.RowsAffected == 0 {
		return false
	}

	return true
}
