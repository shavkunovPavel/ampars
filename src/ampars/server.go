package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

// получить роутер
func getRouter() (router *httprouter.Router) {
	router = httprouter.New()
	router.POST("/requests", handleRequest)
	return
}

// старт сервера
func main() {
	mux := getRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 7777),
		Handler:        mux,
		ReadTimeout:    time.Duration(10 * int64(time.Second)),
		WriteTimeout:   time.Duration(120 * int64(time.Second)),
		IdleTimeout:    time.Duration(120 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatalln(server.ListenAndServe())
}
