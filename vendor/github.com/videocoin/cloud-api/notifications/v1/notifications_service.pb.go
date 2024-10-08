// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: notifications/v1/notifications_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	rpc "github.com/videocoin/cloud-api/rpc"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() {
	proto.RegisterFile("notifications/v1/notifications_service.proto", fileDescriptor_486b70dc91eaa483)
}
func init() {
	golang_proto.RegisterFile("notifications/v1/notifications_service.proto", fileDescriptor_486b70dc91eaa483)
}

var fileDescriptor_486b70dc91eaa483 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0x1e, 0x16, 0xe9, 0x49, 0x2a, 0x88, 0x74, 0x25, 0x07, 0xcf, 0xee, 0x84, 0xea,
	0x1b, 0x08, 0x82, 0x27, 0x3d, 0xec, 0xcd, 0x8b, 0xa4, 0xd9, 0x34, 0x0d, 0x74, 0x33, 0xa1, 0x9d,
	0x06, 0xf4, 0xe8, 0x2b, 0xf8, 0x42, 0x1e, 0xf7, 0x28, 0xf8, 0x02, 0xd2, 0xf5, 0x41, 0xa4, 0x49,
	0x65, 0x77, 0x6f, 0x99, 0xcc, 0xf7, 0xcf, 0xff, 0xff, 0xe9, 0xb5, 0x45, 0x32, 0x95, 0x91, 0x82,
	0x0c, 0xda, 0x8e, 0xfb, 0x82, 0x1f, 0x7d, 0xbc, 0x74, 0xaa, 0xf5, 0x46, 0x2a, 0x70, 0x2d, 0x12,
	0x66, 0xb9, 0x6c, 0xb0, 0x5f, 0x83, 0x70, 0x06, 0x8e, 0x30, 0xf0, 0x45, 0xbe, 0xd0, 0x88, 0xba,
	0x51, 0x3c, 0x90, 0x65, 0x5f, 0x71, 0xb5, 0x71, 0xf4, 0x1a, 0x85, 0xf9, 0xe5, 0xb4, 0x14, 0xce,
	0x70, 0x61, 0x2d, 0xd2, 0xa4, 0x8b, 0xdb, 0xa5, 0x36, 0x54, 0xf7, 0x25, 0x48, 0xdc, 0x70, 0x8d,
	0x1a, 0xf7, 0x37, 0xc6, 0x29, 0x0c, 0xe1, 0x35, 0xe1, 0xfc, 0x00, 0xf7, 0x66, 0xad, 0x50, 0xa2,
	0xb1, 0x3c, 0x44, 0x5b, 0x8e, 0x06, 0xad, 0x93, 0xbc, 0x56, 0xa2, 0xa1, 0x3a, 0x0a, 0x6e, 0xaa,
	0xf4, 0xec, 0xf1, 0x20, 0xee, 0x2a, 0x76, 0xca, 0x9e, 0xd2, 0xf9, 0x43, 0xc0, 0xb2, 0x73, 0x88,
	0xf9, 0xe0, 0xdf, 0x18, 0xee, 0xc7, 0xf0, 0xf9, 0x02, 0xf6, 0x85, 0x5b, 0x27, 0x21, 0xe2, 0x2b,
	0x12, 0xd4, 0x77, 0x57, 0xa7, 0xef, 0xdf, 0xbf, 0x1f, 0xb3, 0x34, 0x3b, 0x99, 0xcc, 0xde, 0xee,
	0x2e, 0xb6, 0x03, 0x4b, 0xbe, 0x06, 0x96, 0xfc, 0x0c, 0x2c, 0xf9, 0xdc, 0xb1, 0x64, 0xbb, 0x63,
	0xc9, 0xf3, 0xcc, 0x17, 0xe5, 0x3c, 0x1c, 0xbe, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xf2,
	0x24, 0xf3, 0x6f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.notifications.v1.NotificationService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	Health(context.Context, *types.Empty) (*rpc.HealthStatus, error)
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.notifications.v1.NotificationService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Health(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.notifications.v1.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _NotificationService_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications/v1/notifications_service.proto",
}
