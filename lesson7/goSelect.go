//在一个协程里处理多个channel,需要用到Select
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1500 * time.Millisecond)
		c1 <- "1.5"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		//结果输出one,1.5; select,谁先进来，先处理谁
		case msg1 := <-c1:
			fmt.Printf("recived: %s\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("recived: %s\n", msg2)
		}
	}
}
