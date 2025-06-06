// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: pagination/v1/pagination.proto

package v1

import (
	_ "github.com/google/gnostic/openapiv3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 分页通用请求
type PagingRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 当前页码
	Page *int32 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	// 每一页的行数
	PageSize *int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	// AND过滤参数，其语法为json格式的字符串，如：{"key1":"val1","key2":"val2"}，具体请参见：https://github.com/tx7do/go-utils/tree/main/entgo/query/README.md
	Query *string `protobuf:"bytes,3,opt,name=query,proto3,oneof" json:"query,omitempty"`
	// OR过滤参数，语法同AND过滤参数。
	OrQuery *string `protobuf:"bytes,4,opt,name=or_query,json=or,proto3,oneof" json:"or_query,omitempty"`
	// 排序条件，其语法为JSON字符串，例如：{"val1", "-val2"}。字段名前加'-'为降序，否则为升序。
	OrderBy []string `protobuf:"bytes,5,rep,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// 是否不分页，如果为true，则page和pageSize参数无效。
	NoPaging *bool `protobuf:"varint,6,opt,name=no_paging,json=noPaging,proto3,oneof" json:"no_paging,omitempty"`
	// 字段掩码，其作用为SELECT中的字段，其语法为使用逗号分隔字段名，例如：id,realName,userName。如果为空则选中所有字段，即SELECT *。
	FieldMask     *fieldmaskpb.FieldMask `protobuf:"bytes,7,opt,name=field_mask,json=fieldMask,proto3,oneof" json:"field_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PagingRequest) Reset() {
	*x = PagingRequest{}
	mi := &file_pagination_v1_pagination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PagingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingRequest) ProtoMessage() {}

func (x *PagingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_v1_pagination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingRequest.ProtoReflect.Descriptor instead.
func (*PagingRequest) Descriptor() ([]byte, []int) {
	return file_pagination_v1_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PagingRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *PagingRequest) GetPageSize() int32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *PagingRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *PagingRequest) GetOrQuery() string {
	if x != nil && x.OrQuery != nil {
		return *x.OrQuery
	}
	return ""
}

