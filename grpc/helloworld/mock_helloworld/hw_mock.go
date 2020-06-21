package mock_helloworld

import (
	"github.com/golang/mock/gomock"
)

// MockGreeterClient : Mock of GreeterClient interface
type MockGreeterClient struct {
	ctrl     *gomock.Controller
	recorder *_MockGreeterClientRecorder
}

// Recorder for MockGreeterClient (not exported)
type _MockGreeterClientRecorder struct {
	mock *MockGreeterClient
}

// NewMockGreeterClient : NewMockGreeterClient
func NewMockGreeterClient(ctrl *gomock.Controller) *MockGreeterClient {
	mock := &MockGreeterClient{ctrl: ctrl}
	mock.recorder = &_MockGreeterClientRecorder{mock}
	return mock
}

// EXPECT : EXPECT
func (_m *MockGreeterClient) EXPECT() *_MockGreeterClientRecorder {
	return _m.recorder
}
