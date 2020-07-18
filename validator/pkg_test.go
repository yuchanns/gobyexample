package validator

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"
)

func TestRegister(t *testing.T) {
	// assume requestData
	requestSvc := "ant"
	requestData := []byte(`{"bankAccountNo":"12345678","certificateNo":"11010119900307571X","mobile":"13694283597","subMerchantName":"yuchanns"}`)
	svcNew, ok := Services[requestSvc]
	if !ok {
		t.Error("service not found")
	}
	var decoder = json.NewDecoder(bytes.NewReader(requestData))
	svc := svcNew()
	register, err := Register(svc, decoder)
	if err != nil {
		t.Error(err)
	}
	resp, err := svc.DoRegister(register)
	if err != nil {
		t.Error(err)
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		t.Error(err)
	}
	t.Log(bytes.NewBuffer(jsonResp).String())
}

func TestInvalidRequest(t *testing.T) {
	// assume requestData
	requestSvc := "ant"
	requestData := []byte(`{"bankAccountNo":"12345678","certificateNo":"11010119900307571X","mobile":"13694283597","subMerchantName":"yuchanns"}`)
	svcNew, ok := Services[requestSvc]
	if !ok {
		t.Error("service not found")
	}
	var decoder = json.NewDecoder(bytes.NewReader(requestData))
	svc := svcNew()
	_, err := Modify(svc, decoder)
	if err == nil {
		t.Error(errors.New("err is not nil"))
	}
	t.Log(err)
}
