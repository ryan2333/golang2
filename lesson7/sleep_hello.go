//使用channel同步
package main

import (
	"fmt"
	"time"
)

func main() {
	s := "的hello"
	first := make(chan int)
	in := first
	for i, n := range s {
		out := make(chan int)
		go func(i int, n rune, in chan int, out chan int) {
			in_num := <-in
			fmt.Printf("in_num: %v, letter: %s\n", in_num, string(n))
			out <- i
		}(i, n, in, out)
		in = out
	}
	fmt.Printf("Begin to first: %d\n", -1)
	first <- -1
	time.Sleep(3 * time.Second)
}
