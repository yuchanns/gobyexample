package stdresp

type DefaultResp struct {
	Code int
	Msg  string
	Data interface{}
}

type msg struct {
	msg string
}

func (m *msg) apply(resp *DefaultResp) {
	resp.Msg = m.msg
}

func WithMsg(m string) IOption {
	opt := &msg{msg: m}

	return opt
}

type code struct {
	code int
}

func (c *code) apply(resp *DefaultResp) {
	resp.Code = c.code
}

func WithCode(c int) IOption {
	opt := &code{code: c}

	return opt
}

type data struct {
	data interface{}
}

func (d *data) apply(resp *DefaultResp) {
	resp.Data = d.data
}

func WithData(d interface{}) IOption {
	opt := &data{data: d}

	return opt
}

func NewStdResp(data interface{}, opts ...IOption) *DefaultResp {
	resp := &DefaultResp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}

	for _, opt := range opts {
		opt.apply(resp)
	}

	return resp
}

func NewStdRespErr(opts ...IOption) *DefaultResp {
	resp := &DefaultResp{
		Code: 1,
		Msg:  "failed",
		Data: nil,
	}

	for _, opt := range opts {
		opt.apply(resp)
	}

	return resp
}
