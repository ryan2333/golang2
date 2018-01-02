package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func mapAction(n int, ch chan int) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	ch <- n * 2
}

func reduceAction(length int, ch chan int) {
	cnt := 0
	sum := 0
	for {
		n := <-ch
		sum += n
		cnt++
		log.Printf("n: %d, sum: %d", n, sum)
		if cnt == length {
			break
		}
	}
	fmt.Println(sum)
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	ch := make(chan int)
	for _, v := range list {
		go mapAction(v, ch)
	}
	reduceAction(len(list), ch)
}
