package main

import (
	"fmt"
	"net"
)

func main() {

	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
	}
	//way :1
	//bs:=[]byte("hello")
	//conn.Write()

	//way:2
	conn.Write([]byte("hello sir from rafi"))

	bs := make([]byte, 1024)
	n, _ := conn.Read(bs)
	fmt.Println(string(bs)[:n])

	conn.Close()

}
