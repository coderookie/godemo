// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: employee.proto

package employeeProto

import (
	user "godemo/models/protobuf_package/user"
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

// getEmployeeInfo方法中rpc返回的proto
type EmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *int64        `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Username *string       `protobuf:"bytes,2,req,name=username" json:"username,omitempty"`
	Company  *user.Company `protobuf:"bytes,3,req,name=company" json:"company,omitempty"`
	Language *string       `protobuf:"bytes,4,opt,name=language" json:"language,omitempty"`
}

func (x *EmployeeResponse) Reset() {
	*x = EmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeResponse) ProtoMessage() {}

func (x *EmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeResponse.ProtoReflect.Descriptor instead.
func (*EmployeeResponse) Descriptor() ([]byte, []int) {
	return file_employee_proto_rawDescGZIP(), []int{0}
}

func (x *EmployeeResponse) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *EmployeeResponse) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *EmployeeResponse) GetCompany() *user.Company {
	if x != nil {
		return x.Company
	}
	return nil
}

func (x *EmployeeResponse) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

// getEmployeeInfo方法中rpc请求的proto
type EmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
}

func (x *EmployeeRequest) Reset() {
	*x = EmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeRequest) ProtoMessage() {}

func (x *EmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeRequest.ProtoReflect.Descriptor instead.
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return file_employee_proto_rawDescGZIP(), []int{1}
}

func (x *EmployeeRequest) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

var File_employee_proto protoreflect.FileDescriptor

var file_employee_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x10,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x0f,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x6f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x3b, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_employee_proto_rawDescOnce sync.Once
	file_employee_proto_rawDescData = file_employee_proto_rawDesc
)

func file_employee_proto_rawDescGZIP() []byte {
	file_employee_proto_rawDescOnce.Do(func() {
		file_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_employee_proto_rawDescData)
	})
	return file_employee_proto_rawDescData
}

var file_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_employee_proto_goTypes = []interface{}{
	(*EmployeeResponse)(nil), // 0: EmployeeResponse
	(*EmployeeRequest)(nil),  // 1: EmployeeRequest
	(*user.Company)(nil),     // 2: Company
}
var file_employee_proto_depIdxs = []int32{
	2, // 0: EmployeeResponse.company:type_name -> Company
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_employee_proto_init() }
func file_employee_proto_init() {
	if File_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeResponse); i {
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
		file_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeRequest); i {
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
			RawDescriptor: file_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_employee_proto_goTypes,
		DependencyIndexes: file_employee_proto_depIdxs,
		MessageInfos:      file_employee_proto_msgTypes,
	}.Build()
	File_employee_proto = out.File
	file_employee_proto_rawDesc = nil
	file_employee_proto_goTypes = nil
	file_employee_proto_depIdxs = nil
}
