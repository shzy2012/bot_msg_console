package main

import (
	"bot_msg_example/service"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var port = flag.String("p", "9010", "http port")

//staticHandler static file server
func staticHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(*r.URL)
		start := time.Now()
		h.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr
		ua := r.UserAgent()
		log.Printf(
			"%s\t%s\t%s\t%s\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			ua,
			time.Since(start),
		)
	})
}

func main() {

	flag.Parse()

	hub := service.NewHub()
	go hub.Run()

	//设置http static
	http.Handle("/", staticHandler(http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		service.ServeWs(hub, w, r)
	})

	//设置bot消息路由/消息投递
	http.HandleFunc("/bot/", func(w http.ResponseWriter, r *http.Request) {
		service.ServeBotMsg(hub, w, r)
	})

	//开启http服务
	addr := fmt.Sprintf(":%s", *port)
	log.Printf("listen on: %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
