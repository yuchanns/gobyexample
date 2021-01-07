package greeter

import (
	"context"
	"fmt"
)

type IDomSvc interface {
	SayHello(ctx context.Context, do *Greeter) string
}

type DomSvc struct {
	greeterRepo IGreeterRepo
}

func NewDomSvc(greeterRepo IGreeterRepo) IDomSvc {
	return &DomSvc{greeterRepo: greeterRepo}
}

func (d *DomSvc) SayHello(ctx context.Context, do *Greeter) string {
	if err := d.greeterRepo.CreateLog(ctx, do); err != nil {
		fmt.Println("failed to create log:", err)
	}
	return fmt.Sprintf("hello, %s", do.Name)
}
