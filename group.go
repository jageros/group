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
	"os"
	"os/signal"
	"syscall"
)

var g_ *Group

type Group struct {
	eg     *errgroup.Group
	ctx    context.Context
	cancel context.CancelFunc
}

func New() *Group {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx2 := errgroup.WithContext(ctx)
	gp := &Group{
		eg:     eg,
		ctx:    ctx2,
		cancel: cancel,
	}
	return gp
}

func NewWithContext(ctx context.Context) *Group {
	ctx_, cancel := context.WithCancel(ctx)
	eg, ctx2 := errgroup.WithContext(ctx_)
	gp := &Group{
		eg:     eg,
		ctx:    ctx2,
		cancel: cancel,
	}
	return gp
}

func NewWithSignal(sig ...os.Signal) *Group {
	ctx, cancel := signal.NotifyContext(context.Background(), sig...)
	eg, ctx2 := errgroup.WithContext(ctx)
	gp := &Group{
		eg:     eg,
		ctx:    ctx2,
		cancel: cancel,
	}
	return gp
}

func Default() *Group {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	eg, ctx2 := errgroup.WithContext(ctx)
	gp := &Group{
		eg:     eg,
		ctx:    ctx2,
		cancel: cancel,
	}
	return gp
}

func (g *Group) Go(f func(ctx context.Context) error) {
	g.eg.Go(func() error {
		return f(g.ctx)
	})
}

func (g *Group) Wait() error {
	return g.eg.Wait()
}

func (g *Group) Cancel() {
	g.cancel()
}

// ============ Global API ============

func init() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	eg, ctx2 := errgroup.WithContext(ctx)
	g_ = &Group{
		eg:     eg,
		ctx:    ctx2,
		cancel: cancel,
	}
}

func Go(f func(ctx context.Context) error) {
	g_.eg.Go(func() error {
		return f(g_.ctx)
	})
}

func Wait() error {
	return g_.eg.Wait()
}

func Cancel() {
	g_.cancel()
}
