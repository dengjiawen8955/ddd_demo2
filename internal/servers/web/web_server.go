package web

import (
	"context"
	"dc2/config"
	"dc2/internal/common/logs"
	"dc2/internal/servers"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebServer struct {
	httpServer *http.Server
	Engin      *gin.Engine
	Apps       *servers.Apps
}

func (s *WebServer) AsyncStart() {
	logs.Debugf("[服务启动] [rpc] 服务地址: %s", s.httpServer.Addr)
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Fatalf("[服务启动] [rpc] 服务异常: %s", zap.Error(err))
		}
	}()
}

func (s *WebServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logs.Debugf("[服务关闭] [rpc] 关闭服务")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		logs.Fatalf("[服务关闭] [rpc] 关闭服务异常: %s", zap.Error(err))
	}
}

func NewWebServer(cfg *config.SugaredConfig, apps *servers.Apps) servers.ServerInterface {
	e := gin.Default()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Web.Port),
		Handler: e,
	}

	s := &WebServer{
		httpServer: httpServer,
		Engin:      e,
		Apps:       apps,
	}

	// 注册路由
	WithRouter(s)

	return s
}
