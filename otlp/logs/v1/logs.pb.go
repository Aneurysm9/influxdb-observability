// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/logs/v1/logs.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v11 "github.com/influxdata/influxdb-observability/otlp/common/v1"
	v1 "github.com/influxdata/influxdb-observability/otlp/resource/v1"
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

// Possible values for LogRecord.SeverityNumber.
type SeverityNumber int32

const (
	// UNSPECIFIED is the default SeverityNumber, it MUST NOT be used.
	SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED SeverityNumber = 0
	SeverityNumber_SEVERITY_NUMBER_TRACE       SeverityNumber = 1
	SeverityNumber_SEVERITY_NUMBER_TRACE2      SeverityNumber = 2
	SeverityNumber_SEVERITY_NUMBER_TRACE3      SeverityNumber = 3
	SeverityNumber_SEVERITY_NUMBER_TRACE4      SeverityNumber = 4
	SeverityNumber_SEVERITY_NUMBER_DEBUG       SeverityNumber = 5
	SeverityNumber_SEVERITY_NUMBER_DEBUG2      SeverityNumber = 6
	SeverityNumber_SEVERITY_NUMBER_DEBUG3      SeverityNumber = 7
	SeverityNumber_SEVERITY_NUMBER_DEBUG4      SeverityNumber = 8
	SeverityNumber_SEVERITY_NUMBER_INFO        SeverityNumber = 9
	SeverityNumber_SEVERITY_NUMBER_INFO2       SeverityNumber = 10
	SeverityNumber_SEVERITY_NUMBER_INFO3       SeverityNumber = 11
	SeverityNumber_SEVERITY_NUMBER_INFO4       SeverityNumber = 12
	SeverityNumber_SEVERITY_NUMBER_WARN        SeverityNumber = 13
	SeverityNumber_SEVERITY_NUMBER_WARN2       SeverityNumber = 14
	SeverityNumber_SEVERITY_NUMBER_WARN3       SeverityNumber = 15
	SeverityNumber_SEVERITY_NUMBER_WARN4       SeverityNumber = 16
	SeverityNumber_SEVERITY_NUMBER_ERROR       SeverityNumber = 17
	SeverityNumber_SEVERITY_NUMBER_ERROR2      SeverityNumber = 18
	SeverityNumber_SEVERITY_NUMBER_ERROR3      SeverityNumber = 19
	SeverityNumber_SEVERITY_NUMBER_ERROR4      SeverityNumber = 20
	SeverityNumber_SEVERITY_NUMBER_FATAL       SeverityNumber = 21
	SeverityNumber_SEVERITY_NUMBER_FATAL2      SeverityNumber = 22
	SeverityNumber_SEVERITY_NUMBER_FATAL3      SeverityNumber = 23
	SeverityNumber_SEVERITY_NUMBER_FATAL4      SeverityNumber = 24
)

var SeverityNumber_name = map[int32]string{
	0:  "SEVERITY_NUMBER_UNSPECIFIED",
	1:  "SEVERITY_NUMBER_TRACE",
	2:  "SEVERITY_NUMBER_TRACE2",
	3:  "SEVERITY_NUMBER_TRACE3",
	4:  "SEVERITY_NUMBER_TRACE4",
	5:  "SEVERITY_NUMBER_DEBUG",
	6:  "SEVERITY_NUMBER_DEBUG2",
	7:  "SEVERITY_NUMBER_DEBUG3",
	8:  "SEVERITY_NUMBER_DEBUG4",
	9:  "SEVERITY_NUMBER_INFO",
	10: "SEVERITY_NUMBER_INFO2",
	11: "SEVERITY_NUMBER_INFO3",
	12: "SEVERITY_NUMBER_INFO4",
	13: "SEVERITY_NUMBER_WARN",
	14: "SEVERITY_NUMBER_WARN2",
	15: "SEVERITY_NUMBER_WARN3",
	16: "SEVERITY_NUMBER_WARN4",
	17: "SEVERITY_NUMBER_ERROR",
	18: "SEVERITY_NUMBER_ERROR2",
	19: "SEVERITY_NUMBER_ERROR3",
	20: "SEVERITY_NUMBER_ERROR4",
	21: "SEVERITY_NUMBER_FATAL",
	22: "SEVERITY_NUMBER_FATAL2",
	23: "SEVERITY_NUMBER_FATAL3",
	24: "SEVERITY_NUMBER_FATAL4",
}

