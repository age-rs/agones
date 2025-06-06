// Copyright 2024 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.
// Copyright 2020 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v6.30.0
// source: beta.proto

package beta

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// I am Empty
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{0}
}

// A representation of a Counter.
type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Counter
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The current count of the Counter
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// The maximum capacity of the Counter
	Capacity int64 `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{1}
}

func (x *Counter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Counter) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Counter) GetCapacity() int64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

// A representation of a Counter Update Request.
type CounterUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Counter to update
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The value to set the Counter Count
	Count *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	// The value to set the Counter Capacity
	Capacity *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	// countDiff tracks if a Counter Update Request is CountIncrement (positive), CountDecrement
	// (negative), 0 if a CountSet or CapacitySet request
	CountDiff int64 `protobuf:"varint,4,opt,name=countDiff,proto3" json:"countDiff,omitempty"`
}

func (x *CounterUpdateRequest) Reset() {
	*x = CounterUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterUpdateRequest) ProtoMessage() {}

func (x *CounterUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterUpdateRequest.ProtoReflect.Descriptor instead.
func (*CounterUpdateRequest) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{2}
}

func (x *CounterUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CounterUpdateRequest) GetCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.Count
	}
	return nil
}

func (x *CounterUpdateRequest) GetCapacity() *wrapperspb.Int64Value {
	if x != nil {
		return x.Capacity
	}
	return nil
}

func (x *CounterUpdateRequest) GetCountDiff() int64 {
	if x != nil {
		return x.CountDiff
	}
	return 0
}

type GetCounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Counter to get
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCounterRequest) Reset() {
	*x = GetCounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCounterRequest) ProtoMessage() {}

func (x *GetCounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCounterRequest.ProtoReflect.Descriptor instead.
func (*GetCounterRequest) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{3}
}

func (x *GetCounterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateCounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The requested update to make to the Counter
	CounterUpdateRequest *CounterUpdateRequest `protobuf:"bytes,1,opt,name=counterUpdateRequest,proto3" json:"counterUpdateRequest,omitempty"`
}

func (x *UpdateCounterRequest) Reset() {
	*x = UpdateCounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCounterRequest) ProtoMessage() {}

func (x *UpdateCounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCounterRequest.ProtoReflect.Descriptor instead.
func (*UpdateCounterRequest) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCounterRequest) GetCounterUpdateRequest() *CounterUpdateRequest {
	if x != nil {
		return x.CounterUpdateRequest
	}
	return nil
}

// A representation of a List.
type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the List
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The maximum capacity of the List
	Capacity int64 `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	// The array of items in the List ["v1", "v2", …]
	Values []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{5}
}

func (x *List) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *List) GetCapacity() int64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *List) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the List to get
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{6}
}

