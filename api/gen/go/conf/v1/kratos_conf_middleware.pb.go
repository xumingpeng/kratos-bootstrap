// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: conf/v1/kratos_conf_middleware.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 中间件
type Middleware struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableLogging        bool                    `protobuf:"varint,1,opt,name=enable_logging,json=enableLogging,proto3" json:"enable_logging,omitempty"`                        // 日志开关
	EnableRecovery       bool                    `protobuf:"varint,2,opt,name=enable_recovery,json=enableRecovery,proto3" json:"enable_recovery,omitempty"`                     // 异常恢复
	EnableTracing        bool                    `protobuf:"varint,3,opt,name=enable_tracing,json=enableTracing,proto3" json:"enable_tracing,omitempty"`                        // 链路追踪开关
	EnableValidate       bool                    `protobuf:"varint,4,opt,name=enable_validate,json=enableValidate,proto3" json:"enable_validate,omitempty"`                     // 参数校验开关
	EnableCircuitBreaker bool                    `protobuf:"varint,5,opt,name=enable_circuit_breaker,json=enableCircuitBreaker,proto3" json:"enable_circuit_breaker,omitempty"` // 熔断器
	Limiter              *Middleware_RateLimiter `protobuf:"bytes,6,opt,name=limiter,proto3" json:"limiter,omitempty"`
	Metrics              *Middleware_Metrics     `protobuf:"bytes,7,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Auth                 *Middleware_Auth        `protobuf:"bytes,8,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *Middleware) Reset() {
	*x = Middleware{}
	mi := &file_conf_v1_kratos_conf_middleware_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware) ProtoMessage() {}

func (x *Middleware) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_middleware_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware.ProtoReflect.Descriptor instead.
func (*Middleware) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_middleware_proto_rawDescGZIP(), []int{0}
}

func (x *Middleware) GetEnableLogging() bool {
	if x != nil {
		return x.EnableLogging
	}
	return false
}

func (x *Middleware) GetEnableRecovery() bool {
	if x != nil {
		return x.EnableRecovery
	}
	return false
}

func (x *Middleware) GetEnableTracing() bool {
	if x != nil {
		return x.EnableTracing
	}
	return false
}

func (x *Middleware) GetEnableValidate() bool {
	if x != nil {
		return x.EnableValidate
	}
	return false
}

func (x *Middleware) GetEnableCircuitBreaker() bool {
	if x != nil {
		return x.EnableCircuitBreaker
	}
	return false
}

func (x *Middleware) GetLimiter() *Middleware_RateLimiter {
	if x != nil {
		return x.Limiter
	}
	return nil
}

