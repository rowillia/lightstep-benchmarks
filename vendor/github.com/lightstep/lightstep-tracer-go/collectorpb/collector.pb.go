// Code generated by protoc-gen-go.
// source: collector.proto
// DO NOT EDIT!

/*
Package collectorpb is a generated protocol buffer package.

It is generated from these files:
	collector.proto

It has these top-level messages:
	SpanContext
	KeyValue
	Log
	Reference
	Span
	Tracer
	MetricsSample
	InternalMetrics
	Auth
	ReportRequest
	Command
	ReportResponse
*/
package collectorpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type Reference_Relationship int32

const (
	Reference_CHILD_OF     Reference_Relationship = 0
	Reference_FOLLOWS_FROM Reference_Relationship = 1
)

var Reference_Relationship_name = map[int32]string{
	0: "CHILD_OF",
	1: "FOLLOWS_FROM",
}
var Reference_Relationship_value = map[string]int32{
	"CHILD_OF":     0,
	"FOLLOWS_FROM": 1,
}

func (x Reference_Relationship) String() string {
	return proto.EnumName(Reference_Relationship_name, int32(x))
}
func (Reference_Relationship) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type SpanContext struct {
	TraceId uint64            `protobuf:"varint,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	SpanId  uint64            `protobuf:"varint,2,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	Baggage map[string]string `protobuf:"bytes,3,rep,name=baggage" json:"baggage,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SpanContext) Reset()                    { *m = SpanContext{} }
func (m *SpanContext) String() string            { return proto.CompactTextString(m) }
func (*SpanContext) ProtoMessage()               {}
func (*SpanContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpanContext) GetBaggage() map[string]string {
	if m != nil {
		return m.Baggage
	}
	return nil
}

type KeyValue struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*KeyValue_StringValue
	//	*KeyValue_IntValue
	//	*KeyValue_DoubleValue
	//	*KeyValue_BoolValue
	Value isKeyValue_Value `protobuf_oneof:"value"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isKeyValue_Value interface {
	isKeyValue_Value()
}

type KeyValue_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,oneof"`
}
type KeyValue_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,oneof"`
}
type KeyValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,oneof"`
}
type KeyValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,oneof"`
}

func (*KeyValue_StringValue) isKeyValue_Value() {}
func (*KeyValue_IntValue) isKeyValue_Value()    {}
func (*KeyValue_DoubleValue) isKeyValue_Value() {}
func (*KeyValue_BoolValue) isKeyValue_Value()   {}