var SeverityNumber_value = map[string]int32{
	"SEVERITY_NUMBER_UNSPECIFIED": 0,
	"SEVERITY_NUMBER_TRACE":       1,
	"SEVERITY_NUMBER_TRACE2":      2,
	"SEVERITY_NUMBER_TRACE3":      3,
	"SEVERITY_NUMBER_TRACE4":      4,
	"SEVERITY_NUMBER_DEBUG":       5,
	"SEVERITY_NUMBER_DEBUG2":      6,
	"SEVERITY_NUMBER_DEBUG3":      7,
	"SEVERITY_NUMBER_DEBUG4":      8,
	"SEVERITY_NUMBER_INFO":        9,
	"SEVERITY_NUMBER_INFO2":       10,
	"SEVERITY_NUMBER_INFO3":       11,
	"SEVERITY_NUMBER_INFO4":       12,
	"SEVERITY_NUMBER_WARN":        13,
	"SEVERITY_NUMBER_WARN2":       14,
	"SEVERITY_NUMBER_WARN3":       15,
	"SEVERITY_NUMBER_WARN4":       16,
	"SEVERITY_NUMBER_ERROR":       17,
	"SEVERITY_NUMBER_ERROR2":      18,
	"SEVERITY_NUMBER_ERROR3":      19,
	"SEVERITY_NUMBER_ERROR4":      20,
	"SEVERITY_NUMBER_FATAL":       21,
	"SEVERITY_NUMBER_FATAL2":      22,
	"SEVERITY_NUMBER_FATAL3":      23,
	"SEVERITY_NUMBER_FATAL4":      24,
}

func (x SeverityNumber) String() string {
	return proto.EnumName(SeverityNumber_name, int32(x))
}

func (SeverityNumber) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{0}
}

// Masks for LogRecord.flags field.
type LogRecordFlags int32

const (
	LogRecordFlags_LOG_RECORD_FLAG_UNSPECIFIED      LogRecordFlags = 0
	LogRecordFlags_LOG_RECORD_FLAG_TRACE_FLAGS_MASK LogRecordFlags = 255
)

var LogRecordFlags_name = map[int32]string{
	0:   "LOG_RECORD_FLAG_UNSPECIFIED",
	255: "LOG_RECORD_FLAG_TRACE_FLAGS_MASK",
}

var LogRecordFlags_value = map[string]int32{
	"LOG_RECORD_FLAG_UNSPECIFIED":      0,
	"LOG_RECORD_FLAG_TRACE_FLAGS_MASK": 255,
}

func (x LogRecordFlags) String() string {
	return proto.EnumName(LogRecordFlags_name, int32(x))
}

func (LogRecordFlags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{1}
}

// A collection of InstrumentationLibraryLogs from a Resource.
type ResourceLogs struct {
	// The resource for the logs in this message.
	// If this field is not set then resource info is unknown.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// A list of InstrumentationLibraryLogs that originate from a resource.
	InstrumentationLibraryLogs []*InstrumentationLibraryLogs `protobuf:"bytes,2,rep,name=instrumentation_library_logs,json=instrumentationLibraryLogs,proto3" json:"instrumentation_library_logs,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                      `json:"-"`
	XXX_unrecognized           []byte                        `json:"-"`
	XXX_sizecache              int32                         `json:"-"`
}

func (m *ResourceLogs) Reset()         { *m = ResourceLogs{} }
func (m *ResourceLogs) String() string { return proto.CompactTextString(m) }
func (*ResourceLogs) ProtoMessage()    {}
func (*ResourceLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{0}
}
func (m *ResourceLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceLogs.Unmarshal(m, b)
}
func (m *ResourceLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceLogs.Marshal(b, m, deterministic)
}
func (m *ResourceLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceLogs.Merge(m, src)
}
func (m *ResourceLogs) XXX_Size() int {
	return xxx_messageInfo_ResourceLogs.Size(m)
}
func (m *ResourceLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceLogs.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceLogs proto.InternalMessageInfo

func (m *ResourceLogs) GetResource() *v1.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceLogs) GetInstrumentationLibraryLogs() []*InstrumentationLibraryLogs {
	if m != nil {
		return m.InstrumentationLibraryLogs
	}
	return nil
}

// A collection of Logs produced by an InstrumentationLibrary.
type InstrumentationLibraryLogs struct {
	// The instrumentation library information for the logs in this message.
	// Semantically when InstrumentationLibrary isn't set, it is equivalent with
	// an empty instrumentation library name (unknown).
	InstrumentationLibrary *v11.InstrumentationLibrary `protobuf:"bytes,1,opt,name=instrumentation_library,json=instrumentationLibrary,proto3" json:"instrumentation_library,omitempty"`
	// A list of log records.
	Logs                 []*LogRecord `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InstrumentationLibraryLogs) Reset()         { *m = InstrumentationLibraryLogs{} }
