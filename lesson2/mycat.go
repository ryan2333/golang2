package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func catFile(name []string){
	for _, fname := range name {
		content, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(content))
	}
}

func main (){
	fileName := os.Args[1:]
	if len(fileName) == 0{
		fmt.Println("Please input fileName; example mycat test.txt")
		os.Exit(0)
	}
	catFile(fileName)
}
