//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"RedisShake/cmd/redis-oms/biz"
	"RedisShake/cmd/redis-oms/server"
	"RedisShake/internal/config"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// initApp init kratos application.
func InitApp(string, *config.Service, *viper.Viper) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, newApp))
}