func (m *InstrumentationLibraryLogs) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibraryLogs) ProtoMessage()    {}
func (*InstrumentationLibraryLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{1}
}
func (m *InstrumentationLibraryLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentationLibraryLogs.Unmarshal(m, b)
}
func (m *InstrumentationLibraryLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentationLibraryLogs.Marshal(b, m, deterministic)
}
func (m *InstrumentationLibraryLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibraryLogs.Merge(m, src)
}
func (m *InstrumentationLibraryLogs) XXX_Size() int {
	return xxx_messageInfo_InstrumentationLibraryLogs.Size(m)
}
func (m *InstrumentationLibraryLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibraryLogs.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibraryLogs proto.InternalMessageInfo

func (m *InstrumentationLibraryLogs) GetInstrumentationLibrary() *v11.InstrumentationLibrary {
	if m != nil {
		return m.InstrumentationLibrary
	}
	return nil
}

func (m *InstrumentationLibraryLogs) GetLogs() []*LogRecord {
	if m != nil {
		return m.Logs
	}
	return nil
}

// A log record according to OpenTelemetry Log Data Model:
// https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md
type LogRecord struct {
	// time_unix_nano is the time when the event occurred.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	TimeUnixNano uint64 `protobuf:"fixed64,1,opt,name=time_unix_nano,json=timeUnixNano,proto3" json:"time_unix_nano,omitempty"`
	// Numerical value of the severity, normalized to values described in Log Data Model.
	// [Optional].
	SeverityNumber SeverityNumber `protobuf:"varint,2,opt,name=severity_number,json=severityNumber,proto3,enum=internal.opentelemetry.proto.logs.v1.SeverityNumber" json:"severity_number,omitempty"`
	// The severity text (also known as log level). The original string representation as
	// it is known at the source. [Optional].
	SeverityText string `protobuf:"bytes,3,opt,name=severity_text,json=severityText,proto3" json:"severity_text,omitempty"`
	// Short event identifier that does not contain varying parts. Name describes
	// what happened (e.g. "ProcessStarted"). Recommended to be no longer than 50
	// characters. Not guaranteed to be unique in any way. [Optional].
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// A value containing the body of the log record. Can be for example a human-readable
	// string message (including multi-line) describing the event in a free form or it can
	// be a structured data composed of arrays and maps of other values. [Optional].
	Body *v11.AnyValue `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// Additional attributes that describe the specific event occurrence. [Optional].
	Attributes             []*v11.KeyValue `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty"`
	DroppedAttributesCount uint32          `protobuf:"varint,7,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// Flags, a bit field. 8 least significant bits are the trace flags as
	// defined in W3C Trace Context specification. 24 most significant bits are reserved
	// and must be set to 0. Readers must not assume that 24 most significant bits
	// will be zero and must correctly mask the bits when reading 8-bit trace flag (use
	// flags & TRACE_FLAGS_MASK). [Optional].
	Flags uint32 `protobuf:"fixed32,8,opt,name=flags,proto3" json:"flags,omitempty"`
	// A unique identifier for a trace. All logs from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid. Can be set for logs that are part of request processing
	// and have an assigned trace id. [Optional].
	TraceId []byte `protobuf:"bytes,9,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid. Can be set for logs that are part of a particular processing span.
	// If span_id is present trace_id SHOULD be also present. [Optional].
	SpanId               []byte   `protobuf:"bytes,10,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRecord) Reset()         { *m = LogRecord{} }
func (m *LogRecord) String() string { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()    {}
func (*LogRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{2}
}
func (m *LogRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRecord.Unmarshal(m, b)
}
func (m *LogRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRecord.Marshal(b, m, deterministic)
}
func (m *LogRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRecord.Merge(m, src)
}
func (m *LogRecord) XXX_Size() int {
	return xxx_messageInfo_LogRecord.Size(m)
}
func (m *LogRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LogRecord proto.InternalMessageInfo

func (m *LogRecord) GetTimeUnixNano() uint64 {
	if m != nil {
		return m.TimeUnixNano
	}
	return 0
}

func (m *LogRecord) GetSeverityNumber() SeverityNumber {
	if m != nil {
		return m.SeverityNumber
	}
	return SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED
}

func (m *LogRecord) GetSeverityText() string {
	if m != nil {
		return m.SeverityText
	}
	return ""
}

func (m *LogRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogRecord) GetBody() *v11.AnyValue {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *LogRecord) GetAttributes() []*v11.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *LogRecord) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func (m *LogRecord) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *LogRecord) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *LogRecord) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func init() {
	proto.RegisterEnum("internal.opentelemetry.proto.logs.v1.SeverityNumber", SeverityNumber_name, SeverityNumber_value)
	proto.RegisterEnum("internal.opentelemetry.proto.logs.v1.LogRecordFlags", LogRecordFlags_name, LogRecordFlags_value)
	proto.RegisterType((*ResourceLogs)(nil), "internal.opentelemetry.proto.logs.v1.ResourceLogs")
	proto.RegisterType((*InstrumentationLibraryLogs)(nil), "internal.opentelemetry.proto.logs.v1.InstrumentationLibraryLogs")
	proto.RegisterType((*LogRecord)(nil), "internal.opentelemetry.proto.logs.v1.LogRecord")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/logs/v1/logs.proto", fileDescriptor_d1c030a3ec7e961e)
}

var fileDescriptor_d1c030a3ec7e961e = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x6e, 0xe2, 0x46,
	0x14, 0xc6, 0xeb, 0x84, 0x40, 0x32, 0x21, 0xec, 0x74, 0x9a, 0x4d, 0xbc, 0x6c, 0xa5, 0xb5, 0xd2,
	0x6d, 0x8b, 0x22, 0xd5, 0x34, 0x86, 0x8b, 0xf6, 0xa6, 0xaa, 0x01, 0x13, 0xa1, 0x65, 0x01, 0x0d,
	0xb0, 0xfd, 0x23, 0x55, 0x96, 0x8d, 0x67, 0xe9, 0x48, 0x66, 0x86, 0xda, 0x63, 0x0a, 0xb7, 0x7d,
	0xb5, 0xbe, 0x45, 0x6f, 0xab, 0x3e, 0x47, 0x2b, 0x0f, 0x7f, 0xba, 0x20, 0x4f, 0xc4, 0x15, 0x33,
	0xe7, 0x77, 0xbe, 0x6f, 0x8e, 0x3f, 0xcc, 0x00, 0xbe, 0xe0, 0x73, 0xc2, 0x04, 0x09, 0xc9, 0x8c,
	0x88, 0x68, 0x55, 0x9d, 0x47, 0x5c, 0xf0, 0x6a, 0xc8, 0xa7, 0x71, 0x75, 0xf1, 0x20, 0x3f, 0x4d,
	0x59, 0x42, 0xaf, 0x29, 0x13, 0x24, 0x62, 0x5e, 0x68, 0xee, 0x09, 0xd6, 0xd4, 0x94, 0x8d, 0x8b,
	0x87, 0xf2, 0x7d, 0x96, 0xdb, 0x84, 0xcf, 0x66, 0x9c, 0xa5, 0x7e, 0xeb, 0xd5, 0x5a, 0x53, 0x36,
	0xb3, 0x7a, 0x23, 0x12, 0xf3, 0x24, 0x9a, 0x90, 0xb4, 0x7b, 0xbb, 0x5e, 0xf7, 0xdf, 0xfd, 0xad,
	0x81, 0x22, 0xde, 0x94, 0xba, 0x7c, 0x1a, 0xa3, 0x1e, 0x38, 0xdf, 0xb6, 0xe8, 0x9a, 0xa1, 0x55,
	0x2e, 0x2d, 0xcb, 0x7c, 0x72, 0xca, 0x9d, 0xe1, 0xe2, 0xc1, 0xdc, 0x3a, 0xe1, 0x9d, 0x07, 0xfa,
	0x43, 0x03, 0x9f, 0x52, 0x16, 0x8b, 0x28, 0x99, 0x11, 0x26, 0x3c, 0x41, 0x39, 0x73, 0x43, 0xea,
	0x47, 0x5e, 0xb4, 0x72, 0xd3, 0x07, 0xd4, 0x4f, 0x8c, 0xd3, 0xca, 0xa5, 0xf5, 0xbd, 0x79, 0x4c,
	0x14, 0x66, 0x67, 0xdf, 0xa9, 0xbb, 0x36, 0x4a, 0x07, 0xc7, 0x65, 0xaa, 0x64, 0x77, 0x7f, 0x69,
	0xa0, 0xac, 0x96, 0xa2, 0xdf, 0xc1, 0xad, 0x62, 0xc4, 0x4d, 0x04, 0xdf, 0x3d, 0x3d, 0xdd, 0xe6,
	0x1b, 0x50, 0xce, 0x87, 0x6f, 0xb2, 0x67, 0x43, 0x4d, 0x90, 0xfb, 0x20, 0x83, 0xea, 0x71, 0x19,
	0x74, 0xf9, 0x14, 0x93, 0x09, 0x8f, 0x02, 0x2c, 0xc5, 0x77, 0xff, 0x9c, 0x82, 0x8b, 0x5d, 0x0d,
	0xbd, 0x06, 0x25, 0x41, 0x67, 0xc4, 0x4d, 0x18, 0x5d, 0xba, 0xcc, 0x63, 0x5c, 0x3e, 0x42, 0x1e,
	0x17, 0xd3, 0xea, 0x98, 0xd1, 0x65, 0xcf, 0x63, 0x1c, 0xfd, 0x02, 0x9e, 0xc5, 0x64, 0x41, 0x22,
	0x2a, 0x56, 0x2e, 0x4b, 0x66, 0x3e, 0x89, 0xf4, 0x13, 0x43, 0xab, 0x94, 0xac, 0xfa, 0x71, 0x33,
	0x0c, 0x37, 0xe2, 0x9e, 0xd4, 0xe2, 0x52, 0xbc, 0xb7, 0x47, 0x9f, 0x81, 0xab, 0x9d, 0xbd, 0x20,
	0x4b, 0xa1, 0x9f, 0x1a, 0x5a, 0xe5, 0x02, 0x17, 0xb7, 0xc5, 0x11, 0x59, 0x0a, 0x84, 0x40, 0x8e,
	0x79, 0x33, 0xa2, 0xe7, 0x24, 0x93, 0x6b, 0xd4, 0x02, 0x39, 0x9f, 0x07, 0x2b, 0xfd, 0x4c, 0xc6,
	0xfe, 0xf5, 0xb1, 0xb1, 0xdb, 0x6c, 0xf5, 0xce, 0x0b, 0x13, 0x82, 0xa5, 0x1a, 0x0d, 0x00, 0xf0,
	0x84, 0x88, 0xa8, 0x9f, 0x08, 0x12, 0xeb, 0x79, 0x19, 0xee, 0xd1, 0x5e, 0x6f, 0xc8, 0xc6, 0xeb,
	0x03, 0x0f, 0xf4, 0x0d, 0xd0, 0x83, 0x88, 0xcf, 0xe7, 0x24, 0x70, 0xff, 0xaf, 0xba, 0x13, 0x9e,
	0x30, 0xa1, 0x17, 0x0c, 0xad, 0x72, 0x85, 0x6f, 0x36, 0xdc, 0xde, 0xe1, 0x66, 0x4a, 0xd1, 0x35,
	0x38, 0x7b, 0x1f, 0x7a, 0xd3, 0x58, 0x3f, 0x37, 0xb4, 0x4a, 0x01, 0xaf, 0x37, 0xe8, 0x05, 0x38,
	0x17, 0x91, 0x37, 0x21, 0x2e, 0x0d, 0xf4, 0x0b, 0x43, 0xab, 0x14, 0x71, 0x41, 0xee, 0x3b, 0x01,
	0xba, 0x05, 0x85, 0x78, 0xee, 0xb1, 0x94, 0x00, 0x49, 0xf2, 0xe9, 0xb6, 0x13, 0xdc, 0xff, 0x79,
	0x06, 0x4a, 0xfb, 0xb9, 0xa3, 0x57, 0xe0, 0xe5, 0xd0, 0x79, 0xe7, 0xe0, 0xce, 0xe8, 0x27, 0xb7,
	0x37, 0x7e, 0xdb, 0x70, 0xb0, 0x3b, 0xee, 0x0d, 0x07, 0x4e, 0xb3, 0xd3, 0xee, 0x38, 0x2d, 0xf8,
	0x11, 0x7a, 0x01, 0x9e, 0x1f, 0x36, 0x8c, 0xb0, 0xdd, 0x74, 0xa0, 0x86, 0xca, 0xe0, 0x26, 0x13,
	0x59, 0xf0, 0x44, 0xc9, 0x6a, 0xf0, 0x54, 0xc9, 0xea, 0x30, 0x97, 0x75, 0x5c, 0xcb, 0x69, 0x8c,
	0x1f, 0xe1, 0x59, 0x96, 0x4c, 0x22, 0x0b, 0xe6, 0x95, 0xac, 0x06, 0x0b, 0x4a, 0x56, 0x87, 0xe7,
	0x48, 0x07, 0xd7, 0x87, 0xac, 0xd3, 0x6b, 0xf7, 0xe1, 0x45, 0xd6, 0x20, 0x29, 0xb1, 0x20, 0x50,
	0xa1, 0x1a, 0xbc, 0x54, 0xa1, 0x3a, 0x2c, 0x66, 0x1d, 0xf5, 0x83, 0x8d, 0x7b, 0xf0, 0x2a, 0x4b,
	0x94, 0x12, 0x0b, 0x96, 0x54, 0xa8, 0x06, 0x9f, 0xa9, 0x50, 0x1d, 0xc2, 0x2c, 0xe4, 0x60, 0xdc,
	0xc7, 0xf0, 0xe3, 0xac, 0x30, 0x24, 0xb2, 0x20, 0x52, 0xb2, 0x1a, 0xfc, 0x44, 0xc9, 0xea, 0xf0,
	0x3a, 0xeb, 0xb8, 0xb6, 0x3d, 0xb2, 0xbb, 0xf0, 0x79, 0x96, 0x4c, 0x22, 0x0b, 0xde, 0x28, 0x59,
	0x0d, 0xde, 0x2a, 0x59, 0x1d, 0xea, 0xf7, 0x3f, 0x82, 0xd2, 0xee, 0xb2, 0x6a, 0xcb, 0xdf, 0xc2,
	0x2b, 0xf0, 0xb2, 0xdb, 0x7f, 0x74, 0xb1, 0xd3, 0xec, 0xe3, 0x96, 0xdb, 0xee, 0xda, 0x8f, 0x07,
	0x2f, 0xf1, 0xe7, 0xc0, 0x38, 0x6c, 0x90, 0x6f, 0x9c, 0x5c, 0x0e, 0xdd, 0xb7, 0xf6, 0xf0, 0x0d,
	0xfc, 0x57, 0x6b, 0xfc, 0x06, 0xbe, 0xa4, 0xfc, 0xa8, 0xeb, 0xab, 0x91, 0xde, 0x97, 0xf1, 0x20,
	0x2d, 0x0d, 0xb4, 0x9f, 0xbf, 0x9d, 0x52, 0xf1, 0x6b, 0xe2, 0xa7, 0x57, 0x40, 0x95, 0xb2, 0xf7,
	0x61, 0xb2, 0x0c, 0x3c, 0xe1, 0x6d, 0x97, 0xfe, 0x57, 0xdc, 0x8f, 0x49, 0xb4, 0xf0, 0x7c, 0x1a,
	0x52, 0xb1, 0xaa, 0x72, 0x11, 0xce, 0xb7, 0x7f, 0xe5, 0x7e, 0x5e, 0xda, 0xd6, 0xfe, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xac, 0x42, 0x63, 0xdb, 0xf0, 0x07, 0x00, 0x00,
}
