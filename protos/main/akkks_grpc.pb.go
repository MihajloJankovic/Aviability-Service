// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: akkks.proto

package pakets

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

const (
	AccommodationAviability_GetAccommodationCheck_FullMethodName      = "/AccommodationAviability/GetAccommodationCheck"
	AccommodationAviability_GetAllforAccomendation_FullMethodName     = "/AccommodationAviability/GetAllforAccomendation"
	AccommodationAviability_SetAccommodationAviability_FullMethodName = "/AccommodationAviability/SetAccommodationAviability"
)

// AccommodationAviabilityClient is the client API for AccommodationAviability service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationAviabilityClient interface {
	GetAccommodationCheck(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*Emptyb, error)
	GetAllforAccomendation(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*DummyLista3, error)
	SetAccommodationAviability(ctx context.Context, in *CheckSet, opts ...grpc.CallOption) (*Emptyb, error)
}

type accommodationAviabilityClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationAviabilityClient(cc grpc.ClientConnInterface) AccommodationAviabilityClient {
	return &accommodationAviabilityClient{cc}
}

func (c *accommodationAviabilityClient) GetAccommodationCheck(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*Emptyb, error) {
	out := new(Emptyb)
	err := c.cc.Invoke(ctx, AccommodationAviability_GetAccommodationCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationAviabilityClient) GetAllforAccomendation(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*DummyLista3, error) {
	out := new(DummyLista3)
	err := c.cc.Invoke(ctx, AccommodationAviability_GetAllforAccomendation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationAviabilityClient) SetAccommodationAviability(ctx context.Context, in *CheckSet, opts ...grpc.CallOption) (*Emptyb, error) {
	out := new(Emptyb)
	err := c.cc.Invoke(ctx, AccommodationAviability_SetAccommodationAviability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationAviabilityServer is the server API for AccommodationAviability service.
// All implementations must embed UnimplementedAccommodationAviabilityServer
// for forward compatibility
type AccommodationAviabilityServer interface {
	GetAccommodationCheck(context.Context, *CheckRequest) (*Emptyb, error)
	GetAllforAccomendation(context.Context, *GetAllRequest) (*DummyLista3, error)
	SetAccommodationAviability(context.Context, *CheckSet) (*Emptyb, error)
	mustEmbedUnimplementedAccommodationAviabilityServer()
}

// UnimplementedAccommodationAviabilityServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationAviabilityServer struct {
}

func (UnimplementedAccommodationAviabilityServer) GetAccommodationCheck(context.Context, *CheckRequest) (*Emptyb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationCheck not implemented")
}
func (UnimplementedAccommodationAviabilityServer) GetAllforAccomendation(context.Context, *GetAllRequest) (*DummyLista3, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllforAccomendation not implemented")
}
func (UnimplementedAccommodationAviabilityServer) SetAccommodationAviability(context.Context, *CheckSet) (*Emptyb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccommodationAviability not implemented")
}
func (UnimplementedAccommodationAviabilityServer) mustEmbedUnimplementedAccommodationAviabilityServer() {
}

// UnsafeAccommodationAviabilityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationAviabilityServer will
// result in compilation errors.
type UnsafeAccommodationAviabilityServer interface {
	mustEmbedUnimplementedAccommodationAviabilityServer()
}

func RegisterAccommodationAviabilityServer(s grpc.ServiceRegistrar, srv AccommodationAviabilityServer) {
	s.RegisterService(&AccommodationAviability_ServiceDesc, srv)
}

func _AccommodationAviability_GetAccommodationCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationAviabilityServer).GetAccommodationCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationAviability_GetAccommodationCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationAviabilityServer).GetAccommodationCheck(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationAviability_GetAllforAccomendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationAviabilityServer).GetAllforAccomendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationAviability_GetAllforAccomendation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationAviabilityServer).GetAllforAccomendation(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationAviability_SetAccommodationAviability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationAviabilityServer).SetAccommodationAviability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationAviability_SetAccommodationAviability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationAviabilityServer).SetAccommodationAviability(ctx, req.(*CheckSet))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationAviability_ServiceDesc is the grpc.ServiceDesc for AccommodationAviability service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationAviability_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccommodationAviability",
	HandlerType: (*AccommodationAviabilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccommodationCheck",
			Handler:    _AccommodationAviability_GetAccommodationCheck_Handler,
		},
		{
			MethodName: "GetAllforAccomendation",
			Handler:    _AccommodationAviability_GetAllforAccomendation_Handler,
		},
		{
			MethodName: "SetAccommodationAviability",
			Handler:    _AccommodationAviability_SetAccommodationAviability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "akkks.proto",
}
