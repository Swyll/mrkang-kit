package kindergarten

import (
	"context"
	"mrkang-kit/log"
	"mrkang-kit/log/zaplogger"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

//WimpyKid 服务
type WimpyKid interface {
	Start() error
	Finalize() error
}

type Kindergaeten struct {
	ctx  context.Context
	sigs []os.Signal
	log  log.Logger
	kids []WimpyKid

	cancel func()
}

type opt func(*Kindergaeten)

func NewKind(opts ...opt) *Kindergaeten {
	k := &Kindergaeten{
		sigs: []os.Signal{
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGQUIT,
		},
		log:  zaplogger.NewDefauLogger(),
		kids: make([]WimpyKid, 0, 4),
	}

	ctx, cancel := context.WithCancel(context.Background())
	k.ctx = ctx
	k.cancel = cancel

	for _, opt := range opts {
		opt(k)
	}

	return k
}

func WithLogger(log log.Logger) opt {
	return func(k *Kindergaeten) {
		k.log = log
	}
}

func WithSigs(sigs ...os.Signal) opt {
	return func(k *Kindergaeten) {
		k.sigs = sigs
	}
}

func (k *Kindergaeten) Register(wk WimpyKid) {
	k.kids = append(k.kids, wk)
}

func (k *Kindergaeten) Start() error {
	eg, ctx := errgroup.WithContext(k.ctx)

	c := make(chan os.Signal)
	signal.Notify(c, k.sigs...)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			return func() error {
				k.Stop()
				return nil
			}()
		}
	})

	for _, kid := range k.kids {
		server := kid
		eg.Go(kid.Start)

		eg.Go(func() error {
			<-ctx.Done()
			return server.Finalize()
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (k *Kindergaeten) Stop() {
	if k.cancel != nil {
		k.cancel()
	}
}
