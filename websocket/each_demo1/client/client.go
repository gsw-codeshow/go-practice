package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9090")
	if nil != err {
		fmt.Println(err)
		return
	}
	conn, err := net.DialTCP(
		"tcp",
		nil,
		tcpAddr)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println("connected server!")
	defer conn.Close()

	go onMessageRecived(conn)

	b := []byte("time\n")
	conn.Write(b)
	<-quitSemaphore
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if nil != err {
			quitSemaphore <- true
			break
		}
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write(b)
	}
}
