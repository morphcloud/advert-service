// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: advert_service.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AdvertReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Advert *Advert `protobuf:"bytes,1,opt,name=advert,proto3" json:"advert,omitempty"`
	Error  *Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AdvertReply) Reset() {
	*x = AdvertReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advert_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvertReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvertReply) ProtoMessage() {}

func (x *AdvertReply) ProtoReflect() protoreflect.Message {
	mi := &file_advert_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvertReply.ProtoReflect.Descriptor instead.
func (*AdvertReply) Descriptor() ([]byte, []int) {
	return file_advert_service_proto_rawDescGZIP(), []int{0}
}

func (x *AdvertReply) GetAdvert() *Advert {
	if x != nil {
		return x.Advert
	}
	return nil
}

func (x *AdvertReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advert_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_advert_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_advert_service_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_advert_service_proto protoreflect.FileDescriptor

var file_advert_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x1a, 0x14, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x0b, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x52,
	0x06, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xe2,
	0x01, 0x0a, 0x0d, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x45, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x72,
	0x70, 0x68, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x19, 0x2e, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x72,
	0x70, 0x68, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_advert_service_proto_rawDescOnce sync.Once
	file_advert_service_proto_rawDescData = file_advert_service_proto_rawDesc
)

func file_advert_service_proto_rawDescGZIP() []byte {
	file_advert_service_proto_rawDescOnce.Do(func() {
		file_advert_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_advert_service_proto_rawDescData)
	})
	return file_advert_service_proto_rawDescData
}

var file_advert_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_advert_service_proto_goTypes = []interface{}{
	(*AdvertReply)(nil), // 0: morphcloud.advert.AdvertReply
	(*Error)(nil),       // 1: morphcloud.advert.Error
	(*Advert)(nil),      // 2: morphcloud.advert.Advert
	(*Filter)(nil),      // 3: morphcloud.advert.Filter
}
var file_advert_service_proto_depIdxs = []int32{
	2, // 0: morphcloud.advert.AdvertReply.advert:type_name -> morphcloud.advert.Advert
	1, // 1: morphcloud.advert.AdvertReply.error:type_name -> morphcloud.advert.Error
	2, // 2: morphcloud.advert.AdvertService.Create:input_type -> morphcloud.advert.Advert
	3, // 3: morphcloud.advert.AdvertService.Read:input_type -> morphcloud.advert.Filter
	2, // 4: morphcloud.advert.AdvertService.Update:input_type -> morphcloud.advert.Advert
	0, // 5: morphcloud.advert.AdvertService.Create:output_type -> morphcloud.advert.AdvertReply
	0, // 6: morphcloud.advert.AdvertService.Read:output_type -> morphcloud.advert.AdvertReply
	0, // 7: morphcloud.advert.AdvertService.Update:output_type -> morphcloud.advert.AdvertReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_advert_service_proto_init() }
func file_advert_service_proto_init() {
	if File_advert_service_proto != nil {
		return
	}
	file_advert_message_proto_init()
	file_filter_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_advert_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvertReply); i {
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
		file_advert_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_advert_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_advert_service_proto_goTypes,
		DependencyIndexes: file_advert_service_proto_depIdxs,
		MessageInfos:      file_advert_service_proto_msgTypes,
	}.Build()
	File_advert_service_proto = out.File
	file_advert_service_proto_rawDesc = nil
	file_advert_service_proto_goTypes = nil
	file_advert_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdvertServiceClient is the client API for AdvertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdvertServiceClient interface {
	Create(ctx context.Context, in *Advert, opts ...grpc.CallOption) (*AdvertReply, error)
	Read(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*AdvertReply, error)
	Update(ctx context.Context, in *Advert, opts ...grpc.CallOption) (*AdvertReply, error)
}

type advertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvertServiceClient(cc grpc.ClientConnInterface) AdvertServiceClient {
	return &advertServiceClient{cc}
}

func (c *advertServiceClient) Create(ctx context.Context, in *Advert, opts ...grpc.CallOption) (*AdvertReply, error) {
	out := new(AdvertReply)
	err := c.cc.Invoke(ctx, "/morphcloud.advert.AdvertService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertServiceClient) Read(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*AdvertReply, error) {
	out := new(AdvertReply)
	err := c.cc.Invoke(ctx, "/morphcloud.advert.AdvertService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertServiceClient) Update(ctx context.Context, in *Advert, opts ...grpc.CallOption) (*AdvertReply, error) {
	out := new(AdvertReply)
	err := c.cc.Invoke(ctx, "/morphcloud.advert.AdvertService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertServiceServer is the server API for AdvertService service.
type AdvertServiceServer interface {
	Create(context.Context, *Advert) (*AdvertReply, error)
	Read(context.Context, *Filter) (*AdvertReply, error)
	Update(context.Context, *Advert) (*AdvertReply, error)
}

// UnimplementedAdvertServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdvertServiceServer struct {
}

func (*UnimplementedAdvertServiceServer) Create(context.Context, *Advert) (*AdvertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAdvertServiceServer) Read(context.Context, *Filter) (*AdvertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedAdvertServiceServer) Update(context.Context, *Advert) (*AdvertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterAdvertServiceServer(s *grpc.Server, srv AdvertServiceServer) {
	s.RegisterService(&_AdvertService_serviceDesc, srv)
}

func _AdvertService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Advert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/morphcloud.advert.AdvertService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertServiceServer).Create(ctx, req.(*Advert))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/morphcloud.advert.AdvertService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertServiceServer).Read(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Advert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/morphcloud.advert.AdvertService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertServiceServer).Update(ctx, req.(*Advert))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdvertService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "morphcloud.advert.AdvertService",
	HandlerType: (*AdvertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AdvertService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _AdvertService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdvertService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advert_service.proto",
}