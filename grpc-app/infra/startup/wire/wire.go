//+build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/yuchanns/gobyexample/grpc-app/app"
	"github.com/yuchanns/gobyexample/grpc-app/domain/greeter"
	"github.com/yuchanns/gobyexample/grpc-app/infra/repository/sqlite3"
	helloworld "github.com/yuchanns/gobyexample/grpc-app/proto/greeter"
	"github.com/yuchanns/gobyexample/grpc-app/server"
)

func InitializeGreeterServer() helloworld.GreeterServer {
	wire.Build(server.NewGreeterServer, app.NewGreeterSvc, greeter.NewDomSvc, sqlite3.NewGreeterRepo)
	return &server.GreeterServer{}
}
