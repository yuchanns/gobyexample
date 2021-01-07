package startup

import (
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/yuchanns/gobyexample/grpc-app/common"
	"github.com/yuchanns/gobyexample/grpc-app/infra/startup/wire"
	helloworld "github.com/yuchanns/gobyexample/grpc-app/proto/greeter"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func BuildGrpcOpentracingMiddlewares() ([]grpc.ServerOption, func(), error) {
	var opts []grpc.ServerOption
	tracer, closer, err := common.NewJaegerTracer("grpc-app", os.Getenv("AGENT_HOST_PORT"))
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts,
		grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())),
		grpc.StreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(tracer, otgrpc.LogPayloads())),
	)
	closeFunc := func() {
		_ = closer.Close()
	}
	return opts, closeFunc, nil
}

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
