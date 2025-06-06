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
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v6.30.0
// source: beta.proto

package beta

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SDKClient is the client API for SDK service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SDKClient interface {
	// Gets a Counter. Returns NOT_FOUND if the Counter does not exist.
	GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*Counter, error)
	// UpdateCounter returns the updated Counter. Returns NOT_FOUND if the Counter does not exist (name cannot be updated).
	// Returns OUT_OF_RANGE if the Count is out of range [0,Capacity].
	UpdateCounter(ctx context.Context, in *UpdateCounterRequest, opts ...grpc.CallOption) (*Counter, error)
	// Gets a List. Returns NOT_FOUND if the List does not exist.
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*List, error)
	// UpdateList returns the updated List. Returns NOT_FOUND if the List does not exist (name cannot be updated).
	// **THIS WILL OVERWRITE ALL EXISTING LIST.VALUES WITH ANY REQUEST LIST.VALUES**
	// Use AddListValue() or RemoveListValue() for modifying the List.Values field.
	// Returns INVALID_ARGUMENT if the field mask path(s) are not field(s) of the List.
	// If a field mask path(s) is specified, but the value is not set in the request List object,
	// then the default value for the variable will be set (i.e. 0 for "capacity", empty list for "values").
	UpdateList(ctx context.Context, in *UpdateListRequest, opts ...grpc.CallOption) (*List, error)
	// Adds a value to a List and returns updated List. Returns NOT_FOUND if the List does not exist.
	// Returns ALREADY_EXISTS if the value is already in the List.
	// Returns OUT_OF_RANGE if the List is already at Capacity.
	AddListValue(ctx context.Context, in *AddListValueRequest, opts ...grpc.CallOption) (*List, error)
	// Removes a value from a List and returns updated List. Returns NOT_FOUND if the List does not exist.
	// Returns NOT_FOUND if the value is not in the List.
	RemoveListValue(ctx context.Context, in *RemoveListValueRequest, opts ...grpc.CallOption) (*List, error)
}

type sDKClient struct {
	cc grpc.ClientConnInterface
}

func NewSDKClient(cc grpc.ClientConnInterface) SDKClient {
	return &sDKClient{cc}
}

func (c *sDKClient) GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*Counter, error) {
	out := new(Counter)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.beta.SDK/GetCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) UpdateCounter(ctx context.Context, in *UpdateCounterRequest, opts ...grpc.CallOption) (*Counter, error) {
	out := new(Counter)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.beta.SDK/UpdateCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.beta.SDK/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) UpdateList(ctx context.Context, in *UpdateListRequest, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.beta.SDK/UpdateList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) AddListValue(ctx context.Context, in *AddListValueRequest, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.beta.SDK/AddListValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) RemoveListValue(ctx context.Context, in *RemoveListValueRequest, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.beta.SDK/RemoveListValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SDKServer is the server API for SDK service.
// All implementations should embed UnimplementedSDKServer
// for forward compatibility
type SDKServer interface {
	// Gets a Counter. Returns NOT_FOUND if the Counter does not exist.
	GetCounter(context.Context, *GetCounterRequest) (*Counter, error)
	// UpdateCounter returns the updated Counter. Returns NOT_FOUND if the Counter does not exist (name cannot be updated).
	// Returns OUT_OF_RANGE if the Count is out of range [0,Capacity].
	UpdateCounter(context.Context, *UpdateCounterRequest) (*Counter, error)
	// Gets a List. Returns NOT_FOUND if the List does not exist.
	GetList(context.Context, *GetListRequest) (*List, error)
	// UpdateList returns the updated List. Returns NOT_FOUND if the List does not exist (name cannot be updated).
	// **THIS WILL OVERWRITE ALL EXISTING LIST.VALUES WITH ANY REQUEST LIST.VALUES**
	// Use AddListValue() or RemoveListValue() for modifying the List.Values field.
	// Returns INVALID_ARGUMENT if the field mask path(s) are not field(s) of the List.
	// If a field mask path(s) is specified, but the value is not set in the request List object,
	// then the default value for the variable will be set (i.e. 0 for "capacity", empty list for "values").
	UpdateList(context.Context, *UpdateListRequest) (*List, error)
	// Adds a value to a List and returns updated List. Returns NOT_FOUND if the List does not exist.
	// Returns ALREADY_EXISTS if the value is already in the List.
	// Returns OUT_OF_RANGE if the List is already at Capacity.
	AddListValue(context.Context, *AddListValueRequest) (*List, error)
	// Removes a value from a List and returns updated List. Returns NOT_FOUND if the List does not exist.
	// Returns NOT_FOUND if the value is not in the List.
	RemoveListValue(context.Context, *RemoveListValueRequest) (*List, error)
}

// UnimplementedSDKServer should be embedded to have forward compatible implementations.
type UnimplementedSDKServer struct {
}

func (UnimplementedSDKServer) GetCounter(context.Context, *GetCounterRequest) (*Counter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCounter not implemented")
}
func (UnimplementedSDKServer) UpdateCounter(context.Context, *UpdateCounterRequest) (*Counter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCounter not implemented")
}
func (UnimplementedSDKServer) GetList(context.Context, *GetListRequest) (*List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedSDKServer) UpdateList(context.Context, *UpdateListRequest) (*List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateList not implemented")
}
func (UnimplementedSDKServer) AddListValue(context.Context, *AddListValueRequest) (*List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddListValue not implemented")
}
func (UnimplementedSDKServer) RemoveListValue(context.Context, *RemoveListValueRequest) (*List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveListValue not implemented")
}

// UnsafeSDKServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SDKServer will
// result in compilation errors.
type UnsafeSDKServer interface {
	mustEmbedUnimplementedSDKServer()
}

func RegisterSDKServer(s grpc.ServiceRegistrar, srv SDKServer) {
	s.RegisterService(&SDK_ServiceDesc, srv)
}

func _SDK_GetCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).GetCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.beta.SDK/GetCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).GetCounter(ctx, req.(*GetCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_UpdateCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).UpdateCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.beta.SDK/UpdateCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).UpdateCounter(ctx, req.(*UpdateCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.beta.SDK/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_UpdateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).UpdateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.beta.SDK/UpdateList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).UpdateList(ctx, req.(*UpdateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_AddListValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddListValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).AddListValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.beta.SDK/AddListValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).AddListValue(ctx, req.(*AddListValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_RemoveListValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveListValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).RemoveListValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.beta.SDK/RemoveListValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).RemoveListValue(ctx, req.(*RemoveListValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SDK_ServiceDesc is the grpc.ServiceDesc for SDK service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SDK_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agones.dev.sdk.beta.SDK",
	HandlerType: (*SDKServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCounter",
			Handler:    _SDK_GetCounter_Handler,
		},
		{
			MethodName: "UpdateCounter",
			Handler:    _SDK_UpdateCounter_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SDK_GetList_Handler,
		},
		{
			MethodName: "UpdateList",
			Handler:    _SDK_UpdateList_Handler,
		},
		{
			MethodName: "AddListValue",
			Handler:    _SDK_AddListValue_Handler,
		},
		{
			MethodName: "RemoveListValue",
			Handler:    _SDK_RemoveListValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beta.proto",
}
