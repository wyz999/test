package main

import (
	"io"
	"log"
	"net"
	"time"
)

//示例: 并发的Clock服务
func handConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println(err)
			//continue
			return
		}
		time.Sleep(time.Second * 1)
	}
}
func conn1() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			//continue
			return
		}
		handConn(conn)
	}
}
func main() {
	conn1()
}