func (x *GetListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The List to update
	List *List `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
	// Required. Mask (list) of fields to update.
	// Fields are specified relative to the List
	// (e.g. `capacity`, `values`; *not* `List.capacity` or `List.values`).
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateListRequest) Reset() {
	*x = UpdateListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateListRequest) ProtoMessage() {}

func (x *UpdateListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateListRequest.ProtoReflect.Descriptor instead.
func (*UpdateListRequest) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateListRequest) GetList() *List {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *UpdateListRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type AddListValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the List to add a value to.
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddListValueRequest) Reset() {
	*x = AddListValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddListValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddListValueRequest) ProtoMessage() {}

func (x *AddListValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddListValueRequest.ProtoReflect.Descriptor instead.
func (*AddListValueRequest) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{8}
}

func (x *AddListValueRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddListValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RemoveListValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the List to remove a value from.
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *RemoveListValueRequest) Reset() {
	*x = RemoveListValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beta_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveListValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveListValueRequest) ProtoMessage() {}

func (x *RemoveListValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beta_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveListValueRequest.ProtoReflect.Descriptor instead.
func (*RemoveListValueRequest) Descriptor() ([]byte, []int) {
	return file_beta_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveListValueRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoveListValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_beta_proto protoreflect.FileDescriptor

var file_beta_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x67,
	0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62, 0x65, 0x74,
	0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x7c, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x3a, 0x2b, 0xea,
	0x41, 0x28, 0x0a, 0x12, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x7d, 0x22, 0x88, 0x02, 0x0a, 0x14, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x66, 0x66,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x66,
	0x66, 0x3a, 0x52, 0xea, 0x41, 0x4f, 0x0a, 0x1f, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64,
	0x65, 0x76, 0x2f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x7b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x7d, 0x22, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41,
	0x14, 0x0a, 0x12, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x87, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x28,
	0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x21, 0x0a, 0x1f, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x14, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x72,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x3a, 0x22,
	0xea, 0x41, 0x1f, 0x0a, 0x0f, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x6c, 0x69, 0x73,
	0x74, 0x7d, 0x22, 0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x11, 0x0a, 0x0f, 0x61, 0x67, 0x6f,
	0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x18, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x11, 0x0a, 0x0f, 0x61, 0x67, 0x6f, 0x6e,
	0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x73, 0x6b, 0x22, 0x5f, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xe2, 0x41, 0x01, 0x02, 0xfa,
	0x41, 0x11, 0x0a, 0x0f, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x62, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xe2,
	0x41, 0x01, 0x02, 0xfa, 0x41, 0x11, 0x0a, 0x0f, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64,
	0x65, 0x76, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xcb, 0x06, 0x0a, 0x03, 0x53, 0x44,
	0x4b, 0x12, 0x7b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x26, 0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x27, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0xbc,
	0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x29, 0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x67,
	0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62, 0x65, 0x74,
	0x61, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x62, 0xda, 0x41, 0x14, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x32, 0x2d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x14, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6f, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65,
	0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x24, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x8c,
	0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e,
	0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x3b, 0xda, 0x41, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x32, 0x1a, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x7e, 0x0a,
	0x0c, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x2e,
	0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x3a, 0x61, 0x64, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01,
	0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x2b, 0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x67, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x51, 0x5a, 0x06, 0x2e, 0x2f, 0x62, 0x65, 0x74,
	0x61, 0x92, 0x41, 0x46, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x0f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x73, 0x65, 0x74, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_beta_proto_rawDescOnce sync.Once
	file_beta_proto_rawDescData = file_beta_proto_rawDesc
)

func file_beta_proto_rawDescGZIP() []byte {
	file_beta_proto_rawDescOnce.Do(func() {
		file_beta_proto_rawDescData = protoimpl.X.CompressGZIP(file_beta_proto_rawDescData)
	})
	return file_beta_proto_rawDescData
}

var file_beta_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_beta_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: agones.dev.sdk.beta.Empty
	(*Counter)(nil),                // 1: agones.dev.sdk.beta.Counter
	(*CounterUpdateRequest)(nil),   // 2: agones.dev.sdk.beta.CounterUpdateRequest
	(*GetCounterRequest)(nil),      // 3: agones.dev.sdk.beta.GetCounterRequest
	(*UpdateCounterRequest)(nil),   // 4: agones.dev.sdk.beta.UpdateCounterRequest
	(*List)(nil),                   // 5: agones.dev.sdk.beta.List
	(*GetListRequest)(nil),         // 6: agones.dev.sdk.beta.GetListRequest
	(*UpdateListRequest)(nil),      // 7: agones.dev.sdk.beta.UpdateListRequest
	(*AddListValueRequest)(nil),    // 8: agones.dev.sdk.beta.AddListValueRequest
	(*RemoveListValueRequest)(nil), // 9: agones.dev.sdk.beta.RemoveListValueRequest
	(*wrapperspb.Int64Value)(nil),  // 10: google.protobuf.Int64Value
	(*fieldmaskpb.FieldMask)(nil),  // 11: google.protobuf.FieldMask
}
var file_beta_proto_depIdxs = []int32{
	10, // 0: agones.dev.sdk.beta.CounterUpdateRequest.count:type_name -> google.protobuf.Int64Value
	10, // 1: agones.dev.sdk.beta.CounterUpdateRequest.capacity:type_name -> google.protobuf.Int64Value
	2,  // 2: agones.dev.sdk.beta.UpdateCounterRequest.counterUpdateRequest:type_name -> agones.dev.sdk.beta.CounterUpdateRequest
	5,  // 3: agones.dev.sdk.beta.UpdateListRequest.list:type_name -> agones.dev.sdk.beta.List
	11, // 4: agones.dev.sdk.beta.UpdateListRequest.update_mask:type_name -> google.protobuf.FieldMask
	3,  // 5: agones.dev.sdk.beta.SDK.GetCounter:input_type -> agones.dev.sdk.beta.GetCounterRequest
	4,  // 6: agones.dev.sdk.beta.SDK.UpdateCounter:input_type -> agones.dev.sdk.beta.UpdateCounterRequest
	6,  // 7: agones.dev.sdk.beta.SDK.GetList:input_type -> agones.dev.sdk.beta.GetListRequest
	7,  // 8: agones.dev.sdk.beta.SDK.UpdateList:input_type -> agones.dev.sdk.beta.UpdateListRequest
	8,  // 9: agones.dev.sdk.beta.SDK.AddListValue:input_type -> agones.dev.sdk.beta.AddListValueRequest
	9,  // 10: agones.dev.sdk.beta.SDK.RemoveListValue:input_type -> agones.dev.sdk.beta.RemoveListValueRequest
	1,  // 11: agones.dev.sdk.beta.SDK.GetCounter:output_type -> agones.dev.sdk.beta.Counter
	1,  // 12: agones.dev.sdk.beta.SDK.UpdateCounter:output_type -> agones.dev.sdk.beta.Counter
	5,  // 13: agones.dev.sdk.beta.SDK.GetList:output_type -> agones.dev.sdk.beta.List
	5,  // 14: agones.dev.sdk.beta.SDK.UpdateList:output_type -> agones.dev.sdk.beta.List
	5,  // 15: agones.dev.sdk.beta.SDK.AddListValue:output_type -> agones.dev.sdk.beta.List
	5,  // 16: agones.dev.sdk.beta.SDK.RemoveListValue:output_type -> agones.dev.sdk.beta.List
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_beta_proto_init() }
func file_beta_proto_init() {
	if File_beta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_beta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
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
		file_beta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterUpdateRequest); i {
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
		file_beta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCounterRequest); i {
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
		file_beta_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCounterRequest); i {
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
		file_beta_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
		file_beta_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_beta_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateListRequest); i {
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
		file_beta_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddListValueRequest); i {
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
		file_beta_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveListValueRequest); i {
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
			RawDescriptor: file_beta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_beta_proto_goTypes,
		DependencyIndexes: file_beta_proto_depIdxs,
		MessageInfos:      file_beta_proto_msgTypes,
	}.Build()
	File_beta_proto = out.File
	file_beta_proto_rawDesc = nil
	file_beta_proto_goTypes = nil
	file_beta_proto_depIdxs = nil
}
