package main

import (
	"RedisShake/cmd/redis-oms/di"
	"RedisShake/internal/config"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/uuid"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// Branch is current branch name the code is built off.
	Branch string
	// Revision is the short commit hash of source tree.
	Revision string
	// BuildDate is the date when the binary was built.
	BuildDate string
)

func main() {
	_ = config.LoadConfig()

	// uuid
	id := uuid.New()

	app, closeFunc, err := di.InitApp(id, config.Opt.Service)
	if err != nil {
		log.Errorf("err: %v", err.Error())
		return
	}
	defer closeFunc()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Errorf("err: %v", err.Error())
		return
	}
}
