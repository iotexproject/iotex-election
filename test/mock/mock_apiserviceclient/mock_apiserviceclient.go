// Code generated by MockGen. DO NOT EDIT.
// Source: ./pb/api/api.pb.go

// Package mock_apiserviceclient is a generated GoMock package.
package mock_apiserviceclient

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	api "github.com/iotexproject/iotex-election/pb/api"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAPIServiceClient is a mock of APIServiceClient interface
type MockAPIServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAPIServiceClientMockRecorder
}

// MockAPIServiceClientMockRecorder is the mock recorder for MockAPIServiceClient
type MockAPIServiceClientMockRecorder struct {
	mock *MockAPIServiceClient
}

// NewMockAPIServiceClient creates a new mock instance
func NewMockAPIServiceClient(ctrl *gomock.Controller) *MockAPIServiceClient {
	mock := &MockAPIServiceClient{ctrl: ctrl}
	mock.recorder = &MockAPIServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIServiceClient) EXPECT() *MockAPIServiceClientMockRecorder {
	return m.recorder
}

// GetMeta mocks base method
func (m *MockAPIServiceClient) GetMeta(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*api.ChainMeta, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMeta", varargs...)
	ret0, _ := ret[0].(*api.ChainMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeta indicates an expected call of GetMeta
func (mr *MockAPIServiceClientMockRecorder) GetMeta(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeta", reflect.TypeOf((*MockAPIServiceClient)(nil).GetMeta), varargs...)
}

// GetCandidates mocks base method
func (m *MockAPIServiceClient) GetCandidates(ctx context.Context, in *api.GetCandidatesRequest, opts ...grpc.CallOption) (*api.CandidateResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCandidates", varargs...)
	ret0, _ := ret[0].(*api.CandidateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidates indicates an expected call of GetCandidates
func (mr *MockAPIServiceClientMockRecorder) GetCandidates(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidates", reflect.TypeOf((*MockAPIServiceClient)(nil).GetCandidates), varargs...)
}

// GetCandidateByName mocks base method
func (m *MockAPIServiceClient) GetCandidateByName(ctx context.Context, in *api.GetCandidateByNameRequest, opts ...grpc.CallOption) (*api.Candidate, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCandidateByName", varargs...)
	ret0, _ := ret[0].(*api.Candidate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidateByName indicates an expected call of GetCandidateByName
func (mr *MockAPIServiceClientMockRecorder) GetCandidateByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidateByName", reflect.TypeOf((*MockAPIServiceClient)(nil).GetCandidateByName), varargs...)
}

// GetBucketsByCandidate mocks base method
func (m *MockAPIServiceClient) GetBucketsByCandidate(ctx context.Context, in *api.GetBucketsByCandidateRequest, opts ...grpc.CallOption) (*api.BucketResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketsByCandidate", varargs...)
	ret0, _ := ret[0].(*api.BucketResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketsByCandidate indicates an expected call of GetBucketsByCandidate
func (mr *MockAPIServiceClientMockRecorder) GetBucketsByCandidate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketsByCandidate", reflect.TypeOf((*MockAPIServiceClient)(nil).GetBucketsByCandidate), varargs...)
}

// IsHealth mocks base method
func (m *MockAPIServiceClient) IsHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*api.HealthCheckResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsHealth", varargs...)
	ret0, _ := ret[0].(*api.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsHealth indicates an expected call of IsHealth
func (mr *MockAPIServiceClientMockRecorder) IsHealth(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHealth", reflect.TypeOf((*MockAPIServiceClient)(nil).IsHealth), varargs...)
}

// MockAPIServiceServer is a mock of APIServiceServer interface
type MockAPIServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAPIServiceServerMockRecorder
}

// MockAPIServiceServerMockRecorder is the mock recorder for MockAPIServiceServer
type MockAPIServiceServerMockRecorder struct {
	mock *MockAPIServiceServer
}

// NewMockAPIServiceServer creates a new mock instance
func NewMockAPIServiceServer(ctrl *gomock.Controller) *MockAPIServiceServer {
	mock := &MockAPIServiceServer{ctrl: ctrl}
	mock.recorder = &MockAPIServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIServiceServer) EXPECT() *MockAPIServiceServerMockRecorder {
	return m.recorder
}

// GetMeta mocks base method
func (m *MockAPIServiceServer) GetMeta(arg0 context.Context, arg1 *empty.Empty) (*api.ChainMeta, error) {
	ret := m.ctrl.Call(m, "GetMeta", arg0, arg1)
	ret0, _ := ret[0].(*api.ChainMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeta indicates an expected call of GetMeta
func (mr *MockAPIServiceServerMockRecorder) GetMeta(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeta", reflect.TypeOf((*MockAPIServiceServer)(nil).GetMeta), arg0, arg1)
}

// GetCandidates mocks base method
func (m *MockAPIServiceServer) GetCandidates(arg0 context.Context, arg1 *api.GetCandidatesRequest) (*api.CandidateResponse, error) {
	ret := m.ctrl.Call(m, "GetCandidates", arg0, arg1)
	ret0, _ := ret[0].(*api.CandidateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidates indicates an expected call of GetCandidates
func (mr *MockAPIServiceServerMockRecorder) GetCandidates(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidates", reflect.TypeOf((*MockAPIServiceServer)(nil).GetCandidates), arg0, arg1)
}

// GetCandidateByName mocks base method
func (m *MockAPIServiceServer) GetCandidateByName(arg0 context.Context, arg1 *api.GetCandidateByNameRequest) (*api.Candidate, error) {
	ret := m.ctrl.Call(m, "GetCandidateByName", arg0, arg1)
	ret0, _ := ret[0].(*api.Candidate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidateByName indicates an expected call of GetCandidateByName
func (mr *MockAPIServiceServerMockRecorder) GetCandidateByName(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidateByName", reflect.TypeOf((*MockAPIServiceServer)(nil).GetCandidateByName), arg0, arg1)
}

// GetBucketsByCandidate mocks base method
func (m *MockAPIServiceServer) GetBucketsByCandidate(arg0 context.Context, arg1 *api.GetBucketsByCandidateRequest) (*api.BucketResponse, error) {
	ret := m.ctrl.Call(m, "GetBucketsByCandidate", arg0, arg1)
	ret0, _ := ret[0].(*api.BucketResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketsByCandidate indicates an expected call of GetBucketsByCandidate
func (mr *MockAPIServiceServerMockRecorder) GetBucketsByCandidate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketsByCandidate", reflect.TypeOf((*MockAPIServiceServer)(nil).GetBucketsByCandidate), arg0, arg1)
}

// IsHealth mocks base method
func (m *MockAPIServiceServer) IsHealth(arg0 context.Context, arg1 *empty.Empty) (*api.HealthCheckResponse, error) {
	ret := m.ctrl.Call(m, "IsHealth", arg0, arg1)
	ret0, _ := ret[0].(*api.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsHealth indicates an expected call of IsHealth
func (mr *MockAPIServiceServerMockRecorder) IsHealth(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHealth", reflect.TypeOf((*MockAPIServiceServer)(nil).IsHealth), arg0, arg1)
}
