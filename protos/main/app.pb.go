// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.19.6
// source: app.proto

package main

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

type DummyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy []*CheckRequest `protobuf:"bytes,1,rep,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *DummyList) Reset() {
	*x = DummyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyList) ProtoMessage() {}

func (x *DummyList) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyList.ProtoReflect.Descriptor instead.
func (*DummyList) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{0}
}

func (x *DummyList) GetDummy() []*CheckRequest {
	if x != nil {
		return x.Dummy
	}
	return nil
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{1}
}

func (x *CheckRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CheckSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Adress   string `protobuf:"bytes,5,opt,name=adress,proto3" json:"adress,omitempty"`
	Email    string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CheckSet) Reset() {
	*x = CheckSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckSet) ProtoMessage() {}

func (x *CheckSet) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckSet.ProtoReflect.Descriptor instead.
func (*CheckSet) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{2}
}

func (x *CheckSet) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CheckSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CheckSet) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CheckSet) GetAdress() string {
	if x != nil {
		return x.Adress
	}
	return ""
}

func (x *CheckSet) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Emptyb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Emptyb) Reset() {
	*x = Emptyb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emptyb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emptyb) ProtoMessage() {}

func (x *Emptyb) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emptyb.ProtoReflect.Descriptor instead.
func (*Emptyb) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{3}
}

var File_app_proto protoreflect.FileDescriptor

var file_app_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x09, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x6d,
	0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x1e, 0x0a,
	0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7a, 0x0a,
	0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x08, 0x0a, 0x06, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x62, 0x32, 0xab, 0x01, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x76, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x62,
	0x12, 0x2d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x66, 0x6f, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x62, 0x1a, 0x0a, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x1a, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x76, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x09, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x1a, 0x07, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x62, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_proto_rawDescOnce sync.Once
	file_app_proto_rawDescData = file_app_proto_rawDesc
)

func file_app_proto_rawDescGZIP() []byte {
	file_app_proto_rawDescOnce.Do(func() {
		file_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proto_rawDescData)
	})
	return file_app_proto_rawDescData
}

var file_app_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_proto_goTypes = []interface{}{
	(*DummyList)(nil),    // 0: DummyList
	(*CheckRequest)(nil), // 1: CheckRequest
	(*CheckSet)(nil),     // 2: CheckSet
	(*Emptyb)(nil),       // 3: Emptyb
}
var file_app_proto_depIdxs = []int32{
	1, // 0: DummyList.dummy:type_name -> CheckRequest
	1, // 1: AccommodationAviability.GetAccommodationCheck:input_type -> CheckRequest
	3, // 2: AccommodationAviability.GetAllforAccomendation:input_type -> Emptyb
	2, // 3: AccommodationAviability.SetAccommodationAviability:input_type -> CheckSet
	3, // 4: AccommodationAviability.GetAccommodationCheck:output_type -> Emptyb
	0, // 5: AccommodationAviability.GetAllforAccomendation:output_type -> DummyList
	3, // 6: AccommodationAviability.SetAccommodationAviability:output_type -> Emptyb
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_proto_init() }
func file_app_proto_init() {
	if File_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyList); i {
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
		file_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckSet); i {
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
		file_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emptyb); i {
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
			RawDescriptor: file_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_proto_goTypes,
		DependencyIndexes: file_app_proto_depIdxs,
		MessageInfos:      file_app_proto_msgTypes,
	}.Build()
	File_app_proto = out.File
	file_app_proto_rawDesc = nil
	file_app_proto_goTypes = nil
	file_app_proto_depIdxs = nil
}