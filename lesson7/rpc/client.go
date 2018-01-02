//向服务端发起rpc调用，调用服务端GetLine函数，向服务端发送命令并接收命令返回的结果
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply string
		err = client.Call("Listener.GetLine", line, &reply) //调用服务端GetLine方法，并获取返回
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
	}
}
