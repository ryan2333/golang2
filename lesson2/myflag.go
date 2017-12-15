package main

import (
	"flag"
	"fmt"
)

func main(){
	linef := flag.Bool("n", false, "line")
	file := flag.String("f", "", "file name")
	num := flag.Int("N", 10, "line number")
	flag.Parse()
	fmt.Println("file", *file)
	fmt.Println("num", *num)
	fmt.Println("linef",*linef)
}