func (m *KeyValue) GetValue() isKeyValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyValue) GetStringValue() string {
	if x, ok := m.GetValue().(*KeyValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *KeyValue) GetIntValue() int64 {
	if x, ok := m.GetValue().(*KeyValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *KeyValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*KeyValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *KeyValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*KeyValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KeyValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _KeyValue_OneofMarshaler, _KeyValue_OneofUnmarshaler, _KeyValue_OneofSizer, []interface{}{
		(*KeyValue_StringValue)(nil),
		(*KeyValue_IntValue)(nil),
		(*KeyValue_DoubleValue)(nil),
		(*KeyValue_BoolValue)(nil),
	}
}

func _KeyValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*KeyValue)
	// value
	switch x := m.Value.(type) {
	case *KeyValue_StringValue:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *KeyValue_IntValue:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *KeyValue_DoubleValue:
		b.EncodeVarint(4<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *KeyValue_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("KeyValue.Value has unexpected type %T", x)
	}
	return nil
}

func _KeyValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*KeyValue)
	switch tag {
	case 2: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &KeyValue_StringValue{x}
		return true, err
	case 3: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &KeyValue_IntValue{int64(x)}
		return true, err
	case 4: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &KeyValue_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 5: // value.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &KeyValue_BoolValue{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _KeyValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*KeyValue)
	// value
	switch x := m.Value.(type) {
	case *KeyValue_StringValue:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *KeyValue_IntValue:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *KeyValue_DoubleValue:
		n += proto.SizeVarint(4<<3 | proto.WireFixed64)
		n += 8
	case *KeyValue_BoolValue:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Log struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Keyvalues []*KeyValue                `protobuf:"bytes,2,rep,name=keyvalues" json:"keyvalues,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Log) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Log) GetKeyvalues() []*KeyValue {
	if m != nil {
		return m.Keyvalues
	}
	return nil
}

type Reference struct {
	Relationship Reference_Relationship `protobuf:"varint,1,opt,name=relationship,enum=lightstep.collector.Reference_Relationship" json:"relationship,omitempty"`
	SpanContext  *SpanContext           `protobuf:"bytes,2,opt,name=span_context,json=spanContext" json:"span_context,omitempty"`
}

func (m *Reference) Reset()                    { *m = Reference{} }
func (m *Reference) String() string            { return proto.CompactTextString(m) }
func (*Reference) ProtoMessage()               {}
func (*Reference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Reference) GetSpanContext() *SpanContext {
	if m != nil {
		return m.SpanContext
	}
	return nil
}

type Span struct {
	SpanContext    *SpanContext               `protobuf:"bytes,1,opt,name=span_context,json=spanContext" json:"span_context,omitempty"`
	OperationName  string                     `protobuf:"bytes,2,opt,name=operation_name,json=operationName" json:"operation_name,omitempty"`
	References     []*Reference               `protobuf:"bytes,3,rep,name=references" json:"references,omitempty"`
	StartTimestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	DurationMicros uint64                     `protobuf:"varint,5,opt,name=duration_micros,json=durationMicros" json:"duration_micros,omitempty"`
	Tags           []*KeyValue                `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	Logs           []*Log                     `protobuf:"bytes,7,rep,name=logs" json:"logs,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Span) GetSpanContext() *SpanContext {
	if m != nil {
		return m.SpanContext
	}
	return nil
}

func (m *Span) GetReferences() []*Reference {
	if m != nil {
		return m.References
	}
	return nil
}

func (m *Span) GetStartTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartTimestamp
	}
	return nil
}

func (m *Span) GetTags() []*KeyValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Span) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type Tracer struct {
	TracerId uint64      `protobuf:"varint,1,opt,name=tracer_id,json=tracerId" json:"tracer_id,omitempty"`
	Tags     []*KeyValue `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
}

func (m *Tracer) Reset()                    { *m = Tracer{} }
func (m *Tracer) String() string            { return proto.CompactTextString(m) }
func (*Tracer) ProtoMessage()               {}
func (*Tracer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Tracer) GetTags() []*KeyValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

type MetricsSample struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*MetricsSample_IntValue
	//	*MetricsSample_DoubleValue
	Value isMetricsSample_Value `protobuf_oneof:"value"`
}

func (m *MetricsSample) Reset()                    { *m = MetricsSample{} }
func (m *MetricsSample) String() string            { return proto.CompactTextString(m) }
func (*MetricsSample) ProtoMessage()               {}
func (*MetricsSample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isMetricsSample_Value interface {
	isMetricsSample_Value()
}

type MetricsSample_IntValue struct {
	IntValue int64 `protobuf:"varint,2,opt,name=int_value,json=intValue,oneof"`
}
type MetricsSample_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,oneof"`
}

func (*MetricsSample_IntValue) isMetricsSample_Value()    {}
func (*MetricsSample_DoubleValue) isMetricsSample_Value() {}

func (m *MetricsSample) GetValue() isMetricsSample_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricsSample) GetIntValue() int64 {
	if x, ok := m.GetValue().(*MetricsSample_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *MetricsSample) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*MetricsSample_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetricsSample) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetricsSample_OneofMarshaler, _MetricsSample_OneofUnmarshaler, _MetricsSample_OneofSizer, []interface{}{
		(*MetricsSample_IntValue)(nil),
		(*MetricsSample_DoubleValue)(nil),
	}
}

func _MetricsSample_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetricsSample)
	// value
	switch x := m.Value.(type) {
	case *MetricsSample_IntValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *MetricsSample_DoubleValue:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case nil:
	default:
		return fmt.Errorf("MetricsSample.Value has unexpected type %T", x)
	}
	return nil
}

