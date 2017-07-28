// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/serverpb/server.proto

/*
Package serverpb is a generated protocol buffer package.

It is generated from these files:
	server/serverpb/server.proto

It has these top-level messages:
	DrawIPRequest
	DrawIPResponse
	GetPrefixIncludingIPRequest
	GetPrefixIncludingIPResponse
	ActivateIPRequest
	ActivateIPResponse
	Tag
*/
package serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DrawIPRequest struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask string `protobuf:"bytes,2,opt,name=mask" json:"mask,omitempty"`
	Tags []*Tag `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
}

func (m *DrawIPRequest) Reset()                    { *m = DrawIPRequest{} }
func (m *DrawIPRequest) String() string            { return proto.CompactTextString(m) }
func (*DrawIPRequest) ProtoMessage()               {}
func (*DrawIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DrawIPRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DrawIPRequest) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *DrawIPRequest) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type DrawIPResponse struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *DrawIPResponse) Reset()                    { *m = DrawIPResponse{} }
func (m *DrawIPResponse) String() string            { return proto.CompactTextString(m) }
func (*DrawIPResponse) ProtoMessage()               {}
func (*DrawIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DrawIPResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GetPrefixIncludingIPRequest struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *GetPrefixIncludingIPRequest) Reset()                    { *m = GetPrefixIncludingIPRequest{} }
func (m *GetPrefixIncludingIPRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPrefixIncludingIPRequest) ProtoMessage()               {}
func (*GetPrefixIncludingIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetPrefixIncludingIPRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type GetPrefixIncludingIPResponse struct {
}

func (m *GetPrefixIncludingIPResponse) Reset()                    { *m = GetPrefixIncludingIPResponse{} }
func (m *GetPrefixIncludingIPResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPrefixIncludingIPResponse) ProtoMessage()               {}
func (*GetPrefixIncludingIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ActivateIPRequest struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *ActivateIPRequest) Reset()                    { *m = ActivateIPRequest{} }
func (m *ActivateIPRequest) String() string            { return proto.CompactTextString(m) }
func (*ActivateIPRequest) ProtoMessage()               {}
func (*ActivateIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ActivateIPRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type ActivateIPResponse struct {
}

func (m *ActivateIPResponse) Reset()                    { *m = ActivateIPResponse{} }
func (m *ActivateIPResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivateIPResponse) ProtoMessage()               {}
func (*ActivateIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Tag struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*DrawIPRequest)(nil), "serverpb.DrawIPRequest")
	proto.RegisterType((*DrawIPResponse)(nil), "serverpb.DrawIPResponse")
	proto.RegisterType((*GetPrefixIncludingIPRequest)(nil), "serverpb.GetPrefixIncludingIPRequest")
	proto.RegisterType((*GetPrefixIncludingIPResponse)(nil), "serverpb.GetPrefixIncludingIPResponse")
	proto.RegisterType((*ActivateIPRequest)(nil), "serverpb.ActivateIPRequest")
	proto.RegisterType((*ActivateIPResponse)(nil), "serverpb.ActivateIPResponse")
	proto.RegisterType((*Tag)(nil), "serverpb.Tag")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PrefixService service

type PrefixServiceClient interface {
	DrawIP(ctx context.Context, in *DrawIPRequest, opts ...grpc.CallOption) (*DrawIPResponse, error)
	GetPrefixIncludingIP(ctx context.Context, in *GetPrefixIncludingIPRequest, opts ...grpc.CallOption) (*GetPrefixIncludingIPResponse, error)
	ActivateIP(ctx context.Context, in *ActivateIPRequest, opts ...grpc.CallOption) (*ActivateIPResponse, error)
}

type prefixServiceClient struct {
	cc *grpc.ClientConn
}

func NewPrefixServiceClient(cc *grpc.ClientConn) PrefixServiceClient {
	return &prefixServiceClient{cc}
}

func (c *prefixServiceClient) DrawIP(ctx context.Context, in *DrawIPRequest, opts ...grpc.CallOption) (*DrawIPResponse, error) {
	out := new(DrawIPResponse)
	err := grpc.Invoke(ctx, "/serverpb.PrefixService/DrawIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prefixServiceClient) GetPrefixIncludingIP(ctx context.Context, in *GetPrefixIncludingIPRequest, opts ...grpc.CallOption) (*GetPrefixIncludingIPResponse, error) {
	out := new(GetPrefixIncludingIPResponse)
	err := grpc.Invoke(ctx, "/serverpb.PrefixService/GetPrefixIncludingIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prefixServiceClient) ActivateIP(ctx context.Context, in *ActivateIPRequest, opts ...grpc.CallOption) (*ActivateIPResponse, error) {
	out := new(ActivateIPResponse)
	err := grpc.Invoke(ctx, "/serverpb.PrefixService/ActivateIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PrefixService service

type PrefixServiceServer interface {
	DrawIP(context.Context, *DrawIPRequest) (*DrawIPResponse, error)
	GetPrefixIncludingIP(context.Context, *GetPrefixIncludingIPRequest) (*GetPrefixIncludingIPResponse, error)
	ActivateIP(context.Context, *ActivateIPRequest) (*ActivateIPResponse, error)
}

func RegisterPrefixServiceServer(s *grpc.Server, srv PrefixServiceServer) {
	s.RegisterService(&_PrefixService_serviceDesc, srv)
}

func _PrefixService_DrawIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrawIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefixServiceServer).DrawIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.PrefixService/DrawIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefixServiceServer).DrawIP(ctx, req.(*DrawIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrefixService_GetPrefixIncludingIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrefixIncludingIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefixServiceServer).GetPrefixIncludingIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.PrefixService/GetPrefixIncludingIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefixServiceServer).GetPrefixIncludingIP(ctx, req.(*GetPrefixIncludingIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrefixService_ActivateIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefixServiceServer).ActivateIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.PrefixService/ActivateIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefixServiceServer).ActivateIP(ctx, req.(*ActivateIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrefixService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.PrefixService",
	HandlerType: (*PrefixServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DrawIP",
			Handler:    _PrefixService_DrawIP_Handler,
		},
		{
			MethodName: "GetPrefixIncludingIP",
			Handler:    _PrefixService_GetPrefixIncludingIP_Handler,
		},
		{
			MethodName: "ActivateIP",
			Handler:    _PrefixService_ActivateIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/server.proto",
}

func init() { proto.RegisterFile("server/serverpb/server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0x8a, 0xda, 0x40,
	0x14, 0x86, 0x31, 0xb1, 0xd2, 0x9e, 0xa2, 0xb4, 0xa7, 0x8a, 0x21, 0xa6, 0x62, 0x53, 0x2a, 0xde,
	0x68, 0xc0, 0x3e, 0x41, 0xa1, 0x50, 0xbc, 0x13, 0x2b, 0xbd, 0x1f, 0xe3, 0x18, 0x06, 0x35, 0x99,
	0x66, 0xc6, 0xb8, 0x22, 0xde, 0xf8, 0x0a, 0xfb, 0x68, 0xfb, 0x04, 0x0b, 0xfb, 0x20, 0x8b, 0x33,
	0x13, 0xb2, 0xa2, 0xee, 0x5e, 0x79, 0x9c, 0xf3, 0xe7, 0xff, 0xbf, 0xf9, 0x13, 0xf0, 0x04, 0x4d,
	0x33, 0x9a, 0x06, 0xfa, 0x87, 0xcf, 0xcc, 0x30, 0xe0, 0x69, 0x22, 0x13, 0x7c, 0x9f, 0x1f, 0xbb,
	0x5e, 0x94, 0x24, 0xd1, 0x8a, 0x06, 0x84, 0xb3, 0x80, 0xc4, 0x71, 0x22, 0x89, 0x64, 0x49, 0x2c,
	0xb4, 0xce, 0xff, 0x07, 0xd5, 0xdf, 0x29, 0xd9, 0x8e, 0xc6, 0x13, 0xfa, 0x7f, 0x43, 0x85, 0xc4,
	0x1a, 0x58, 0x8c, 0x3b, 0xa5, 0x4e, 0xa9, 0xf7, 0x61, 0x62, 0x31, 0x8e, 0x08, 0xe5, 0x35, 0x11,
	0x4b, 0xc7, 0x52, 0x27, 0x6a, 0xc6, 0x6f, 0x50, 0x96, 0x24, 0x12, 0x8e, 0xdd, 0xb1, 0x7b, 0x1f,
	0x87, 0xd5, 0x41, 0x9e, 0x35, 0x98, 0x92, 0x68, 0xa2, 0x56, 0xbe, 0x0f, 0xb5, 0xdc, 0x57, 0xf0,
	0x24, 0x16, 0x14, 0x3f, 0x81, 0xbd, 0x16, 0x91, 0x71, 0x3e, 0x8d, 0x7e, 0x1f, 0x5a, 0x7f, 0xa8,
	0x1c, 0xa7, 0x74, 0xc1, 0xee, 0x46, 0x71, 0xb8, 0xda, 0xcc, 0x59, 0x1c, 0xdd, 0x24, 0xf1, 0xdb,
	0xe0, 0x5d, 0x97, 0xeb, 0x00, 0xff, 0x3b, 0x7c, 0xfe, 0x15, 0x4a, 0x96, 0x11, 0x49, 0x6f, 0x9b,
	0xd4, 0x01, 0x5f, 0x8a, 0xcc, 0xa3, 0x7d, 0xb0, 0xa7, 0x24, 0x3a, 0x21, 0x2e, 0xe9, 0x2e, 0x47,
	0x5c, 0xd2, 0x1d, 0xd6, 0xe1, 0x5d, 0x46, 0x56, 0x1b, 0x6a, 0xae, 0xaf, 0xff, 0x0c, 0x1f, 0x2d,
	0xa8, 0x6a, 0x8e, 0xbf, 0x34, 0xcd, 0x58, 0x48, 0x31, 0x84, 0x8a, 0xbe, 0x2e, 0x36, 0x8b, 0x36,
	0xce, 0x8a, 0x75, 0x9d, 0xcb, 0x85, 0x49, 0xef, 0x1e, 0x1f, 0x9e, 0xee, 0xad, 0x0e, 0xb6, 0xd5,
	0x3b, 0xe2, 0xca, 0x38, 0xd8, 0x33, 0x7e, 0x08, 0xf6, 0xa7, 0xba, 0x0f, 0xc1, 0x3c, 0x25, 0x5b,
	0xc6, 0xf1, 0x58, 0x82, 0xfa, 0xb5, 0x06, 0xf0, 0x47, 0x61, 0xfd, 0x4a, 0xa1, 0x6e, 0xf7, 0x2d,
	0x99, 0xe1, 0x69, 0x29, 0x9e, 0x06, 0x7e, 0x51, 0x3c, 0x8c, 0x6b, 0x16, 0xcd, 0x85, 0x0b, 0x80,
	0xa2, 0x40, 0x6c, 0x15, 0x96, 0x17, 0xdd, 0xbb, 0xde, 0xf5, 0xa5, 0x49, 0xf9, 0xaa, 0x52, 0x9a,
	0x7e, 0xe3, 0x2c, 0x85, 0x18, 0xe1, 0xac, 0xa2, 0xbe, 0xcf, 0x9f, 0xcf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xb1, 0x6e, 0xab, 0x29, 0xe7, 0x02, 0x00, 0x00,
}
