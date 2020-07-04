package net

// Server IServer接口实现，定义一个 Server 服务器模块
type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的 ip 版本
	IPVersion string
	// 服务器监听的 ip
	IP string
	// 服务器监听的端口
	Port string
}

// Start 实现接口中的启动服务方法
func (s *Server) Start() {

}

// Stop 实现接口中的停止服务方法
func (s *Server) Stop() {

}

// Server 实现接口中的运行服务方法
func (s *Server) Server() {

}
