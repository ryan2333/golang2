package main

import (
	"fmt"
	"time"
)

func main() {
	s := []byte("hello\n")
	for i := 0; i < len(s); i++ {
		go func(i int) {
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("%s", string(s[i]))
		}(i)
	}
	fmt.Println("\n")
	time.Sleep(10 * time.Second)
}
