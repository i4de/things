// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/rule.proto

package rule

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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{0}
}

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{1}
}

func (x *PageInfo) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SceneInfoReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //场景id
}

func (x *SceneInfoReadReq) Reset() {
	*x = SceneInfoReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneInfoReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneInfoReadReq) ProtoMessage() {}

func (x *SceneInfoReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneInfoReadReq.ProtoReflect.Descriptor instead.
func (*SceneInfoReadReq) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{2}
}

func (x *SceneInfoReadReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SceneInfoIndexReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page        *PageInfo `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`               //分页信息 只获取一个则不填
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               //场景名 模糊查询
	State       int64     `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`            //状态: 1启用 2禁用
	TriggerType string    `protobuf:"bytes,4,opt,name=triggerType,proto3" json:"triggerType,omitempty"` //触发类型 device: 设备触发 timer: 定时触发 manual:手动触发
}

func (x *SceneInfoIndexReq) Reset() {
	*x = SceneInfoIndexReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneInfoIndexReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneInfoIndexReq) ProtoMessage() {}

func (x *SceneInfoIndexReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneInfoIndexReq.ProtoReflect.Descriptor instead.
func (*SceneInfoIndexReq) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{3}
}

func (x *SceneInfoIndexReq) GetPage() *PageInfo {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *SceneInfoIndexReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SceneInfoIndexReq) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *SceneInfoIndexReq) GetTriggerType() string {
	if x != nil {
		return x.TriggerType
	}
	return ""
}

type SceneInfoIndexResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*SceneInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`    //设备信息
	Total int64        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` //总数(只有分页的时候会返回)
}

func (x *SceneInfoIndexResp) Reset() {
	*x = SceneInfoIndexResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneInfoIndexResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneInfoIndexResp) ProtoMessage() {}

func (x *SceneInfoIndexResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneInfoIndexResp.ProtoReflect.Descriptor instead.
func (*SceneInfoIndexResp) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{4}
}

func (x *SceneInfoIndexResp) GetList() []*SceneInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SceneInfoIndexResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SceneInfoDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SceneInfoDeleteReq) Reset() {
	*x = SceneInfoDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneInfoDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneInfoDeleteReq) ProtoMessage() {}

func (x *SceneInfoDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneInfoDeleteReq.ProtoReflect.Descriptor instead.
func (*SceneInfoDeleteReq) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{5}
}

func (x *SceneInfoDeleteReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SceneInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                   //场景id
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                //场景名
	State       int64  `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`             //状态: 1启用 2禁用
	Desc        string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`                //描述
	CreatedTime int64  `protobuf:"varint,5,opt,name=createdTime,proto3" json:"createdTime,omitempty"` //创建时间 秒级时间戳 只读
	Trigger     string `protobuf:"bytes,6,opt,name=trigger,proto3" json:"trigger,omitempty"`          //触发器
	When        string `protobuf:"bytes,7,opt,name=when,proto3" json:"when,omitempty"`                //触发条件
	Then        string `protobuf:"bytes,8,opt,name=then,proto3" json:"then,omitempty"`                //满足条件时执行的动作
}

func (x *SceneInfo) Reset() {
	*x = SceneInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneInfo) ProtoMessage() {}

func (x *SceneInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneInfo.ProtoReflect.Descriptor instead.
func (*SceneInfo) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{6}
}

func (x *SceneInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SceneInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SceneInfo) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *SceneInfo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *SceneInfo) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *SceneInfo) GetTrigger() string {
	if x != nil {
		return x.Trigger
	}
	return ""
}

func (x *SceneInfo) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

func (x *SceneInfo) GetThen() string {
	if x != nil {
		return x.Then
	}
	return ""
}

type FlowInfoReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FlowInfoReadReq) Reset() {
	*x = FlowInfoReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowInfoReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowInfoReadReq) ProtoMessage() {}

func (x *FlowInfoReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowInfoReadReq.ProtoReflect.Descriptor instead.
func (*FlowInfoReadReq) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{7}
}

func (x *FlowInfoReadReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FlowInfoIndexReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *PageInfo `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"` //分页信息 只获取一个则不填
}

