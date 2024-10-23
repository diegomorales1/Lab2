// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: Digimon.proto

package digimon

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
	DigimonService_SendDigimon_FullMethodName      = "/digimon.DigimonService/SendDigimon"
	DigimonService_NotifyCompletion_FullMethodName = "/digimon.DigimonService/NotifyCompletion"
)

// DigimonServiceClient is the client API for DigimonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DigimonServiceClient interface {
	// Método para enviar Digimons desde el Primary Node a los nodos regionales
	SendDigimon(ctx context.Context, in *Digimon, opts ...grpc.CallOption) (*Ack, error)
	// Método para notificar a los nodos regionales que el Primary Node ha terminado de enviar Digimons y enviar PS
	NotifyCompletion(ctx context.Context, in *CompletionInfo, opts ...grpc.CallOption) (*Ack, error)
}

type digimonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDigimonServiceClient(cc grpc.ClientConnInterface) DigimonServiceClient {
	return &digimonServiceClient{cc}
}

func (c *digimonServiceClient) SendDigimon(ctx context.Context, in *Digimon, opts ...grpc.CallOption) (*Ack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ack)
	err := c.cc.Invoke(ctx, DigimonService_SendDigimon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *digimonServiceClient) NotifyCompletion(ctx context.Context, in *CompletionInfo, opts ...grpc.CallOption) (*Ack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ack)
	err := c.cc.Invoke(ctx, DigimonService_NotifyCompletion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DigimonServiceServer is the server API for DigimonService service.
// All implementations must embed UnimplementedDigimonServiceServer
// for forward compatibility.
type DigimonServiceServer interface {
	// Método para enviar Digimons desde el Primary Node a los nodos regionales
	SendDigimon(context.Context, *Digimon) (*Ack, error)
	// Método para notificar a los nodos regionales que el Primary Node ha terminado de enviar Digimons y enviar PS
	NotifyCompletion(context.Context, *CompletionInfo) (*Ack, error)
	mustEmbedUnimplementedDigimonServiceServer()
}

// UnimplementedDigimonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDigimonServiceServer struct{}

func (UnimplementedDigimonServiceServer) SendDigimon(context.Context, *Digimon) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDigimon not implemented")
}
func (UnimplementedDigimonServiceServer) NotifyCompletion(context.Context, *CompletionInfo) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyCompletion not implemented")
}
func (UnimplementedDigimonServiceServer) mustEmbedUnimplementedDigimonServiceServer() {}
func (UnimplementedDigimonServiceServer) testEmbeddedByValue()                        {}

// UnsafeDigimonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DigimonServiceServer will
// result in compilation errors.
type UnsafeDigimonServiceServer interface {
	mustEmbedUnimplementedDigimonServiceServer()
}

func RegisterDigimonServiceServer(s grpc.ServiceRegistrar, srv DigimonServiceServer) {
	// If the following call pancis, it indicates UnimplementedDigimonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DigimonService_ServiceDesc, srv)
}

func _DigimonService_SendDigimon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Digimon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigimonServiceServer).SendDigimon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DigimonService_SendDigimon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigimonServiceServer).SendDigimon(ctx, req.(*Digimon))
	}
	return interceptor(ctx, in, info, handler)
}

func _DigimonService_NotifyCompletion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigimonServiceServer).NotifyCompletion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DigimonService_NotifyCompletion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigimonServiceServer).NotifyCompletion(ctx, req.(*CompletionInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// DigimonService_ServiceDesc is the grpc.ServiceDesc for DigimonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DigimonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "digimon.DigimonService",
	HandlerType: (*DigimonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendDigimon",
			Handler:    _DigimonService_SendDigimon_Handler,
		},
		{
			MethodName: "NotifyCompletion",
			Handler:    _DigimonService_NotifyCompletion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Digimon.proto",
}

const (
	ResultService_SendSacrificioResult_FullMethodName = "/digimon.ResultService/SendSacrificioResult"
)

// ResultServiceClient is the client API for ResultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResultServiceClient interface {
	// Método para que los nodos regionales envíen los resultados de sacrificio al Primary Node
	SendSacrificioResult(ctx context.Context, in *SacrificioResult, opts ...grpc.CallOption) (*Ack, error)
}

type resultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResultServiceClient(cc grpc.ClientConnInterface) ResultServiceClient {
	return &resultServiceClient{cc}
}

func (c *resultServiceClient) SendSacrificioResult(ctx context.Context, in *SacrificioResult, opts ...grpc.CallOption) (*Ack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ack)
	err := c.cc.Invoke(ctx, ResultService_SendSacrificioResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultServiceServer is the server API for ResultService service.
// All implementations must embed UnimplementedResultServiceServer
// for forward compatibility.
type ResultServiceServer interface {
	// Método para que los nodos regionales envíen los resultados de sacrificio al Primary Node
	SendSacrificioResult(context.Context, *SacrificioResult) (*Ack, error)
	mustEmbedUnimplementedResultServiceServer()
}

// UnimplementedResultServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedResultServiceServer struct{}

func (UnimplementedResultServiceServer) SendSacrificioResult(context.Context, *SacrificioResult) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSacrificioResult not implemented")
}
func (UnimplementedResultServiceServer) mustEmbedUnimplementedResultServiceServer() {}
func (UnimplementedResultServiceServer) testEmbeddedByValue()                       {}

// UnsafeResultServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResultServiceServer will
// result in compilation errors.
type UnsafeResultServiceServer interface {
	mustEmbedUnimplementedResultServiceServer()
}

func RegisterResultServiceServer(s grpc.ServiceRegistrar, srv ResultServiceServer) {
	// If the following call pancis, it indicates UnimplementedResultServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ResultService_ServiceDesc, srv)
}

func _ResultService_SendSacrificioResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SacrificioResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServiceServer).SendSacrificioResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResultService_SendSacrificioResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServiceServer).SendSacrificioResult(ctx, req.(*SacrificioResult))
	}
	return interceptor(ctx, in, info, handler)
}

// ResultService_ServiceDesc is the grpc.ServiceDesc for ResultService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResultService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "digimon.ResultService",
	HandlerType: (*ResultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSacrificioResult",
			Handler:    _ResultService_SendSacrificioResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Digimon.proto",
}