func _MetricsSample_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetricsSample)
	switch tag {
	case 2: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &MetricsSample_IntValue{int64(x)}
		return true, err
	case 3: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &MetricsSample_DoubleValue{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _MetricsSample_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetricsSample)
	// value
	switch x := m.Value.(type) {
	case *MetricsSample_IntValue:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *MetricsSample_DoubleValue:
		n += proto.SizeVarint(3<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InternalMetrics struct {
	StartTimestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	DurationMicros uint64                     `protobuf:"varint,2,opt,name=duration_micros,json=durationMicros" json:"duration_micros,omitempty"`
	Logs           []*Log                     `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
	Counts         []*MetricsSample           `protobuf:"bytes,4,rep,name=counts" json:"counts,omitempty"`
	Gauges         []*MetricsSample           `protobuf:"bytes,5,rep,name=gauges" json:"gauges,omitempty"`
}

func (m *InternalMetrics) Reset()                    { *m = InternalMetrics{} }
func (m *InternalMetrics) String() string            { return proto.CompactTextString(m) }
func (*InternalMetrics) ProtoMessage()               {}
func (*InternalMetrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *InternalMetrics) GetStartTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartTimestamp
	}
	return nil
}

func (m *InternalMetrics) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *InternalMetrics) GetCounts() []*MetricsSample {
	if m != nil {
		return m.Counts
	}
	return nil
}

func (m *InternalMetrics) GetGauges() []*MetricsSample {
	if m != nil {
		return m.Gauges
	}
	return nil
}

type Auth struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ReportRequest struct {
	Tracer                *Tracer          `protobuf:"bytes,1,opt,name=tracer" json:"tracer,omitempty"`
	Auth                  *Auth            `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
	Spans                 []*Span          `protobuf:"bytes,3,rep,name=spans" json:"spans,omitempty"`
	TimestampOffsetMicros uint32           `protobuf:"varint,5,opt,name=timestamp_offset_micros,json=timestampOffsetMicros" json:"timestamp_offset_micros,omitempty"`
	InternalMetrics       *InternalMetrics `protobuf:"bytes,6,opt,name=internal_metrics,json=internalMetrics" json:"internal_metrics,omitempty"`
}

func (m *ReportRequest) Reset()                    { *m = ReportRequest{} }
func (m *ReportRequest) String() string            { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()               {}
func (*ReportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReportRequest) GetTracer() *Tracer {
	if m != nil {
		return m.Tracer
	}
	return nil
}

func (m *ReportRequest) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *ReportRequest) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *ReportRequest) GetInternalMetrics() *InternalMetrics {
	if m != nil {
		return m.InternalMetrics
	}
	return nil
}

type Command struct {
	Disable bool `protobuf:"varint,1,opt,name=disable" json:"disable,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type ReportResponse struct {
	Commands          []*Command                 `protobuf:"bytes,1,rep,name=commands" json:"commands,omitempty"`
	ReceiveTimestamp  *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=receive_timestamp,json=receiveTimestamp" json:"receive_timestamp,omitempty"`
	TransmitTimestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=transmit_timestamp,json=transmitTimestamp" json:"transmit_timestamp,omitempty"`
	Errors            []string                   `protobuf:"bytes,4,rep,name=errors" json:"errors,omitempty"`
}

func (m *ReportResponse) Reset()                    { *m = ReportResponse{} }
func (m *ReportResponse) String() string            { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()               {}
func (*ReportResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReportResponse) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *ReportResponse) GetReceiveTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.ReceiveTimestamp
	}
	return nil
}

