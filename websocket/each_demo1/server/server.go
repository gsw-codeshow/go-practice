package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr
	var err error
	tcpAddr, err = net.ResolveTCPAddr("tcp", "127.0.0.1:9090")
	if nil != err {
		fmt.Println(err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if nil != err {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected : " + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if nil != err {
			return
		}
		fmt.Println(string(message))
		msg := time.Now().String() + "\n"
		msgByte := []byte(msg)
		conn.Write(msgByte)
	}
}
