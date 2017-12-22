package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./imgs"))) //监听文件服务，根目录在http服务的当前目录
	log.Fatal(http.ListenAndServe(os.Args[1], nil))       //加参数启动http服务到指定端口 :8888
}
