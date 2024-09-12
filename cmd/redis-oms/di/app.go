package di

import (
	"RedisShake/cmd/redis-oms/biz"
	"RedisShake/internal/config"
	"github.com/go-bamboo/pkg/kratos"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/rpc"
)

func newApp(id string, srv *config.Service, gs *rpc.Server, w *biz.Watcher) *kratos.App {
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(srv.Name),
		kratos.Version(srv.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(log.NewLogger(log.GetCore(), 2)),
		kratos.Server(
			gs,
			w,
		),
		//kratos.Registrar(r),
	)
	return app
}
