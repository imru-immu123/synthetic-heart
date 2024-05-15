// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: syntest.proto

package proto

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

// SynTestPluginClient is the client API for SynTestPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SynTestPluginClient interface {
	// Called once at the start - for setup
	Initialise(ctx context.Context, in *SynTestConfig, opts ...grpc.CallOption) (*Empty, error)
	// Called periodically or when another test finishes - as configured in the synthetic test config
	// TestRun contains details from another test that the current test might depend on, otherwise empty
	PerformTest(ctx context.Context, in *Trigger, opts ...grpc.CallOption) (*TestResult, error)
	// Called once before the plugin is killed
	Finish(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type synTestPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewSynTestPluginClient(cc grpc.ClientConnInterface) SynTestPluginClient {
	return &synTestPluginClient{cc}
}

func (c *synTestPluginClient) Initialise(ctx context.Context, in *SynTestConfig, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.syntest.SynTestPlugin/Initialise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synTestPluginClient) PerformTest(ctx context.Context, in *Trigger, opts ...grpc.CallOption) (*TestResult, error) {
	out := new(TestResult)
	err := c.cc.Invoke(ctx, "/proto.syntest.SynTestPlugin/PerformTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synTestPluginClient) Finish(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.syntest.SynTestPlugin/Finish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SynTestPluginServer is the server API for SynTestPlugin service.
// All implementations must embed UnimplementedSynTestPluginServer
// for forward compatibility
type SynTestPluginServer interface {
	// Called once at the start - for setup
	Initialise(context.Context, *SynTestConfig) (*Empty, error)
	// Called periodically or when another test finishes - as configured in the synthetic test config
	// TestRun contains details from another test that the current test might depend on, otherwise empty
	PerformTest(context.Context, *Trigger) (*TestResult, error)
	// Called once before the plugin is killed
	Finish(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedSynTestPluginServer()
}

// UnimplementedSynTestPluginServer must be embedded to have forward compatible implementations.
type UnimplementedSynTestPluginServer struct {
}

func (UnimplementedSynTestPluginServer) Initialise(context.Context, *SynTestConfig) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initialise not implemented")
}
func (UnimplementedSynTestPluginServer) PerformTest(context.Context, *Trigger) (*TestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformTest not implemented")
}
func (UnimplementedSynTestPluginServer) Finish(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finish not implemented")
}
func (UnimplementedSynTestPluginServer) mustEmbedUnimplementedSynTestPluginServer() {}

// UnsafeSynTestPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SynTestPluginServer will
// result in compilation errors.
type UnsafeSynTestPluginServer interface {
	mustEmbedUnimplementedSynTestPluginServer()
}

func RegisterSynTestPluginServer(s grpc.ServiceRegistrar, srv SynTestPluginServer) {
	s.RegisterService(&SynTestPlugin_ServiceDesc, srv)
}

func _SynTestPlugin_Initialise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SynTestConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynTestPluginServer).Initialise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.syntest.SynTestPlugin/Initialise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynTestPluginServer).Initialise(ctx, req.(*SynTestConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynTestPlugin_PerformTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trigger)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynTestPluginServer).PerformTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.syntest.SynTestPlugin/PerformTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynTestPluginServer).PerformTest(ctx, req.(*Trigger))
	}
	return interceptor(ctx, in, info, handler)
}

func _SynTestPlugin_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynTestPluginServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.syntest.SynTestPlugin/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynTestPluginServer).Finish(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SynTestPlugin_ServiceDesc is the grpc.ServiceDesc for SynTestPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SynTestPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.syntest.SynTestPlugin",
	HandlerType: (*SynTestPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Initialise",
			Handler:    _SynTestPlugin_Initialise_Handler,
		},
		{
			MethodName: "PerformTest",
			Handler:    _SynTestPlugin_PerformTest_Handler,
		},
		{
			MethodName: "Finish",
			Handler:    _SynTestPlugin_Finish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syntest.proto",
}
