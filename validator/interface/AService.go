package _interface

import "github.com/pkg/errors"

// IService is an interface that requires validators and actions implemented
type IService interface {
	// Get Data Structure And Validator
	GetRegisterStruct() (interface{}, error)
	GetModifyStruct() (interface{}, error)
	GetBindStruct() (interface{}, error)
	GetConfirmStruct() (interface{}, error)
	GetConsumeStruct() (interface{}, error)
	GetRefundStruct() (interface{}, error)
	GetQueryConsumeStruct() (interface{}, error)
	GetQueryRefundStruct() (interface{}, error)
	GetBalanceStruct() (interface{}, error)
	GetUnbindStruct() (interface{}, error)

	// Action Do Request
	DoRegister(interface{}) (map[string]interface{}, error)
	DoModify(interface{}) (map[string]interface{}, error)
	DoBind(interface{}) (map[string]interface{}, error)
	DoConfirm(interface{}) (map[string]interface{}, error)
	DoConsume(interface{}) (map[string]interface{}, error)
	DoRefund(interface{}) (map[string]interface{}, error)
	DoQueryConsume(interface{}) (map[string]interface{}, error)
	DoQueryRefund(interface{}) (map[string]interface{}, error)
	DoBalance(interface{}) (map[string]interface{}, error)
	DoUnbind(interface{}) (map[string]interface{}, error)
}

// AService is for embedded as the default implementation into custom service implementing IService
type AService struct{}

func (s *AService) GetRegisterStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) GetModifyStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) GetBindStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) GetConfirmStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) GetConsumeStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) GetRefundStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) GetQueryConsumeStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) GetQueryRefundStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) GetBalanceStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) GetUnbindStruct() (interface{}, error) {
	return nil, errors.New("invalid validator")
}

func (s *AService) DoRegister(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
func (s *AService) DoModify(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
func (s *AService) DoBind(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
func (s *AService) DoConfirm(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
func (s *AService) DoConsume(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
func (s *AService) DoRefund(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
func (s *AService) DoQueryConsume(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
func (s *AService) DoQueryRefund(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
func (s *AService) DoBalance(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
func (s *AService) DoUnbind(interface{}) (map[string]interface{}, error) {
	return nil, errors.New("invalid route")
}
