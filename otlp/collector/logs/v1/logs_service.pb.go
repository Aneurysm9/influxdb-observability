// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/collector/logs/v1/logs_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/influxdata/influxdb-observability/otlp/logs/v1"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExportLogsServiceRequest struct {
	// An array of ResourceLogs.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceLogs         []*v1.ResourceLogs `protobuf:"bytes,1,rep,name=resource_logs,json=resourceLogs,proto3" json:"resource_logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExportLogsServiceRequest) Reset()         { *m = ExportLogsServiceRequest{} }
func (m *ExportLogsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportLogsServiceRequest) ProtoMessage()    {}
func (*ExportLogsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e3bf87aaa43acd4, []int{0}
}
func (m *ExportLogsServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportLogsServiceRequest.Unmarshal(m, b)
}
func (m *ExportLogsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportLogsServiceRequest.Marshal(b, m, deterministic)
}
func (m *ExportLogsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportLogsServiceRequest.Merge(m, src)
}
func (m *ExportLogsServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ExportLogsServiceRequest.Size(m)
}
func (m *ExportLogsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportLogsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportLogsServiceRequest proto.InternalMessageInfo

func (m *ExportLogsServiceRequest) GetResourceLogs() []*v1.ResourceLogs {
	if m != nil {
		return m.ResourceLogs
	}
	return nil
}

type ExportLogsServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportLogsServiceResponse) Reset()         { *m = ExportLogsServiceResponse{} }
func (m *ExportLogsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportLogsServiceResponse) ProtoMessage()    {}
func (*ExportLogsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e3bf87aaa43acd4, []int{1}
}
func (m *ExportLogsServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportLogsServiceResponse.Unmarshal(m, b)
}
func (m *ExportLogsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportLogsServiceResponse.Marshal(b, m, deterministic)
}
func (m *ExportLogsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportLogsServiceResponse.Merge(m, src)
}
func (m *ExportLogsServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ExportLogsServiceResponse.Size(m)
}
func (m *ExportLogsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportLogsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportLogsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportLogsServiceRequest)(nil), "internal.opentelemetry.proto.collector.logs.v1.ExportLogsServiceRequest")
	proto.RegisterType((*ExportLogsServiceResponse)(nil), "internal.opentelemetry.proto.collector.logs.v1.ExportLogsServiceResponse")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/collector/logs/v1/logs_service.proto", fileDescriptor_8e3bf87aaa43acd4)
}

var fileDescriptor_8e3bf87aaa43acd4 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x2d, 0xc2, 0x0e, 0x99, 0x82, 0xf4, 0x34, 0xe7, 0x45, 0x76, 0x10, 0x2f, 0x26, 0xb4,
	0x5e, 0x3c, 0x4f, 0x04, 0x05, 0x0f, 0xa3, 0x1e, 0x04, 0x2f, 0xa3, 0x8d, 0x9f, 0x35, 0x90, 0xe5,
	0x8b, 0xc9, 0xd7, 0xb2, 0xfd, 0x07, 0xff, 0x83, 0x3f, 0xc1, 0xbf, 0x28, 0x69, 0xa6, 0x54, 0x9c,
	0xc2, 0xc0, 0x53, 0x9b, 0xe4, 0x7d, 0x9f, 0xf7, 0x4d, 0xf8, 0xd8, 0x05, 0x5a, 0x30, 0x04, 0x1a,
	0x16, 0x40, 0x6e, 0x25, 0xac, 0x43, 0x42, 0x21, 0x51, 0x6b, 0x90, 0x84, 0x4e, 0x68, 0xac, 0xbd,
	0x68, 0xb3, 0xee, 0x3b, 0xf7, 0xe0, 0x5a, 0x25, 0x81, 0x77, 0xa2, 0x94, 0x2b, 0x43, 0xe0, 0x4c,
	0xa9, 0xf9, 0x37, 0x44, 0x3c, 0xe5, 0x5f, 0x08, 0x1e, 0xac, 0xbc, 0xcd, 0xc6, 0x27, 0x9b, 0x92,
	0xfa, 0xfc, 0xe8, 0x9c, 0x78, 0x36, 0xba, 0x5a, 0x5a, 0x74, 0x74, 0x8b, 0xb5, 0xbf, 0x8b, 0x91,
	0x05, 0xbc, 0x34, 0xe0, 0x29, 0xbd, 0x67, 0xfb, 0x0e, 0x3c, 0x36, 0x4e, 0xc2, 0x3c, 0x58, 0x46,
	0xc9, 0xf1, 0xee, 0xe9, 0x30, 0xcf, 0xff, 0xee, 0xb2, 0x6e, 0xc0, 0x8b, 0xb5, 0x35, 0x80, 0x8b,
	0x3d, 0xd7, 0x5b, 0x4d, 0x8e, 0xd8, 0xe1, 0x86, 0x50, 0x6f, 0xd1, 0x78, 0xc8, 0xdf, 0x13, 0x36,
	0xec, 0xed, 0xa7, 0x6f, 0x09, 0x1b, 0x44, 0x75, 0x7a, 0xbd, 0xe5, 0x2b, 0xf0, 0xdf, 0xae, 0x36,
	0xbe, 0xf9, 0x07, 0x52, 0xec, 0x3b, 0xd9, 0x99, 0xbe, 0x26, 0x2c, 0x53, 0xb8, 0x25, 0x71, 0x7a,
	0xd0, 0x83, 0xcd, 0x82, 0x66, 0x96, 0x3c, 0x5c, 0xd6, 0x8a, 0x9e, 0x9b, 0x8a, 0x4b, 0x5c, 0x08,
	0x65, 0x9e, 0x74, 0xb3, 0x7c, 0x2c, 0xa9, 0xfc, 0xfc, 0xad, 0xce, 0xb0, 0x0a, 0x53, 0x51, 0x56,
	0x4a, 0x2b, 0x5a, 0x09, 0x24, 0x6d, 0x7f, 0x4e, 0x4f, 0x35, 0xe8, 0x12, 0xcf, 0x3f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x17, 0x4b, 0x5b, 0x18, 0x6d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogsServiceClient is the client API for LogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *ExportLogsServiceRequest, opts ...grpc.CallOption) (*ExportLogsServiceResponse, error)
}

type logsServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogsServiceClient(cc *grpc.ClientConn) LogsServiceClient {
	return &logsServiceClient{cc}
}

func (c *logsServiceClient) Export(ctx context.Context, in *ExportLogsServiceRequest, opts ...grpc.CallOption) (*ExportLogsServiceResponse, error) {
	out := new(ExportLogsServiceResponse)
	err := c.cc.Invoke(ctx, "/internal.opentelemetry.proto.collector.logs.v1.LogsService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogsServiceServer is the server API for LogsService service.
type LogsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *ExportLogsServiceRequest) (*ExportLogsServiceResponse, error)
}

// UnimplementedLogsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogsServiceServer struct {
}

func (*UnimplementedLogsServiceServer) Export(ctx context.Context, req *ExportLogsServiceRequest) (*ExportLogsServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterLogsServiceServer(s *grpc.Server, srv LogsServiceServer) {
	s.RegisterService(&_LogsService_serviceDesc, srv)
}

func _LogsService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportLogsServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.opentelemetry.proto.collector.logs.v1.LogsService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServiceServer).Export(ctx, req.(*ExportLogsServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.opentelemetry.proto.collector.logs.v1.LogsService",
	HandlerType: (*LogsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _LogsService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/logs/v1/logs_service.proto",
}
