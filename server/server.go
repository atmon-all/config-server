package server

import (
	"fmt"
	"net"

	"github.com/atmom/configserver/config"
	configserver "github.com/atmom/configserver/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// start server.
func Start(config *config.ServerConfig) {
	if config.Port < 1024 || config.Port > 65536 {
		logrus.Fatalf("Config server start error, invalid server port %d.", config.Port)
	}

	// listen tcp.
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.Fatalf("Config server start error %v.", err)
	}

	// create server.
	s := grpc.NewServer()
	configserver.RegisterConfigServerServer(s, &service{Config: config})

	logrus.Infof("Config server start success, listening at %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		logrus.Fatalf("Config server start error %v.", err)
	}
}
