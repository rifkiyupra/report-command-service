// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: protos/report_aggregated.proto

package pb

import (
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

type AggregatedMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count           int32  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	SnortDstAddress string `protobuf:"bytes,2,opt,name=snort_dst_address,json=snortDstAddress,proto3" json:"snort_dst_address,omitempty"`
	SnortDstPort    int32  `protobuf:"varint,3,opt,name=snort_dst_port,json=snortDstPort,proto3" json:"snort_dst_port,omitempty"`
	SnortSrcAddress string `protobuf:"bytes,4,opt,name=snort_src_address,json=snortSrcAddress,proto3" json:"snort_src_address,omitempty"`
	SnortSrcPort    int32  `protobuf:"varint,5,opt,name=snort_src_port,json=snortSrcPort,proto3" json:"snort_src_port,omitempty"`
}

func (x *AggregatedMetric) Reset() {
	*x = AggregatedMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_report_aggregated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregatedMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregatedMetric) ProtoMessage() {}

func (x *AggregatedMetric) ProtoReflect() protoreflect.Message {
	mi := &file_protos_report_aggregated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregatedMetric.ProtoReflect.Descriptor instead.
func (*AggregatedMetric) Descriptor() ([]byte, []int) {
	return file_protos_report_aggregated_proto_rawDescGZIP(), []int{0}
}

func (x *AggregatedMetric) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AggregatedMetric) GetSnortDstAddress() string {
	if x != nil {
		return x.SnortDstAddress
	}
	return ""
}

func (x *AggregatedMetric) GetSnortDstPort() int32 {
	if x != nil {
		return x.SnortDstPort
	}
	return 0
}

func (x *AggregatedMetric) GetSnortSrcAddress() string {
	if x != nil {
		return x.SnortSrcAddress
	}
	return ""
}

func (x *AggregatedMetric) GetSnortSrcPort() int32 {
	if x != nil {
		return x.SnortSrcPort
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventMetricsCount   int32               `protobuf:"varint,1,opt,name=event_metrics_count,json=eventMetricsCount,proto3" json:"event_metrics_count,omitempty"`
	Metrics             []*AggregatedMetric `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	SensorId            string              `protobuf:"bytes,3,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	SnortClassification string              `protobuf:"bytes,4,opt,name=snort_classification,json=snortClassification,proto3" json:"snort_classification,omitempty"`
	SnortMessage        string              `protobuf:"bytes,5,opt,name=snort_message,json=snortMessage,proto3" json:"snort_message,omitempty"`
	SnortPriority       int32               `protobuf:"varint,6,opt,name=snort_priority,json=snortPriority,proto3" json:"snort_priority,omitempty"`
	SnortSeconds        int64               `protobuf:"varint,7,opt,name=snort_seconds,json=snortSeconds,proto3" json:"snort_seconds,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_report_aggregated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_protos_report_aggregated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_protos_report_aggregated_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetEventMetricsCount() int32 {
	if x != nil {
		return x.EventMetricsCount
	}
	return 0
}

func (x *Event) GetMetrics() []*AggregatedMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *Event) GetSensorId() string {
	if x != nil {
		return x.SensorId
	}
	return ""
}

func (x *Event) GetSnortClassification() string {
	if x != nil {
		return x.SnortClassification
	}
	return ""
}

func (x *Event) GetSnortMessage() string {
	if x != nil {
		return x.SnortMessage
	}
	return ""
}

func (x *Event) GetSnortPriority() int32 {
	if x != nil {
		return x.SnortPriority
	}
	return 0
}

func (x *Event) GetSnortSeconds() int64 {
	if x != nil {
		return x.SnortSeconds
	}
	return 0
}

type ReportAggregated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ReportAggregated) Reset() {
	*x = ReportAggregated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_report_aggregated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportAggregated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportAggregated) ProtoMessage() {}

func (x *ReportAggregated) ProtoReflect() protoreflect.Message {
	mi := &file_protos_report_aggregated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportAggregated.ProtoReflect.Descriptor instead.
func (*ReportAggregated) Descriptor() ([]byte, []int) {
	return file_protos_report_aggregated_proto_rawDescGZIP(), []int{2}
}

func (x *ReportAggregated) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_protos_report_aggregated_proto protoreflect.FileDescriptor

var file_protos_report_aggregated_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0xcc, 0x01, 0x0a, 0x10, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2a, 0x0a, 0x11, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6e, 0x6f, 0x72,
	0x74, 0x44, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x73,
	0x6e, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x44, 0x73, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x72, 0x63, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6e,
	0x6f, 0x72, 0x74, 0x53, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a,
	0x0e, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x53, 0x72, 0x63, 0x50,
	0x6f, 0x72, 0x74, 0x22, 0xa8, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x14, 0x73, 0x6e,
	0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x6e, 0x6f, 0x72,
	0x74, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6e, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x73, 0x6e, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x35,
	0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_report_aggregated_proto_rawDescOnce sync.Once
	file_protos_report_aggregated_proto_rawDescData = file_protos_report_aggregated_proto_rawDesc
)

func file_protos_report_aggregated_proto_rawDescGZIP() []byte {
	file_protos_report_aggregated_proto_rawDescOnce.Do(func() {
		file_protos_report_aggregated_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_report_aggregated_proto_rawDescData)
	})
	return file_protos_report_aggregated_proto_rawDescData
}

var file_protos_report_aggregated_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_report_aggregated_proto_goTypes = []interface{}{
	(*AggregatedMetric)(nil), // 0: pb.AggregatedMetric
	(*Event)(nil),            // 1: pb.Event
	(*ReportAggregated)(nil), // 2: pb.ReportAggregated
}
var file_protos_report_aggregated_proto_depIdxs = []int32{
	0, // 0: pb.Event.metrics:type_name -> pb.AggregatedMetric
	1, // 1: pb.ReportAggregated.events:type_name -> pb.Event
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_report_aggregated_proto_init() }
func file_protos_report_aggregated_proto_init() {
	if File_protos_report_aggregated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_report_aggregated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregatedMetric); i {
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
		file_protos_report_aggregated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_protos_report_aggregated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportAggregated); i {
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
			RawDescriptor: file_protos_report_aggregated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_report_aggregated_proto_goTypes,
		DependencyIndexes: file_protos_report_aggregated_proto_depIdxs,
		MessageInfos:      file_protos_report_aggregated_proto_msgTypes,
	}.Build()
	File_protos_report_aggregated_proto = out.File
	file_protos_report_aggregated_proto_rawDesc = nil
	file_protos_report_aggregated_proto_goTypes = nil
	file_protos_report_aggregated_proto_depIdxs = nil
}
