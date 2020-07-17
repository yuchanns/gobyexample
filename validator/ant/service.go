package ant

import "github.com/pkg/errors"

type Register struct {
	BankAccountNo   string `json:"bankAccountNo" validate:"required"`
	CertificateNo   string `json:"certificateNo" validate:"required"`
	Mobile          string `json:"mobile" validate:"required"`
	SubMerchantName string `json:"subMerchantName" validate:"required"`
}

type Service struct {
	//
}

func (s *Service) GetRegisterStruct() (interface{}, error) {
	return &Register{}, nil
}

func (s *Service) GetModifyStruct() (interface{}, error) {
	return nil, errors.New("invalid route")
}

func (s *Service) GetBindStruct() (interface{}, error) {
	return nil, errors.New("invalid route")
}

func (s *Service) GetConfirmStruct() (interface{}, error) {
	return nil, errors.New("invalid route")
}

func (s *Service) GetConsumeStruct() (interface{}, error) {
	return nil, errors.New("invalid route")
}

func (s *Service) GetRefundStruct() (interface{}, error) {
	return nil, errors.New("invalid route")
}

func (s *Service) GetQueryConsumeStruct() (interface{}, error) {
	return nil, errors.New("invalid route")
}

func (s *Service) GetQueryRefundStruct() (interface{}, error) {
	return nil, errors.New("invalid route")
}

func (s *Service) GetBalanceStruct() (interface{}, error) {
	return nil, errors.New("invalid route")
}

func (s *Service) GetUnbindStruct() (interface{}, error) {
	return &Register{}, nil
}

func (s *Service) DoRegister(data interface{}) (map[string]interface{}, error) {
	register, ok := data.(*Register)
	if !ok {
		return nil, errors.New("cannot assert data as ant.Register")
	}
	// do actual request things
	resp := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{
			"bankAccountNo": register.BankAccountNo,
			"status":        1,
		},
	}

	return resp, nil
}