func (x *PagingRequest) GetOrderBy() []string {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

func (x *PagingRequest) GetNoPaging() bool {
	if x != nil && x.NoPaging != nil {
		return *x.NoPaging
	}
	return false
}

func (x *PagingRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

// 分页通用结果
type PagingResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 总数
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 分页数据
	Items         [][]byte `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PagingResponse) Reset() {
	*x = PagingResponse{}
	mi := &file_pagination_v1_pagination_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PagingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingResponse) ProtoMessage() {}

func (x *PagingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_v1_pagination_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingResponse.ProtoReflect.Descriptor instead.
func (*PagingResponse) Descriptor() ([]byte, []int) {
	return file_pagination_v1_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PagingResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PagingResponse) GetItems() [][]byte {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_pagination_v1_pagination_proto protoreflect.FileDescriptor

var file_pagination_v1_pagination_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x08, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x1e, 0xba, 0x47, 0x1b, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0xf0, 0x3f, 0x92, 0x02, 0x0c, 0xe5, 0xbd, 0x93, 0xe5, 0x89, 0x8d, 0xe9, 0xa1,
	0xb5, 0xe7, 0xa0, 0x81, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x46, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x24, 0xba, 0x47, 0x21, 0x8a, 0x02, 0x09, 0x09, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x24, 0x40, 0x92, 0x02, 0x12, 0xe6, 0xaf, 0x8f, 0xe4, 0xb8, 0x80, 0xe9, 0xa1, 0xb5, 0xe7,
	0x9a, 0x84, 0xe8, 0xa1, 0x8c, 0xe6, 0x95, 0xb0, 0x48, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0xf5, 0x01, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0xd9, 0x01, 0xba, 0x47, 0xd5, 0x01, 0x3a, 0x1f,
	0x12, 0x1d, 0x7b, 0x22, 0x6b, 0x65, 0x79, 0x31, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x31, 0x22,
	0x2c, 0x22, 0x6b, 0x65, 0x79, 0x32, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x32, 0x22, 0x7d, 0x92,
	0x02, 0xb0, 0x01, 0x41, 0x4e, 0x44, 0xe8, 0xbf, 0x87, 0xe6, 0xbb, 0xa4, 0xe5, 0x8f, 0x82, 0xe6,
	0x95, 0xb0, 0xef, 0xbc, 0x8c, 0xe5, 0x85, 0xb6, 0xe8, 0xaf, 0xad, 0xe6, 0xb3, 0x95, 0xe4, 0xb8,
	0xba, 0x6a, 0x73, 0x6f, 0x6e, 0xe6, 0xa0, 0xbc, 0xe5, 0xbc, 0x8f, 0xe7, 0x9a, 0x84, 0xe5, 0xad,
	0x97, 0xe7, 0xac, 0xa6, 0xe4, 0xb8, 0xb2, 0xef, 0xbc, 0x8c, 0xe5, 0xa6, 0x82, 0xef, 0xbc, 0x9a,
	0x7b, 0x22, 0x6b, 0x65, 0x79, 0x31, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x31, 0x22, 0x2c, 0x22,
	0x6b, 0x65, 0x79, 0x32, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x32, 0x22, 0x7d, 0xef, 0xbc, 0x8c,
	0xe5, 0x85, 0xb7, 0xe4, 0xbd, 0x93, 0xe8, 0xaf, 0xb7, 0xe5, 0x8f, 0x82, 0xe8, 0xa7, 0x81, 0xef,
	0xbc, 0x9a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x78, 0x37, 0x64, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x75, 0x74,
	0x69, 0x6c, 0x73, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x6e,
	0x74, 0x67, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x52, 0x45, 0x41, 0x44, 0x4d, 0x45,
	0x2e, 0x6d, 0x64, 0x48, 0x02, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x50, 0x0a, 0x08, 0x6f, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x35, 0xba, 0x47, 0x32, 0x3a, 0x1f, 0x12, 0x1d, 0x7b, 0x22, 0x6b, 0x65, 0x79, 0x31,
	0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x31, 0x22, 0x2c, 0x22, 0x6b, 0x65, 0x79, 0x32, 0x22, 0x3a,
	0x22, 0x76, 0x61, 0x6c, 0x32, 0x22, 0x7d, 0x92, 0x02, 0x0e, 0x4f, 0x52, 0xe8, 0xbf, 0x87, 0xe6,
	0xbb, 0xa4, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0x48, 0x03, 0x52, 0x02, 0x6f, 0x72, 0x88, 0x01,
	0x01, 0x12, 0xb0, 0x01, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x94, 0x01, 0xba, 0x47, 0x90, 0x01, 0x3a, 0x13, 0x12, 0x11, 0x7b,
	0x22, 0x76, 0x61, 0x6c, 0x31, 0x22, 0x2c, 0x20, 0x22, 0x2d, 0x76, 0x61, 0x6c, 0x32, 0x22, 0x7d,
	0x92, 0x02, 0x78, 0xe6, 0x8e, 0x92, 0xe5, 0xba, 0x8f, 0xe6, 0x9d, 0xa1, 0xe4, 0xbb, 0xb6, 0xef,
	0xbc, 0x8c, 0xe5, 0x85, 0xb6, 0xe8, 0xaf, 0xad, 0xe6, 0xb3, 0x95, 0xe4, 0xb8, 0xba, 0x4a, 0x53,
	0x4f, 0x4e, 0xe5, 0xad, 0x97, 0xe7, 0xac, 0xa6, 0xe4, 0xb8, 0xb2, 0xef, 0xbc, 0x8c, 0xe4, 0xbe,
	0x8b, 0xe5, 0xa6, 0x82, 0xef, 0xbc, 0x9a, 0x7b, 0x22, 0x76, 0x61, 0x6c, 0x31, 0x22, 0x2c, 0x20,
	0x22, 0x2d, 0x76, 0x61, 0x6c, 0x32, 0x22, 0x7d, 0xe3, 0x80, 0x82, 0xe5, 0xad, 0x97, 0xe6, 0xae,
	0xb5, 0xe5, 0x90, 0x8d, 0xe5, 0x89, 0x8d, 0xe5, 0x8a, 0xa0, 0x27, 0x2d, 0x27, 0xe4, 0xb8, 0xba,
	0xe9, 0x99, 0x8d, 0xe5, 0xba, 0x8f, 0xef, 0xbc, 0x8c, 0xe5, 0x90, 0xa6, 0xe5, 0x88, 0x99, 0xe4,
	0xb8, 0xba, 0xe5, 0x8d, 0x87, 0xe5, 0xba, 0x8f, 0xe3, 0x80, 0x82, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x12, 0x6b, 0x0a, 0x09, 0x6e, 0x6f, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x49, 0xba, 0x47, 0x46, 0x92, 0x02, 0x43, 0xe6,
	0x98, 0xaf, 0xe5, 0x90, 0xa6, 0xe4, 0xb8, 0x8d, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xef, 0xbc,
	0x8c, 0xe5, 0xa6, 0x82, 0xe6, 0x9e, 0x9c, 0xe4, 0xb8, 0xba, 0x74, 0x72, 0x75, 0x65, 0xef, 0xbc,
	0x8c, 0xe5, 0x88, 0x99, 0x70, 0x61, 0x67, 0x65, 0xe5, 0x92, 0x8c, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0, 0xe6, 0x97, 0xa0, 0xe6, 0x95, 0x88, 0xe3,
	0x80, 0x82, 0x48, 0x04, 0x52, 0x08, 0x6e, 0x6f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x88, 0x01,
	0x01, 0x12, 0x8d, 0x02, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x42, 0xcc, 0x01, 0xba, 0x47, 0xc8, 0x01, 0x3a, 0x16, 0x12, 0x14, 0x69, 0x64, 0x2c,
	0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x2c, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x92, 0x02, 0xac, 0x01, 0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5, 0xe6, 0x8e, 0xa9, 0xe7, 0xa0,
	0x81, 0xef, 0xbc, 0x8c, 0xe5, 0x85, 0xb6, 0xe4, 0xbd, 0x9c, 0xe7, 0x94, 0xa8, 0xe4, 0xb8, 0xba,
	0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0xe4, 0xb8, 0xad, 0xe7, 0x9a, 0x84, 0xe5, 0xad, 0x97, 0xe6,
	0xae, 0xb5, 0xef, 0xbc, 0x8c, 0xe5, 0x85, 0xb6, 0xe8, 0xaf, 0xad, 0xe6, 0xb3, 0x95, 0xe4, 0xb8,
	0xba, 0xe4, 0xbd, 0xbf, 0xe7, 0x94, 0xa8, 0xe9, 0x80, 0x97, 0xe5, 0x8f, 0xb7, 0xe5, 0x88, 0x86,
	0xe9, 0x9a, 0x94, 0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5, 0xe5, 0x90, 0x8d, 0xef, 0xbc, 0x8c, 0xe4,
	0xbe, 0x8b, 0xe5, 0xa6, 0x82, 0xef, 0xbc, 0x9a, 0x69, 0x64, 0x2c, 0x72, 0x65, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x2c, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0xe3, 0x80, 0x82, 0xe5,
	0xa6, 0x82, 0xe6, 0x9e, 0x9c, 0xe4, 0xb8, 0xba, 0xe7, 0xa9, 0xba, 0xe5, 0x88, 0x99, 0xe9, 0x80,
	0x89, 0xe4, 0xb8, 0xad, 0xe6, 0x89, 0x80, 0xe6, 0x9c, 0x89, 0xe5, 0xad, 0x97, 0xe6, 0xae, 0xb5,
	0xef, 0xbc, 0x8c, 0xe5, 0x8d, 0xb3, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x20, 0x2a, 0xe3, 0x80,
	0x82, 0x48, 0x05, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x6f, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x3c, 0x0a, 0x0e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0xa5, 0x01, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x78, 0x37,
	0x64, 0x6f, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50,
	0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xca,
	0x02, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x16, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_pagination_v1_pagination_proto_rawDescOnce sync.Once
	file_pagination_v1_pagination_proto_rawDescData []byte
)

func file_pagination_v1_pagination_proto_rawDescGZIP() []byte {
	file_pagination_v1_pagination_proto_rawDescOnce.Do(func() {
		file_pagination_v1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pagination_v1_pagination_proto_rawDesc), len(file_pagination_v1_pagination_proto_rawDesc)))
	})
	return file_pagination_v1_pagination_proto_rawDescData
}

var file_pagination_v1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pagination_v1_pagination_proto_goTypes = []any{
	(*PagingRequest)(nil),         // 0: pagination.PagingRequest
	(*PagingResponse)(nil),        // 1: pagination.PagingResponse
	(*fieldmaskpb.FieldMask)(nil), // 2: google.protobuf.FieldMask
}
var file_pagination_v1_pagination_proto_depIdxs = []int32{
	2, // 0: pagination.PagingRequest.field_mask:type_name -> google.protobuf.FieldMask
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pagination_v1_pagination_proto_init() }
func file_pagination_v1_pagination_proto_init() {
	if File_pagination_v1_pagination_proto != nil {
		return
	}
	file_pagination_v1_pagination_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pagination_v1_pagination_proto_rawDesc), len(file_pagination_v1_pagination_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pagination_v1_pagination_proto_goTypes,
		DependencyIndexes: file_pagination_v1_pagination_proto_depIdxs,
		MessageInfos:      file_pagination_v1_pagination_proto_msgTypes,
	}.Build()
	File_pagination_v1_pagination_proto = out.File
	file_pagination_v1_pagination_proto_goTypes = nil
	file_pagination_v1_pagination_proto_depIdxs = nil
}