func (x *FlowInfoIndexReq) Reset() {
	*x = FlowInfoIndexReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowInfoIndexReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowInfoIndexReq) ProtoMessage() {}

func (x *FlowInfoIndexReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowInfoIndexReq.ProtoReflect.Descriptor instead.
func (*FlowInfoIndexReq) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{8}
}

func (x *FlowInfoIndexReq) GetPage() *PageInfo {
	if x != nil {
		return x.Page
	}
	return nil
}

type FlowInfoIndexResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*FlowInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`    //设备信息
	Total int64       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` //总数(只有分页的时候会返回)
}

func (x *FlowInfoIndexResp) Reset() {
	*x = FlowInfoIndexResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowInfoIndexResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowInfoIndexResp) ProtoMessage() {}

func (x *FlowInfoIndexResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowInfoIndexResp.ProtoReflect.Descriptor instead.
func (*FlowInfoIndexResp) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{9}
}

func (x *FlowInfoIndexResp) GetList() []*FlowInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *FlowInfoIndexResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type FlowInfoDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FlowInfoDeleteReq) Reset() {
	*x = FlowInfoDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowInfoDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowInfoDeleteReq) ProtoMessage() {}

func (x *FlowInfoDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowInfoDeleteReq.ProtoReflect.Descriptor instead.
func (*FlowInfoDeleteReq) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{10}
}

