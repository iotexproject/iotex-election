// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIServiceClient interface {
	// get the blockchain meta data
	GetMeta(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChainMeta, error)
	// get candidates
	GetCandidates(ctx context.Context, in *GetCandidatesRequest, opts ...grpc.CallOption) (*CandidateResponse, error)
	// get candidate by name
	GetCandidateByName(ctx context.Context, in *GetCandidateByNameRequest, opts ...grpc.CallOption) (*Candidate, error)
	// get buckets by candidate
	GetBucketsByCandidate(ctx context.Context, in *GetBucketsByCandidateRequest, opts ...grpc.CallOption) (*BucketResponse, error)
	// get Buckets
	GetBuckets(ctx context.Context, in *GetBucketsRequest, opts ...grpc.CallOption) (*BucketResponse, error)
	// health endpoint
	IsHealth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	// get raw data by height
	GetRawData(ctx context.Context, in *GetRawDataRequest, opts ...grpc.CallOption) (*RawDataResponse, error)
	// get proof for a given account
	GetProof(ctx context.Context, in *ProofRequest, opts ...grpc.CallOption) (*ProofResponse, error)
}

type aPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIServiceClient(cc grpc.ClientConnInterface) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) GetMeta(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChainMeta, error) {
	out := new(ChainMeta)
	err := c.cc.Invoke(ctx, "/api.APIService/getMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetCandidates(ctx context.Context, in *GetCandidatesRequest, opts ...grpc.CallOption) (*CandidateResponse, error) {
	out := new(CandidateResponse)
	err := c.cc.Invoke(ctx, "/api.APIService/getCandidates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetCandidateByName(ctx context.Context, in *GetCandidateByNameRequest, opts ...grpc.CallOption) (*Candidate, error) {
	out := new(Candidate)
	err := c.cc.Invoke(ctx, "/api.APIService/getCandidateByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetBucketsByCandidate(ctx context.Context, in *GetBucketsByCandidateRequest, opts ...grpc.CallOption) (*BucketResponse, error) {
	out := new(BucketResponse)
	err := c.cc.Invoke(ctx, "/api.APIService/getBucketsByCandidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetBuckets(ctx context.Context, in *GetBucketsRequest, opts ...grpc.CallOption) (*BucketResponse, error) {
	out := new(BucketResponse)
	err := c.cc.Invoke(ctx, "/api.APIService/getBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) IsHealth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/api.APIService/isHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetRawData(ctx context.Context, in *GetRawDataRequest, opts ...grpc.CallOption) (*RawDataResponse, error) {
	out := new(RawDataResponse)
	err := c.cc.Invoke(ctx, "/api.APIService/getRawData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) GetProof(ctx context.Context, in *ProofRequest, opts ...grpc.CallOption) (*ProofResponse, error) {
	out := new(ProofResponse)
	err := c.cc.Invoke(ctx, "/api.APIService/getProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
// All implementations must embed UnimplementedAPIServiceServer
// for forward compatibility
type APIServiceServer interface {
	// get the blockchain meta data
	GetMeta(context.Context, *emptypb.Empty) (*ChainMeta, error)
	// get candidates
	GetCandidates(context.Context, *GetCandidatesRequest) (*CandidateResponse, error)
	// get candidate by name
	GetCandidateByName(context.Context, *GetCandidateByNameRequest) (*Candidate, error)
	// get buckets by candidate
	GetBucketsByCandidate(context.Context, *GetBucketsByCandidateRequest) (*BucketResponse, error)
	// get Buckets
	GetBuckets(context.Context, *GetBucketsRequest) (*BucketResponse, error)
	// health endpoint
	IsHealth(context.Context, *emptypb.Empty) (*HealthCheckResponse, error)
	// get raw data by height
	GetRawData(context.Context, *GetRawDataRequest) (*RawDataResponse, error)
	// get proof for a given account
	GetProof(context.Context, *ProofRequest) (*ProofResponse, error)
	mustEmbedUnimplementedAPIServiceServer()
}

// UnimplementedAPIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServiceServer struct {
}

func (UnimplementedAPIServiceServer) GetMeta(context.Context, *emptypb.Empty) (*ChainMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeta not implemented")
}
func (UnimplementedAPIServiceServer) GetCandidates(context.Context, *GetCandidatesRequest) (*CandidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCandidates not implemented")
}
func (UnimplementedAPIServiceServer) GetCandidateByName(context.Context, *GetCandidateByNameRequest) (*Candidate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCandidateByName not implemented")
}
func (UnimplementedAPIServiceServer) GetBucketsByCandidate(context.Context, *GetBucketsByCandidateRequest) (*BucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBucketsByCandidate not implemented")
}
func (UnimplementedAPIServiceServer) GetBuckets(context.Context, *GetBucketsRequest) (*BucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuckets not implemented")
}
func (UnimplementedAPIServiceServer) IsHealth(context.Context, *emptypb.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHealth not implemented")
}
func (UnimplementedAPIServiceServer) GetRawData(context.Context, *GetRawDataRequest) (*RawDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRawData not implemented")
}
func (UnimplementedAPIServiceServer) GetProof(context.Context, *ProofRequest) (*ProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProof not implemented")
}
func (UnimplementedAPIServiceServer) mustEmbedUnimplementedAPIServiceServer() {}

// UnsafeAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServiceServer will
// result in compilation errors.
type UnsafeAPIServiceServer interface {
	mustEmbedUnimplementedAPIServiceServer()
}

func RegisterAPIServiceServer(s *grpc.Server, srv APIServiceServer) {
	s.RegisterService(&_APIService_serviceDesc, srv)
}

func _APIService_GetMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/GetMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetMeta(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetCandidates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCandidatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetCandidates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/GetCandidates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetCandidates(ctx, req.(*GetCandidatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetCandidateByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCandidateByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetCandidateByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/GetCandidateByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetCandidateByName(ctx, req.(*GetCandidateByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetBucketsByCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBucketsByCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetBucketsByCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/GetBucketsByCandidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetBucketsByCandidate(ctx, req.(*GetBucketsByCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/GetBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetBuckets(ctx, req.(*GetBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_IsHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).IsHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/IsHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).IsHealth(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetRawData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRawDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetRawData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/GetRawData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetRawData(ctx, req.(*GetRawDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_GetProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/GetProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetProof(ctx, req.(*ProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _APIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMeta",
			Handler:    _APIService_GetMeta_Handler,
		},
		{
			MethodName: "getCandidates",
			Handler:    _APIService_GetCandidates_Handler,
		},
		{
			MethodName: "getCandidateByName",
			Handler:    _APIService_GetCandidateByName_Handler,
		},
		{
			MethodName: "getBucketsByCandidate",
			Handler:    _APIService_GetBucketsByCandidate_Handler,
		},
		{
			MethodName: "getBuckets",
			Handler:    _APIService_GetBuckets_Handler,
		},
		{
			MethodName: "isHealth",
			Handler:    _APIService_IsHealth_Handler,
		},
		{
			MethodName: "getRawData",
			Handler:    _APIService_GetRawData_Handler,
		},
		{
			MethodName: "getProof",
			Handler:    _APIService_GetProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/api/api.proto",
}
