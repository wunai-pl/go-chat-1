package process

import (
	"context"
	"reflect"
	"sync"

	"go-chat/internal/websocket/internal/process/server"
	"golang.org/x/sync/errgroup"
)

var once sync.Once

type IServer interface {
	Setup(ctx context.Context) error
}

// SubServers 订阅的服务列表
type SubServers struct {
	Health    *server.Health      // 注册健康上报
	Subscribe *server.WsSubscribe // 注册消息订阅
}

type Server struct {
	items []IServer
}

func NewServer(routines *SubServers) *Server {
	s := &Server{}

	s.binds(routines)

	return s
}

func (c *Server) binds(routines *SubServers) {
	elem := reflect.ValueOf(routines).Elem()
	for i := 0; i < elem.NumField(); i++ {
		if v, ok := elem.Field(i).Interface().(IServer); ok {
			c.items = append(c.items, v)
		}
	}
}

// Start 启动服务
func (c *Server) Start(eg *errgroup.Group, ctx context.Context) {
	once.Do(func() {
		for _, process := range c.items {
			func(obj IServer) {
				eg.Go(func() error {
					return obj.Setup(ctx)
				})
			}(process)
		}
	})
}
