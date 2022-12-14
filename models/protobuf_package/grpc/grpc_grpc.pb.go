// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: grpc.proto

package godemoGrpc

import (
	context "context"
	employee "godemo/models/protobuf_package/employee"
	search "godemo/models/protobuf_package/search"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GodemoServiceClient is the client API for GodemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GodemoServiceClient interface {
	// 获取职员信息
	GetEmployeeInfo(ctx context.Context, in *employee.EmployeeRequest, opts ...grpc.CallOption) (*employee.EmployeeResponse, error)
	// 搜索请求
	Search(ctx context.Context, in *search.SearchRequest, opts ...grpc.CallOption) (*search.SearchResponse, error)
}

type godemoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGodemoServiceClient(cc grpc.ClientConnInterface) GodemoServiceClient {
	return &godemoServiceClient{cc}
}

func (c *godemoServiceClient) GetEmployeeInfo(ctx context.Context, in *employee.EmployeeRequest, opts ...grpc.CallOption) (*employee.EmployeeResponse, error) {
	out := new(employee.EmployeeResponse)
	err := c.cc.Invoke(ctx, "/GodemoService/getEmployeeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *godemoServiceClient) Search(ctx context.Context, in *search.SearchRequest, opts ...grpc.CallOption) (*search.SearchResponse, error) {
	out := new(search.SearchResponse)
	err := c.cc.Invoke(ctx, "/GodemoService/search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GodemoServiceServer is the server API for GodemoService service.
// All implementations must embed UnimplementedGodemoServiceServer
// for forward compatibility
type GodemoServiceServer interface {
	// 获取职员信息
	GetEmployeeInfo(context.Context, *employee.EmployeeRequest) (*employee.EmployeeResponse, error)
	// 搜索请求
	Search(context.Context, *search.SearchRequest) (*search.SearchResponse, error)
	mustEmbedUnimplementedGodemoServiceServer()
}

// UnimplementedGodemoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGodemoServiceServer struct {
}

func (UnimplementedGodemoServiceServer) GetEmployeeInfo(context.Context, *employee.EmployeeRequest) (*employee.EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeInfo not implemented")
}
func (UnimplementedGodemoServiceServer) Search(context.Context, *search.SearchRequest) (*search.SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedGodemoServiceServer) mustEmbedUnimplementedGodemoServiceServer() {}

// UnsafeGodemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GodemoServiceServer will
// result in compilation errors.
type UnsafeGodemoServiceServer interface {
	mustEmbedUnimplementedGodemoServiceServer()
}

func RegisterGodemoServiceServer(s grpc.ServiceRegistrar, srv GodemoServiceServer) {
	s.RegisterService(&GodemoService_ServiceDesc, srv)
}

func _GodemoService_GetEmployeeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(employee.EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GodemoServiceServer).GetEmployeeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GodemoService/getEmployeeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GodemoServiceServer).GetEmployeeInfo(ctx, req.(*employee.EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GodemoService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(search.SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GodemoServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GodemoService/search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GodemoServiceServer).Search(ctx, req.(*search.SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GodemoService_ServiceDesc is the grpc.ServiceDesc for GodemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GodemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GodemoService",
	HandlerType: (*GodemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getEmployeeInfo",
			Handler:    _GodemoService_GetEmployeeInfo_Handler,
		},
		{
			MethodName: "search",
			Handler:    _GodemoService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
