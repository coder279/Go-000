package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

type indexHandler struct {
	content string
}

// 信号处理
func dealSignal(ctx context.Context) error {

	c := make(chan os.Signal)
	signal.Notify(c)

	fmt.Println("start")
	for {
		select {
		case s := <-c:
			return fmt.Errorf("get %v signal", s)
		case <-ctx.Done():
			return fmt.Errorf("signal routine：other work done")
		}
	}
}

// http请求处理
func startServer(ctx context.Context, addr string, hander http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: hander,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println("process done")
		s.Shutdown(context.Background())

	}(ctx)

	fmt.Println("start。。")
	return s.ListenAndServe()
}

// 写响应
func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ih.content)
}

func main() {
	ctx := context.Background()
	g, cancelCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return dealSignal(cancelCtx)
	})

	g.Go(func() error {
		return startServer(cancelCtx, ":8080", &indexHandler{content: "test success!"})
	})

	if err := g.Wait(); err != nil {
		fmt.Println("err:", err.Error())
	}

	fmt.Println("end")
}