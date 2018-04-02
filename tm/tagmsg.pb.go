// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tm/tagmsg.proto

/*
Package tm is a generated protocol buffer package.

Put Service

Put Service API consists of a single service which returns a message.

It is generated from these files:
	tm/tagmsg.proto

It has these top-level messages:
	PutRequest
	PutResponse
*/
package tm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Message represents a simple message sent to the Put service.
type PutRequest struct {
	// Client IP represents the message identifier.
	ClientIp string `protobuf:"bytes,1,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"`
	// Server IP represents the message identifier.
	ServerIp string `protobuf:"bytes,2,opt,name=server_ip,json=serverIp" json:"server_ip,omitempty"`
	// Tags.
	Tags map[string]string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The message to be saved.
	Message string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PutRequest) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *PutRequest) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *PutRequest) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PutRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PutResponse struct {
	// The response message.
	ResponseMsg string `protobuf:"bytes,1,opt,name=response_msg,json=responseMsg" json:"response_msg,omitempty"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PutResponse) GetResponseMsg() string {
	if m != nil {
		return m.ResponseMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*PutRequest)(nil), "tagmsg.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "tagmsg.PutResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TagMsgService service

type TagMsgServiceClient interface {
	// Put method receives a simple message and returns it.
	// The message posted as the id parameter will also be returned.
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
}

type tagMsgServiceClient struct {
	cc *grpc.ClientConn
}

func NewTagMsgServiceClient(cc *grpc.ClientConn) TagMsgServiceClient {
	return &tagMsgServiceClient{cc}
}

func (c *tagMsgServiceClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/tagmsg.TagMsgService/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TagMsgService service

type TagMsgServiceServer interface {
	// Put method receives a simple message and returns it.
	// The message posted as the id parameter will also be returned.
	Put(context.Context, *PutRequest) (*PutResponse, error)
}

func RegisterTagMsgServiceServer(s *grpc.Server, srv TagMsgServiceServer) {
	s.RegisterService(&_TagMsgService_serviceDesc, srv)
}

func _TagMsgService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagMsgServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tagmsg.TagMsgService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagMsgServiceServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagMsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tagmsg.TagMsgService",
	HandlerType: (*TagMsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _TagMsgService_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tm/tagmsg.proto",
}

func init() { proto.RegisterFile("tm/tagmsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xcd, 0x1f, 0xab, 0x99, 0x28, 0xca, 0xe8, 0x21, 0x54, 0x0f, 0x35, 0xa7, 0x9e, 0x62,
	0xa8, 0x07, 0xc5, 0x9b, 0x82, 0x87, 0x1e, 0x0a, 0x25, 0xf6, 0xe4, 0xa5, 0xac, 0x65, 0x58, 0x82,
	0xcd, 0x1f, 0x77, 0x26, 0x81, 0x7e, 0x46, 0xbf, 0x94, 0x24, 0x9b, 0x2a, 0xe2, 0x6d, 0xde, 0x7b,
	0xc3, 0x6f, 0xde, 0x2e, 0x9c, 0x49, 0x71, 0x2b, 0x4a, 0x17, 0xac, 0x93, 0xda, 0x54, 0x52, 0xe1,
	0xc8, 0xaa, 0xf8, 0xcb, 0x01, 0x58, 0x36, 0x92, 0xd1, 0x67, 0x43, 0x2c, 0x78, 0x05, 0xc1, 0x66,
	0x9b, 0x53, 0x29, 0xeb, 0xbc, 0x8e, 0x9c, 0x89, 0x33, 0x0d, 0xb2, 0x63, 0x6b, 0xcc, 0xeb, 0x2e,
	0x64, 0x32, 0x2d, 0x99, 0x2e, 0x74, 0x6d, 0x68, 0x8d, 0x79, 0x8d, 0x29, 0xf8, 0xa2, 0x34, 0x47,
	0xde, 0xc4, 0x9b, 0x86, 0xb3, 0xeb, 0x64, 0xb8, 0xf6, 0xcb, 0x4e, 0x56, 0x4a, 0xf3, 0x4b, 0x29,
	0x66, 0x97, 0xf5, 0x9b, 0x18, 0xc1, 0x51, 0x41, 0xcc, 0x4a, 0x53, 0xe4, 0xf7, 0xb0, 0xbd, 0x1c,
	0xdf, 0x43, 0xf0, 0xb3, 0x8c, 0xe7, 0xe0, 0x7d, 0xd0, 0x6e, 0x28, 0xd3, 0x8d, 0x78, 0x09, 0x87,
	0xad, 0xda, 0x36, 0x34, 0x74, 0xb0, 0xe2, 0xd1, 0x7d, 0x70, 0xe2, 0x14, 0xc2, 0xfe, 0x20, 0xd7,
	0x55, 0xc9, 0x84, 0x37, 0x70, 0x62, 0x86, 0x79, 0x5d, 0xb0, 0x1e, 0x18, 0xe1, 0xde, 0x5b, 0xb0,
	0x9e, 0x3d, 0xc1, 0xe9, 0x4a, 0xe9, 0x05, 0xeb, 0x57, 0x32, 0x6d, 0xbe, 0x21, 0x4c, 0xc1, 0x5b,
	0x36, 0x82, 0xf8, 0xff, 0x01, 0xe3, 0x8b, 0x3f, 0x9e, 0x65, 0xc4, 0x07, 0xcf, 0xfe, 0x9b, 0x2b,
	0xc5, 0xfb, 0xa8, 0xff, 0xd7, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x7d, 0x87, 0x44,
	0x6a, 0x01, 0x00, 0x00,
}