func (x *FlowInfoDeleteReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FlowInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                //流的名称
	IsDisabled  int64  `protobuf:"varint,3,opt,name=isDisabled,proto3" json:"isDisabled,omitempty"`   //是否禁用 1:是 2:否
	Desc        string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`                //描述
	CreatedTime int64  `protobuf:"varint,5,opt,name=createdTime,proto3" json:"createdTime,omitempty"` //创建时间 秒级时间戳 只读
}

func (x *FlowInfo) Reset() {
	*x = FlowInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowInfo) ProtoMessage() {}

func (x *FlowInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowInfo.ProtoReflect.Descriptor instead.
func (*FlowInfo) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{11}
}

func (x *FlowInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FlowInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlowInfo) GetIsDisabled() int64 {
	if x != nil {
		return x.IsDisabled
	}
	return 0
}

func (x *FlowInfo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *FlowInfo) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

var File_proto_rule_proto protoreflect.FileDescriptor

var file_proto_rule_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x83, 0x01, 0x0a,
	0x11, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x71, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x4f, 0x0a, 0x12, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x09, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x77, 0x68, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x68, 0x65, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x68, 0x65, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x46, 0x6c, 0x6f,
	0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x10,
	0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71,
	0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x11, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x46,
	0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x23, 0x0a, 0x11, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x08, 0x46, 0x6c, 0x6f,
	0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69,
	0x73, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32,
	0xa4, 0x02, 0x0a, 0x0a, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x30,
	0x0a, 0x0e, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x0e, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77,
	0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e,
	0x72, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x0d, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x46, 0x6c,
	0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x35, 0x0a, 0x0c, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x15, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x46, 0x6c,
	0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb2, 0x02, 0x0a, 0x0c, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x0f, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x72, 0x75,
	0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0f, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x0e, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x0f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0e,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x38, 0x0a, 0x0d, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x16, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rule_proto_rawDescOnce sync.Once
	file_proto_rule_proto_rawDescData = file_proto_rule_proto_rawDesc
)

func file_proto_rule_proto_rawDescGZIP() []byte {
	file_proto_rule_proto_rawDescOnce.Do(func() {
		file_proto_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rule_proto_rawDescData)
	})
	return file_proto_rule_proto_rawDescData
}

var file_proto_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_rule_proto_goTypes = []interface{}{
	(*Response)(nil),           // 0: rule.Response
	(*PageInfo)(nil),           // 1: rule.PageInfo
	(*SceneInfoReadReq)(nil),   // 2: rule.SceneInfoReadReq
	(*SceneInfoIndexReq)(nil),  // 3: rule.SceneInfoIndexReq
	(*SceneInfoIndexResp)(nil), // 4: rule.SceneInfoIndexResp
	(*SceneInfoDeleteReq)(nil), // 5: rule.SceneInfoDeleteReq
	(*SceneInfo)(nil),          // 6: rule.SceneInfo
	(*FlowInfoReadReq)(nil),    // 7: rule.FlowInfoReadReq
	(*FlowInfoIndexReq)(nil),   // 8: rule.FlowInfoIndexReq
	(*FlowInfoIndexResp)(nil),  // 9: rule.FlowInfoIndexResp
	(*FlowInfoDeleteReq)(nil),  // 10: rule.FlowInfoDeleteReq
	(*FlowInfo)(nil),           // 11: rule.FlowInfo
}
var file_proto_rule_proto_depIdxs = []int32{
	1,  // 0: rule.SceneInfoIndexReq.page:type_name -> rule.PageInfo
	6,  // 1: rule.SceneInfoIndexResp.list:type_name -> rule.SceneInfo
	1,  // 2: rule.FlowInfoIndexReq.page:type_name -> rule.PageInfo
	11, // 3: rule.FlowInfoIndexResp.list:type_name -> rule.FlowInfo
	11, // 4: rule.RuleEngine.flowInfoCreate:input_type -> rule.FlowInfo
	11, // 5: rule.RuleEngine.flowInfoUpdate:input_type -> rule.FlowInfo
	10, // 6: rule.RuleEngine.flowInfoDelete:input_type -> rule.FlowInfoDeleteReq
	8,  // 7: rule.RuleEngine.flowInfoIndex:input_type -> rule.FlowInfoIndexReq
	7,  // 8: rule.RuleEngine.flowInfoRead:input_type -> rule.FlowInfoReadReq
	6,  // 9: rule.SceneLinkage.sceneInfoCreate:input_type -> rule.SceneInfo
	6,  // 10: rule.SceneLinkage.sceneInfoUpdate:input_type -> rule.SceneInfo
	5,  // 11: rule.SceneLinkage.sceneInfoDelete:input_type -> rule.SceneInfoDeleteReq
	3,  // 12: rule.SceneLinkage.sceneInfoIndex:input_type -> rule.SceneInfoIndexReq
	2,  // 13: rule.SceneLinkage.sceneInfoRead:input_type -> rule.SceneInfoReadReq
	0,  // 14: rule.RuleEngine.flowInfoCreate:output_type -> rule.Response
	0,  // 15: rule.RuleEngine.flowInfoUpdate:output_type -> rule.Response
	0,  // 16: rule.RuleEngine.flowInfoDelete:output_type -> rule.Response
	9,  // 17: rule.RuleEngine.flowInfoIndex:output_type -> rule.FlowInfoIndexResp
	11, // 18: rule.RuleEngine.flowInfoRead:output_type -> rule.FlowInfo
	0,  // 19: rule.SceneLinkage.sceneInfoCreate:output_type -> rule.Response
	0,  // 20: rule.SceneLinkage.sceneInfoUpdate:output_type -> rule.Response
	0,  // 21: rule.SceneLinkage.sceneInfoDelete:output_type -> rule.Response
	4,  // 22: rule.SceneLinkage.sceneInfoIndex:output_type -> rule.SceneInfoIndexResp
	6,  // 23: rule.SceneLinkage.sceneInfoRead:output_type -> rule.SceneInfo
	14, // [14:24] is the sub-list for method output_type
	4,  // [4:14] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_rule_proto_init() }
func file_proto_rule_proto_init() {
	if File_proto_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_proto_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneInfoReadReq); i {
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
		file_proto_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneInfoIndexReq); i {
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
		file_proto_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneInfoIndexResp); i {
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
		file_proto_rule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneInfoDeleteReq); i {
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
		file_proto_rule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneInfo); i {
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
		file_proto_rule_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowInfoReadReq); i {
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
		file_proto_rule_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowInfoIndexReq); i {
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
		file_proto_rule_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowInfoIndexResp); i {
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
		file_proto_rule_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowInfoDeleteReq); i {
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
		file_proto_rule_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowInfo); i {
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
			RawDescriptor: file_proto_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_rule_proto_goTypes,
		DependencyIndexes: file_proto_rule_proto_depIdxs,
		MessageInfos:      file_proto_rule_proto_msgTypes,
	}.Build()
	File_proto_rule_proto = out.File
	file_proto_rule_proto_rawDesc = nil
	file_proto_rule_proto_goTypes = nil
	file_proto_rule_proto_depIdxs = nil
}
