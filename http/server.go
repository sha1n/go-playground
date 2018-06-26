package http

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type server struct {
	stopChan   chan bool
	httpServer *http.Server
	logger     *logrus.Entry
}

func (server *server) StartAsync() {
	server.logger.Printf("Staring HTTP Server %s", server.httpServer.Addr)

	go func() {
		err := server.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			server.logger.Fatalf("listen: %s\n", err)
		}
	}()

	go func() {
		stop := <-server.stopChan

		server.logger.Debug("received stop signal", stop)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := server.httpServer.Shutdown(ctx); err != nil {
			server.stopChan <- false
			server.logger.Warn("Server Shutdown:", err)
		}
		server.stopChan <- true
	}()
}

func (server *server) StopAsync() {
	server.stopChan <- true
}

func (server *server) StopNow(timeout time.Duration) (err error) {
	server.StopAsync()
	timer := time.NewTimer(timeout)
	select {
	case stopped := <-server.stopChan:
		if !stopped {
			err = errors.New("failed to stop server")
		}
	case <-timer.C:
		err = errors.New("timeout waiting for server to stop")
	}
	return err
}
