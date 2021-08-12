/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    group
 * @Date:    2021/8/12 1:38 下午
 * @package: group
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package group

import (
	"context"
	"golang.org/x/sync/errgroup"
	"os/signal"
	"syscall"
)

var g *group

type group struct {
	eg     *errgroup.Group
	ctx    context.Context
	cancel context.CancelFunc
}

func init() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	eg, ctx2 := errgroup.WithContext(ctx)
	g = &group{
		eg:     eg,
		ctx:    ctx2,
		cancel: cancel,
	}
}

func Go(f func(ctx context.Context) error) {
	g.eg.Go(func() error {
		return f(g.ctx)
	})
}

func Wait() error {
	return g.eg.Wait()
}

func Cancel() {
	g.cancel()
}
