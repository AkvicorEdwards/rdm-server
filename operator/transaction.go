package operator

import (
	"encoding/json"
	"rdm/dam"
	"strconv"
)

// 获取所有交易记录的json格式
func GetAllTransactionJson() string {
	transaction := dam.GetAllTransaction()
	data, err := json.Marshal(transaction)
	if err != nil {
		return ""
	}
	return string(data)
}

// 获取用户所有交易记录的json格式
func GetUserTransactionJson(uuid int) string {
	transaction := dam.GetUserTransactions(uuid)
	data, err := json.Marshal(transaction)
	if err != nil {
		return ""
	}
	return string(data)
}

func GetTransaction(utid int) *dam.Transaction {
	return dam.GetTransaction(utid)
}

// 更新一条交易记录
func UpdateTransaction(utid, wuid int, unixTime int64, rmb int64) bool {
	transaction := dam.GetTransaction(utid)
	work := dam.GetWork(wuid, transaction.Uuid, true)
	// 回滚此记录的影响
	if transaction.Rmb < 0 { // 原记录为提现
		if !dam.UpdateUserCoin(transaction.Uuid, 0, -transaction.Rmb, -transaction.Coin) {
			return false
		}
	} else { // 原记录为充值
		if !dam.UpdateUserCoin(transaction.Uuid, -transaction.Rmb, 0, -transaction.Coin) {
			return false
		}
	}
	var rmbIn int64 = 0
	var rmbOut int64 = 0
	var coin int64 = 0
	if rmb < 0 { // 修改后的记录为提现
		rmbIn = 0
		rmbOut = abs64(rmb)
		coin = abs64(rmb)* work.Coin
	} else { // 修改后的记录为充值
		rmbIn = abs64(rmb)
		rmbOut = 0
		coin = abs64(rmb)* work.Coin
	}

	if !dam.UpdateTransaction(dam.Transaction{
		Utid:     utid,
		Wuid:     wuid,
		UnixTime: unixTime,
		Rmb:	  rmb,
		Coin:     coin,
	}) {
		// 取消回滚
		if transaction.Rmb < 0 { // 原记录为提现
			dam.UpdateUserCoin(transaction.Uuid, 0, transaction.Rmb, -transaction.Coin)
		} else { // 原记录为充值
			dam.UpdateUserCoin(transaction.Uuid, transaction.Rmb, 0, -transaction.Coin)
		}
		return false
	}
	return dam.UpdateUserCoin(transaction.Uuid, rmbIn, rmbOut, coin)
}

// 删除一组交易记录
func DeleteTransactions(utids []string) bool {
	utid := make([]int, len(utids))
	var err error
	for k, v := range utids {
		utid[k], err = strconv.Atoi(v)
		if err != nil {
			return false
		}
		transaction := dam.GetTransaction(utid[k])
		// 回滚此记录的影响
		if transaction.Rmb < 0 { // 原记录为提现
			if !dam.UpdateUserCoin(transaction.Uuid, 0, -transaction.Rmb, -transaction.Coin) {
				return false
			}
		} else { // 原记录为充值
			if !dam.UpdateUserCoin(transaction.Uuid, -transaction.Rmb, 0, -transaction.Coin) {
				return false
			}
		}
	}
	return dam.DeleteTransactions(utid)
}
