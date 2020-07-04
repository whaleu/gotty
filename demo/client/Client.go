package main

import (
	"fmt"
	"net"
	"time"
)

// 模拟客户端
func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	// 链接远程服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start error ", err)
		return
	}

	for {
		// 链接调用 Write 方法，写数据
		_, err := conn.Write([]byte("Hello gotty"))
		if err != nil {
			fmt.Println("write conn err ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("Message from Server: %s, cnt = %d\n", buf, cnt)

		time.Sleep(5 * time.Second)

	}

}
