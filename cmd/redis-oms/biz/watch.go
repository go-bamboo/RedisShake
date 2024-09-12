package biz

import (
	"RedisShake/internal/config"
	"RedisShake/internal/filter"
	"RedisShake/internal/log"
	"RedisShake/internal/reader"
	"RedisShake/internal/utils"
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
	"sync"
	"time"
)

var _ transport.Server = (*Watcher)(nil)

type Watcher struct {
	ctx context.Context
	cf  context.CancelFunc
	v   *viper.Viper
	wg  sync.WaitGroup
}

func NewWatcher(v *viper.Viper) *Watcher {
	ctx, cancel := context.WithCancel(context.Background())
	return &Watcher{
		ctx: ctx,
		cf:  cancel,
		v:   v,
	}
}

func (w *Watcher) Start(context.Context) error {
	w.wg.Add(1)
	go w.run()
	return nil
}
func (w *Watcher) Stop(context.Context) error {
	w.cf()
	w.wg.Wait()
	return nil
}

func (w *Watcher) run() {
	defer w.wg.Done()
	utils.ChdirAndAcquireFileLock()
	luaRuntime := filter.NewFunctionFilter(config.Opt.Filter.Function)

	opts := new(reader.SyncReaderOptions)
	defaults.SetDefaults(opts)
	err := w.v.UnmarshalKey("sync_reader", opts)
	if err != nil {
		log.Panicf("failed to read the SyncReader config entry. err: %v", err)
	}
	var theReader reader.Reader
	if opts.Cluster {
		theReader = reader.NewSyncClusterReader(w.ctx, opts)
		log.Infof("create SyncClusterReader: %v", opts.Address)
	} else {
		theReader = reader.NewSyncStandaloneReader(w.ctx, opts)
		log.Infof("create SyncStandaloneReader: %v", opts.Address)
	}

	log.Infof("start syncing...")

	ch := theReader.StartRead(w.ctx)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
Loop:
	for {
		select {
		case e, ok := <-ch:
			if !ok {
				// ch has been closed, exit the loop
				break Loop
			}
			// calc arguments
			e.Parse()

			// filter
			if !filter.Filter(e) {
				log.Debugf("skip command: %v", e)
				continue
			}
			log.Debugf("function before: %v", e)
			entries := luaRuntime.RunFunction(e)
			log.Debugf("function after: %v", entries)

			for _, theEntry := range entries {
				theEntry.Parse()
			}
		case <-ticker.C:
		}
	}

	utils.ReleaseFileLock() // Release file lock
	log.Infof("all done")
}
