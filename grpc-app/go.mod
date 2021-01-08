module github.com/yuchanns/gobyexample/grpc-app

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-resty/resty/v2 v2.3.0
	github.com/golang/protobuf v1.4.3
	github.com/google/wire v0.4.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.uber.org/atomic v1.7.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.34.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.9
)
