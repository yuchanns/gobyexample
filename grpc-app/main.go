package main

import (
	"github.com/yuchanns/gobyexample/grpc-app/common"
	"github.com/yuchanns/gobyexample/grpc-app/infra/startup"
	"os"
)

func main() {
	app := &common.GrpcApplication{
		Endpoint:           ":9090",
		GatewayAddr:        ":8080",
		AppName:            "grpc-app",
		AgentHostPort:      os.Getenv("AGENT_HOST_PORT"),
		RegisterGrpcServer: startup.RegisterGrpcServer,
		RegisterGateway:    startup.RegisterGateway,
		RegisterVars:       startup.RegisterVars,
	}
	server := common.NewGrpcServer(app)
	server.Run()
	defer server.Close()
}
