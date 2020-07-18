package ant

import (
	"github.com/pkg/errors"
	_interface "github.com/yuchanns/gobyexample/validator/interface"
	"net/http"
)

func NewService() _interface.IService {
	return &Service{
		Client: &http.Client{},
	}
}

// ant.Service is one of those gateway services implementing IService
type Service struct {
	*_interface.AService
	Client *http.Client
}

func (s *Service) DoRegister(data interface{}) (map[string]interface{}, error) {
	register, ok := data.(*Register)
	if !ok {
		return nil, errors.New("cannot assert data as ant.Register")
	}
	// do actual request things
	resp := map[string]interface{}{
		"bankAccountNo": register.BankAccountNo,
		"status":        1,
	}

	return resp, nil
}
