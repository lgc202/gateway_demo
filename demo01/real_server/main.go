package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rs1 := &RealServer{Addr: "127.0.0.1:2003"}
	rs1.Run()
	rs2 := &RealServer{Addr: "127.0.0.1:2004"}
	rs2.Run()

	//监听关闭信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

type RealServer struct {
	Addr string
}

func (r *RealServer) Run() {
	log.Println("Starting httpserver at " + r.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.HelloHandler)
	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request) {
	// 127.0.0.1:8008/abc?sdsdsa=11
	// r.Addr=127.0.0.1:8008
	// req.URL.Path=/abc
	// fmt.Println(req.Host)
	upath := fmt.Sprintf("http://%s%s\n", r.Addr, req.URL.Path)
	io.WriteString(w, upath)
}
