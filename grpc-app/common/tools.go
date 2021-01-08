package common

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/opentracing/opentracing-go"
)

func MustConvert(from, to interface{}) {
	jsonByte, err := json.Marshal(from)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(jsonByte, to); err != nil {
		panic(err)
	}
}

// NewRequest returns a well configured *resty.Request
// opentracing.SpanContext will be injected into request header if exists
func NewRequest(ctx context.Context) *resty.Request {
	req := resty.New().R().SetContext(ctx).EnableTrace()
	// Inject span
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		parentCtx := parent.Context()
		_ = opentracing.GlobalTracer().Inject(
			parentCtx,
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(req.Header))
	}

	return req
}
