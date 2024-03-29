// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package answerepb

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

// AnswereServiceClient is the client API for AnswereService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnswereServiceClient interface {
	AnswereCreate(ctx context.Context, in *AnswereCreateRequest, opts ...grpc.CallOption) (*AnswereCreateResponse, error)
	AnswereList(ctx context.Context, in *AnswereListRequest, opts ...grpc.CallOption) (*AnswereListResponse, error)
	AnswereDelete(ctx context.Context, in *AnswereDeleteRequest, opts ...grpc.CallOption) (*AnswereDeleteResponse, error)
	AnswereEdit(ctx context.Context, in *AnswereEditRequest, opts ...grpc.CallOption) (*AnswereEditResponse, error)
	AnswereUpdate(ctx context.Context, in *AnswereUpdateRequest, opts ...grpc.CallOption) (*AnswereUpdateResponse, error)
	CorrectAnswere(ctx context.Context, in *CorrectAnswereRequest, opts ...grpc.CallOption) (*CorrectAnswereResponse, error)
}

type answereServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnswereServiceClient(cc grpc.ClientConnInterface) AnswereServiceClient {
	return &answereServiceClient{cc}
}

func (c *answereServiceClient) AnswereCreate(ctx context.Context, in *AnswereCreateRequest, opts ...grpc.CallOption) (*AnswereCreateResponse, error) {
	out := new(AnswereCreateResponse)
	err := c.cc.Invoke(ctx, "/answerepb.AnswereService/AnswereCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answereServiceClient) AnswereList(ctx context.Context, in *AnswereListRequest, opts ...grpc.CallOption) (*AnswereListResponse, error) {
	out := new(AnswereListResponse)
	err := c.cc.Invoke(ctx, "/answerepb.AnswereService/AnswereList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answereServiceClient) AnswereDelete(ctx context.Context, in *AnswereDeleteRequest, opts ...grpc.CallOption) (*AnswereDeleteResponse, error) {
	out := new(AnswereDeleteResponse)
	err := c.cc.Invoke(ctx, "/answerepb.AnswereService/AnswereDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answereServiceClient) AnswereEdit(ctx context.Context, in *AnswereEditRequest, opts ...grpc.CallOption) (*AnswereEditResponse, error) {
	out := new(AnswereEditResponse)
	err := c.cc.Invoke(ctx, "/answerepb.AnswereService/AnswereEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answereServiceClient) AnswereUpdate(ctx context.Context, in *AnswereUpdateRequest, opts ...grpc.CallOption) (*AnswereUpdateResponse, error) {
	out := new(AnswereUpdateResponse)
	err := c.cc.Invoke(ctx, "/answerepb.AnswereService/AnswereUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answereServiceClient) CorrectAnswere(ctx context.Context, in *CorrectAnswereRequest, opts ...grpc.CallOption) (*CorrectAnswereResponse, error) {
	out := new(CorrectAnswereResponse)
	err := c.cc.Invoke(ctx, "/answerepb.AnswereService/CorrectAnswere", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnswereServiceServer is the server API for AnswereService service.
// All implementations must embed UnimplementedAnswereServiceServer
// for forward compatibility
type AnswereServiceServer interface {
	AnswereCreate(context.Context, *AnswereCreateRequest) (*AnswereCreateResponse, error)
	AnswereList(context.Context, *AnswereListRequest) (*AnswereListResponse, error)
	AnswereDelete(context.Context, *AnswereDeleteRequest) (*AnswereDeleteResponse, error)
	AnswereEdit(context.Context, *AnswereEditRequest) (*AnswereEditResponse, error)
	AnswereUpdate(context.Context, *AnswereUpdateRequest) (*AnswereUpdateResponse, error)
	CorrectAnswere(context.Context, *CorrectAnswereRequest) (*CorrectAnswereResponse, error)
	mustEmbedUnimplementedAnswereServiceServer()
}

// UnimplementedAnswereServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnswereServiceServer struct {
}

func (UnimplementedAnswereServiceServer) AnswereCreate(context.Context, *AnswereCreateRequest) (*AnswereCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswereCreate not implemented")
}
func (UnimplementedAnswereServiceServer) AnswereList(context.Context, *AnswereListRequest) (*AnswereListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswereList not implemented")
}
func (UnimplementedAnswereServiceServer) AnswereDelete(context.Context, *AnswereDeleteRequest) (*AnswereDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswereDelete not implemented")
}
func (UnimplementedAnswereServiceServer) AnswereEdit(context.Context, *AnswereEditRequest) (*AnswereEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswereEdit not implemented")
}
func (UnimplementedAnswereServiceServer) AnswereUpdate(context.Context, *AnswereUpdateRequest) (*AnswereUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswereUpdate not implemented")
}
func (UnimplementedAnswereServiceServer) CorrectAnswere(context.Context, *CorrectAnswereRequest) (*CorrectAnswereResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CorrectAnswere not implemented")
}
func (UnimplementedAnswereServiceServer) mustEmbedUnimplementedAnswereServiceServer() {}

// UnsafeAnswereServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnswereServiceServer will
// result in compilation errors.
type UnsafeAnswereServiceServer interface {
	mustEmbedUnimplementedAnswereServiceServer()
}

func RegisterAnswereServiceServer(s grpc.ServiceRegistrar, srv AnswereServiceServer) {
	s.RegisterService(&AnswereService_ServiceDesc, srv)
}

func _AnswereService_AnswereCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswereCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswereServiceServer).AnswereCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answerepb.AnswereService/AnswereCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswereServiceServer).AnswereCreate(ctx, req.(*AnswereCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswereService_AnswereList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswereListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswereServiceServer).AnswereList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answerepb.AnswereService/AnswereList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswereServiceServer).AnswereList(ctx, req.(*AnswereListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswereService_AnswereDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswereDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswereServiceServer).AnswereDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answerepb.AnswereService/AnswereDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswereServiceServer).AnswereDelete(ctx, req.(*AnswereDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswereService_AnswereEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswereEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswereServiceServer).AnswereEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answerepb.AnswereService/AnswereEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswereServiceServer).AnswereEdit(ctx, req.(*AnswereEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswereService_AnswereUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswereUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswereServiceServer).AnswereUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answerepb.AnswereService/AnswereUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswereServiceServer).AnswereUpdate(ctx, req.(*AnswereUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswereService_CorrectAnswere_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CorrectAnswereRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswereServiceServer).CorrectAnswere(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answerepb.AnswereService/CorrectAnswere",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswereServiceServer).CorrectAnswere(ctx, req.(*CorrectAnswereRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnswereService_ServiceDesc is the grpc.ServiceDesc for AnswereService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnswereService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "answerepb.AnswereService",
	HandlerType: (*AnswereServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnswereCreate",
			Handler:    _AnswereService_AnswereCreate_Handler,
		},
		{
			MethodName: "AnswereList",
			Handler:    _AnswereService_AnswereList_Handler,
		},
		{
			MethodName: "AnswereDelete",
			Handler:    _AnswereService_AnswereDelete_Handler,
		},
		{
			MethodName: "AnswereEdit",
			Handler:    _AnswereService_AnswereEdit_Handler,
		},
		{
			MethodName: "AnswereUpdate",
			Handler:    _AnswereService_AnswereUpdate_Handler,
		},
		{
			MethodName: "CorrectAnswere",
			Handler:    _AnswereService_CorrectAnswere_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stackoverflow/gunk/v1/answere/all.proto",
}
