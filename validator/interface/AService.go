package _interface

type AService interface {
	// Get Data Structure
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
}
