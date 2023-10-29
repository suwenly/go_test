package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println(" client start..")
	// 创建链接， 得到一个conn链接

	time.Sleep(1 * time.Second)

	dialConn, err := net.Dial("tcp4", "127.0.0.1:8999")
	if err != nil {

	}

	for {
		_, err := dialConn.Write([]byte("Hello, V0.1"))
		if err != nil {

		}
		buf := make([]byte, 512)

		//将服务端返回的数据写入 buf
		cnt, err2 := dialConn.Read(buf)
		if err2 != nil {
			fmt.Println(" read buf error")
			return
		}

		fmt.Printf("Server call back: %s, cnt=%d", buf, cnt)

		time.Sleep(1 * time.Second)

	}

}
