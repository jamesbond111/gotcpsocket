/**
 * @note
 *
 * @author: libi
 * @date: 2020/4/1
 */

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		str = strings.Trim(str,"\n")
		if str == "exit" {
			return
		}
		n, err := conn.Write([]byte(str))
		if err != nil {
			return
		}
		fmt.Printf("客户端发送了%d个字节", n)
	}

}