func (x *Middleware) GetMetrics() *Middleware_Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *Middleware) GetAuth() *Middleware_Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// JWT校验
type Middleware_Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"` // JWT签名的算法，支持算法：HS256
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`       // JWT 秘钥
}

func (x *Middleware_Auth) Reset() {
	*x = Middleware_Auth{}
	mi := &file_conf_v1_kratos_conf_middleware_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_Auth) ProtoMessage() {}

func (x *Middleware_Auth) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_middleware_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_Auth.ProtoReflect.Descriptor instead.
func (*Middleware_Auth) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_middleware_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Middleware_Auth) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Middleware_Auth) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// 限流器
type Middleware_RateLimiter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 限流器名字，支持：bbr。
}

func (x *Middleware_RateLimiter) Reset() {
	*x = Middleware_RateLimiter{}
	mi := &file_conf_v1_kratos_conf_middleware_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_RateLimiter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_RateLimiter) ProtoMessage() {}

func (x *Middleware_RateLimiter) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_middleware_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_RateLimiter.ProtoReflect.Descriptor instead.
func (*Middleware_RateLimiter) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_middleware_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Middleware_RateLimiter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 性能指标
type Middleware_Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Histogram bool `protobuf:"varint,1,opt,name=histogram,proto3" json:"histogram,omitempty"` // 直方图
	Counter   bool `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`     // 计数器
	Gauge     bool `protobuf:"varint,3,opt,name=gauge,proto3" json:"gauge,omitempty"`         // 仪表盘
	Summary   bool `protobuf:"varint,4,opt,name=summary,proto3" json:"summary,omitempty"`     // 摘要
}

func (x *Middleware_Metrics) Reset() {
	*x = Middleware_Metrics{}
	mi := &file_conf_v1_kratos_conf_middleware_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Middleware_Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_Metrics) ProtoMessage() {}

func (x *Middleware_Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_middleware_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_Metrics.ProtoReflect.Descriptor instead.
func (*Middleware_Metrics) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_middleware_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Middleware_Metrics) GetHistogram() bool {
	if x != nil {
		return x.Histogram
	}
	return false
}

func (x *Middleware_Metrics) GetCounter() bool {
	if x != nil {
		return x.Counter
	}
	return false
}

func (x *Middleware_Metrics) GetGauge() bool {
	if x != nil {
		return x.Gauge
	}
	return false
}

func (x *Middleware_Metrics) GetSummary() bool {
	if x != nil {
		return x.Summary
	}
	return false
}

var File_conf_v1_kratos_conf_middleware_proto protoreflect.FileDescriptor

var file_conf_v1_kratos_conf_middleware_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x04, 0x0a,
	0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x5f, 0x62, 0x72,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65,
	0x72, 0x12, 0x36, 0x0a, 0x07, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x29, 0x0a,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x30, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x1a, 0x21, 0x0a, 0x0b, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x71, 0x0a,
	0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x42, 0x8b, 0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x42, 0x19, 0x4b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x78, 0x37, 0x64, 0x6f, 0x2f, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2d, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0xca, 0x02, 0x04,
	0x43, 0x6f, 0x6e, 0x66, 0xe2, 0x02, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_v1_kratos_conf_middleware_proto_rawDescOnce sync.Once
	file_conf_v1_kratos_conf_middleware_proto_rawDescData = file_conf_v1_kratos_conf_middleware_proto_rawDesc
)

func file_conf_v1_kratos_conf_middleware_proto_rawDescGZIP() []byte {
	file_conf_v1_kratos_conf_middleware_proto_rawDescOnce.Do(func() {
		file_conf_v1_kratos_conf_middleware_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_v1_kratos_conf_middleware_proto_rawDescData)
	})
	return file_conf_v1_kratos_conf_middleware_proto_rawDescData
}

var file_conf_v1_kratos_conf_middleware_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conf_v1_kratos_conf_middleware_proto_goTypes = []any{
	(*Middleware)(nil),             // 0: conf.Middleware
	(*Middleware_Auth)(nil),        // 1: conf.Middleware.Auth
	(*Middleware_RateLimiter)(nil), // 2: conf.Middleware.RateLimiter
	(*Middleware_Metrics)(nil),     // 3: conf.Middleware.Metrics
}
var file_conf_v1_kratos_conf_middleware_proto_depIdxs = []int32{
	2, // 0: conf.Middleware.limiter:type_name -> conf.Middleware.RateLimiter
	3, // 1: conf.Middleware.metrics:type_name -> conf.Middleware.Metrics
	1, // 2: conf.Middleware.auth:type_name -> conf.Middleware.Auth
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_conf_v1_kratos_conf_middleware_proto_init() }
func file_conf_v1_kratos_conf_middleware_proto_init() {
	if File_conf_v1_kratos_conf_middleware_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_v1_kratos_conf_middleware_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_v1_kratos_conf_middleware_proto_goTypes,
		DependencyIndexes: file_conf_v1_kratos_conf_middleware_proto_depIdxs,
		MessageInfos:      file_conf_v1_kratos_conf_middleware_proto_msgTypes,
	}.Build()
	File_conf_v1_kratos_conf_middleware_proto = out.File
	file_conf_v1_kratos_conf_middleware_proto_rawDesc = nil
	file_conf_v1_kratos_conf_middleware_proto_goTypes = nil
	file_conf_v1_kratos_conf_middleware_proto_depIdxs = nil
}
