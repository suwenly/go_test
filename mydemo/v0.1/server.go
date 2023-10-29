package main

import "go_test/inet"

func main() {
	//创建 server
	s := inet.NewServer("v0.1")

	s.Server()

}
