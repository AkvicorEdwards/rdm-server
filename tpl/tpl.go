package tpl

import (
	"fmt"
	"html/template"
	"log"
)

var err error
var errCnt = 0

var Index *template.Template
var Work *template.Template
var User *template.Template
var Record *template.Template
var Transaction *template.Template
var ToDo *template.Template
var Forest *template.Template
var Login *template.Template
var Register *template.Template
var TopUp *template.Template
var Auto *template.Template
var PermissionCode *template.Template
var Analyze *template.Template

func Parse() {
	err = nil
	errCnt = 0

	Index = addFromFile("./tpl/index.html")
	Work = addFromFile("./tpl/work.html")
	User = addFromFile("./tpl/user.html")
	Record = addFromFile("./tpl/record.html")
	Transaction = addFromFile("./tpl/transaction.html")
	ToDo = addFromFile("./tpl/todo.html")
	Forest = addFromFile("./tpl/forest.html")
	Login = addFromFile("./tpl/login.html")
	Register = addFromFile("./tpl/register.html")
	TopUp = addFromFile("./tpl/topup.html")
	Auto = addFromFile("./tpl/auto.html")
	PermissionCode = addFromFile("./tpl/permission_code.html")
	Analyze = addFromFile("./tpl/analyze.html")

	log.Printf("Parsing the html template was completed with %d errors\n", errCnt)
}

func add(name, tpl string) (t *template.Template) {
	t, err = template.New(name).Parse(tpl)
	if err != nil {
		errCnt++
		fmt.Println(err)
	}
	return
}

func addFromFile(file string) (t *template.Template) {
	t, err = template.ParseFiles(file)
	if err != nil {
		errCnt++
		fmt.Println(err)
	}
	return
}
