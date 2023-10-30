package inet

import (
	"fmt"
	"go_test/iface"
	"net"
)

type ZConn struct {
	Conn      *net.TCPConn //当前链接的socket套接字
	ConnID    uint32       //链接的id
	isClosed  bool         //当前的链接状态
	handleApi iface.HandleFun
	ExistChan chan bool
}

// 初始化
func NewZConn(conn *net.TCPConn, connId uint32, callback iface.HandleFun) *ZConn {
	c := &ZConn{
		Conn:      conn,
		ConnID:    connId,
		handleApi: callback,
		isClosed:  false,
		ExistChan: make(chan bool, 1),
	}
	return c
}

func (c *ZConn) StartReader() {
	fmt.Println("Reader Start().., ConnId=", c.ConnID)
	defer fmt.Println("ConnId=", c.ConnID, "Reader is exist, remote addr=", c.GetRemoteAddr())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		readCnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println(" recv buf err", err)
			continue
		}
		//调用当前链接所绑定的handleApi

		if err := c.handleApi(c.Conn, buf, readCnt); err != nil {
			fmt.Println("ConnId,", c.ConnID, "handle is err", err)
			break
		}

	}

}

func (c *ZConn) Start() {
	fmt.Println("Conn Start().., ConnId=", c.ConnID)
	//启动从当前链接的读数据 业务
	go c.StartReader()
}

func (c *ZConn) Stop() {
	fmt.Println("Conn Stop().. ConnId=", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.Conn.Close()

	close(c.ExistChan)
}

func (c *ZConn) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c *ZConn) GetRemoteAddr() *net.Addr {
	return c.c
} //获取远程客户端的tcp状态， IP 和 Port

func (c *ZConn) Send(data []byte) error {

} //发送数据， 将数据发送给远程的客户端
