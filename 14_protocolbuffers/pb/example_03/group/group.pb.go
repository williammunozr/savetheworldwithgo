// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: group.proto

package group

import (
	proto "github.com/golang/protobuf/proto"
	user "github.com/juanmanuel-tirado/savetheworldwithgo/14_protocolbuffers/pb/example_03/user"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Category int32

const (
	Category_DEVELOPER Category = 0
	Category_OPERATOR  Category = 1
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "DEVELOPER",
		1: "OPERATOR",
	}
	Category_value = map[string]int32{
		"DEVELOPER": 0,
		"OPERATOR":  1,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_group_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_group_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category Category     `protobuf:"varint,2,opt,name=category,proto3,enum=group.Category" json:"category,omitempty"`
	Score    float32      `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	Users    []*user.User `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Group) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_DEVELOPER
}

func (x *Group) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Group) GetUsers() []*user.User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_group_proto protoreflect.FileDescriptor

var file_group_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7c, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2a, 0x27,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45,
	0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x61, 0x6e, 0x6d, 0x61, 0x6e, 0x75, 0x65, 0x6c,
	0x2d, 0x74, 0x69, 0x72, 0x61, 0x64, 0x6f, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x74, 0x68, 0x65, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x77, 0x69, 0x74, 0x68, 0x67, 0x6f, 0x2f, 0x31, 0x32, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x62,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x30, 0x33, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_proto_rawDescOnce sync.Once
	file_group_proto_rawDescData = file_group_proto_rawDesc
)

func file_group_proto_rawDescGZIP() []byte {
	file_group_proto_rawDescOnce.Do(func() {
		file_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_proto_rawDescData)
	})
	return file_group_proto_rawDescData
}

var file_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_group_proto_goTypes = []interface{}{
	(Category)(0),     // 0: group.Category
	(*Group)(nil),     // 1: group.Group
	(*user.User)(nil), // 2: user.User
}
var file_group_proto_depIdxs = []int32{
	0, // 0: group.Group.category:type_name -> group.Category
	2, // 1: group.Group.users:type_name -> user.User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_group_proto_init() }
func file_group_proto_init() {
	if File_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
			RawDescriptor: file_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_group_proto_goTypes,
		DependencyIndexes: file_group_proto_depIdxs,
		EnumInfos:         file_group_proto_enumTypes,
		MessageInfos:      file_group_proto_msgTypes,
	}.Build()
	File_group_proto = out.File
	file_group_proto_rawDesc = nil
	file_group_proto_goTypes = nil
	file_group_proto_depIdxs = nil
}
