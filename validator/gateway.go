package validator

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/yuchanns/gobyexample/validator/ant"
	_interface "github.com/yuchanns/gobyexample/validator/interface"
)

var Services = map[string]func() _interface.IService{
	"ant": ant.NewService,
}

// the unified entrance for register
func Register(svc _interface.IService, decoder *json.Decoder) (interface{}, error) {
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

// the unified entrance for modify
func Modify(svc _interface.IService, decoder *json.Decoder) (interface{}, error) {
	modify, err := svc.GetModifyStruct()
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(modify); err != nil {
		return nil, err
	}
	validate := validator.New()
	if err := validate.Struct(modify); err != nil {
		return nil, err
	}

	return modify, nil
}
