package validator

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/yuchanns/gobyexample/validator/ant"
	_interface "github.com/yuchanns/gobyexample/validator/interface"
)

var Services = map[string]func() _interface.AService{
	"ant": func() _interface.AService {
		return &ant.Service{}
	},
}

func Register(svc _interface.AService, decoder *json.Decoder) (interface{}, error) {
	register, err := svc.GetRegisterStruct()
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(register); err != nil {
		return nil, err
	}
	validate := validator.New()
	if err := validate.Struct(register); err != nil {
		return nil, err
	}

	return register, nil
}
