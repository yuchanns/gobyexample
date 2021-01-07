package common

import (
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
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

	return tracer, closer, nil
}

func BuildGrpcOpentracingMiddlewares(name, agentHostPort string) ([]grpc.ServerOption, func(), error) {
	var opts []grpc.ServerOption
	tracer, closer, err := NewJaegerTracer(name, agentHostPort)
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts,
		grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())),
		grpc.StreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(tracer, otgrpc.LogPayloads())),
	)
	closeFunc := func() {
		_ = closer.Close()
	}
	return opts, closeFunc, nil
}
