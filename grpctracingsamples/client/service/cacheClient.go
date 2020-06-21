package service

import (
	"context"
	pb "learn-go/grpctracingsamples"
)

// CacheClient : grpc client
type CacheClient struct{}

// CallGet : call get function
func (cc *CacheClient) CallGet(ctx context.Context, key string, csc pb.CacheServiceClient) ([]byte, error) {
	getReq := &pb.GetReq{Key: key}
	getResp, err := csc.Get(ctx, getReq)
	if err != nil {
		return nil, err
	}
	value := getResp.Value
	return value, err
}

// CallStore : call store function
func (cc *CacheClient) CallStore(key string, value []byte, client pb.CacheServiceClient) (*pb.StoreResp, error) {
	storeReq := &pb.StoreReq{Key: key, Value: value}
	storeResp, err := client.Store(context.Background(), storeReq)
	if err != nil {
		return nil, err
	}
	return storeResp, err
}
