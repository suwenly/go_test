package inet

import (
	"fmt"
	"go_test/iface"
	"net"
)

type Server struct {
	Name     string
	IPServer string
	IP       string
	Port     int
}

func (s *Server) Start() {
	fmt.Println("Start Server Listenner")

	go func() {
		// 获取一个 tcp addr
		addr, err := net.ResolveTCPAddr(s.IPServer, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {

		}
		//监听服务器地址
		tcpListen, err := net.ListenTCP(s.IPServer, addr)
		if err != nil {

		}
		fmt.Println("Start Server succ, Listenning..")

		// 阻塞等待客户端链接，
		for {
			//如果有客户端链接过来，阻塞会返回
			tcpConn, err := tcpListen.AcceptTCP()
			if err != nil {

			}
			go func() {
				for {
					buf := make([]byte, 512)
					readCnt, err2 := tcpConn.Read(buf)
					if err2 != nil {
						fmt.Println(" recv buf err", err2)
						continue
					}

					//回复功能
					if _, err3 := tcpConn.Write(buf[:readCnt]); err3 != nil {
						fmt.Println(" write back buf err", err3)
						continue
					}
				}
			}()
		}

	}()

}

func (s *Server) Stop() {

}

func (s *Server) Server() {
	s.Start()

	//阻塞状态
	select {}
}

func NewServer(name string) iface.IServer {
	s := &Server{
		Name:     name,
		IPServer: "tcp4",
		IP:       "127.0.0.1",
		Port:     8999,
	}
	return s
}
