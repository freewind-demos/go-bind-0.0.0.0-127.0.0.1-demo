package main

import (
	"net"
	"fmt"
	"time"
)

func main() {

	go bind("0.0.0.0:9999", "0.0.0.0 ")
	go func() {
		time.Sleep(2 * time.Second)
		bind("127.0.0.1:9999", "127.0.0.1 ")
	}()

	time.Sleep(time.Hour)
}

func bind(address string, content string) {
	fmt.Println("-------------", address, "-----------------")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(listener.Addr().String())

	conn, _ := listener.Accept()
	for {
		_, err := conn.Write([]byte(content))
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
