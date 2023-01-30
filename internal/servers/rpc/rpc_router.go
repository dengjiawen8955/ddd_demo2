package rpc

import (
	pb_user "dc2/internal/servers/rpc/protos/in/user"
	"dc2/internal/user"
)

func WithRouter(s *RpcServer) {
	// 新建 server
	userServer := user.NewUserServer(s.Apps.UserApp)

	// 注册路由
	pb_user.RegisterUserServer(s.srv, userServer)
}
