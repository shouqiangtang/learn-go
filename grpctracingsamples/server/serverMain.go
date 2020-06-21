package main

import (
	"fmt"
	"net"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"

	pb "learn-go/grpctracingsamples"
	"learn-go/grpctracingsamples/server/service"
)

const (
	// endpointURL            = "http://localhost:9411/api/v1/spans"
	jaegerAgentAddr        = "localhost:6831"
	hostURL                = "localhost:5051"
	serviceNameCacheServer = "cache server"
	network                = "tcp"
)

func newTracer() (opentracing.Tracer, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}
	cfg.ServiceName = serviceNameCacheServer
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter.LocalAgentHostPort = "127.0.0.1:6831"

	tracer, _, err := cfg.NewTracer()
	if err != nil {
		return nil, err
	}

	return tracer, nil
}

func main() {
	fmt.Println("starting server...")
	connection, err := net.Listen(network, hostURL)
	if err != nil {
		panic(err)
	}
	tracer, err := newTracer()
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
		),
	}
	srv := grpc.NewServer(opts...)
	cs := initCache()
	pb.RegisterCacheServiceServer(srv, cs)

	fmt.Println("server listening on port 5051")
	err = srv.Serve(connection)
	if err != nil {
		panic(err)
	}
}

func initCache() pb.CacheServiceServer {
	s := make(map[string][]byte)
	s["123"] = []byte{123}
	s["231"] = []byte{231}
	cs := &service.CacheService{Storage: s}
	return cs
}
