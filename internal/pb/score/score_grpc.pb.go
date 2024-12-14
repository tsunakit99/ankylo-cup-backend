// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: score.proto

package score

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ScoreService_RecordScore_FullMethodName          = "/score.ScoreService/RecordScore"
	ScoreService_GetTop10ScoresByGame_FullMethodName = "/score.ScoreService/GetTop10ScoresByGame"
)

// ScoreServiceClient is the client API for ScoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScoreServiceClient interface {
	RecordScore(ctx context.Context, in *RecordScoreRequest, opts ...grpc.CallOption) (*RecordScoreResponse, error)
	GetTop10ScoresByGame(ctx context.Context, in *GetTop10ScoresByGameRequest, opts ...grpc.CallOption) (*GetTop10ScoresByGameResponse, error)
}

type scoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScoreServiceClient(cc grpc.ClientConnInterface) ScoreServiceClient {
	return &scoreServiceClient{cc}
}

func (c *scoreServiceClient) RecordScore(ctx context.Context, in *RecordScoreRequest, opts ...grpc.CallOption) (*RecordScoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecordScoreResponse)
	err := c.cc.Invoke(ctx, ScoreService_RecordScore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreServiceClient) GetTop10ScoresByGame(ctx context.Context, in *GetTop10ScoresByGameRequest, opts ...grpc.CallOption) (*GetTop10ScoresByGameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTop10ScoresByGameResponse)
	err := c.cc.Invoke(ctx, ScoreService_GetTop10ScoresByGame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoreServiceServer is the server API for ScoreService service.
// All implementations must embed UnimplementedScoreServiceServer
// for forward compatibility.
type ScoreServiceServer interface {
	RecordScore(context.Context, *RecordScoreRequest) (*RecordScoreResponse, error)
	GetTop10ScoresByGame(context.Context, *GetTop10ScoresByGameRequest) (*GetTop10ScoresByGameResponse, error)
	mustEmbedUnimplementedScoreServiceServer()
}

// UnimplementedScoreServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedScoreServiceServer struct{}

func (UnimplementedScoreServiceServer) RecordScore(context.Context, *RecordScoreRequest) (*RecordScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordScore not implemented")
}
func (UnimplementedScoreServiceServer) GetTop10ScoresByGame(context.Context, *GetTop10ScoresByGameRequest) (*GetTop10ScoresByGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTop10ScoresByGame not implemented")
}
func (UnimplementedScoreServiceServer) mustEmbedUnimplementedScoreServiceServer() {}
func (UnimplementedScoreServiceServer) testEmbeddedByValue()                      {}

// UnsafeScoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScoreServiceServer will
// result in compilation errors.
type UnsafeScoreServiceServer interface {
	mustEmbedUnimplementedScoreServiceServer()
}

func RegisterScoreServiceServer(s grpc.ServiceRegistrar, srv ScoreServiceServer) {
	// If the following call pancis, it indicates UnimplementedScoreServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ScoreService_ServiceDesc, srv)
}

func _ScoreService_RecordScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServiceServer).RecordScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreService_RecordScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServiceServer).RecordScore(ctx, req.(*RecordScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoreService_GetTop10ScoresByGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTop10ScoresByGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServiceServer).GetTop10ScoresByGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScoreService_GetTop10ScoresByGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServiceServer).GetTop10ScoresByGame(ctx, req.(*GetTop10ScoresByGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScoreService_ServiceDesc is the grpc.ServiceDesc for ScoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "score.ScoreService",
	HandlerType: (*ScoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordScore",
			Handler:    _ScoreService_RecordScore_Handler,
		},
		{
			MethodName: "GetTop10ScoresByGame",
			Handler:    _ScoreService_GetTop10ScoresByGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "score.proto",
}