// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: market.proto

package market

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Calendar int32

const (
	Calendar_COMPLETE Calendar = 0
	Calendar_NYSE     Calendar = 1
	Calendar_TSX      Calendar = 2
)

// Enum value maps for Calendar.
var (
	Calendar_name = map[int32]string{
		0: "COMPLETE",
		1: "NYSE",
		2: "TSX",
	}
	Calendar_value = map[string]int32{
		"COMPLETE": 0,
		"NYSE":     1,
		"TSX":      2,
	}
)

func (x Calendar) Enum() *Calendar {
	p := new(Calendar)
	*p = x
	return p
}

func (x Calendar) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Calendar) Descriptor() protoreflect.EnumDescriptor {
	return file_market_proto_enumTypes[0].Descriptor()
}

func (Calendar) Type() protoreflect.EnumType {
	return &file_market_proto_enumTypes[0]
}

func (x Calendar) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Calendar.Descriptor instead.
func (Calendar) EnumDescriptor() ([]byte, []int) {
	return file_market_proto_rawDescGZIP(), []int{0}
}

type RangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol   string               `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	First    *timestamp.Timestamp `protobuf:"bytes,2,opt,name=first,proto3" json:"first,omitempty"`
	Last     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last,proto3" json:"last,omitempty"`
	Calendar Calendar             `protobuf:"varint,4,opt,name=calendar,proto3,enum=market.Calendar" json:"calendar,omitempty"`
}

func (x *RangeRequest) Reset() {
	*x = RangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeRequest) ProtoMessage() {}

func (x *RangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_market_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeRequest.ProtoReflect.Descriptor instead.
func (*RangeRequest) Descriptor() ([]byte, []int) {
	return file_market_proto_rawDescGZIP(), []int{0}
}

func (x *RangeRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *RangeRequest) GetFirst() *timestamp.Timestamp {
	if x != nil {
		return x.First
	}
	return nil
}

func (x *RangeRequest) GetLast() *timestamp.Timestamp {
	if x != nil {
		return x.Last
	}
	return nil
}

func (x *RangeRequest) GetCalendar() Calendar {
	if x != nil {
		return x.Calendar
	}
	return Calendar_COMPLETE
}

type DataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Double    float64              `protobuf:"fixed64,2,opt,name=double,proto3" json:"double,omitempty"`
}

func (x *DataPoint) Reset() {
	*x = DataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPoint) ProtoMessage() {}

func (x *DataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_market_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPoint.ProtoReflect.Descriptor instead.
func (*DataPoint) Descriptor() ([]byte, []int) {
	return file_market_proto_rawDescGZIP(), []int{1}
}

func (x *DataPoint) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DataPoint) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

type TimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*DataPoint `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *TimeSeries) Reset() {
	*x = TimeSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeries) ProtoMessage() {}

func (x *TimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_market_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeries.ProtoReflect.Descriptor instead.
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return file_market_proto_rawDescGZIP(), []int{2}
}

func (x *TimeSeries) GetData() []*DataPoint {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_market_proto protoreflect.FileDescriptor

var file_market_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x30, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x6c, 0x61,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x22, 0x5d, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x22,
	0x33, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x25, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x2a, 0x2b, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x59, 0x53, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x53, 0x58, 0x10,
	0x02, 0x32, 0x7d, 0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x33, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x75, 0x63, 0x6b, 0x6c, 0x65, 0x73, 0x73, 0x2d, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_market_proto_rawDescOnce sync.Once
	file_market_proto_rawDescData = file_market_proto_rawDesc
)

func file_market_proto_rawDescGZIP() []byte {
	file_market_proto_rawDescOnce.Do(func() {
		file_market_proto_rawDescData = protoimpl.X.CompressGZIP(file_market_proto_rawDescData)
	})
	return file_market_proto_rawDescData
}

var file_market_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_market_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_market_proto_goTypes = []interface{}{
	(Calendar)(0),               // 0: market.Calendar
	(*RangeRequest)(nil),        // 1: market.RangeRequest
	(*DataPoint)(nil),           // 2: market.DataPoint
	(*TimeSeries)(nil),          // 3: market.TimeSeries
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_market_proto_depIdxs = []int32{
	4, // 0: market.RangeRequest.first:type_name -> google.protobuf.Timestamp
	4, // 1: market.RangeRequest.last:type_name -> google.protobuf.Timestamp
	0, // 2: market.RangeRequest.calendar:type_name -> market.Calendar
	4, // 3: market.DataPoint.timestamp:type_name -> google.protobuf.Timestamp
	2, // 4: market.TimeSeries.data:type_name -> market.DataPoint
	1, // 5: market.MarketData.Query:input_type -> market.RangeRequest
	1, // 6: market.MarketData.QueryStream:input_type -> market.RangeRequest
	3, // 7: market.MarketData.Query:output_type -> market.TimeSeries
	2, // 8: market.MarketData.QueryStream:output_type -> market.DataPoint
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_market_proto_init() }
func file_market_proto_init() {
	if File_market_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_market_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeries); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_market_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_market_proto_goTypes,
		DependencyIndexes: file_market_proto_depIdxs,
		EnumInfos:         file_market_proto_enumTypes,
		MessageInfos:      file_market_proto_msgTypes,
	}.Build()
	File_market_proto = out.File
	file_market_proto_rawDesc = nil
	file_market_proto_goTypes = nil
	file_market_proto_depIdxs = nil
}
