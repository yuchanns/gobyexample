//+build wireinject

package wire

import (
	"github.com/google/wire"
)

func InitialEvent() *Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)

	return &Event{}
}
