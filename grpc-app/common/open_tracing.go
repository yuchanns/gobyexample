package common

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"io"
)

func NewJaegerTracer(name, hostPort string) (opentracing.Tracer, io.Closer, error) {
	tp, err := jaeger.NewUDPTransport(hostPort, 0)
	if err != nil {
		return nil, nil, err
	}
	reporter := jaeger.NewRemoteReporter(tp)

	var sampler jaeger.Sampler
	sampler = jaeger.NewConstSampler(true)

	tracer, closer := jaeger.NewTracer(name,
		sampler,
		reporter,
	)
	opentracing.SetGlobalTracer(tracer)

	return tracer, closer, nil
}
