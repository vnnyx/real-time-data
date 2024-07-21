// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: pb/vector/vector.proto

package vector

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataVectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *DataVectorRequest) Reset() {
	*x = DataVectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_vector_vector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataVectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataVectorRequest) ProtoMessage() {}

func (x *DataVectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_vector_vector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataVectorRequest.ProtoReflect.Descriptor instead.
func (*DataVectorRequest) Descriptor() ([]byte, []int) {
	return file_pb_vector_vector_proto_rawDescGZIP(), []int{0}
}

func (x *DataVectorRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *DataVectorRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_pb_vector_vector_proto protoreflect.FileDescriptor

var file_pb_vector_vector_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x62, 0x2e, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x32, 0x57, 0x0a, 0x06, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x4d, 0x0a, 0x0f, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x54, 0x6f, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x42, 0x12, 0x1c, 0x2e,
	0x70, 0x62, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6e, 0x6e, 0x79, 0x78, 0x2f, 0x72, 0x65,
	0x61, 0x6c, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x62, 0x2f,
	0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_vector_vector_proto_rawDescOnce sync.Once
	file_pb_vector_vector_proto_rawDescData = file_pb_vector_vector_proto_rawDesc
)

func file_pb_vector_vector_proto_rawDescGZIP() []byte {
	file_pb_vector_vector_proto_rawDescOnce.Do(func() {
		file_pb_vector_vector_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_vector_vector_proto_rawDescData)
	})
	return file_pb_vector_vector_proto_rawDescData
}

var file_pb_vector_vector_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_vector_vector_proto_goTypes = []interface{}{
	(*DataVectorRequest)(nil),    // 0: pb.vector.DataVectorRequest
	(*wrapperspb.BoolValue)(nil), // 1: google.protobuf.BoolValue
}
var file_pb_vector_vector_proto_depIdxs = []int32{
	0, // 0: pb.vector.Vector.StoreToVectorDB:input_type -> pb.vector.DataVectorRequest
	1, // 1: pb.vector.Vector.StoreToVectorDB:output_type -> google.protobuf.BoolValue
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_vector_vector_proto_init() }
func file_pb_vector_vector_proto_init() {
	if File_pb_vector_vector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_vector_vector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataVectorRequest); i {
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
			RawDescriptor: file_pb_vector_vector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_vector_vector_proto_goTypes,
		DependencyIndexes: file_pb_vector_vector_proto_depIdxs,
		MessageInfos:      file_pb_vector_vector_proto_msgTypes,
	}.Build()
	File_pb_vector_vector_proto = out.File
	file_pb_vector_vector_proto_rawDesc = nil
	file_pb_vector_vector_proto_goTypes = nil
	file_pb_vector_vector_proto_depIdxs = nil
}
