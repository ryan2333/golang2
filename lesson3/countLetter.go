package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	sum := 0
	content, _ := ioutil.ReadAll(os.Stdin)
	for _, v := range content {
		if string(v) == "\n" {
			sum += 1
		}
	}
	// fmt.Println(len(content))
	fmt.Println(sum)
}
