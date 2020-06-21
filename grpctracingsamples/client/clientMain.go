package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "learn-go/grpctracingsamples"
	"learn-go/grpctracingsamples/client/service"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
)

const (
	jaegerAgentAddr        = "127.0.0.1:6831"
	hostURL                = "localhost:5051"
	serviceNameCacheClient = "cache client"
)

func newTracer() (opentracing.Tracer, io.Closer, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, nil, err
	}
	cfg.ServiceName = serviceNameCacheClient
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter.LocalAgentHostPort = jaegerAgentAddr

	// TODO(ys) a quick hack to ensure random generators get different seeds, which are based on current time.
	time.Sleep(100 * time.Millisecond)

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}

	return tracer, closer, nil
}

func main() {
	fmt.Println("client start")
	tracer, closer, err := newTracer()
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	conn, err := grpc.Dial(hostURL, grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	csc := pb.NewCacheServiceClient(conn)
	callServer(csc)
}

func callServer(csc pb.CacheServiceClient) {
	ctx := context.Background()
	key := "123"
	cc := service.CacheClient{}
	value, err := cc.CallGet(ctx, key, csc)
	if err != nil {
		fmt.Println("error call get:", err)
	} else {
		fmt.Printf("value=%v for key=%v\n", value, key)
	}
	key = "231"
	value, err = cc.CallGet(ctx, key, csc)
	if err != nil {
		fmt.Println("error call get:", err)
	} else {
		fmt.Printf("value=%v for key=%v\n", value, key)
	}
}
