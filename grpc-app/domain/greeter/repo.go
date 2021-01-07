package greeter

import "context"

type IGreeterRepo interface {
	CreateLog(ctx context.Context, do *Greeter) error
}
