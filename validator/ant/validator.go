package ant

type Register struct {
	BankAccountNo   string `json:"bankAccountNo" validate:"required"`
	CertificateNo   string `json:"certificateNo" validate:"required"`
	Mobile          string `json:"mobile" validate:"required"`
	SubMerchantName string `json:"subMerchantName" validate:"required"`
}

func (s *Service) GetRegisterStruct() (interface{}, error) {
	return &Register{}, nil
}
