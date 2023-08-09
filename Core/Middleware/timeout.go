package Middleware

import (
	"context"
	"net/http"
	"time"
)

const (
	timeout       = 300 * time.Second
	timeout_error = "Request Timeout"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (self *TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if self.Next == nil {
		self.Next = http.DefaultServeMux
	}

	oldContext := r.Context()
	ctx, cancel := context.WithTimeout(oldContext, timeout)
	defer cancel()
	r.WithContext(ctx)

	ch := make(chan struct{}, 1)

	processTimeout(w, r, self, ch, ctx)

}

func processTimeout(w http.ResponseWriter, r *http.Request, m *TimeoutMiddleware, ch chan struct{}, ctx context.Context) {
	go func() {
		m.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte(timeout_error))
		return
	case <-ch:
		return

	}
}
