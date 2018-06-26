package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sha1n/go-playground/logging"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
}

type Server interface {
	StartAsync()
	StopAsync()
	StopNow(timeout time.Duration) error
}

type ServerBuilder interface {
	Build() Server
	WithGetHandler(path string, handler func(c *gin.Context)) ServerBuilder
	WithPostHandler(path string, handler func(c *gin.Context)) ServerBuilder
}

type serverBuilder struct {
	port   int
	engine *gin.Engine
	logger *logrus.Entry
}

func (sb *serverBuilder) Build() Server {
	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(int(sb.port)),
		Handler: sb.engine,
	}

	s := &server{
		stopChan:   make(chan bool, 1),
		httpServer: httpServer,
		logger:     sb.logger,
	}

	return s
}

func (sb *serverBuilder) WithGetHandler(path string, handler func(c *gin.Context)) ServerBuilder {
	sb.engine.GET(path, handlerWrapperFor(sb.logger, path, handler))
	return sb
}

func (sb *serverBuilder) WithPostHandler(path string, handler func(c *gin.Context)) ServerBuilder {
	sb.engine.POST(path, handlerWrapperFor(sb.logger, path, handler))
	return sb
}

func NewServer(port int) ServerBuilder {
	logger := logging.NewEntryFor("http-server").WithField("port", port)

	router := gin.Default()
	router.Use(gin.Recovery())
	router.HandleMethodNotAllowed = true

	sb := &serverBuilder{
		port:   port,
		engine: router,
		logger: logger,
	}

	return sb
}

func handlerWrapperFor(logger *logrus.Entry, path string, handler func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		trace := logging.NewTrace(strings.TrimPrefix(path, "/"), logger)
		c.Set("trace", trace)

		traceSegment := trace.StartSegment("handleRequest")
		defer traceSegment.End()

		// calling actual handler
		handler(c)

	}
}
