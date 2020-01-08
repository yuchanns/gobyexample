package stdresp

type IOption interface {
	apply(*DefaultResp)
}
