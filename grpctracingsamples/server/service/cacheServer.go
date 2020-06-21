package service

import (
	"context"
	"fmt"
	pb "learn-go/grpctracingsamples"
	"time"

	"github.com/opentracing/opentracing-go"
)

const serviceNameDbQueryUser = "db query user"

// CacheService : CacheService struct
type CacheService struct {
	Storage map[string][]byte
}

// Get : Get function
func (cs *CacheService) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	time.Sleep(5 * time.Millisecond)
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		pctx := parent.Context()
		if tracer := opentracing.GlobalTracer(); tracer != nil {
			mysqlSpan := tracer.StartSpan(serviceNameDbQueryUser, opentracing.ChildOf(pctx))
			defer mysqlSpan.Finish()
			// do some options
			time.Sleep(time.Millisecond * 10)
		}
	}
	key := req.GetKey()
	value := cs.Storage[key]
	fmt.Println("get called with return of value: ", value)
	resp := &pb.GetResp{Value: value}
	return resp, nil
}

// Store : Store function
func (cs *CacheService) Store(ctx context.Context, req *pb.StoreReq) (*pb.StoreResp, error) {
	key := req.Key
	value := req.Value
	if oldValue, ok := cs.Storage[key]; ok {
		cs.Storage[key] = value
		fmt.Printf(" key=%v already exist, old vale=%v|replaced with new value=%v\n", key, oldValue, cs.Storage)
	} else {
		cs.Storage[key] = value
		fmt.Printf(" key=%v not existing, add new value=%v\n", key, cs.Storage)
	}
	r := &pb.StoreResp{}
	return r, nil
}

// Dump : Dump function
func (cs *CacheService) Dump(req *pb.DumpReq, resp pb.CacheService_DumpServer) error {
	return nil
}
