package web

import (
	"ddd_demo2/internal/user"
	"ddd_demo2/internal/youke"

	"github.com/gin-gonic/gin"
)

func WithRouter(s *WebServer) {
	// 新建 handler
	userHandler := user.NewUserHandler(s.Apps.UserApp)
	authMiddleware := user.NewAuthMiddleware(s.Apps.UserApp)
	youkeHandler := youke.NewYoukeHandler(s.Apps.YoukeApp)

	// ====== 测试 ======
	s.Engin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	baseApi := s.Engin.Group("/dc2")

	// 鉴权
	auth := baseApi.Group("/auth")
	auth.POST("/login", userHandler.Login)
	auth.POST("/register", userHandler.Register)

	// api
	api := baseApi.Group("/api")

	// 中间件
	api.Use(authMiddleware.Auth)

	// 路由
	api.GET("/user_info", userHandler.UserInfo)
	api.POST("/transfer", userHandler.Transfer)

	// =====优客====

	youkeApi := baseApi.Group("/youke")
	youkeApi.POST("/create_order", youkeHandler.CreateOrder)
	youkeApi.POST("/get_user_orders", youkeHandler.GetUserOrders)
	youkeApi.POST("/get_user", youkeHandler.GetUser)
	youkeApi.POST("/save_user", youkeHandler.SaveUser)

}

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
