package server

import (
	"context"

	"github.com/atmom/configserver/config"
	configserver "github.com/atmom/configserver/grpc"
	"github.com/sirupsen/logrus"
)

type service struct {
	configserver.UnimplementedConfigServerServer
	Config *config.ServerConfig
}

// Get Flow configuration from config server.
func (service *service) GetFlow(ctx context.Context, request *configserver.GetFlowRequest) (*configserver.GetFlowResponse, error) {
	logrus.Infof("Received: %v", request)
	return &configserver.GetFlowResponse{Code: 0, Message: "success"}, nil
}
