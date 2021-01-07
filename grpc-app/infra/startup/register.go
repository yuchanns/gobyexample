package startup

import (
	"github.com/yuchanns/gobyexample/grpc-app/common"
	"github.com/yuchanns/gobyexample/grpc-app/infra/startup/wire"
	helloworld "github.com/yuchanns/gobyexample/grpc-app/proto/greeter"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RegisterGrpcServer(srv *grpc.Server) error {
	greeterSrv := wire.InitializeGreeterServer()
	helloworld.RegisterGreeterServer(srv, greeterSrv)
	return nil
}

func RegisterVars() error {
	var err error
	common.DB, err = newGorm()
	return err
}

func newGorm() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}
