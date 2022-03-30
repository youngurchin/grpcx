package grpcx

import (
	"github.com/fagongzi/log"
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	addr   string
	server *gin.Engine
}

func newHTTPServer(addr string, httpSetup func(*gin.Engine)) *httpServer {
	server := gin.Default()
	httpSetup(server)

	return &httpServer{
		addr:   addr,
		server: server,
	}
}

func (s *httpServer) start() error {
	log.Infof("rpc: start a grpc http proxy server at %s", s.addr)
	return s.server.Run(s.addr)
}

func (s *httpServer) stop() error {
	return nil
}
