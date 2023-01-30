package main

import (
	"dc2/config"
	"dc2/internal/common/logs"
	"dc2/internal/common/signals"
	"dc2/internal/servers"
	"dc2/internal/servers/rpc"
	"dc2/internal/servers/web"
)

func main() {
	// 初始化配置
	cfg := config.NewConfig("./config.yaml")

	// 初始化日志
	logs.Init(cfg.Log)

	// 获取 servers, 比如 WebServer, RpcServer
	servers := NewServers(cfg)

	// 启动 servers
	servers.AsyncStart()

	// 优雅退出
	signals.WaitWith(servers.Stop)
}

// NewServers 通过配置文件初始化 Repo 依赖, 然后初始化 App, 最后组装为 Server
// 比如 UserRepo -> UserApp -> WebServer
func NewServers(cfg *config.SugaredConfig) servers.ServerInterface {
	repos := servers.NewRepos(cfg)
	apps := servers.NewApps(repos)

	servers := servers.NewServers()
	servers.AddServer(web.NewWebServer(cfg, apps))
	servers.AddServer(rpc.NewRpcServer(cfg, apps))

	return servers
}
