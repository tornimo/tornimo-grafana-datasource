// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/mobile_device_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Request message for [MobileDeviceConstantService.GetMobileDeviceConstant][google.ads.googleads.v2.services.MobileDeviceConstantService.GetMobileDeviceConstant].
type GetMobileDeviceConstantRequest struct {
	// Resource name of the mobile device to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMobileDeviceConstantRequest) Reset()         { *m = GetMobileDeviceConstantRequest{} }
func (m *GetMobileDeviceConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetMobileDeviceConstantRequest) ProtoMessage()    {}
func (*GetMobileDeviceConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b671ef47eb7057e, []int{0}
}

func (m *GetMobileDeviceConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Unmarshal(m, b)
}
func (m *GetMobileDeviceConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetMobileDeviceConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMobileDeviceConstantRequest.Merge(m, src)
}
func (m *GetMobileDeviceConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Size(m)
}
func (m *GetMobileDeviceConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMobileDeviceConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMobileDeviceConstantRequest proto.InternalMessageInfo

func (m *GetMobileDeviceConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMobileDeviceConstantRequest)(nil), "google.ads.googleads.v2.services.GetMobileDeviceConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/mobile_device_constant_service.proto", fileDescriptor_5b671ef47eb7057e)
}

var fileDescriptor_5b671ef47eb7057e = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcb, 0x4a, 0xeb, 0x40,
	0x18, 0x26, 0x39, 0x70, 0xe0, 0x84, 0xe3, 0x26, 0x9b, 0x96, 0x54, 0x24, 0xd4, 0x2e, 0x44, 0x71,
	0x06, 0xd2, 0x85, 0x30, 0xa2, 0x98, 0x6a, 0xa9, 0x1b, 0xa5, 0x54, 0xe8, 0x42, 0x02, 0x61, 0x9a,
	0x0c, 0x21, 0x90, 0xcc, 0xd4, 0xfc, 0x69, 0x37, 0xe2, 0xc6, 0x85, 0x3e, 0x80, 0x6f, 0xe0, 0xd2,
	0x37, 0xb1, 0x5b, 0x5f, 0xc1, 0x95, 0x4f, 0x21, 0xb9, 0x4c, 0xaa, 0xd0, 0xb4, 0xbb, 0x2f, 0xf3,
	0x7f, 0x97, 0xff, 0x12, 0xad, 0x1f, 0x08, 0x11, 0x44, 0x0c, 0x53, 0x1f, 0x70, 0x01, 0x33, 0x34,
	0xb7, 0x30, 0xb0, 0x64, 0x1e, 0x7a, 0x0c, 0x70, 0x2c, 0x26, 0x61, 0xc4, 0x5c, 0x9f, 0x65, 0x9f,
	0xae, 0x27, 0x38, 0xa4, 0x94, 0xa7, 0x6e, 0x59, 0x47, 0xd3, 0x44, 0xa4, 0x42, 0x37, 0x0b, 0x2d,
	0xa2, 0x3e, 0xa0, 0xca, 0x06, 0xcd, 0x2d, 0x24, 0x6d, 0x8c, 0xd3, 0xba, 0xa0, 0x84, 0x81, 0x98,
	0x25, 0xf5, 0x49, 0x45, 0x82, 0xb1, 0x2d, 0xf5, 0xd3, 0x10, 0x53, 0xce, 0x45, 0x4a, 0xd3, 0x50,
	0x70, 0x28, 0xab, 0x8d, 0x1f, 0x55, 0x2f, 0x0a, 0x99, 0x94, 0xb5, 0xfb, 0xda, 0xce, 0x80, 0xa5,
	0x57, 0xb9, 0xf3, 0x45, 0x6e, 0x7c, 0x5e, 0xfa, 0x8e, 0xd8, 0xdd, 0x8c, 0x41, 0xaa, 0xef, 0x6a,
	0x5b, 0xb2, 0x05, 0x97, 0xd3, 0x98, 0x35, 0x15, 0x53, 0xd9, 0xfb, 0x37, 0xfa, 0x2f, 0x1f, 0xaf,
	0x69, 0xcc, 0xac, 0x27, 0x55, 0x6b, 0xad, 0x32, 0xb9, 0x29, 0xc6, 0xd3, 0xdf, 0x15, 0xad, 0x51,
	0x93, 0xa3, 0x9f, 0xa1, 0x4d, 0xcb, 0x41, 0xeb, 0x5b, 0x34, 0x8e, 0x6a, 0x1d, 0xaa, 0xe5, 0xa1,
	0x55, 0xfa, 0x76, 0xf7, 0xf1, 0xe3, 0xf3, 0x45, 0x3d, 0xd4, 0x0f, 0xb2, 0x45, 0xdf, 0xff, 0x1a,
	0xf3, 0x24, 0x5e, 0x21, 0x00, 0xbc, 0xff, 0x60, 0xb4, 0x16, 0x76, 0x73, 0x19, 0x52, 0xa2, 0x69,
	0x08, 0xc8, 0x13, 0x71, 0xef, 0x59, 0xd5, 0x3a, 0x9e, 0x88, 0x37, 0x8e, 0xd4, 0x33, 0xd7, 0xac,
	0x6b, 0x98, 0x9d, 0x66, 0xa8, 0xdc, 0x5e, 0x96, 0x2e, 0x81, 0x88, 0x28, 0x0f, 0x90, 0x48, 0x02,
	0x1c, 0x30, 0x9e, 0x1f, 0x0e, 0x2f, 0x73, 0xeb, 0xff, 0xcd, 0x63, 0x09, 0x5e, 0xd5, 0x3f, 0x03,
	0xdb, 0x7e, 0x53, 0xcd, 0x41, 0x61, 0x68, 0xfb, 0x80, 0x0a, 0x98, 0xa1, 0xb1, 0x85, 0xca, 0x60,
	0x58, 0x48, 0x8a, 0x63, 0xfb, 0xe0, 0x54, 0x14, 0x67, 0x6c, 0x39, 0x92, 0xf2, 0xa5, 0x76, 0x8a,
	0x77, 0x42, 0x6c, 0x1f, 0x08, 0xa9, 0x48, 0x84, 0x8c, 0x2d, 0x42, 0x24, 0x6d, 0xf2, 0x37, 0xef,
	0xb3, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x2e, 0x00, 0xeb, 0x42, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MobileDeviceConstantServiceClient is the client API for MobileDeviceConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MobileDeviceConstantServiceClient interface {
	// Returns the requested mobile device constant in full detail.
	GetMobileDeviceConstant(ctx context.Context, in *GetMobileDeviceConstantRequest, opts ...grpc.CallOption) (*resources.MobileDeviceConstant, error)
}

type mobileDeviceConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewMobileDeviceConstantServiceClient(cc *grpc.ClientConn) MobileDeviceConstantServiceClient {
	return &mobileDeviceConstantServiceClient{cc}
}

func (c *mobileDeviceConstantServiceClient) GetMobileDeviceConstant(ctx context.Context, in *GetMobileDeviceConstantRequest, opts ...grpc.CallOption) (*resources.MobileDeviceConstant, error) {
	out := new(resources.MobileDeviceConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.MobileDeviceConstantService/GetMobileDeviceConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileDeviceConstantServiceServer is the server API for MobileDeviceConstantService service.
type MobileDeviceConstantServiceServer interface {
	// Returns the requested mobile device constant in full detail.
	GetMobileDeviceConstant(context.Context, *GetMobileDeviceConstantRequest) (*resources.MobileDeviceConstant, error)
}

func RegisterMobileDeviceConstantServiceServer(s *grpc.Server, srv MobileDeviceConstantServiceServer) {
	s.RegisterService(&_MobileDeviceConstantService_serviceDesc, srv)
}

func _MobileDeviceConstantService_GetMobileDeviceConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMobileDeviceConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileDeviceConstantServiceServer).GetMobileDeviceConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.MobileDeviceConstantService/GetMobileDeviceConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileDeviceConstantServiceServer).GetMobileDeviceConstant(ctx, req.(*GetMobileDeviceConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MobileDeviceConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.MobileDeviceConstantService",
	HandlerType: (*MobileDeviceConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMobileDeviceConstant",
			Handler:    _MobileDeviceConstantService_GetMobileDeviceConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/mobile_device_constant_service.proto",
}
