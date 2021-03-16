package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//net short form of network
	nl, err := net.Listen("tcp", ":8888") //ip:port 127.0.0.1                                                                                       ")//ip:port
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1
	}

	conn, err := nl.Accept() //layer-5 session layer
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1
	}
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr)
	//byte
	conn.Write([]byte("welcome to success the code "))
	conn.Close()
	nl.Close()

}
