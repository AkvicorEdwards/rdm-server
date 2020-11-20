package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"rdm/def"
	"rdm/session"
	"regexp"
)

type str2func map[string]func(http.ResponseWriter, *http.Request)

var public str2func

func ParsePrefix() {
	public = make(str2func)

	// 界面
	public["/"] = index
	public["/login"] = login
	public["/register"] = register
	public["/todo"] = todo
	public["/forest"] = forest
	public["/topup"] = topUp
	public["/work"] = work
	public["/user"] = user
	public["/record"] = record
	public["/transaction"] = transaction
	public["/auto"] = auto
	public["/permission_code"] = permissionCode
	public["/analyze"] = analyze

	// 用户 api
	public["/add_user"] = addUser
	public["/delete_users"] = deleteUsers
	public["/update_user"] = updateUser
	public["/user_list"] = getUserList

	// 项 api
	public["/add_work"] = addWork
	public["/get_work"] = getWork
	public["/delete_works"] = deleteWorks
	public["/update_work"] = updateWork
	public["/work_list"] = getWorkList
	public["/work_time"] = getWorkTime
	public["/work_once"] = getWorkOnce
	public["/work_custom"] = getWorkCustom
	public["/work_auto"] = getWorkAuto
	public["/work_todo"] = getWorkToDo
	public["/get_user_auto"] = getUserWorkAuto
	public["/auto_remain"] = autoRemain

	// 普通记录 api
	public["/get_record"] = getRecord
	public["/update_record"] = updateRecord
	public["/delete_records"] = deleteRecords
	public["/record_list"] = getRecordList

	// 交易记录 api
	public["/update_transaction"] = updateTransaction
	public["/delete_transactions"] = deleteTransactions
	public["/transaction_list"] = getTransactionList

	// 任务的开始/完成 api
	public["/plant"] = plant

	// 注册码 api
	public["/add_permission_code"] = addPermissionCode
	public["/delete_permission_codes"] = deletePermissionCodes
	public["/update_permission_code"] = updatePermissionCode
	public["/permission_code_list"] = getPermissionCodeList

	public["/get_analyze"] = getAnalyze
}

type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/register" {
		public["/register"](w, r)
		return
	}

	if session.GetUUID(r) <= 0 {
		public["/login"](w, r)
		return
	}

	if h, ok := public[r.URL.Path]; ok {
		h(w, r)
		return
	}

	if ok, _ := regexp.MatchString("/favicon.ico", r.URL.Path); ok {
		download(w, "./favicon.ico")
	}

}

func CheckUserPermission(r *http.Request, permission int64) bool {
	p := session.GetPermission(r)
	return def.CheckPermission(p, permission)
}

func download(w http.ResponseWriter, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		_, _ = fmt.Fprintln(w, "File Not Found")
		return
	}
	defer func() { _ = file.Close() }()
	data := make([]byte, 1024)
	for {
		n, err1 := file.Read(data)
		if err1 != nil && err1 != io.EOF {
			_, _ = fmt.Fprintln(w, "File Read Error")
			return
		}
		nn, err2 := w.Write(data[:n])
		if err2 != nil || nn != n {
			_, _ = fmt.Fprintln(w, "File Write Error")
			return
		}
		if err1 == io.EOF {
			return
		}
	}
}

func Fprint(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprint(w, a...)
}

func Fprintf(w http.ResponseWriter, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, format, a...)
}

func Fprintln(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprintln(w, a...)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
