package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
)

var _ transport.Server = (*Watcher)(nil)

type Watcher struct{}

func NewWatcher() *Watcher {
	return &Watcher{}
}

func (w *Watcher) Start(context.Context) error {
	return nil
}
func (w *Watcher) Stop(context.Context) error {
	return nil
}
