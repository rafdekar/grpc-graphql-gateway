// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: starwars/starwars.proto

package starwars

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
	StartwarsService_GetHero_FullMethodName    = "/startwars.StartwarsService/GetHero"
	StartwarsService_GetHuman_FullMethodName   = "/startwars.StartwarsService/GetHuman"
	StartwarsService_GetDroid_FullMethodName   = "/startwars.StartwarsService/GetDroid"
	StartwarsService_ListHumans_FullMethodName = "/startwars.StartwarsService/ListHumans"
	StartwarsService_ListDroids_FullMethodName = "/startwars.StartwarsService/ListDroids"
)

// StartwarsServiceClient is the client API for StartwarsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StartwarsServiceClient interface {
	GetHero(ctx context.Context, in *GetHeroRequest, opts ...grpc.CallOption) (*Character, error)
	GetHuman(ctx context.Context, in *GetHumanRequest, opts ...grpc.CallOption) (*Character, error)
	GetDroid(ctx context.Context, in *GetDroidRequest, opts ...grpc.CallOption) (*Character, error)
	ListHumans(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListHumansResponse, error)
	ListDroids(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListDroidsResponse, error)
}

type startwarsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStartwarsServiceClient(cc grpc.ClientConnInterface) StartwarsServiceClient {
	return &startwarsServiceClient{cc}
}

func (c *startwarsServiceClient) GetHero(ctx context.Context, in *GetHeroRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, StartwarsService_GetHero_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) GetHuman(ctx context.Context, in *GetHumanRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, StartwarsService_GetHuman_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) GetDroid(ctx context.Context, in *GetDroidRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, StartwarsService_GetDroid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) ListHumans(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListHumansResponse, error) {
	out := new(ListHumansResponse)
	err := c.cc.Invoke(ctx, StartwarsService_ListHumans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) ListDroids(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListDroidsResponse, error) {
	out := new(ListDroidsResponse)
	err := c.cc.Invoke(ctx, StartwarsService_ListDroids_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StartwarsServiceServer is the server API for StartwarsService service.
// All implementations must embed UnimplementedStartwarsServiceServer
// for forward compatibility
type StartwarsServiceServer interface {
	GetHero(context.Context, *GetHeroRequest) (*Character, error)
	GetHuman(context.Context, *GetHumanRequest) (*Character, error)
	GetDroid(context.Context, *GetDroidRequest) (*Character, error)
	ListHumans(context.Context, *ListEmptyRequest) (*ListHumansResponse, error)
	ListDroids(context.Context, *ListEmptyRequest) (*ListDroidsResponse, error)
	mustEmbedUnimplementedStartwarsServiceServer()
}

// UnimplementedStartwarsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStartwarsServiceServer struct {
}

func (UnimplementedStartwarsServiceServer) GetHero(context.Context, *GetHeroRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHero not implemented")
}
func (UnimplementedStartwarsServiceServer) GetHuman(context.Context, *GetHumanRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHuman not implemented")
}
func (UnimplementedStartwarsServiceServer) GetDroid(context.Context, *GetDroidRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDroid not implemented")
}
func (UnimplementedStartwarsServiceServer) ListHumans(context.Context, *ListEmptyRequest) (*ListHumansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHumans not implemented")
}
func (UnimplementedStartwarsServiceServer) ListDroids(context.Context, *ListEmptyRequest) (*ListDroidsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDroids not implemented")
}
func (UnimplementedStartwarsServiceServer) mustEmbedUnimplementedStartwarsServiceServer() {}

// UnsafeStartwarsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StartwarsServiceServer will
// result in compilation errors.
type UnsafeStartwarsServiceServer interface {
	mustEmbedUnimplementedStartwarsServiceServer()
}

func RegisterStartwarsServiceServer(s grpc.ServiceRegistrar, srv StartwarsServiceServer) {
	s.RegisterService(&StartwarsService_ServiceDesc, srv)
}

func _StartwarsService_GetHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).GetHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StartwarsService_GetHero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).GetHero(ctx, req.(*GetHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_GetHuman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHumanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).GetHuman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StartwarsService_GetHuman_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).GetHuman(ctx, req.(*GetHumanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_GetDroid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDroidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).GetDroid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StartwarsService_GetDroid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).GetDroid(ctx, req.(*GetDroidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_ListHumans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).ListHumans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StartwarsService_ListHumans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).ListHumans(ctx, req.(*ListEmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_ListDroids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).ListDroids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StartwarsService_ListDroids_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).ListDroids(ctx, req.(*ListEmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StartwarsService_ServiceDesc is the grpc.ServiceDesc for StartwarsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StartwarsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "startwars.StartwarsService",
	HandlerType: (*StartwarsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHero",
			Handler:    _StartwarsService_GetHero_Handler,
		},
		{
			MethodName: "GetHuman",
			Handler:    _StartwarsService_GetHuman_Handler,
		},
		{
			MethodName: "GetDroid",
			Handler:    _StartwarsService_GetDroid_Handler,
		},
		{
			MethodName: "ListHumans",
			Handler:    _StartwarsService_ListHumans_Handler,
		},
		{
			MethodName: "ListDroids",
			Handler:    _StartwarsService_ListDroids_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "starwars/starwars.proto",
}
