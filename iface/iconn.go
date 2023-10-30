package iface

import "net"

type IConn interface {
	Start() //启动链接

	Stop() //停止链接， 结束当前链接的工作

	GetTCPConn() *net.TCPConn //获取当前链接模块的链接Id

	GetRemoteAddr() *net.Addr //获取远程客户端的tcp状态， IP 和 Port

	Send(data []byte) error //发送数据， 将数据发送给远程的客户端
}

type HandleFun func(*net.TCPConn, []byte, int) error
