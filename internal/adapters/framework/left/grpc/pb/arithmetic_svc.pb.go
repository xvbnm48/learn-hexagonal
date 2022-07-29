// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arithmetic_svc.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("arithmetic_svc.proto", fileDescriptor_05d056e18e5f0c80)
}

var fileDescriptor_05d056e18e5f0c80 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2c, 0xca, 0x2c,
	0xc9, 0xc8, 0x4d, 0x2d, 0xc9, 0x4c, 0x8e, 0x2f, 0x2e, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2a, 0x48, 0x92, 0x12, 0xc8, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x8a, 0xcf, 0x2d, 0x4e,
	0x87, 0x88, 0x1a, 0x7d, 0x61, 0xe4, 0x12, 0x74, 0x84, 0x2b, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x15, 0x32, 0xe1, 0xe2, 0x76, 0x4f, 0x2d, 0x71, 0x4c, 0x49, 0xc9, 0x2c, 0xc9, 0xcc, 0xcf,
	0x13, 0x12, 0xd7, 0x2b, 0x48, 0xd2, 0xf3, 0x2f, 0x48, 0x2d, 0x4a, 0x04, 0x71, 0x03, 0x12, 0x8b,
	0x12, 0x73, 0x53, 0x4b, 0x52, 0x8b, 0x8a, 0xa5, 0xb8, 0x40, 0x12, 0x8e, 0x79, 0xc5, 0xe5, 0xa9,
	0x45, 0x4a, 0x0c, 0x42, 0xe6, 0x5c, 0x7c, 0xee, 0xa9, 0x25, 0xc1, 0xa5, 0x49, 0x25, 0x45, 0x89,
	0xc9, 0xa4, 0x68, 0xb4, 0xe2, 0x12, 0x74, 0x4f, 0x2d, 0xf1, 0x2d, 0xcd, 0x29, 0xc9, 0x2c, 0xc8,
	0xc9, 0x4c, 0x4e, 0x24, 0x45, 0x2f, 0xc4, 0xa9, 0x2e, 0x99, 0x65, 0x99, 0xc5, 0xc4, 0xeb, 0x72,
	0x62, 0x8b, 0x62, 0xd1, 0xd3, 0x2f, 0x48, 0x4a, 0x62, 0x03, 0x87, 0x82, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0x78, 0xed, 0x2a, 0x33, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ArithmeticServiceClient is the client API for ArithmeticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArithmeticServiceClient interface {
	GetAddition(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetSubtraction(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetMultiplication(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetDivision(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
}

type arithmeticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArithmeticServiceClient(cc grpc.ClientConnInterface) ArithmeticServiceClient {
	return &arithmeticServiceClient{cc}
}

func (c *arithmeticServiceClient) GetAddition(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetAddition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetSubtraction(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetSubtraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetMultiplication(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetMultiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetDivision(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetDivision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticServiceServer is the server API for ArithmeticService service.
type ArithmeticServiceServer interface {
	GetAddition(context.Context, *OperationParameters) (*Answer, error)
	GetSubtraction(context.Context, *OperationParameters) (*Answer, error)
	GetMultiplication(context.Context, *OperationParameters) (*Answer, error)
	GetDivision(context.Context, *OperationParameters) (*Answer, error)
}

// UnimplementedArithmeticServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArithmeticServiceServer struct {
}

func (*UnimplementedArithmeticServiceServer) GetAddition(ctx context.Context, req *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddition not implemented")
}
func (*UnimplementedArithmeticServiceServer) GetSubtraction(ctx context.Context, req *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubtraction not implemented")
}
func (*UnimplementedArithmeticServiceServer) GetMultiplication(ctx context.Context, req *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiplication not implemented")
}
func (*UnimplementedArithmeticServiceServer) GetDivision(ctx context.Context, req *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDivision not implemented")
}

func RegisterArithmeticServiceServer(s *grpc.Server, srv ArithmeticServiceServer) {
	s.RegisterService(&_ArithmeticService_serviceDesc, srv)
}

func _ArithmeticService_GetAddition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetAddition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetAddition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetAddition(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetSubtraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetSubtraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetSubtraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetSubtraction(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetMultiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetMultiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetMultiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetMultiplication(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetDivision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetDivision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetDivision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetDivision(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArithmeticService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArithmeticService",
	HandlerType: (*ArithmeticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddition",
			Handler:    _ArithmeticService_GetAddition_Handler,
		},
		{
			MethodName: "GetSubtraction",
			Handler:    _ArithmeticService_GetSubtraction_Handler,
		},
		{
			MethodName: "GetMultiplication",
			Handler:    _ArithmeticService_GetMultiplication_Handler,
		},
		{
			MethodName: "GetDivision",
			Handler:    _ArithmeticService_GetDivision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arithmetic_svc.proto",
}