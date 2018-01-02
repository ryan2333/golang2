package main

import (
	"fmt"
	"time"
)

func preOrder(ch chan string) {
	var str string
	str = <-ch
	fmt.Printf("got phone: %s\n", str)
}

func main() {
	var ch chan string //定义channel
	ch = make(chan string)
	// ch = make(chan string, 1) //初始化channel,长度为1

	// var str string
	// str = <-ch //从channel读出数据，此处会报错，不能自己写自己读

	go preOrder(ch)  //启动协程
	ch <- "iphone x" //向channel写入数据，箭头一直朝左，写入阻塞
	time.Sleep(time.Second)
}
