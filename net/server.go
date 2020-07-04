package net

import (
	"fmt"
	"gotty/intfc"
	"net"
)

// Server IServer接口实现，定义一个 Server 服务器模块
type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的 ip 版本
	IPVersion string
	// 服务器监听的 ip
	IP string
	// 服务器监听的端口
	Port int
}

// Start 实现接口中的启动服务方法
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listerner at IP: %s, Port: %d, is starting\n", s.IP, s.Port)
	// 获取一个 TCP 的 Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error: ", err)
		return
	}
	// 监听服务器的地址
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("Listen ", s.IPVersion, " error", err)
		return
	}
	// 阻塞的等待客户端链接，处理客户端链接业务
	fmt.Println("start Gotty server successful, ", s.Name, " succ, Listernning...")
	for {
		// 如果有客户端链接，阻塞会返回
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept error", err)
			continue
		}
		// 与客户端成功建立链接
		go func() {
			for {
				buf := make([]byte, 512)
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("recv buf err", err)
					continue
				}
				fmt.Printf("receive message from client buf %s, cnt: %d\n", buf, cnt)
				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Println("Write back buf err ", err)
					continue
				}

			}
		}()
	}
}

// Stop 实现接口中的停止服务方法
func (s *Server) Stop() {

}

// Serve 实现接口中的运行服务方法
func (s *Server) Serve() {
	// 启动 server 的服务功能
	s.Start()

	// 阻塞
	select {}
}

// NewServer 能椒
func NewServer(name string) intfc.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
