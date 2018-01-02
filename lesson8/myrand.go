package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //生成随机种子，以unix时间
	for i := 0; i < 5; i++ {
		log.Println(rand.Intn(10)) //打印10以为5个随机数
	}
	fmt.Println(time.Now().Unix())
}
