package dam

import "log"

// 向 transaction 表中添加一条记录
func AddTransaction(r Transaction) bool {
	if !Connected {
		Connect()
	}
	lockTransaction.Lock()
	defer lockTransaction.Unlock()

	r.Utid = GetInc("transaction") + 1

	if err := db.Table("transaction").Create(&r).Error; err != nil {
		log.Println(err)
		return false
	}

	UpdateInc("transaction", r.Utid)

	return true
}

// 从 transaction 表中删除一条记录
func DeleteTransaction(utid int) bool {
	if !Connected {
		Connect()
	}
	lockTransaction.Lock()
	defer lockTransaction.Unlock()

	if err := db.Table("transaction").Where("utid=?", utid).Delete(&Transaction{Utid: utid}); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 从 transaction 表中删除一组记录
func DeleteTransactions(utid []int) bool {
	if !Connected {
		Connect()
	}
	lockTransaction.Lock()
	defer lockTransaction.Unlock()

	if err := db.Table("transaction").Delete(Transaction{}, "utid IN (?)", utid).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 更新 transaction 表中的一条记录
func UpdateTransaction(r Transaction) bool {
	if !Connected {
		Connect()
	}
	lockTransaction.Lock()
	defer lockTransaction.Unlock()

	res := db.Table("transaction").Where("utid=?", r.Utid).Updates(map[string]interface{}{
		"wuid":      r.Wuid,
		"unix_time": r.UnixTime,
		"rmb":  	 r.Rmb,
		"coin":      r.Coin,
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

// 获取 transaction 表中的一条记录
func GetTransaction(utid int) *Transaction {
	if !Connected {
		Connect()
	}

	work := &Transaction{}
	db.Table("transaction").Where("utid=?", utid).First(work)

	return work
}

// 获取 transaction 表中用户的所有记录
func GetUserTransactions(uuid int) []Transaction {
	if !Connected {
		Connect()
	}

	transactions := make([]Transaction, 0)
	db.Table("transaction").Where("uuid=?", uuid).Find(&transactions)

	return transactions
}

// 获取 transaction 表中的所有记录
func GetAllTransaction() []Transaction {
	if !Connected {
		Connect()
	}
	transactions := make([]Transaction, 0)
	db.Table("transaction").Find(&transactions)
	return transactions
}
