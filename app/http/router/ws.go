package router

import (
	"github.com/gin-gonic/gin"
	"go-chat/app/http/handler"
	"go-chat/app/http/middleware"
	"go-chat/config"
)

// RegisterWsRoute 注册 Websocket 路由
func RegisterWsRoute(conf *config.Config, router *gin.Engine, handler *handler.Handler) {
	router.GET("/wss/socket.io", middleware.JwtAuth(conf, "api"), handler.Ws.SocketIo)
	router.GET("/wss/admin.io", middleware.JwtAuth(conf, "api"), handler.Ws.AdminIo)
}
