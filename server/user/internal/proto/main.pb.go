// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: main.proto

package proto

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

type ArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article    *Article    `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	UserConfig *UserConfig `protobuf:"bytes,2,opt,name=userConfig,proto3" json:"userConfig,omitempty"`
}

func (x *ArticleResponse) Reset() {
	*x = ArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleResponse) ProtoMessage() {}

func (x *ArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleResponse.ProtoReflect.Descriptor instead.
func (*ArticleResponse) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

func (x *ArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

func (x *ArticleResponse) GetUserConfig() *UserConfig {
	if x != nil {
		return x.UserConfig
	}
	return nil
}

var File_main_proto protoreflect.FileDescriptor

var file_main_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61,
	0x69, 0x6e, 0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x71, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x32, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x32, 0xd3, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x29, 0x0a, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x32, 0x44, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0e, 0x5a, 0x0c, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_proto_rawDescOnce sync.Once
	file_main_proto_rawDescData = file_main_proto_rawDesc
)

func file_main_proto_rawDescGZIP() []byte {
	file_main_proto_rawDescOnce.Do(func() {
		file_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_proto_rawDescData)
	})
	return file_main_proto_rawDescData
}

var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_main_proto_goTypes = []interface{}{
	(*ArticleResponse)(nil), // 0: main.ArticleResponse
	(*Article)(nil),         // 1: content.Article
	(*UserConfig)(nil),      // 2: config.UserConfig
	(*Empty)(nil),           // 3: auth.Empty
	(*Username)(nil),        // 4: auth.Username
	(*EntryName)(nil),       // 5: content.EntryName
	(*AuthResponse)(nil),    // 6: auth.AuthResponse
	(*User)(nil),            // 7: auth.User
}
var file_main_proto_depIdxs = []int32{
	1, // 0: main.ArticleResponse.article:type_name -> content.Article
	2, // 1: main.ArticleResponse.userConfig:type_name -> config.UserConfig
	3, // 2: main.Auth.Auth:input_type -> auth.Empty
	4, // 3: main.Auth.UpdateUsername:input_type -> auth.Username
	2, // 4: main.Auth.UpdateUserConfig:input_type -> config.UserConfig
	3, // 5: main.Auth.GetUserConfig:input_type -> auth.Empty
	5, // 6: main.Content.GetArticle:input_type -> content.EntryName
	6, // 7: main.Auth.Auth:output_type -> auth.AuthResponse
	7, // 8: main.Auth.UpdateUsername:output_type -> auth.User
	2, // 9: main.Auth.UpdateUserConfig:output_type -> config.UserConfig
	2, // 10: main.Auth.GetUserConfig:output_type -> config.UserConfig
	0, // 11: main.Content.GetArticle:output_type -> main.ArticleResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_main_proto_init() }
func file_main_proto_init() {
	if File_main_proto != nil {
		return
	}
	file_auth_proto_init()
	file_content_proto_init()
	file_user_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleResponse); i {
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
			RawDescriptor: file_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_main_proto_goTypes,
		DependencyIndexes: file_main_proto_depIdxs,
		MessageInfos:      file_main_proto_msgTypes,
	}.Build()
	File_main_proto = out.File
	file_main_proto_rawDesc = nil
	file_main_proto_goTypes = nil
	file_main_proto_depIdxs = nil
}
