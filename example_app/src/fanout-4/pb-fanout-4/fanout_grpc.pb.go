// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: pb-fanout-4/fanout.proto

package pb_fanout_4

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

// Fanout_4Client is the client API for Fanout_4 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Fanout_4Client interface {
	GetIds(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type fanout_4Client struct {
	cc grpc.ClientConnInterface
}

func NewFanout_4Client(cc grpc.ClientConnInterface) Fanout_4Client {
	return &fanout_4Client{cc}
}

func (c *fanout_4Client) GetIds(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/fanout_4.fanout_4/GetIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Fanout_4Server is the server API for Fanout_4 service.
// All implementations must embed UnimplementedFanout_4Server
// for forward compatibility
type Fanout_4Server interface {
	GetIds(context.Context, *Req) (*Res, error)
	mustEmbedUnimplementedFanout_4Server()
}

// UnimplementedFanout_4Server must be embedded to have forward compatible implementations.
type UnimplementedFanout_4Server struct {
}

func (UnimplementedFanout_4Server) GetIds(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIds not implemented")
}
func (UnimplementedFanout_4Server) mustEmbedUnimplementedFanout_4Server() {}

// UnsafeFanout_4Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Fanout_4Server will
// result in compilation errors.
type UnsafeFanout_4Server interface {
	mustEmbedUnimplementedFanout_4Server()
}

func RegisterFanout_4Server(s grpc.ServiceRegistrar, srv Fanout_4Server) {
	s.RegisterService(&Fanout_4_ServiceDesc, srv)
}

func _Fanout_4_GetIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fanout_4Server).GetIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fanout_4.fanout_4/GetIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fanout_4Server).GetIds(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Fanout_4_ServiceDesc is the grpc.ServiceDesc for Fanout_4 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fanout_4_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fanout_4.fanout_4",
	HandlerType: (*Fanout_4Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIds",
			Handler:    _Fanout_4_GetIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb-fanout-4/fanout.proto",
}