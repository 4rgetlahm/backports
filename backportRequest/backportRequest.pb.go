// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: backportRequest.proto

package backportRequest

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

type BackportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reference    string   `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	Image        string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	BaseBranch   string   `protobuf:"bytes,3,opt,name=base_branch,json=baseBranch,proto3" json:"base_branch,omitempty"`
	TargetBranch string   `protobuf:"bytes,4,opt,name=target_branch,json=targetBranch,proto3" json:"target_branch,omitempty"`
	Commits      []string `protobuf:"bytes,5,rep,name=commits,proto3" json:"commits,omitempty"`
}

func (x *BackportRequest) Reset() {
	*x = BackportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backportRequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackportRequest) ProtoMessage() {}

func (x *BackportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backportRequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackportRequest.ProtoReflect.Descriptor instead.
func (*BackportRequest) Descriptor() ([]byte, []int) {
	return file_backportRequest_proto_rawDescGZIP(), []int{0}
}

func (x *BackportRequest) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *BackportRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *BackportRequest) GetBaseBranch() string {
	if x != nil {
		return x.BaseBranch
	}
	return ""
}

func (x *BackportRequest) GetTargetBranch() string {
	if x != nil {
		return x.TargetBranch
	}
	return ""
}

func (x *BackportRequest) GetCommits() []string {
	if x != nil {
		return x.Commits
	}
	return nil
}

type BackportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
}

func (x *BackportResponse) Reset() {
	*x = BackportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backportRequest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackportResponse) ProtoMessage() {}

func (x *BackportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backportRequest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackportResponse.ProtoReflect.Descriptor instead.
func (*BackportResponse) Descriptor() ([]byte, []int) {
	return file_backportRequest_proto_rawDescGZIP(), []int{1}
}

func (x *BackportResponse) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

var File_backportRequest_proto protoreflect.FileDescriptor

var file_backportRequest_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x22,
	0x24, 0x0a, 0x10, 0x42, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x70, 0x6f, 0x64, 0x32, 0x4c, 0x0a, 0x16, 0x42, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x32, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10,
	0x2e, 0x42, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x34, 0x72, 0x67, 0x65, 0x74, 0x6c, 0x61, 0x68, 0x6d, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backportRequest_proto_rawDescOnce sync.Once
	file_backportRequest_proto_rawDescData = file_backportRequest_proto_rawDesc
)

func file_backportRequest_proto_rawDescGZIP() []byte {
	file_backportRequest_proto_rawDescOnce.Do(func() {
		file_backportRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_backportRequest_proto_rawDescData)
	})
	return file_backportRequest_proto_rawDescData
}

var file_backportRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_backportRequest_proto_goTypes = []interface{}{
	(*BackportRequest)(nil),  // 0: BackportRequest
	(*BackportResponse)(nil), // 1: BackportResponse
}
var file_backportRequest_proto_depIdxs = []int32{
	0, // 0: BackportRequestService.RunBackport:input_type -> BackportRequest
	1, // 1: BackportRequestService.RunBackport:output_type -> BackportResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_backportRequest_proto_init() }
func file_backportRequest_proto_init() {
	if File_backportRequest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backportRequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackportRequest); i {
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
		file_backportRequest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackportResponse); i {
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
			RawDescriptor: file_backportRequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backportRequest_proto_goTypes,
		DependencyIndexes: file_backportRequest_proto_depIdxs,
		MessageInfos:      file_backportRequest_proto_msgTypes,
	}.Build()
	File_backportRequest_proto = out.File
	file_backportRequest_proto_rawDesc = nil
	file_backportRequest_proto_goTypes = nil
	file_backportRequest_proto_depIdxs = nil
}