func (m *ReportResponse) GetTransmitTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.TransmitTimestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*SpanContext)(nil), "lightstep.collector.SpanContext")
	proto.RegisterType((*KeyValue)(nil), "lightstep.collector.KeyValue")
	proto.RegisterType((*Log)(nil), "lightstep.collector.Log")
	proto.RegisterType((*Reference)(nil), "lightstep.collector.Reference")
	proto.RegisterType((*Span)(nil), "lightstep.collector.Span")
	proto.RegisterType((*Tracer)(nil), "lightstep.collector.Tracer")
	proto.RegisterType((*MetricsSample)(nil), "lightstep.collector.MetricsSample")
	proto.RegisterType((*InternalMetrics)(nil), "lightstep.collector.InternalMetrics")
	proto.RegisterType((*Auth)(nil), "lightstep.collector.Auth")
	proto.RegisterType((*ReportRequest)(nil), "lightstep.collector.ReportRequest")
	proto.RegisterType((*Command)(nil), "lightstep.collector.Command")
	proto.RegisterType((*ReportResponse)(nil), "lightstep.collector.ReportResponse")
	proto.RegisterEnum("lightstep.collector.Reference_Relationship", Reference_Relationship_name, Reference_Relationship_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CollectorService service

type CollectorServiceClient interface {
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type collectorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCollectorServiceClient(cc *grpc.ClientConn) CollectorServiceClient {
	return &collectorServiceClient{cc}
}

func (c *collectorServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := grpc.Invoke(ctx, "/lightstep.collector.CollectorService/Report", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CollectorService service

type CollectorServiceServer interface {
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

func RegisterCollectorServiceServer(s *grpc.Server, srv CollectorServiceServer) {
	s.RegisterService(&_CollectorService_serviceDesc, srv)
}

func _CollectorService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lightstep.collector.CollectorService/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServiceServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lightstep.collector.CollectorService",
	HandlerType: (*CollectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _CollectorService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector.proto",
}

func init() { proto.RegisterFile("collector.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 955 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0x5d, 0xc7, 0x6e, 0x3e, 0x6e, 0xd2, 0x24, 0x3b, 0x7c, 0x6c, 0xda, 0x65, 0xd9, 0xe2, 0x82,
	0x58, 0x04, 0x9b, 0x8a, 0xac, 0x84, 0x56, 0x45, 0x42, 0xa2, 0x81, 0xd2, 0x8a, 0x96, 0x88, 0x69,
	0x05, 0x88, 0x17, 0xcb, 0x71, 0xa6, 0xae, 0xb5, 0x8e, 0xc7, 0x8c, 0x27, 0x15, 0x95, 0x78, 0xe1,
	0xc7, 0xf0, 0x88, 0xf8, 0x17, 0xbc, 0xf3, 0x67, 0x78, 0xe6, 0x7a, 0xc6, 0x5f, 0x05, 0x6f, 0x53,
	0xa1, 0x7d, 0xf3, 0xdc, 0x39, 0x67, 0xe6, 0xce, 0xb9, 0xe7, 0x5e, 0xc3, 0xc0, 0xe3, 0x61, 0xc8,
	0x3c, 0xc9, 0xc5, 0x38, 0x16, 0x5c, 0x72, 0xf2, 0x5a, 0x18, 0xf8, 0x97, 0x32, 0x91, 0x2c, 0x1e,
	0x17, 0x5b, 0xdb, 0x8f, 0x7d, 0xce, 0xfd, 0x90, 0xed, 0x29, 0xc8, 0x7c, 0x75, 0xb1, 0x27, 0x83,
	0x25, 0x4b, 0xa4, 0xbb, 0x8c, 0x35, 0xcb, 0xfe, 0xd3, 0x80, 0xee, 0x59, 0xec, 0x46, 0x53, 0x1e,
	0x49, 0xf6, 0xb3, 0x24, 0x5b, 0xd0, 0x96, 0xc2, 0xf5, 0x98, 0x13, 0x2c, 0x46, 0xc6, 0x8e, 0xf1,
	0xc4, 0xa2, 0x2d, 0xb5, 0x3e, 0x5e, 0x90, 0x07, 0xd0, 0x4a, 0x10, 0x99, 0xee, 0x34, 0xd4, 0x4e,
	0x33, 0x5d, 0xe2, 0xc6, 0x57, 0xd0, 0x9a, 0xbb, 0xbe, 0xef, 0xfa, 0x6c, 0x64, 0xee, 0x98, 0x4f,
	0xba, 0x93, 0xa7, 0xe3, 0x9a, 0x5c, 0xc6, 0x95, 0x6b, 0xc6, 0x07, 0x1a, 0xff, 0x65, 0x24, 0xc5,
	0x35, 0xcd, 0xd9, 0xdb, 0xfb, 0xd0, 0xab, 0x6e, 0x90, 0x21, 0x98, 0x2f, 0xd8, 0xb5, 0xca, 0xa3,
	0x43, 0xd3, 0x4f, 0xf2, 0x3a, 0x6c, 0x5c, 0xb9, 0xe1, 0x8a, 0xa9, 0x0c, 0x3a, 0x54, 0x2f, 0xf6,
	0x1b, 0xcf, 0x0d, 0xfb, 0x0f, 0x03, 0xda, 0x5f, 0xb3, 0xeb, 0xef, 0xd2, 0x40, 0x0d, 0x71, 0x17,
	0x7a, 0x89, 0x14, 0x41, 0xe4, 0x3b, 0x15, 0xfe, 0xd1, 0x3d, 0xda, 0xd5, 0x51, 0x4d, 0x7b, 0x04,
	0x9d, 0x20, 0x92, 0x19, 0xc2, 0x44, 0x84, 0x89, 0x88, 0x36, 0x86, 0xf4, 0x36, 0x9e, 0xb1, 0xe0,
	0xab, 0x79, 0xc8, 0x32, 0x84, 0x85, 0x08, 0x23, 0x3d, 0x43, 0x47, 0x35, 0xe8, 0x31, 0xc0, 0x9c,
	0xf3, 0x30, 0x83, 0x6c, 0x20, 0xa4, 0x8d, 0x90, 0x4e, 0x1a, 0x53, 0x80, 0x83, 0x56, 0xf6, 0x04,
	0xfb, 0x17, 0x30, 0x4f, 0xb8, 0x4f, 0x9e, 0x43, 0xa7, 0x28, 0x8a, 0xca, 0xb8, 0x3b, 0xd9, 0x1e,
	0xeb, 0xb2, 0x8d, 0xf3, 0xb2, 0x8d, 0xcf, 0x73, 0x04, 0x2d, 0xc1, 0xe4, 0x53, 0xe8, 0xe0, 0xd3,
	0xd4, 0x61, 0x09, 0x3e, 0x28, 0x55, 0xfe, 0x51, 0xad, 0xf2, 0xb9, 0x2e, 0xb4, 0xc4, 0xdb, 0x7f,
	0x19, 0xd0, 0xa1, 0xec, 0x82, 0x09, 0x16, 0x79, 0x8c, 0xcc, 0xa0, 0x27, 0x58, 0xe8, 0xca, 0x80,
	0x47, 0xc9, 0x65, 0xa0, 0xf3, 0xe8, 0x4f, 0x3e, 0xac, 0x3d, 0xad, 0x60, 0xe1, 0x57, 0x49, 0xa1,
	0x37, 0x0e, 0x20, 0x53, 0xd4, 0x3b, 0x35, 0x8b, 0xa7, 0x0b, 0xae, 0xf4, 0xee, 0x4e, 0x76, 0xd6,
	0x19, 0x03, 0xeb, 0x51, 0x2e, 0xec, 0x31, 0xf4, 0xaa, 0x57, 0x90, 0x1e, 0xb4, 0xa7, 0x47, 0xc7,
	0x27, 0x5f, 0x38, 0xb3, 0xc3, 0xe1, 0x3d, 0x2c, 0x72, 0xef, 0x70, 0x76, 0x72, 0x32, 0xfb, 0xfe,
	0xcc, 0x39, 0xa4, 0xb3, 0xd3, 0xa1, 0x61, 0xff, 0x6a, 0x82, 0x95, 0x1e, 0xf6, 0x9f, 0xdb, 0x8d,
	0xff, 0x71, 0x3b, 0x79, 0x0f, 0xfa, 0x3c, 0x66, 0x42, 0x5d, 0xef, 0x44, 0xee, 0x32, 0x37, 0xdd,
	0x66, 0x11, 0xfd, 0x06, 0x83, 0xe4, 0x33, 0x00, 0x91, 0x2b, 0x92, 0x64, 0x0d, 0xf0, 0xf6, 0xed,
	0xc2, 0xd1, 0x0a, 0x03, 0x73, 0x1d, 0x60, 0x39, 0x85, 0x74, 0x4a, 0x17, 0x58, 0x6b, 0x5d, 0xd0,
	0x57, 0x94, 0x62, 0x4d, 0xde, 0x87, 0xc1, 0x62, 0x95, 0xa5, 0xba, 0x0c, 0x3c, 0xc1, 0x13, 0x65,
	0x3d, 0x8b, 0xf6, 0xf3, 0xf0, 0xa9, 0x8a, 0x92, 0x8f, 0xc1, 0x92, 0xae, 0x9f, 0x8c, 0x9a, 0x77,
	0xb1, 0x8b, 0x82, 0x92, 0x8f, 0xc0, 0x0a, 0x39, 0x52, 0x5a, 0x8a, 0x32, 0xaa, 0xa5, 0xa0, 0x91,
	0xa9, 0x42, 0xd9, 0x3f, 0x40, 0xf3, 0x3c, 0x1d, 0x18, 0x82, 0x3c, 0x44, 0x63, 0xab, 0xaf, 0x72,
	0x96, 0xe8, 0xd9, 0x22, 0x70, 0x66, 0xe4, 0x79, 0x58, 0x77, 0xce, 0xc3, 0x8e, 0x61, 0xf3, 0x94,
	0x61, 0xb7, 0x7a, 0xc9, 0x19, 0x3e, 0x39, 0x64, 0x84, 0x80, 0xa5, 0xca, 0xa2, 0xdb, 0x5c, 0x7d,
	0xdf, 0x6c, 0xe1, 0xc6, 0xda, 0x16, 0x36, 0x6b, 0x5a, 0xb8, 0xec, 0xd0, 0xdf, 0x1a, 0x30, 0x38,
	0x46, 0x33, 0x88, 0xc8, 0x0d, 0xb3, 0xab, 0xeb, 0xca, 0x65, 0xbc, 0x8a, 0x72, 0x35, 0x6a, 0xcb,
	0x95, 0x6b, 0x6f, 0xde, 0x45, 0x7b, 0xb2, 0x0f, 0x4d, 0x8f, 0xaf, 0x22, 0x99, 0xcb, 0x6a, 0xd7,
	0xe2, 0x6f, 0x88, 0x48, 0x33, 0x46, 0xca, 0xf5, 0xdd, 0x95, 0xcf, 0x52, 0xe3, 0xdc, 0x99, 0xab,
	0x19, 0xf6, 0x07, 0x60, 0x7d, 0xbe, 0x92, 0x97, 0xe4, 0x1d, 0xe8, 0xb9, 0x1e, 0x7a, 0x3a, 0x71,
	0x24, 0x7f, 0xc1, 0xa2, 0xac, 0x30, 0x5d, 0x1d, 0x3b, 0x4f, 0x43, 0xf6, 0xef, 0x0d, 0xd8, 0xa4,
	0x2c, 0xe6, 0x42, 0x52, 0xf6, 0x13, 0xce, 0x21, 0x49, 0x9e, 0x41, 0x53, 0xbb, 0x22, 0xd3, 0xf1,
	0x61, 0xed, 0xc5, 0xda, 0x53, 0x34, 0x83, 0x92, 0xa7, 0x60, 0xb9, 0x78, 0x63, 0x36, 0x56, 0xb6,
	0x6a, 0x29, 0x69, 0x4a, 0x54, 0xc1, 0xc8, 0x1e, 0x6c, 0xa4, 0x9d, 0x9d, 0xeb, 0xb8, 0xf5, 0xd2,
	0x41, 0x40, 0x35, 0x8e, 0x7c, 0x02, 0x0f, 0x8a, 0xfa, 0x3a, 0xfc, 0xe2, 0x22, 0x61, 0xb2, 0xda,
	0x57, 0x9b, 0xf4, 0x8d, 0x62, 0x7b, 0xa6, 0x76, 0xb3, 0x7a, 0xcd, 0x60, 0x18, 0x64, 0x86, 0x71,
	0x96, 0x5a, 0x2b, 0x6c, 0xb5, 0x34, 0xc7, 0x77, 0x6b, 0xef, 0xfc, 0x97, 0xbb, 0xe8, 0x20, 0xb8,
	0x19, 0xb0, 0x77, 0xa1, 0x35, 0xe5, 0xcb, 0xa5, 0x1b, 0x2d, 0xc8, 0x08, 0x5a, 0x8b, 0x20, 0x71,
	0xd1, 0xa6, 0x4a, 0xa9, 0x36, 0xcd, 0x97, 0xf6, 0xdf, 0x06, 0xf4, 0x73, 0x51, 0x93, 0x18, 0x87,
	0x25, 0xc3, 0xbf, 0x4a, 0xdb, 0xd3, 0xbc, 0x04, 0xd1, 0xe9, 0xa3, 0xdf, 0xaa, 0x4d, 0x20, 0x3b,
	0x9c, 0x16, 0x68, 0xfc, 0x9b, 0xdf, 0x17, 0xcc, 0x63, 0xc1, 0x15, 0xab, 0x58, 0xbc, 0xb1, 0xd6,
	0xe2, 0xc3, 0x8c, 0x54, 0x9a, 0xfc, 0x18, 0x08, 0x56, 0x2b, 0x4a, 0x96, 0x41, 0xb5, 0x59, 0xcc,
	0xb5, 0x27, 0xdd, 0xcf, 0x59, 0xe5, 0x51, 0x6f, 0x42, 0x93, 0x09, 0xc1, 0x85, 0x36, 0x76, 0x87,
	0x66, 0xab, 0x09, 0x83, 0xe1, 0x34, 0x7f, 0xca, 0x19, 0x13, 0x57, 0x01, 0xfe, 0xca, 0xbe, 0x85,
	0xa6, 0xd6, 0x82, 0xd8, 0x2f, 0x99, 0xc2, 0x15, 0xf7, 0x6d, 0xef, 0xde, 0x8a, 0xd1, 0x62, 0x1e,
	0x0c, 0x8e, 0x1a, 0x3f, 0x76, 0x8b, 0xdd, 0x78, 0x3e, 0x6f, 0xaa, 0xb4, 0x9f, 0xfd, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xd1, 0xd5, 0xfa, 0x7e, 0x85, 0x09, 0x00, 0x00,
}
