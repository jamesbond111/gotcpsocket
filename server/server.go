/**
 * @note
 *
 * @author: libi
 * @date: 2020/4/1
 */

package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		fmt.Println("等待客户端发送消息...")
		buf := make([]byte, 2014)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Printf("读取到了%d个字节\n", n)
		fmt.Print(string(buf[:n]))
	}

}

func main() {
	listener, e := net.Listen("tcp", "0.0.0.0:8888")
	if e != nil {
		return
	}
	defer listener.Close()
	fmt.Println("监听客户端的连接...")
	for {
		conn, e := listener.Accept()
		if e != nil {
			return
		}
		go process(conn)
	}

}
