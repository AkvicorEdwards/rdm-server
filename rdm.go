package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rdm/handler"
	"rdm/maintenance"
	"rdm/tpl"
	"time"
)

const port = 8080

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "init" {
		maintenance.InitDatabase()
	}

	go maintenance.ShutDownListener()
	tpl.Parse()
	handler.ParsePrefix()

	addr := fmt.Sprintf(":%d", port)
	server := http.Server{
		Addr:              addr,
		Handler:           &handler.MyHandler{},
		ReadTimeout:       20 * time.Minute,
	}
	log.Printf("http://127.0.0.1%s", addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}