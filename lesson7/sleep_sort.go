package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{2, 7, 4, 1}
	for _, n := range s {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond) //只能使用秒做单位，使用其它单位会因sleep时间过短出现排序不正确的顺序
			fmt.Println(n)
		}(n) //n 传参数
	}
	time.Sleep(1 * time.Second)
}
