package main

import (
	"gotty/net"
)

func main() {
	s := net.NewServer("MyServer")
	s.Serve()
}
