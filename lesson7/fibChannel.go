package main

import "fmt"

func fibChannel(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	c := make(chan int, 2)
	go fibChannel(20, c)
	for i := range c {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")
}
