package main

import (
	"fmt"
	"time"
)

func say(s string, c int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func main() {
	go say("world!", 10) //大概率会打印5次，因为跑完5次，主协程就会退出
	say("hello", 5)
	// go say("world!", 10) //不打印，此处是起协程执行，协程创建完就主程序退出了
}
