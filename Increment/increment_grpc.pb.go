// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package increment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IncrementServiceClient is the client API for IncrementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IncrementServiceClient interface {
	// Client methods.
	Increment(ctx context.Context, in *VoidMessage, opts ...grpc.CallOption) (*IncrementMessage, error)
	GetLeader(ctx context.Context, in *VoidMessage, opts ...grpc.CallOption) (*LeaderMessage, error)
	GetReplicas(ctx context.Context, in *VoidMessage, opts ...grpc.CallOption) (*ReplicaListMessage, error)
	//  Replica methods.
	HeartBeat(ctx context.Context, in *VoidMessage, opts ...grpc.CallOption) (*VoidMessage, error)
	// Sub-replica methods.
	Join(ctx context.Context, in *IpMessage, opts ...grpc.CallOption) (*VoidMessage, error)
	// Leader-replica methods.
	SendReplicas(ctx context.Context, in *ReplicaListMessage, opts ...grpc.CallOption) (*VoidMessage, error)
	SendValue(ctx context.Context, in *IncrementMessage, opts ...grpc.CallOption) (*VoidMessage, error)
}

type incrementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIncrementServiceClient(cc grpc.ClientConnInterface) IncrementServiceClient {
	return &incrementServiceClient{cc}
}

func (c *incrementServiceClient) Increment(ctx context.Context, in *VoidMessage, opts ...grpc.CallOption) (*IncrementMessage, error) {
	out := new(IncrementMessage)
	err := c.cc.Invoke(ctx, "/increment.IncrementService/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementServiceClient) GetLeader(ctx context.Context, in *VoidMessage, opts ...grpc.CallOption) (*LeaderMessage, error) {
	out := new(LeaderMessage)
	err := c.cc.Invoke(ctx, "/increment.IncrementService/GetLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementServiceClient) GetReplicas(ctx context.Context, in *VoidMessage, opts ...grpc.CallOption) (*ReplicaListMessage, error) {
	out := new(ReplicaListMessage)
	err := c.cc.Invoke(ctx, "/increment.IncrementService/GetReplicas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementServiceClient) HeartBeat(ctx context.Context, in *VoidMessage, opts ...grpc.CallOption) (*VoidMessage, error) {
	out := new(VoidMessage)
	err := c.cc.Invoke(ctx, "/increment.IncrementService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementServiceClient) Join(ctx context.Context, in *IpMessage, opts ...grpc.CallOption) (*VoidMessage, error) {
	out := new(VoidMessage)
	err := c.cc.Invoke(ctx, "/increment.IncrementService/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementServiceClient) SendReplicas(ctx context.Context, in *ReplicaListMessage, opts ...grpc.CallOption) (*VoidMessage, error) {
	out := new(VoidMessage)
	err := c.cc.Invoke(ctx, "/increment.IncrementService/SendReplicas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementServiceClient) SendValue(ctx context.Context, in *IncrementMessage, opts ...grpc.CallOption) (*VoidMessage, error) {
	out := new(VoidMessage)
	err := c.cc.Invoke(ctx, "/increment.IncrementService/SendValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncrementServiceServer is the server API for IncrementService service.
// All implementations must embed UnimplementedIncrementServiceServer
// for forward compatibility
type IncrementServiceServer interface {
	// Client methods.
	Increment(context.Context, *VoidMessage) (*IncrementMessage, error)
	GetLeader(context.Context, *VoidMessage) (*LeaderMessage, error)
	GetReplicas(context.Context, *VoidMessage) (*ReplicaListMessage, error)
	//  Replica methods.
	HeartBeat(context.Context, *VoidMessage) (*VoidMessage, error)
	// Sub-replica methods.
	Join(context.Context, *IpMessage) (*VoidMessage, error)
	// Leader-replica methods.
	SendReplicas(context.Context, *ReplicaListMessage) (*VoidMessage, error)
	SendValue(context.Context, *IncrementMessage) (*VoidMessage, error)
	mustEmbedUnimplementedIncrementServiceServer()
}

// UnimplementedIncrementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIncrementServiceServer struct {
}

func (UnimplementedIncrementServiceServer) Increment(context.Context, *VoidMessage) (*IncrementMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedIncrementServiceServer) GetLeader(context.Context, *VoidMessage) (*LeaderMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeader not implemented")
}
func (UnimplementedIncrementServiceServer) GetReplicas(context.Context, *VoidMessage) (*ReplicaListMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplicas not implemented")
}
func (UnimplementedIncrementServiceServer) HeartBeat(context.Context, *VoidMessage) (*VoidMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedIncrementServiceServer) Join(context.Context, *IpMessage) (*VoidMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedIncrementServiceServer) SendReplicas(context.Context, *ReplicaListMessage) (*VoidMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReplicas not implemented")
}
func (UnimplementedIncrementServiceServer) SendValue(context.Context, *IncrementMessage) (*VoidMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendValue not implemented")
}
func (UnimplementedIncrementServiceServer) mustEmbedUnimplementedIncrementServiceServer() {}

// UnsafeIncrementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IncrementServiceServer will
// result in compilation errors.
type UnsafeIncrementServiceServer interface {
	mustEmbedUnimplementedIncrementServiceServer()
}

func RegisterIncrementServiceServer(s grpc.ServiceRegistrar, srv IncrementServiceServer) {
	s.RegisterService(&IncrementService_ServiceDesc, srv)
}

func _IncrementService_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementServiceServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementService/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementServiceServer).Increment(ctx, req.(*VoidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementService_GetLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementServiceServer).GetLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementService/GetLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementServiceServer).GetLeader(ctx, req.(*VoidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementService_GetReplicas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementServiceServer).GetReplicas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementService/GetReplicas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementServiceServer).GetReplicas(ctx, req.(*VoidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementServiceServer).HeartBeat(ctx, req.(*VoidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementService/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementServiceServer).Join(ctx, req.(*IpMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementService_SendReplicas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicaListMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementServiceServer).SendReplicas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementService/SendReplicas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementServiceServer).SendReplicas(ctx, req.(*ReplicaListMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementService_SendValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementServiceServer).SendValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementService/SendValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementServiceServer).SendValue(ctx, req.(*IncrementMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// IncrementService_ServiceDesc is the grpc.ServiceDesc for IncrementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IncrementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "increment.IncrementService",
	HandlerType: (*IncrementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increment",
			Handler:    _IncrementService_Increment_Handler,
		},
		{
			MethodName: "GetLeader",
			Handler:    _IncrementService_GetLeader_Handler,
		},
		{
			MethodName: "GetReplicas",
			Handler:    _IncrementService_GetReplicas_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _IncrementService_HeartBeat_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _IncrementService_Join_Handler,
		},
		{
			MethodName: "SendReplicas",
			Handler:    _IncrementService_SendReplicas_Handler,
		},
		{
			MethodName: "SendValue",
			Handler:    _IncrementService_SendValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Increment/increment.proto",
}
