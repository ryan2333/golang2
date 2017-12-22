package main

import "fmt"

func fibSelect(c, quit chan int) {
	x, y := 0, 1
	for { //一直向channel写入数字
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit: //当quit这个channel有数据时，则退出此协程
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ { //从channel读取10个数
			fmt.Println(<-c)
		}
		quit <- 0 //循环10次后向quit channel里写入0
	}()

	fibSelect(c, quit) //将斐波那契协程当作主协程运行
}
