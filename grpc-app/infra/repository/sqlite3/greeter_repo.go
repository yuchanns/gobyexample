package sqlite3

import (
	"context"
	"github.com/yuchanns/gobyexample/grpc-app/common"
	"github.com/yuchanns/gobyexample/grpc-app/domain/greeter"
	"github.com/yuchanns/gobyexample/grpc-app/infra/repository/sqlite3/models"
)

type GreeterRepo struct {
	//
}

func NewGreeterRepo() greeter.IGreeterRepo {
	return &GreeterRepo{}
}

func (g *GreeterRepo) CreateLog(ctx context.Context, do *greeter.Greeter) error {
	visitor := models.DoToVisitor(do)
	return common.DB.WithContext(ctx).Create(visitor).Error
}
