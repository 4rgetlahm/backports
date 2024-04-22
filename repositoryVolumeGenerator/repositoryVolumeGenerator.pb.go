// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: repositoryVolumeGenerator.proto

package repositoryVolumeGenerator

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

type GenerateRepositoryVolumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeName string `protobuf:"bytes,1,opt,name=volume_name,json=volumeName,proto3" json:"volume_name,omitempty"`
}

func (x *GenerateRepositoryVolumeRequest) Reset() {
	*x = GenerateRepositoryVolumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repositoryVolumeGenerator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRepositoryVolumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRepositoryVolumeRequest) ProtoMessage() {}

func (x *GenerateRepositoryVolumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repositoryVolumeGenerator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRepositoryVolumeRequest.ProtoReflect.Descriptor instead.
func (*GenerateRepositoryVolumeRequest) Descriptor() ([]byte, []int) {
	return file_repositoryVolumeGenerator_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateRepositoryVolumeRequest) GetVolumeName() string {
	if x != nil {
		return x.VolumeName
	}
	return ""
}

type GenerateVolumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeName  string `protobuf:"bytes,1,opt,name=volume_name,json=volumeName,proto3" json:"volume_name,omitempty"`
	ContainerId string `protobuf:"bytes,2,opt,name=containerId,proto3" json:"containerId,omitempty"`
}

func (x *GenerateVolumeResponse) Reset() {
	*x = GenerateVolumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repositoryVolumeGenerator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateVolumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateVolumeResponse) ProtoMessage() {}

func (x *GenerateVolumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repositoryVolumeGenerator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateVolumeResponse.ProtoReflect.Descriptor instead.
func (*GenerateVolumeResponse) Descriptor() ([]byte, []int) {
	return file_repositoryVolumeGenerator_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateVolumeResponse) GetVolumeName() string {
	if x != nil {
		return x.VolumeName
	}
	return ""
}

func (x *GenerateVolumeResponse) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

var File_repositoryVolumeGenerator_proto protoreflect.FileDescriptor

var file_repositoryVolumeGenerator_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x42, 0x0a, 0x1f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x32, 0x6a, 0x0a, 0x21, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x34, 0x72, 0x67,
	0x65, 0x74, 0x6c, 0x61, 0x68, 0x6d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_repositoryVolumeGenerator_proto_rawDescOnce sync.Once
	file_repositoryVolumeGenerator_proto_rawDescData = file_repositoryVolumeGenerator_proto_rawDesc
)

func file_repositoryVolumeGenerator_proto_rawDescGZIP() []byte {
	file_repositoryVolumeGenerator_proto_rawDescOnce.Do(func() {
		file_repositoryVolumeGenerator_proto_rawDescData = protoimpl.X.CompressGZIP(file_repositoryVolumeGenerator_proto_rawDescData)
	})
	return file_repositoryVolumeGenerator_proto_rawDescData
}

var file_repositoryVolumeGenerator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_repositoryVolumeGenerator_proto_goTypes = []interface{}{
	(*GenerateRepositoryVolumeRequest)(nil), // 0: GenerateRepositoryVolumeRequest
	(*GenerateVolumeResponse)(nil),          // 1: GenerateVolumeResponse
}
var file_repositoryVolumeGenerator_proto_depIdxs = []int32{
	0, // 0: RepositoryVolumeGenerationService.Generate:input_type -> GenerateRepositoryVolumeRequest
	1, // 1: RepositoryVolumeGenerationService.Generate:output_type -> GenerateVolumeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_repositoryVolumeGenerator_proto_init() }
func file_repositoryVolumeGenerator_proto_init() {
	if File_repositoryVolumeGenerator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repositoryVolumeGenerator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRepositoryVolumeRequest); i {
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
		file_repositoryVolumeGenerator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateVolumeResponse); i {
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
			RawDescriptor: file_repositoryVolumeGenerator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repositoryVolumeGenerator_proto_goTypes,
		DependencyIndexes: file_repositoryVolumeGenerator_proto_depIdxs,
		MessageInfos:      file_repositoryVolumeGenerator_proto_msgTypes,
	}.Build()
	File_repositoryVolumeGenerator_proto = out.File
	file_repositoryVolumeGenerator_proto_rawDesc = nil
	file_repositoryVolumeGenerator_proto_goTypes = nil
	file_repositoryVolumeGenerator_proto_depIdxs = nil
}
